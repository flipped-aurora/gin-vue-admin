import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'

const DEFAULT_LLM_TIMEOUT = 1000 * 60 * 10

const LLM_AUTO_URL = '/autoCode/llmAuto'
const LLM_AUTO_SSE_URL = '/autoCode/llmAutoSSE'

const timerHost = typeof window !== 'undefined' ? window : globalThis

// 取首个非空字符串字段
const firstText = (...values) =>
  values.find((item) => typeof item === 'string' && item.trim()) || ''

// 合并流式文本，兼容「增量拼接」与「快照覆盖」两种上游风格
const mergeStreamText = (current, incoming) => {
  const next = String(incoming || '')
  if (!next) return current
  if (!current) return next
  if (next === current) return current
  if (next.length > current.length && next.startsWith(current)) return next
  return `${current}${next}`
}

// 从 SSE 事件 payload 里挑出文本字段（兼容多种字段名）
const pickStreamText = (payload) => {
  if (typeof payload === 'string') return payload
  if (!payload || typeof payload !== 'object') return ''
  return firstText(
    payload.answer,
    payload.text,
    payload.content,
    payload.output,
    payload.delta,
    payload.chunk
  )
}

// 构造带鉴权的 SSE 请求头（EventSource 不支持自定义 header，故用 fetch）
const createLLMFetchHeaders = (extraHeaders = {}) => {
  const userStore = useUserStore()
  return {
    'Content-Type': 'application/json',
    Accept: 'text/event-stream',
    'x-token': userStore.token || '',
    'x-user-id': userStore.userInfo?.ID || userStore.userInfo?.id || '',
    ...extraHeaders
  }
}

// 解析非 SSE 响应体为 JSON
const parseFetchBody = async (response) => {
  const text = await response.text()
  if (!text) return null
  try {
    return JSON.parse(text)
  } catch (e) {
    return { code: -1, msg: text }
  }
}

// 由 fetch 响应构造带上游错误信息的 Error
const buildFetchError = async (response) => {
  const body = await parseFetchBody(response)
  const message =
    (body && (body.msg || body.message)) ||
    `LLM request failed: HTTP ${response.status}`
  return new Error(message)
}

// 组装流式结果
const buildLLMStreamResult = (state) => {
  const payload =
    state.lastPayload && typeof state.lastPayload === 'object'
      ? { ...state.lastPayload }
      : {}
  if (state.answerText) payload.answer = state.answerText
  if (state.conversationId) payload.conversation_id = state.conversationId
  if (state.messageId) payload.message_id = state.messageId
  return payload
}

// 流式请求核心：fetch + ReadableStream + TextDecoder 手动解析 SSE
const streamLLMRequest = async (url, data, options = {}) => {
  const controller = options.controller || new AbortController()
  const timeout = options.timeout ?? DEFAULT_LLM_TIMEOUT
  let timeoutId = null
  let timeoutTriggered = false

  if (timeout > 0) {
    timeoutId = timerHost.setTimeout(() => {
      timeoutTriggered = true
      controller.abort()
    }, timeout)
  }

  const state = {
    answerText: '',
    conversationId: '',
    messageId: '',
    lastPayload: null
  }

  const handleSSEEvent = (event) => {
    const eventName = firstText(event.event, event.type) || 'message'
    const text = pickStreamText(event)

    if (event && typeof event === 'object') {
      state.lastPayload = event
      state.conversationId =
        firstText(event.conversation_id, event.conversationId) ||
        state.conversationId
      state.messageId =
        firstText(event.message_id, event.messageId) || state.messageId
    }
    if (text) {
      state.answerText = mergeStreamText(state.answerText, text)
    }

    if (eventName === 'error' || event?.error || event?.status === 'error') {
      throw new Error(
        firstText(event?.message, event?.msg, event?.error) ||
          'LLM stream failed'
      )
    }

    options.onMessage?.({
      event: eventName,
      data: event,
      rawData: JSON.stringify(event),
      text: state.answerText,
      conversationId: state.conversationId,
      messageId: state.messageId
    })
  }

  try {
    const fetchUrl = `${import.meta.env.VITE_BASE_API || ''}${url}`
    console.debug('[SSE] fetch →', fetchUrl)

    const response = await fetch(fetchUrl, {
      method: 'POST',
      headers: createLLMFetchHeaders(options.headers),
      body: JSON.stringify({
        ...data,
        response_mode: data?.response_mode || 'streaming'
      }),
      signal: controller.signal
    })

    if (!response.ok) {
      throw await buildFetchError(response)
    }

    const contentType = String(
      response.headers.get('content-type') || ''
    ).toLowerCase()
    // 上游非 SSE：降级为普通 JSON 解析
    if (!contentType.includes('text/event-stream') || !response.body) {
      console.warn(
        '[SSE] 响应非 SSE 格式，Content-Type:',
        contentType,
        '| 降级为普通 JSON 解析'
      )
      const body = await parseFetchBody(response)
      if (typeof body?.code !== 'undefined' && body.code !== 0) {
        throw new Error(body.msg || 'LLM request failed')
      }
      return body
    }
    console.debug('[SSE] 已进入流式读取模式')

    const reader = response.body.getReader()
    const decoder = new TextDecoder()
    let buffer = ''
    let done = false

    while (!done) {
      const result = await reader.read()
      done = result.done
      if (done) break

      buffer += decoder.decode(result.value, { stream: true })
      const lines = buffer.split('\n')
      // 保留最后一条可能不完整的行
      buffer = lines.pop() || ''

      for (const line of lines) {
        const trimmed = line.trim()
        if (!trimmed) continue
        if (trimmed.startsWith('data:')) {
          const dataStr = trimmed.slice(5).trim()
          if (!dataStr || dataStr === '[DONE]') continue
          try {
            const event = JSON.parse(dataStr)
            handleSSEEvent(event)
          } catch (e) {
            // 非 JSON data 行：当作裸文本累计
            if (dataStr) {
              state.answerText += dataStr
              options.onMessage?.({
                event: 'message',
                data: dataStr,
                rawData: dataStr,
                text: state.answerText,
                conversationId: state.conversationId,
                messageId: state.messageId
              })
            }
          }
        }
      }
    }

    // EOF 后冲刷残留
    if (buffer.trim()) {
      const trimmed = buffer.trim()
      if (trimmed.startsWith('data:')) {
        const dataStr = trimmed.slice(5).trim()
        if (dataStr && dataStr !== '[DONE]') {
          try {
            const event = JSON.parse(dataStr)
            handleSSEEvent(event)
          } catch {
            state.answerText += dataStr
          }
        }
      }
    }

    return buildLLMStreamResult(state)
  } catch (error) {
    if (controller.signal.aborted && timeoutTriggered) {
      throw new Error('LLM stream request timed out')
    }
    throw error
  } finally {
    if (timeoutId) timerHost.clearTimeout(timeoutId)
  }
}

// 通用 SSE 流式请求导出
export const llmAutoSSEStream = (data, options = {}) =>
  streamLLMRequest(LLM_AUTO_SSE_URL, data, options)

// picture 页专用：预设 mode 为 newCreateWeb
export const createWebStream = (data, options = {}) =>
  llmAutoSSEStream({ mode: 'newCreateWeb', ...data }, options)

export const preview = (data) => {
  return service({
    url: '/autoCode/preview',
    method: 'post',
    data
  })
}

export const createTemp = (data) => {
  return service({
    url: '/autoCode/createTemp',
    method: 'post',
    data
  })
}

export const getDB = (params) => {
  return service({
    url: '/autoCode/getDB',
    method: 'get',
    params
  })
}

export const getTable = (params) => {
  return service({
    url: '/autoCode/getTables',
    method: 'get',
    params
  })
}

export const getColumn = (params) => {
  return service({
    url: '/autoCode/getColumn',
    method: 'get',
    params
  })
}

export const getSysHistory = (data) => {
  return service({
    url: '/autoCode/getSysHistory',
    method: 'post',
    data
  })
}

export const rollback = (data) => {
  return service({
    url: '/autoCode/rollback',
    method: 'post',
    data
  })
}

export const getMeta = (data) => {
  return service({
    url: '/autoCode/getMeta',
    method: 'post',
    data
  })
}

export const delSysHistory = (data) => {
  return service({
    url: '/autoCode/delSysHistory',
    method: 'post',
    data
  })
}

export const createPackageApi = (data) => {
  return service({
    url: '/autoCode/createPackage',
    method: 'post',
    data
  })
}

export const getPackageApi = () => {
  return service({
    url: '/autoCode/getPackage',
    method: 'post'
  })
}

export const deletePackageApi = (data) => {
  return service({
    url: '/autoCode/delPackage',
    method: 'post',
    data
  })
}

export const getTemplatesApi = () => {
  return service({
    url: '/autoCode/getTemplates',
    method: 'get'
  })
}

export const installPlug = (data) => {
  return service({
    url: '/autoCode/installPlug',
    method: 'post',
    data
  })
}

export const pubPlug = (params) => {
  return service({
    url: '/autoCode/pubPlug',
    method: 'post',
    params
  })
}

export const llmAuto = (data, options = {}) => {
  return service({
    url: LLM_AUTO_URL,
    method: 'post',
    data: { ...data },
    timeout: options.timeout ?? DEFAULT_LLM_TIMEOUT,
    loadingOption: {
      lock: true,
      fullscreen: true,
      persistLoading: true,
      text: '小淼正在思考，请稍候...'
    },
    donNotShowLoading: options.donNotShowLoading ?? false
  })
}

export const addFunc = (data) => {
  return service({
    url: '/autoCode/addFunc',
    method: 'post',
    data
  })
}

export const initMenu = (data) => {
  return service({
    url: '/autoCode/initMenu',
    method: 'post',
    data
  })
}

export const initAPI = (data) => {
  return service({
    url: '/autoCode/initAPI',
    method: 'post',
    data
  })
}

export const initDictionary = (data) => {
  return service({
    url: '/autoCode/initDictionary',
    method: 'post',
    data
  })
}

export const mcp = (data) => {
  return service({
    url: '/autoCode/mcp',
    method: 'post',
    data
  })
}

export const mcpStatus = () => {
  return service({
    url: '/autoCode/mcpStatus',
    method: 'post'
  })
}

export const mcpStart = () => {
  return service({
    url: '/autoCode/mcpStart',
    method: 'post'
  })
}

export const mcpStop = () => {
  return service({
    url: '/autoCode/mcpStop',
    method: 'post'
  })
}

export const mcpList = (data) => {
  return service({
    url: '/autoCode/mcpList',
    method: 'post',
    data
  })
}

export const mcpTest = (data) => {
  return service({
    url: '/autoCode/mcpTest',
    method: 'post',
    data
  })
}

export const getPluginList = (params) => {
  return service({
    url: '/autoCode/getPluginList',
    method: 'get',
    params
  })
}

export const removePlugin = (params) => {
  return service({
    url: '/autoCode/removePlugin',
    method: 'post',
    params
  })
}

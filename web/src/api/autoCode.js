import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'

const DEFAULT_LLM_TIMEOUT = 1000 * 60 * 10

const LLM_AUTO_URL = '/autoCode/llmAuto'
const LLM_AUTO_SSE_URL = '/autoCode/llmAutoSSE'
const timerHost = typeof window !== 'undefined' ? window : globalThis

const firstText = (...values) =>
  values.find((item) => typeof item === 'string' && item.trim()) || ''

const mergeStreamText = (current, incoming) => {
  const next = String(incoming || '')
  if (!next) return current
  if (!current) return next
  if (next === current) return current
  if (next.length > current.length && next.startsWith(current)) return next
  return `${current}${next}`
}

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

const parseFetchBody = async (response) => {
  const contentType = String(response.headers.get('content-type') || '').toLowerCase()
  if (contentType.includes('application/json')) {
    return response.json()
  }
  return response.text()
}

const buildFetchError = async (response) => {
  const body = await parseFetchBody(response)
  const message =
    firstText(body?.msg, body?.message, body?.error) ||
    (typeof body === 'string' ? body.trim() : '') ||
    response.statusText ||
    'Request failed'
  return new Error(message)
}

const buildLLMStreamResult = (state) => {
  const payload =
    state.lastPayload && typeof state.lastPayload === 'object'
      ? { ...state.lastPayload }
      : {}

  if (state.answerText) {
    if (!firstText(payload.answer, payload.text, payload.content, payload.output)) {
      payload.answer = state.answerText
    } else {
      payload.answer = state.answerText
    }
  }
  if (state.conversationId && !firstText(payload.conversation_id, payload.conversationId)) {
    payload.conversation_id = state.conversationId
  }
  if (state.messageId && !firstText(payload.message_id, payload.messageId)) {
    payload.message_id = state.messageId
  }
  return payload
}

const createLLMFetchHeaders = (extraHeaders = {}) => {
  const userStore = useUserStore()
  return {
    'Content-Type': 'application/json',
    Accept: 'text/event-stream',
    'x-token': userStore.token || '',
    'x-user-id': userStore.userInfo.ID || userStore.userInfo.id || '',
    ...extraHeaders
  }
}

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

export const saveAIWorkflowSession = (data) => {
  return service({
    url: '/autoCode/saveAIWorkflowSession',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}

export const getAIWorkflowSessionList = (data) => {
  return service({
    url: '/autoCode/getAIWorkflowSessionList',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}

export const getAIWorkflowSessionDetail = (data) => {
  return service({
    url: '/autoCode/getAIWorkflowSessionDetail',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}

export const deleteAIWorkflowSession = (data) => {
  return service({
    url: '/autoCode/deleteAIWorkflowSession',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}

export const dumpAIWorkflowMarkdown = (data) => {
  return service({
    url: '/autoCode/dumpAIWorkflowMarkdown',
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

    if (
      eventName === 'error' ||
      event?.error ||
      event?.status === 'error'
    ) {
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

    const contentType = String(response.headers.get('content-type') || '').toLowerCase()
    if (!contentType.includes('text/event-stream') || !response.body) {
      console.warn('[SSE] 响应非 SSE 格式，Content-Type:', contentType, '| 降级为普通 JSON 解析')
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
      // 最后一个元素可能是不完整的行，保留到下次
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
            // 非 JSON 数据，作为文本处理
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

    // 处理 buffer 中剩余的数据
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

export const llmAutoStream = async (data, options = {}) =>
  streamLLMRequest(LLM_AUTO_URL, data, options)

export const llmAutoSSEStream = async (data, options = {}) =>
  streamLLMRequest(LLM_AUTO_SSE_URL, data, options)

export const analyzeRequirementByAI = (data, options = {}) => {
  return llmAuto(
    {
      mode: 'analysisChat',
      ...data
    },
    {
      ...options,
      donNotShowLoading: options.donNotShowLoading ?? true
    }
  )
}

export const analyzeRequirementByAIStream = (data, options = {}) => {
  return llmAutoStream(
    {
      mode: 'analysisChat',
      ...data
    },
    options
  )
}

export const analyzeRequirementByAISSEStream = (data, options = {}) => {
  return llmAutoSSEStream(
    {
      mode: 'analysisChat',
      ...data
    },
    options
  )
}

export const generatePromptFlowByAI = (data, options = {}) => {
  return llmAuto(
    {
      mode: 'workflowPromptChat',
      ...data
    },
    {
      ...options,
      donNotShowLoading: options.donNotShowLoading ?? true
    }
  )
}

export const generatePromptFlowByAIStream = (data, options = {}) => {
  return llmAutoStream(
    {
      mode: 'workflowPromptChat',
      ...data
    },
    options
  )
}

export const generatePromptFlowByAISSEStream = (data, options = {}) => {
  return llmAutoSSEStream(
    {
      mode: 'workflowPromptChat',
      ...data
    },
    options
  )
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

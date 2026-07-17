// 把配置对象序列化成缩进 2 空格的 JSON 文本，供 <pre> 直接展示 / 复制。
const jsonBlock = (obj) => JSON.stringify(obj, null, 2)

export const buildClientConfigs = ({ serverName, url, authHeader, token } = {}) => {
  const name = serverName || 'gva'
  const serverUrl = url || 'http://127.0.0.1:8889/mcp'
  const header = authHeader || 'x-token'
  const tokenValue = token || 'YOUR_GVA_TOKEN'

  return [
    {
      key: 'claude-code',
      label: 'Claude Code',
      location: '项目级 .mcp.json，或用户级 ~/.claude.json',
      note: '原生支持 HTTP 传输，无需桥接；也可用命令：claude mcp add --transport http <名称> <url> --header "x-token: <token>"。修改后重启会话或用 /mcp 查看状态。',
      lang: 'json',
      content: jsonBlock({
        mcpServers: {
          [name]: {
            type: 'http',
            url: serverUrl,
            headers: { [header]: tokenValue }
          }
        }
      })
    },
    {
      key: 'claude-desktop',
      label: 'Claude Desktop',
      location:
        'Windows：%APPDATA%\\Claude\\claude_desktop_config.json ；macOS：~/Library/Application Support/Claude/claude_desktop_config.json',
      note: '桌面版配置文件仅支持 stdio，远程 HTTP 服务需通过 npx mcp-remote 桥接（需先装 Node）。修改后要完全退出并重启 Claude Desktop。',
      lang: 'json',
      content: jsonBlock({
        mcpServers: {
          [name]: {
            command: 'npx',
            args: [
              '-y',
              'mcp-remote',
              serverUrl,
              '--transport',
              'http-only',
              '--header',
              `${header}:${tokenValue}`
            ]
          }
        }
      })
    },
    {
      key: 'codex',
      label: 'Codex CLI',
      location: '~/.codex/config.toml（Windows：C:\\Users\\<你>\\.codex\\config.toml）',
      note: 'TOML 格式，顶层键是 mcp_servers（下划线）。自定义请求头必须通过 http_headers 传递（bearer_token 只能发 Authorization: Bearer）。修改后重启 Codex。',
      lang: 'toml',
      content: [
        '# 启用 Codex 原生 Streamable-HTTP（rmcp）客户端；此行须位于 [mcp_servers.*] 之上',
        'experimental_use_rmcp_client = true',
        '',
        `[mcp_servers.${name}]`,
        `url = "${serverUrl}"`,
        '# 自定义静态请求头（非 Authorization），只能通过 http_headers 传递',
        `http_headers = { "${header}" = "${tokenValue}" }`
      ].join('\n')
    },
    {
      key: 'cursor',
      label: 'Cursor',
      location: '项目级 .cursor/mcp.json，或全局 ~/.cursor/mcp.json',
      note: '顶层键是 mcpServers；有 url 即视为远程 HTTP 服务，无需 type 字段。保存后到 Settings → MCP 确认为已启用状态。',
      lang: 'json',
      content: jsonBlock({
        mcpServers: {
          [name]: {
            url: serverUrl,
            headers: { [header]: tokenValue }
          }
        }
      })
    },
    {
      key: 'trae',
      label: 'Trae',
      location: 'Trae 设置 → MCP → 手动配置（写入 …/Trae/User/globalStorage/mcp.json）',
      note: '与 Cursor 同格式，顶层键 mcpServers，有 url 即远程 HTTP，自定义头直接写在 headers 里。需 Trae v1.3.0+。',
      lang: 'json',
      content: jsonBlock({
        mcpServers: {
          [name]: {
            url: serverUrl,
            headers: { [header]: tokenValue }
          }
        }
      })
    },
    {
      key: 'vscode',
      label: 'VS Code',
      location: '工作区 .vscode/mcp.json（GitHub Copilot 智能体模式，需 VS Code 1.102+）',
      note: '注意顶层键是 servers（不是 mcpServers）；远程 HTTP 需显式声明 type: http。保存后点条目上方的 Start 启动。',
      lang: 'json',
      content: jsonBlock({
        servers: {
          [name]: {
            type: 'http',
            url: serverUrl,
            headers: { [header]: tokenValue }
          }
        }
      })
    },
    {
      key: 'cline',
      label: 'Cline',
      location: 'Cline 侧栏 → MCP Servers → Configure MCP Servers（cline_mcp_settings.json）',
      note: 'type 必须为 streamableHttp（驼峰）；省略会退回旧版 SSE 传输，连不上仅支持 Streamable HTTP 的服务。保存即热加载。',
      lang: 'json',
      content: jsonBlock({
        mcpServers: {
          [name]: {
            type: 'streamableHttp',
            url: serverUrl,
            headers: { [header]: tokenValue },
            disabled: false,
            autoApprove: []
          }
        }
      })
    }
  ]
}

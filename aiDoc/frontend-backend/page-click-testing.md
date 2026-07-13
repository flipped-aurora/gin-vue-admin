# 页面点触测试(AI 驱动浏览器验证)

前端页面改动需要真实浏览器验证(点菜单、填表单、断言渲染)时,AI 按本文档获取登录态并执行点触。适用于 Playwright、chrome-devtools MCP 等任何浏览器自动化环境。

## 环境前置:浏览器自动化能力(没有就建议安装,经确认再装)

点触测试需要 AI 能驱动浏览器。开始前先自检当前环境有没有可用能力(按所用 AI 工具而定:MCP 浏览器服务、内置浏览器技能、项目内 playwright 依赖等)。**没有时主动建议用户安装**(首选 Playwright),用以下话术,经用户确认后再执行安装:

> 点触测试需要浏览器自动化能力,当前环境没有检测到。**建议安装**,推荐按顺序选:
> **A.(推荐)** Playwright MCP(`@playwright/mcp`,适合支持 MCP 的 AI 工具,配置到你的工具即可)
> **B.** 项目本地 Playwright 依赖(`cd web && npm i -D playwright && npx playwright install chromium`,写脚本驱动,跨工具通用)
> **C.** 你所用 AI 工具自带的浏览器技能/插件(按该工具的安装方式)
> **D.** 都不装——我给你一份人工目测清单,你自己过一遍

- 用户选 B 时注意:`web/package.json` 会因此变更,提醒用户这属于开发依赖改动,是否入库由用户定。
- 用户选 D(或拒绝)时,回退为输出「人工目测清单」:逐条列出要看什么、怎么操作、预期是什么(参照本仓库以往验收清单的粒度)。

## 登录态原理(代码事实)

```
localStorage['token'] ──→ userStore.token(pinia/modules/user.js, useStorage 启动时读初值)
                            ├─→ permission.js 路由守卫: 有 token → GetUserInfo → 拉菜单 → 放行
                            └─→ utils/request.js: 每个请求自动带 'x-token' 请求头
响应头 new-token(临期换发)→ request.js 自动写回 localStorage → 长会话不掉线
```

只要 localStorage 有一个有效 JWT,前端即完整登录态;路由是 hash 模式(`/#/...`)。

## 方式一:token 注入(日常点触,首选)

### token 获取优先级(AI 按序执行,每一步拿到后都先验活再用)

> 验活方法:带 `x-token` 请求头调一次 `GET /user/getUserInfo`,code===0 即有效。

1. **读本地约定文件 `<项目根>/.local/gva-test-token`**(整个文件就是一段裸 JWT,可有首尾空白)。有效则静默使用,不打扰用户。
   - `.local/gva-test-token` 是本项目**跨 AI 工具的唯一约定源**。各 AI 工具自己的 MCP/配置文件(如 `.mcp.json`、`.cursor/mcp.json`)里可能恰好有 x-token,但那是工具私有约定、形态各异,**不作为流程依据**,也不要主动去翻。
2. **自主登录获取**(需「安全配置 → 验证码阈值」已调大,见方式二):用测试账号直接 `POST /base/login {username, password}`(阈值内免验证码),响应 `data.token` 即全新 token。全自动,无需用户参与。
3. 以上全部不可用 → **在会话里用以下标准话术向用户索取**(不要自己编话术,保持一致体验):

> 我需要一个登录 token 做页面点触测试,两种取法(任选):
> **A.** 已登录的浏览器里:F12 → Application → Local Storage → 当前站点 → 复制 `token` 的值
> **B.** 系统页面「系统工具 → API Token」签发一个(建议 1~7 天有效期)
> 拿到后直接贴给我。想以后免打扰,可以把它存成项目根 `.local/gva-test-token` 文件(该目录已 gitignore,不会入库),我会自动读取。

4. 用户贴出 token 后,建议(经用户同意)替用户写入 `.local/gva-test-token`,下次免问。

补充:若自动化环境挂接的是用户**正在使用的浏览器**(remote-debugging 连接真实 Chrome),AI 可直接 `evaluate localStorage.getItem('token')` 读取当前登录态;常态下自动化是独立浏览器实例,读不到用户会话——这是安全隔离,按上面优先级取 token 即可。

### AI 注入操作

```js
// Playwright: 在应用 JS 执行前写入(useStorage 只在启动时读初值)
await context.addInitScript((tok) => localStorage.setItem('token', tok), token)
await page.goto('http://localhost:8080/#/layout/dashboard')
```

无 addInitScript 能力的环境等价两步:先开一次站点任意页 → `evaluate` 执行 `localStorage.setItem('token', '<jwt>')` → 再跳目标路由或刷新。

注入后守卫自动 GetUserInfo 进入页面,登录页不会出现;接口请求自动携带 `x-token`,与真人会话无差别。

### 局限

- 后端必须在跑(GetUserInfo 要通);
- 权限视角 = 该 token 所属账号的角色(要测普通角色就用普通账号的 token);
- **覆盖不到登录链路本身**(验证码、账号锁定、首次登录改密)→ 用方式二。

## 方式二:验证码阈值直登(测登录链路时)

- 后台「系统设置 → 安全配置」里的**验证码阈值**(`CaptchaOpen`)语义:`0 = 每次都要验证码;N>0 = 同 IP 错 N 次后才要`(server/api/v1/system/sys_user.go 登录逻辑)。
- 本地开发临时把阈值调大(如 99):登录页不再出现验证码输入框(`/base/captcha` 返回 OpenCaptcha=false),AI 可全真实路径点触登录:输入账号密码 → 点登录。
- **测完改回 0**。该配置存 DB(sys_security_config),只影响当前环境;生产环境严禁调大。

## 安全红线

- token 是真实凭证,等同账号密码:**不写入任何会提交的文件、不出现在截图/录屏/日志/commit message 里**;只允许存在 `.local/`(已 gitignore)或会话内。
- 疑似泄露:改该账号密码或等 JWT 自然过期(默认 7 天,见 config.yaml jwt.expires-time)。
- 点触测试造成的数据写入发生在真实库:破坏性操作(删除、重置)先问用户,或用专门的测试账号/测试数据。

# 主题类名参考（语义 token）

本文件是项目「主题化 UnoCSS 类名」的唯一参考。凡在 `.vue` / `core/componentLibrary` 里写颜色、
背景、描边、阴影，**优先用这里的语义 token 类名**，不要再散写 `bg-white dark:bg-slate-900`、
`text-gray-500 dark:text-gray-400` 这类明暗成对的裸色阶。

## 为什么用语义 token 而不是裸色阶

- **自带明暗**：每个 token 编译成 CSS 变量（如 `--container-bg-color`），`:root` 存亮色值、
  `html.dark` 存暗色值。写一次 `bg-container`，明暗自动切换——**不需要再写 `dark:` 变体**。
- **跟随换肤**：语义色板与中性/表面色全部由 Prism（`@simple-prism/core`）在 `theme/token.js`
  的 `generateTheme` 里按当前主色实时求解；用户换主题色 / 换预设时，这些类名的取色一起变。
- **单一事实源**：取色逻辑集中在 `theme/token.js`，页面只引用语义名，改配色不用全站搜替换。

> 落地链路：`token.js` 生成 hex → `shared.js` 转成 `r g b` 裸通道写入 `--x-color` 变量 →
> `vars.js` 声明 `token → rgb(var(--x-color))` 映射 → UnoCSS 据此生成 `bg-/text-/border-` 原子类。

## 一、语义色板（primary / info / success / warning / error）

每个语义色各有 **11 阶**（`50 100 200 300 400 500 600 700 800 900 950`），另有一个不带阶号的
「裸名」等于该色的 **500 阶**。

| 类名 | 取色来源 | 明暗表现 | 推荐场景 |
| --- | --- | --- | --- |
| `bg-primary` / `text-primary` / `border-primary` | 主色种子色（= `primary-500`） | Prism 把品牌色钉在 500，**明暗两态一致**，与 `--el-color-primary` 同源 | 主按钮、选中态、强调色、菜单项高亮实底 |
| `bg-primary-50` … `bg-primary-950` | Prism 明/暗两套 `scales.primary` 的对应阶 | 同一阶号在亮/暗下各取「该模式的正确色」，**不是**简单反色 | 需要浅底/深底层次时（如 hover 浅底 `primary-50`、描边 `primary-600`） |
| `*-info` `*-success` `*-warning` `*-error` 同上 | 各语义 `scales[name]` | 同上 | info=信息蓝、success=成功绿、warning=警告黄、error=危险红（删除/登出/校验错） |

要点：
- `info` 默认可「跟随主色」（外观设置里 `信息色跟随主色`），开启后 `info` 系列 = `primary` 系列。
- 阶号语义：**数字越小越浅**。亮色下 `50` 近乎白、`950` 近乎黑；暗色场景由暗色 scale 保证观感一致。
- 危险语义统一用 `error`（`text-error` / `bg-error` / `bg-error/10`），不要写裸 `text-red-500`。

## 二、中性 / 表面色（container / layout / border / muted / 文本）

这一组来自 Prism 的 **neutral 色阶**（镜像色阶，同一阶号在明暗各给正确色），在 `token.js`
的 `pickSurface` 里挑选。**只有亮色 `container` 是字面纯白**（neutral 最浅阶仍偏灰，抬升表面需纯白纸感）。

| 类名 | 取色来源（亮 / 暗） | 明暗表现 | 推荐场景 |
| --- | --- | --- | --- |
| `bg-container` | 亮=纯白 `#fff` / 暗=neutral 50（最深面） | 抬升的「纸面」 | 卡片、侧栏、抽屉、下拉浮层等**浮在页面之上**的容器底 |
| `bg-layout` / `bg-main` | 亮=neutral 50 / 暗=neutral 100 | 比 container 低一层的画布 | 整体页面背景、内容区画布（`bg-main` 是 `bg-layout` 的别名） |
| `bg-muted` | neutral 100 | 弱化底 | 分段控件槽、次级区块底色（注意亮色≈slate-50，白底上对比很弱，见下方「叠加高亮」例外） |
| `border-border` | neutral 200 | 描边色 | 卡片/输入框/分隔线的边框（`border-border`、`bg-border` 做 1px 分隔） |
| `bg-control-track` | neutral 300 | 控件未激活轨道 | 开关关闭态、滑块未填充轨道 |
| `text-base-text` | neutral `textContrast`（APCA 高对比求解） | 主文本，明暗都保证达标对比 | 正文、标题、主要文字 |
| `text-muted-foreground` | neutral `text`（APCA 低对比求解） | 次文本，比主文本弱 | 辅助说明、时间、占位、次级标签 |

层次关系（从下到上）：`layout/main`（画布） → `muted`（弱化块） → `container`（抬升卡片/浮层）。
一句话选择：**页面底用 `bg-main`，卡片/浮层用 `bg-container`，主文字 `text-base-text`，次文字 `text-muted-foreground`，边框 `border-border`。**

## 三、进度条 / 主色别名

| 类名 | 取色来源 | 说明 |
| --- | --- | --- |
| `text-active` / `shadow-active` | `--primary-color`（= 主色 500） | 主色别名，语义偏「激活/高亮」，如链接 hover `hover:text-active` |
| `nprogress`（`bg-nprogress`） | 主色 500 | 顶部加载进度条专用 |
| `border-table-border` | `--el-border-color-lighter` | 表格线，桥接 Element Plus 变量 |

## 四、阴影档位

来自 `theme/settings.js` 的 `tokens.light/dark.boxShadow`（与颜色无关，明暗两套值）。

| 类名 | 场景 |
| --- | --- |
| `shadow-header` | 顶栏底部投影 |
| `shadow-sider` | 侧栏 / 浮层投影（下拉菜单面板也用它盖过 EP 浮层） |
| `shadow-tab` | 标签栏投影 |
| `shadow-card` | 卡片投影 |

## 五、组合快捷类（shortcuts，`gva-*`）

定义在 `uno.config.js`，把常用的 token 串固化下来，避免复制粘贴：

| 快捷类 | 等价内容 / 用途 |
| --- | --- |
| `gva-surface` | `p-4 my-2 bg-container text-base-text`——业务卡片盒子的统一表面 |
| `gva-tool-btn` | 顶栏工具图标按钮（36px 扁平点击区，hover 叠加灰底） |
| `gva-menu-item` | 顶栏用户菜单项（图标+文案，hover / 键盘聚焦 / 展开高亮） |
| `gva-menu-item-danger` | 菜单项危险态（退出登录等，`text-error` + 红色淡底） |
| `gva-menu-sep` | 菜单分隔线（0.5px `bg-border`） |
| `gva-menu-label` | 菜单内不可交互的身份信息 label |

## 六、例外：hover / 高亮用「半透明叠加」而非 `bg-muted`

顶栏工具按钮、菜单项的 hover / 高亮底色**刻意用 `hover:bg-black/10 dark:hover:bg-white/10`**
（半透明黑/白叠加），这是**唯一推荐写 `dark:` 变体的场景**。原因：`muted` 亮色值≈slate-50，
与白色顶栏/面板几乎同色，hover 底看不出来；而黑/白叠加与底色无关，明暗两态都清晰可见。

同理，`data-[highlighted]:bg-primary`（下拉菜单项高亮）走主色实底 + 白字，是 reka 组件高亮的既定做法。

## 七、使用约定（强制）

1. **配置 / 框架外壳层**（顶栏、菜单、布局、设置抽屉、仪表盘等）：颜色一律用本文件的语义 token，
   **禁止**新写 `bg-white dark:bg-slate-900`、`text-gray-500 dark:text-gray-400` 这类裸色阶成对写法。
2. **不要手写 `dark:` 颜色变体**去覆盖 token——token 自带明暗；只有「半透明叠加高亮」例外。
3. 业务页面（CRUD / 业务表单弹窗）**不强制**，但新写时同样鼓励用语义 token，减少后续换肤成本。
4. 装饰性固定设计（如仪表盘深色推广卡）、纯示意的迷你预览图（布局/标签预览 wireframe）可保留
   字面色，它们不参与换肤，属既定视觉。
5. 新增语义 token 时：改 `theme/token.js`（取色）+ `theme/vars.js`（声明映射），不要在页面里硬编码。

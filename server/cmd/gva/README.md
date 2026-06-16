# GVA 动态 CLI

这是 Gin-Vue-Admin 的命令行客户端（基于 Cobra）。它根据一份 **manifest** 动态生成子命令，每个子命令对应后台的一个 API，调用后把返回结果直接打印到终端。

CLI 的程序名、可用命令都由 manifest 决定，所以同一份程序可以服务不同的 CLI 配置。

## 一、获取可执行程序

### 方式一：管理后台一键编译（推荐）

系统工具 → AI CLI管理 → 选择某个 CLI →「预览命令」→ 右上角选择**目标平台 / 架构** → 点击「编译下载」。

后台会把该 CLI 的 manifest 用 `go:embed` 内嵌进本目录的源码副本，交叉编译出一个**单二进制**：

- 下载即用，**不需要** `--manifest`、不需要额外文件。
- 文件名和程序名都是 manifest 里的主命令（例如 `opsctl.exe`）。

> 前提：运行后台的服务器需安装 Go 工具链（`go` 在 PATH 中），且后端工作目录在 `server/`。

### 方式二：手动编译

```bash
cd server
go build -o gva ./cmd/gva
```

手动编译的版本**不带内嵌 manifest**，运行时需要用 `--manifest` 指定清单文件（可从「预览命令 → 下载 Manifest」获取）。

## 二、使用

下面以程序名 `gva` 为例；后台编译版的程序名是你 CLI 的主命令（如 `opsctl`），把 `gva` 换成它即可。

### 1. 登录（保存登录态）

```bash
gva --manifest xxx.manifest.json login --token <你的JWT>
```

JWT 在 Gin-Vue-Admin 登录后获得（也可用系统工具里的 API Token）。登录信息写入 `~/.gva/config.json`，**只需执行一次**。

> 后台编译版（内嵌 manifest）不用带 `--manifest`：`opsctl login --token <JWT>`

### 2. 查看可用命令

```bash
gva --help            # 列出所有动态命令
gva <命令> --help      # 查看某个命令需要哪些参数
```

### 3. 调用命令

```bash
gva <命令> [--参数 值 ...]
```

例如：

```bash
gva user-list --page 1 --pageSize 10
```

命令会按 manifest 里定义的 `method + path` 去请求配置好的服务地址，返回体原样打印到终端。

## 三、配置文件

登录和运行参数都存在 `~/.gva/config.json`：

| 字段           | 说明                                   |
| ------------ | ------------------------------------ |
| `baseURL`     | 后台服务地址（优先取 manifest 里的 server.baseURL） |
| `token`       | JWT，`login` 时写入                        |
| `authHeader`  | 认证头名称，默认 `x-token`                    |
| `manifestPath`| manifest 文件路径（文件模式用，`login` 时写入）      |

## 四、全局参数

| 参数             | 作用                 |
| -------------- | ------------------ |
| `--manifest`   | 指定 manifest 文件路径    |
| `--base-url`   | 覆盖服务地址             |
| `--auth-header`| 覆盖认证头名称            |

## 五、两种运行模式

- **内嵌模式**（后台编译版）：manifest 已编进二进制，直接 `gva login` → `gva <命令>`。
- **文件模式**（手动编译版）：需 `--manifest xxx.manifest.json` 加载清单；第一次带 `--manifest login` 后路径会写入配置，之后可省略。

## 六、典型完整流程

```bash
# 1. 从管理后台「预览命令 → 编译下载」得到 opsctl.exe

# 2. 登录
opsctl.exe login --token eyJhbGciOi...

# 3. 用命令
opsctl.exe --help
opsctl.exe user-list --page 1
```

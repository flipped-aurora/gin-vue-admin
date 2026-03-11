---
name: mysql-debugger
description: MySQL 数据库调试助手，用于连接测试、执行 SQL 查询、查看表结构和数据调试。当用户需要测试数据库连接、查询数据、查看表结构或调试数据库问题时使用此 skill。
---

# MySQL 数据库调试助手

提供 MySQL 数据库连接测试、SQL 查询执行、表结构查看和数据调试功能。

## 快速使用（已预编译）

Windows 可执行文件已预编译好，直接使用：

```bash
# 测试连接（自动读取 server/config.debug.yaml）
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -test

# 执行查询
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query "SELECT * FROM sys_users LIMIT 10" -json

# 交互模式
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe
```

### 常用命令

```bash
# 查看所有表（-query 模式支持）
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query ".tables" -json

# 查看表结构（-query 模式支持）
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query ".schema sys_users" -json

# 查看表索引
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query ".indexes sys_users" -json

# 指定配置文件
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -config server/config.debug.yaml -test
```

> **提示**: GORM 默认使用复数形式的表名（如 `hr_employees` 而非 `hr_employee`）。查询前建议先用 `.tables` 确认实际表名。

### 对话直接使用
> "帮我测试数据库连接"
> "查询 sys_users 表前10条数据"

## 功能

1. **连接测试** - 验证数据库连接并显示基本信息
2. **SQL 查询** - 执行 SELECT/INSERT/UPDATE/DELETE 等语句
3. **表结构查看** - 查看表字段、类型、索引等信息
4. **数据调试** - 快速查看和修改数据

## 参数说明

```bash
mysql-debugger [选项]

  -config string    配置文件路径 (默认自动检测 server/config.debug.yaml)
  -host string      数据库主机 (默认 "127.0.0.1")
  -port int         数据库端口 (默认 3306)
  -user string      用户名 (默认 "root")
  -password string  密码
  -database string  数据库名 (默认 "gva")
  -query string     执行单个查询后退出
  -json             JSON 格式输出
  -test             测试连接模式
```

### 交互模式命令

| 命令 | 说明 |
|------|------|
| `.tables` | 显示所有表 |
| `.schema <表名>` | 显示表结构 |
| `.indexes <表名>` | 显示表索引 |
| `.quit` | 退出 |
| SQL 语句 | 直接执行 SQL |

## 配置优先级

1. 命令行参数（最高优先级）
2. `-config` 指定的配置文件
3. 自动检测配置文件（按以下顺序查找）：
   - 当前目录: `server/config.debug.yaml` → `server/config.yaml` → `config.debug.yaml` → `config.yaml`
   - 可执行文件所在目录及其子目录
   - 可执行文件上级目录（项目根目录）
   - 用户主目录: `~/.mysql-debugger/config.yaml`
4. 默认本地连接参数（最低优先级）

## 配置示例

YAML 配置文件格式（config.debug.yaml）：

```yaml
mysql:
  path: 47.121.178.244      # 主机地址
  port: "3306"              # 端口
  config: charset=utf8mb4&parseTime=True&loc=Local
  db-name: gva              # 数据库名
  username: root            # 用户名
  password: your_password   # 密码
```

## 使用示例

### 测试数据库连接

```bash
# 使用默认配置
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -test

# 指定配置文件
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -config server/config.debug.yaml -test
```

### 执行 SQL 查询

```bash
# 查询数据（JSON 输出）
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query "SELECT * FROM sys_users LIMIT 10" -json

# 查看表结构
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query "DESCRIBE sys_users"

# 手动指定连接参数
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -host 47.121.178.244 -port 3306 -user root -password xxx -database gva -query "SELECT * FROM sys_users"
```

### 交互模式

```bash
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe

# 进入交互式命令行
mysql> .tables
mysql> .schema sys_users
mysql> SELECT * FROM sys_users WHERE id = 1
mysql> .quit
```

## 重新编译（如需修改）

```bash
cd .claude/skills/mysql-debugger/scripts
go build -o mysql-debugger.exe main.go
```

## 依赖

- 无需额外依赖，单文件可执行
- 如需重新编译：Go 1.21+, gorm.io/gorm, gorm.io/driver/mysql

## GORM 表名规则

**重要提示**: GORM 默认将结构体名转换为复数形式的表名

```go
// 模型定义
type Employee struct {
    // ...
}

// 实际表名: employees (复数形式)
```

### 常见转换示例

| 结构体名 | 默认表名 | 说明 |
|---------|---------|------|
| `User` | `users` | 规则 +s |
| `Employee` | `employees` | 规则 +es |
| `HrEmployee` | `hr_employees` | 蛇形命名 + 复数 |
| `SysUser` | `sys_users` | 蛇形命名 + 复数 |
| `Person` | `people` | 不规则变化 |
| `News` | `news` | 不可数名词保持单数 |

### 查询技巧

```bash
# 1. 先查看所有表，确认实际表名
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query ".tables" -json

# 2. 使用 LIKE 模糊匹配查找
.claude/skills/mysql-debugger/scripts/mysql-debugger.exe -query "SHOW TABLES LIKE '%employee%'" -json

# 3. 常见错误示例
# ❌ SELECT * FROM hr_employee    # 单数形式（错误）
# ✅ SELECT * FROM hr_employees   # 复数形式（正确）
```

### 自定义表名

如果模型使用了 `TableName()` 方法，则表名可能是单数：

```go
func (Employee) TableName() string {
    return "hr_employee"  // 显式指定为单数
}
```

在这种情况下，表名是 `hr_employee` 而不是 `hr_employees`。

## 安全提示

- 密码通过命令行参数传递，注意 shell 历史记录
- 生产环境谨慎执行 UPDATE/DELETE，建议先 SELECT 确认
- 脚本默认限制返回 100 行数据

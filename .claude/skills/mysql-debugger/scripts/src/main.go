package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config 数据库配置
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// getExecutableDir 获取可执行文件所在目录
func getExecutableDir() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(ex)
}

// findConfigFile 查找配置文件，尝试多个可能的路径
func findConfigFile() string {
	// 可能的配置文件路径（按优先级排序）
	possiblePaths := []string{
		// 当前工作目录下
		"server/config.debug.yaml",
		"server/config.yaml",
		"config.debug.yaml",
		"config.yaml",
		// 可执行文件所在目录下（技能目录）
		filepath.Join(getExecutableDir(), "server/config.debug.yaml"),
		filepath.Join(getExecutableDir(), "server/config.yaml"),
		filepath.Join(getExecutableDir(), "config.debug.yaml"),
		filepath.Join(getExecutableDir(), "config.yaml"),
		// 可执行文件上级目录（项目根目录）
		filepath.Join(getExecutableDir(), "../server/config.debug.yaml"),
		filepath.Join(getExecutableDir(), "../server/config.yaml"),
		filepath.Join(getExecutableDir(), "../config.debug.yaml"),
		filepath.Join(getExecutableDir(), "../config.yaml"),
		// 用户主目录下
		filepath.Join(os.Getenv("USERPROFILE"), ".mysql-debugger/config.yaml"),
		filepath.Join(os.Getenv("HOME"), ".mysql-debugger/config.yaml"),
	}

	for _, path := range possiblePaths {
		if path == "" {
			continue
		}
		// 转换为绝对路径并清理
		absPath, err := filepath.Abs(path)
		if err != nil {
			continue
		}
		if _, err := os.Stat(absPath); err == nil {
			return absPath
		}
	}
	return ""
}

func main() {
	var (
		configPath = flag.String("config", "", "配置文件路径")
		host       = flag.String("host", "127.0.0.1", "数据库主机")
		port       = flag.Int("port", 3306, "数据库端口")
		user       = flag.String("user", "root", "用户名")
		password   = flag.String("password", "", "密码")
		database   = flag.String("database", "gva", "数据库名")
		query      = flag.String("query", "", "执行单个查询")
		jsonOutput = flag.Bool("json", false, "JSON格式输出")
		testMode   = flag.Bool("test", false, "测试连接模式")
	)
	flag.Parse()

	// 加载配置
	cfg := &Config{
		Host:     *host,
		Port:     *port,
		User:     *user,
		Password: *password,
		Database: *database,
	}

	// 从配置文件读取
	if *configPath != "" {
		// 使用用户指定的配置文件
		absPath, err := filepath.Abs(*configPath)
		if err == nil {
			if c := loadConfig(absPath); c != nil {
				cfg = c
			}
		}
	} else {
		// 自动查找配置文件
		if foundPath := findConfigFile(); foundPath != "" {
			if c := loadConfig(foundPath); c != nil {
				fmt.Printf("使用配置文件: %s\n", foundPath)
				cfg = c
			}
		}
	}

	// 命令行参数覆盖配置文件
	if *host != "127.0.0.1" {
		cfg.Host = *host
	}
	if *port != 3306 {
		cfg.Port = *port
	}
	if *user != "root" {
		cfg.User = *user
	}
	if *password != "" {
		cfg.Password = *password
	}
	if *database != "gva" {
		cfg.Database = *database
	}

	if cfg.Password == "" {
		fmt.Println("错误: 未提供密码")
		os.Exit(1)
	}

	// 连接数据库
	db, err := connectDB(cfg)
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		os.Exit(1)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 根据模式执行
	if *testMode {
		testConnection(db, cfg)
	} else if *query != "" {
		executeQuery(db, *query, *jsonOutput)
	} else {
		interactiveMode(db)
	}
}

// connectDB 建立数据库连接
func connectDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// loadConfig 从 YAML 文件加载配置
func loadConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	cfg := &Config{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
	}

	content := string(data)
	inMySQLSection := false
	baseIndent := 0

	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "mysql:") {
			inMySQLSection = true
			baseIndent = len(line) - len(strings.TrimLeft(line, " \t"))
			continue
		}

		if inMySQLSection && trimmed != "" {
			currentIndent := len(line) - len(strings.TrimLeft(line, " \t"))
			if currentIndent <= baseIndent {
				break
			}

			if idx := strings.Index(trimmed, ":"); idx > 0 {
				key := trimmed[:idx]
				value := strings.TrimSpace(trimmed[idx+1:])
				value = strings.Trim(value, "\"'")

				switch key {
				case "path":
					cfg.Host = value
				case "port":
					fmt.Sscanf(value, "%d", &cfg.Port)
				case "username":
					cfg.User = value
				case "password":
					cfg.Password = value
				case "db-name":
					cfg.Database = value
				}
			}
		}
	}

	return cfg
}

// testConnection 测试数据库连接
func testConnection(db *gorm.DB, cfg *Config) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("       MySQL 数据库连接测试")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("\n主机: %s\n", cfg.Host)
	fmt.Printf("端口: %d\n", cfg.Port)
	fmt.Printf("数据库: %s\n", cfg.Database)
	fmt.Printf("用户: %s\n\n", cfg.User)

	// 测试连接
	fmt.Println("[1] 测试数据库连接...")
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("   ❌ 获取连接失败: %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	if err := sqlDB.Ping(); err != nil {
		fmt.Printf("   ❌ 连接失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("   ✅ 连接成功 (耗时: %.3fs)\n", time.Since(start).Seconds())

	// 获取版本
	fmt.Println("\n[2] MySQL 版本信息:")
	var version string
	db.Raw("SELECT VERSION()").Scan(&version)
	fmt.Printf("   📦 %s\n", version)

	// 获取当前数据库
	fmt.Println("\n[3] 当前数据库:")
	var currentDB string
	db.Raw("SELECT DATABASE()").Scan(&currentDB)
	fmt.Printf("   📁 %s\n", currentDB)

	// 获取所有表
	fmt.Println("\n[4] 数据库表列表:")
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)
	if len(tables) > 0 {
		for i, table := range tables {
			if i >= 20 {
				fmt.Printf("   ... 还有 %d 个表\n", len(tables)-20)
				break
			}
			fmt.Printf("   %2d. %s\n", i+1, table)
		}
	} else {
		fmt.Println("   (无表)")
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("       测试完成")
	fmt.Println(strings.Repeat("=", 50))
}

// executeQuery 执行单个查询
func executeQuery(db *gorm.DB, query string, jsonOutput bool) {
	query = strings.TrimSpace(query)
	upperQuery := strings.ToUpper(query)

	// 处理内部命令（即使在 -query 模式下也支持）
	switch {
	case query == ".tables":
		executeShowTables(db, jsonOutput)
		return
	case strings.HasPrefix(query, ".schema "):
		tableName := strings.TrimSpace(query[8:])
		executeShowSchema(db, tableName, jsonOutput)
		return
	case strings.HasPrefix(query, ".indexes "):
		tableName := strings.TrimSpace(query[9:])
		executeShowIndexes(db, tableName, jsonOutput)
		return
	}

	result := map[string]interface{}{
		"success": true,
	}

	if strings.HasPrefix(upperQuery, "SELECT") {
		var rows []map[string]interface{}
		tx := db.Raw(query).Scan(&rows)
		if tx.Error != nil {
			result["success"] = false
			result["error"] = tx.Error.Error()
		} else {
			result["data"] = rows
			result["row_count"] = len(rows)
			if len(rows) > 0 {
				columns := make([]string, 0, len(rows[0]))
				for k := range rows[0] {
					columns = append(columns, k)
				}
				result["columns"] = columns
			}
		}
	} else {
		tx := db.Exec(query)
		if tx.Error != nil {
			result["success"] = false
			result["error"] = tx.Error.Error()
		} else {
			result["affected_rows"] = tx.RowsAffected
		}
	}

	if jsonOutput {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	} else {
		formatResult(result)
	}
}

// formatResult 格式化输出结果
func formatResult(result map[string]interface{}) {
	if !result["success"].(bool) {
		fmt.Printf("查询失败: %v\n", result["error"])
		return
	}

	if _, ok := result["affected_rows"]; ok {
		fmt.Printf("执行成功，影响行数: %d\n", result["affected_rows"])
		return
	}

	rows, ok := result["data"].([]map[string]interface{})
	if !ok || len(rows) == 0 {
		fmt.Println("查询成功，返回 0 行数据")
		return
	}

	// 获取列名
	columns := result["columns"].([]string)

	// 计算列宽
	colWidths := make(map[string]int)
	for _, col := range columns {
		colWidths[col] = len(col)
	}

	for _, row := range rows {
		for _, col := range columns {
			valStr := fmt.Sprintf("%v", row[col])
			if len(valStr) > 50 {
				valStr = valStr[:50]
			}
			if len(valStr) > colWidths[col] {
				colWidths[col] = len(valStr)
			}
		}
	}

	// 输出表头
	var headerParts []string
	for _, col := range columns {
		headerParts = append(headerParts, padRight(col, colWidths[col]))
	}
	header := strings.Join(headerParts, " | ")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	// 输出数据
	limit := 100
	if len(rows) < limit {
		limit = len(rows)
	}
	for i := 0; i < limit; i++ {
		var rowParts []string
		for _, col := range columns {
			valStr := fmt.Sprintf("%v", rows[i][col])
			if len(valStr) > 50 {
				valStr = valStr[:50]
			}
			rowParts = append(rowParts, padRight(valStr, colWidths[col]))
		}
		fmt.Println(strings.Join(rowParts, " | "))
	}

	if len(rows) > 100 {
		fmt.Printf("\n... 还有 %d 行数据\n", len(rows)-100)
	}
	fmt.Printf("\n共 %d 行\n", len(rows))
}

// interactiveMode 交互模式
func interactiveMode(db *gorm.DB) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("    MySQL 查询执行器 - 交互模式")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("命令:")
	fmt.Println("  .tables          - 显示所有表")
	fmt.Println("  .schema <表名>   - 显示表结构")
	fmt.Println("  .indexes <表名>  - 显示表索引")
	fmt.Println("  .quit            - 退出")
	fmt.Println("  或直接输入 SQL 语句")
	fmt.Println(strings.Repeat("=", 50) + "\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("mysql> ")
		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		if command == "" {
			continue
		}

		switch {
		case command == ".quit":
			return
		case command == ".tables":
			showTables(db)
		case strings.HasPrefix(command, ".schema "):
			tableName := strings.TrimSpace(command[8:])
			showSchema(db, tableName)
		case strings.HasPrefix(command, ".indexes "):
			tableName := strings.TrimSpace(command[9:])
			showIndexes(db, tableName)
		default:
			executeQuery(db, command, false)
		}
		fmt.Println()
	}
}

// executeShowTables 执行显示所有表（支持 JSON 输出）
func executeShowTables(db *gorm.DB, jsonOutput bool) {
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)

	result := map[string]interface{}{
		"success":   true,
		"tables":    tables,
		"row_count": len(tables),
	}

	if jsonOutput {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Printf("\n数据库中的表 (%d 个):\n", len(tables))
		for i, table := range tables {
			fmt.Printf("  %2d. %s\n", i+1, table)
		}
		fmt.Println()
	}
}

// executeShowSchema 执行显示表结构（支持 JSON 输出）
func executeShowSchema(db *gorm.DB, tableName string, jsonOutput bool) {
	type ColumnInfo struct {
		ColumnName string `gorm:"column:COLUMN_NAME" json:"column_name"`
		DataType   string `gorm:"column:DATA_TYPE" json:"data_type"`
		IsNullable string `gorm:"column:IS_NULLABLE" json:"is_nullable"`
		Default    string `gorm:"column:COLUMN_DEFAULT" json:"default_value"`
		Comment    string `gorm:"column:COLUMN_COMMENT" json:"comment"`
		Extra      string `gorm:"column:EXTRA" json:"extra"`
	}

	var columns []ColumnInfo
	db.Raw(`
		SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT, EXTRA
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_NAME = ?
		ORDER BY ORDINAL_POSITION
	`, tableName).Scan(&columns)

	result := map[string]interface{}{
		"success":    true,
		"table_name": tableName,
		"columns":    columns,
		"row_count":  len(columns),
	}

	if len(columns) == 0 {
		result["success"] = false
		result["error"] = fmt.Sprintf("表 %s 不存在或没有列", tableName)
	}

	if jsonOutput {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	} else {
		if len(columns) == 0 {
			fmt.Printf("表 %s 不存在或没有列\n", tableName)
			return
		}
		fmt.Printf("\n表 %s 的结构:\n", tableName)
		fmt.Printf("%-20s %-15s %-10s %-20s\n", "字段名", "类型", "可空", "默认值")
		fmt.Println(strings.Repeat("-", 70))
		for _, col := range columns {
			fmt.Printf("%-20s %-15s %-10s %-20s\n",
				col.ColumnName, col.DataType, col.IsNullable, col.Default)
		}
		fmt.Println()
	}
}

// executeShowIndexes 执行显示表索引（支持 JSON 输出）
func executeShowIndexes(db *gorm.DB, tableName string, jsonOutput bool) {
	type IndexInfo struct {
		IndexName string `gorm:"column:INDEX_NAME" json:"index_name"`
		Column    string `gorm:"column:COLUMN_NAME" json:"column"`
		NonUnique int    `gorm:"column:NON_UNIQUE" json:"non_unique"`
	}

	var indexes []IndexInfo
	db.Raw(`
		SELECT INDEX_NAME, COLUMN_NAME, NON_UNIQUE
		FROM INFORMATION_SCHEMA.STATISTICS
		WHERE TABLE_NAME = ?
		ORDER BY INDEX_NAME, SEQ_IN_INDEX
	`, tableName).Scan(&indexes)

	result := map[string]interface{}{
		"success":    true,
		"table_name": tableName,
		"indexes":    indexes,
		"row_count":  len(indexes),
	}

	if len(indexes) == 0 {
		result["success"] = false
		result["error"] = fmt.Sprintf("表 %s 没有索引", tableName)
	}

	if jsonOutput {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	} else {
		if len(indexes) == 0 {
			fmt.Printf("表 %s 没有索引\n", tableName)
			return
		}
		fmt.Printf("\n表 %s 的索引:\n", tableName)
		fmt.Printf("%-20s %-20s %-10s\n", "索引名", "字段", "唯一")
		fmt.Println(strings.Repeat("-", 55))
		for _, idx := range indexes {
			unique := "否"
			if idx.NonUnique == 0 {
				unique = "是"
			}
			fmt.Printf("%-20s %-20s %-10s\n", idx.IndexName, idx.Column, unique)
		}
		fmt.Println()
	}
}

// showTables 显示所有表（交互模式，调用新的函数）
func showTables(db *gorm.DB) {
	executeShowTables(db, false)
}

// showSchema 显示表结构
func showSchema(db *gorm.DB, tableName string) {
	type ColumnInfo struct {
		ColumnName string `gorm:"column:COLUMN_NAME"`
		DataType   string `gorm:"column:DATA_TYPE"`
		IsNullable string `gorm:"column:IS_NULLABLE"`
		Default    string `gorm:"column:COLUMN_DEFAULT"`
		Comment    string `gorm:"column:COLUMN_COMMENT"`
		Extra      string `gorm:"column:EXTRA"`
	}

	var columns []ColumnInfo
	db.Raw(`
		SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT, EXTRA
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_NAME = ?
		ORDER BY ORDINAL_POSITION
	`, tableName).Scan(&columns)

	if len(columns) == 0 {
		fmt.Printf("表 %s 不存在或没有列\n", tableName)
		return
	}

	fmt.Printf("\n表 %s 的结构:\n", tableName)
	fmt.Printf("%-20s %-15s %-10s %-20s\n", "字段名", "类型", "可空", "默认值")
	fmt.Println(strings.Repeat("-", 70))
	for _, col := range columns {
		fmt.Printf("%-20s %-15s %-10s %-20s\n",
			col.ColumnName, col.DataType, col.IsNullable, col.Default)
	}
	fmt.Println()
}

// showIndexes 显示表索引
func showIndexes(db *gorm.DB, tableName string) {
	type IndexInfo struct {
		IndexName string `gorm:"column:INDEX_NAME"`
		Column    string `gorm:"column:COLUMN_NAME"`
		NonUnique int    `gorm:"column:NON_UNIQUE"`
	}

	var indexes []IndexInfo
	db.Raw(`
		SELECT INDEX_NAME, COLUMN_NAME, NON_UNIQUE
		FROM INFORMATION_SCHEMA.STATISTICS
		WHERE TABLE_NAME = ?
		ORDER BY INDEX_NAME, SEQ_IN_INDEX
	`, tableName).Scan(&indexes)

	if len(indexes) == 0 {
		fmt.Printf("表 %s 没有索引\n", tableName)
		return
	}

	fmt.Printf("\n表 %s 的索引:\n", tableName)
	fmt.Printf("%-20s %-20s %-10s\n", "索引名", "字段", "唯一")
	fmt.Println(strings.Repeat("-", 55))
	for _, idx := range indexes {
		unique := "否"
		if idx.NonUnique == 0 {
			unique = "是"
		}
		fmt.Printf("%-20s %-20s %-10s\n", idx.IndexName, idx.Column, unique)
	}
	fmt.Println()
}

// padRight 右填充字符串
func padRight(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

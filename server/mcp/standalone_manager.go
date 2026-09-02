package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const (
	mcpRuntimeDirName     = ".tmp"
	mcpRuntimeSubDir      = "mcp"
	mcpRuntimeMetaName    = "managed-process.json"
	mcpRuntimeLogName     = "mcp.log"
	mcpHealthCheckTimeout = 2 * time.Second
	mcpStartWaitTimeout   = 20 * time.Second
	mcpStopWaitTimeout    = 8 * time.Second
	mcpBuildTimeout       = 2 * time.Minute
)

type ManagedStandaloneStatus struct {
	State      string `json:"state"`
	Managed    bool   `json:"managed"`
	Reachable  bool   `json:"reachable"`
	Starting   bool   `json:"starting"`
	BaseURL    string `json:"baseURL"`
	HealthURL  string `json:"healthURL"`
	ListenAddr string `json:"listenAddr"`
	Path       string `json:"path"`
	AuthHeader string `json:"authHeader"`
	PID        int    `json:"pid,omitempty"`
	LogPath    string `json:"logPath,omitempty"`
	StartedAt  string `json:"startedAt,omitempty"`
	LastError  string `json:"lastError,omitempty"`
	Message    string `json:"message,omitempty"`
}

type managedProcessMeta struct {
	PID        int      `json:"pid"`
	StartedAt  string   `json:"startedAt"`
	LogPath    string   `json:"logPath"`
	ConfigPath string   `json:"configPath"`
	WorkDir    string   `json:"workDir"`
	Command    string   `json:"command"`
	Args       []string `json:"args"`
}

func ResolveMCPListenAddr() string {
	addr := global.GVA_CONFIG.MCP.Addr
	if addr <= 0 {
		addr = 8889
	}
	return fmt.Sprintf(":%d", addr)
}

func ResolveMCPPath() string {
	path := strings.TrimSpace(global.GVA_CONFIG.MCP.Path)
	if path == "" {
		path = "/mcp"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}

func ResolveMCPHealthURL() string {
	baseURL, err := url.Parse(ResolveMCPServiceURL())
	if err != nil || baseURL.Scheme == "" || baseURL.Host == "" {
		return fmt.Sprintf("http://127.0.0.1%s/health", ResolveMCPListenAddr())
	}
	baseURL.Path = "/health"
	baseURL.RawQuery = ""
	baseURL.Fragment = ""
	return baseURL.String()
}

func GetManagedStandaloneStatus(ctx context.Context) ManagedStandaloneStatus {
	reachable, reachErr := checkMCPHealth(ctx)
	meta, _ := readManagedProcessMeta()

	if meta != nil && meta.PID > 0 && !processExists(meta.PID) {
		_ = removeManagedProcessMeta()
		meta = nil
	}

	status := ManagedStandaloneStatus{
		Managed:    meta != nil,
		Reachable:  reachable,
		BaseURL:    ResolveMCPServiceURL(),
		HealthURL:  ResolveMCPHealthURL(),
		ListenAddr: ResolveMCPListenAddr(),
		Path:       ResolveMCPPath(),
		AuthHeader: ConfiguredAuthHeader(),
	}

	if meta != nil {
		status.PID = meta.PID
		status.LogPath = meta.LogPath
		status.StartedAt = meta.StartedAt
	}

	switch {
	case meta != nil && reachable:
		status.State = "running"
		status.Message = "MCP 独立服务运行中"
	case meta != nil && processExists(meta.PID):
		status.State = "starting"
		status.Managed = true
		status.Starting = true
		status.Message = "MCP 独立服务启动中"
	case reachable:
		status.State = "external"
		status.Managed = false
		status.Message = "检测到 MCP 服务已运行，但不是由页面启动的托管进程"
	case meta != nil:
		status.State = "stopped"
		status.Managed = false
		status.Message = "上次托管的 MCP 进程已退出"
	default:
		status.State = "stopped"
		status.Message = "MCP 独立服务未启动"
	}

	if !reachable && reachErr != nil && status.State != "stopped" {
		status.LastError = reachErr.Error()
	}

	return status
}

func StartManagedStandalone(ctx context.Context) (ManagedStandaloneStatus, error) {
	current := GetManagedStandaloneStatus(ctx)
	if current.Reachable {
		return current, nil
	}

	meta, _ := readManagedProcessMeta()
	if meta != nil && meta.PID > 0 && processExists(meta.PID) {
		return waitForManagedProcess(ctx, meta)
	}

	commandPath, commandArgs, workDir, configPath, err := resolveManagedStartCommand()
	if err != nil {
		return GetManagedStandaloneStatus(context.Background()), err
	}

	runtimeDir, err := ensureMCPRuntimeDir()
	if err != nil {
		return GetManagedStandaloneStatus(context.Background()), err
	}

	logPath := filepath.Join(runtimeDir, mcpRuntimeLogName)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return GetManagedStandaloneStatus(context.Background()), err
	}
	defer logFile.Close()

	cmd := exec.Command(commandPath, commandArgs...)
	cmd.Dir = workDir
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.Env = append(os.Environ(), "GVA_MCP_CONFIG="+configPath)
	prepareDetachedProcess(cmd)

	if err := cmd.Start(); err != nil {
		return GetManagedStandaloneStatus(context.Background()), fmt.Errorf("启动 MCP 独立服务失败: %w", err)
	}

	pid := cmd.Process.Pid
	_ = cmd.Process.Release()

	meta = &managedProcessMeta{
		PID:        pid,
		StartedAt:  time.Now().Format(time.RFC3339),
		LogPath:    logPath,
		ConfigPath: configPath,
		WorkDir:    workDir,
		Command:    commandPath,
		Args:       append([]string{}, commandArgs...),
	}
	if err := writeManagedProcessMeta(meta); err != nil {
		return GetManagedStandaloneStatus(context.Background()), err
	}

	return waitForManagedProcess(ctx, meta)
}

func StopManagedStandalone(ctx context.Context) (ManagedStandaloneStatus, error) {
	meta, err := readManagedProcessMeta()
	if err != nil {
		status := GetManagedStandaloneStatus(ctx)
		if status.Reachable {
			return status, errors.New("当前 MCP 服务不是由页面启动的，无法自动停用")
		}
		return status, nil
	}

	if meta.PID > 0 && processExists(meta.PID) {
		if err := terminateProcess(meta.PID); err != nil {
			return GetManagedStandaloneStatus(context.Background()), fmt.Errorf("停用 MCP 独立服务失败: %w", err)
		}
	}

	deadline := time.NewTimer(mcpStopWaitTimeout)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer deadline.Stop()
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return GetManagedStandaloneStatus(context.Background()), ctx.Err()
		case <-deadline.C:
			_ = removeManagedProcessMeta()
			return GetManagedStandaloneStatus(context.Background()), nil
		case <-ticker.C:
			if meta.PID <= 0 || !processExists(meta.PID) {
				_ = removeManagedProcessMeta()
				status := GetManagedStandaloneStatus(context.Background())
				if status.Reachable {
					status.State = "external"
					status.Message = "托管进程已停止，但检测到还有其他 MCP 服务在运行"
				}
				return status, nil
			}
		}
	}
}

func checkMCPHealth(ctx context.Context) (bool, error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), mcpHealthCheckTimeout)
	defer cancel()

	if ctx != nil {
		if deadline, ok := ctx.Deadline(); ok {
			timeoutCtx, cancel = context.WithDeadline(context.Background(), deadline)
			defer cancel()
		}
	}

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, ResolveMCPHealthURL(), nil)
	if err != nil {
		return false, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		return true, nil
	}

	return false, fmt.Errorf("MCP 健康检查失败: %s", resp.Status)
}

func waitForManagedProcess(ctx context.Context, meta *managedProcessMeta) (ManagedStandaloneStatus, error) {
	deadline := time.NewTimer(mcpStartWaitTimeout)
	ticker := time.NewTicker(300 * time.Millisecond)
	defer deadline.Stop()
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return GetManagedStandaloneStatus(context.Background()), ctx.Err()
		case <-deadline.C:
			return GetManagedStandaloneStatus(context.Background()), fmt.Errorf("等待 MCP 独立服务启动超时，请查看日志: %s", meta.LogPath)
		case <-ticker.C:
			current := GetManagedStandaloneStatus(context.Background())
			if current.Reachable {
				return current, nil
			}
			if meta.PID > 0 && !processExists(meta.PID) {
				return current, fmt.Errorf("MCP 独立进程已退出，请查看日志: %s", meta.LogPath)
			}
		}
	}
}

func resolveManagedStartCommand() (string, []string, string, string, error) {
	serverRoot := resolveMCPServerRoot()
	if serverRoot == "" {
		return "", nil, "", "", errors.New("未找到 server 根目录，无法启动 MCP 独立服务")
	}

	configPath, err := resolveMCPConfigPath(serverRoot)
	if err != nil {
		return "", nil, "", "", err
	}

	if explicit := strings.TrimSpace(os.Getenv("GVA_MCP_BIN")); explicit != "" {
		if !fileExists(explicit) {
			return "", nil, "", "", fmt.Errorf("GVA_MCP_BIN 指向的文件不存在: %s", explicit)
		}
		return explicit, []string{"-config", configPath}, filepath.Dir(explicit), configPath, nil
	}

	binaryPath, err := ensureManagedBinary(serverRoot)
	if err != nil {
		return "", nil, "", "", err
	}

	return binaryPath, []string{"-config", configPath}, serverRoot, configPath, nil
}

func ensureManagedBinary(serverRoot string) (string, error) {
	runtimeDir, err := ensureMCPRuntimeDir()
	if err != nil {
		return "", err
	}

	binaryPath := filepath.Join(runtimeDir, managedBinaryName())
	sourceDir := filepath.Join(serverRoot, "cmd", "mcp")

	goBin, lookErr := exec.LookPath("go")
	if lookErr == nil && isDir(sourceDir) {
		buildCtx, cancel := context.WithTimeout(context.Background(), mcpBuildTimeout)
		defer cancel()

		cmd := exec.CommandContext(buildCtx, goBin, "build", "-o", binaryPath, "./cmd/mcp")
		cmd.Dir = serverRoot
		output, err := cmd.CombinedOutput()
		if err != nil {
			message := strings.TrimSpace(string(output))
			if message != "" {
				return "", fmt.Errorf("构建 MCP 独立服务失败: %w, 输出: %s", err, message)
			}
			return "", fmt.Errorf("构建 MCP 独立服务失败: %w", err)
		}
		return binaryPath, nil
	}

	if fileExists(binaryPath) {
		return binaryPath, nil
	}

	return "", errors.New("未检测到可用的 Go 环境，且本地没有可复用的 MCP 独立二进制")
}

func resolveMCPServerRoot() string {
	root := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Root)
	serverDir := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Server)
	if serverDir == "" {
		serverDir = "server"
	}

	candidates := []string{}
	if root != "" {
		candidates = append(candidates, filepath.Join(root, filepath.FromSlash(serverDir)))
		candidates = append(candidates, root)
	}

	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates, cwd)
		candidates = append(candidates, filepath.Join(cwd, "server"))
	}

	for _, candidate := range candidates {
		if isDir(filepath.Join(candidate, "cmd", "mcp")) {
			return candidate
		}
	}

	if len(candidates) > 0 {
		return candidates[0]
	}

	return ""
}

func resolveMCPConfigPath(serverRoot string) (string, error) {
	candidates := []string{
		filepath.Join(serverRoot, "cmd", "mcp", "config.yaml"),
		filepath.Join(serverRoot, "config.yaml"),
	}

	for _, candidate := range candidates {
		if fileExists(candidate) {
			return candidate, nil
		}
	}

	return "", errors.New("未找到 MCP 配置文件，请确认 cmd/mcp/config.yaml 或 server/config.yaml 存在")
}

func ensureMCPRuntimeDir() (string, error) {
	runtimeDir := filepath.Join(resolveMCPProjectRoot(), mcpRuntimeDirName, mcpRuntimeSubDir)
	if err := os.MkdirAll(runtimeDir, 0o755); err != nil {
		return "", err
	}
	return runtimeDir, nil
}

func resolveMCPProjectRoot() string {
	root := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Root)
	if root != "" {
		return root
	}

	serverRoot := resolveMCPServerRoot()
	if serverRoot == "" {
		return "."
	}

	if fileExists(filepath.Join(serverRoot, "go.mod")) {
		return filepath.Dir(serverRoot)
	}

	return serverRoot
}

func readManagedProcessMeta() (*managedProcessMeta, error) {
	data, err := os.ReadFile(managedProcessMetaPath())
	if err != nil {
		return nil, err
	}

	var meta managedProcessMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func writeManagedProcessMeta(meta *managedProcessMeta) error {
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(managedProcessMetaPath(), data, 0o644)
}

func removeManagedProcessMeta() error {
	err := os.Remove(managedProcessMetaPath())
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

func managedProcessMetaPath() string {
	runtimeDir, err := ensureMCPRuntimeDir()
	if err != nil {
		return filepath.Join(resolveMCPProjectRoot(), mcpRuntimeDirName, mcpRuntimeSubDir, mcpRuntimeMetaName)
	}
	return filepath.Join(runtimeDir, mcpRuntimeMetaName)
}

func managedBinaryName() string {
	if runtime.GOOS == "windows" {
		return "gva-mcp-standalone.exe"
	}
	return "gva-mcp-standalone"
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

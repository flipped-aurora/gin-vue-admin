package system

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	autoCodeStagingPrefix = ".autocode-staging-"
	autoCodeTaskBackend   = "backend"
	autoCodeTaskFrontend  = "frontend"
)

var (
	errAutoCodeFileConflict    = errors.New("自动代码目标文件已被外部修改")
	errAutoCodeDuplicateTarget = errors.New("自动代码任务包含重复目标")
)

type autoCodeTaskLayout struct {
	root       string
	serverRoot string
	webRoot    string
}

type autoCodeFileTask struct {
	layout     autoCodeTaskLayout
	stagingDir string
	files      []autoCodeTaskFile
}

type autoCodeTaskFile struct {
	TargetPath    string
	Kind          string
	Existed       bool
	Mode          fs.FileMode
	BeforeHash    string
	AfterHash     string
	BeforeContent []byte
	StagedPath    string
}

type autoCodeFilePublisher func(file *autoCodeTaskFile) (published bool, err error)

func newAutoCodeTaskLayout(root, server, web string) (autoCodeTaskLayout, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return autoCodeTaskLayout{}, fmt.Errorf("解析自动代码根目录失败: %w", err)
	}
	serverRoot, err := pathWithin(root, server)
	if err != nil {
		return autoCodeTaskLayout{}, fmt.Errorf("服务端目录无效: %w", err)
	}
	webRoot, err := pathWithin(root, web)
	if err != nil {
		return autoCodeTaskLayout{}, fmt.Errorf("前端目录无效: %w", err)
	}
	return autoCodeTaskLayout{
		root:       filepath.Clean(root),
		serverRoot: serverRoot,
		webRoot:    webRoot,
	}, nil
}

func prepareAutoCodeFileTask(layout autoCodeTaskLayout, files map[string][]byte) (_ *autoCodeFileTask, err error) {
	stagingDir, err := os.MkdirTemp(layout.root, autoCodeStagingPrefix)
	if err != nil {
		return nil, fmt.Errorf("创建自动代码 staging 目录失败: %w", err)
	}
	task := &autoCodeFileTask{
		layout:     layout,
		stagingDir: stagingDir,
		files:      make([]autoCodeTaskFile, 0, len(files)),
	}
	defer func() {
		if err != nil {
			task.cleanup()
		}
	}()

	targets := make([]string, 0, len(files))
	normalizedFiles := make(map[string][]byte, len(files))
	for target, content := range files {
		absoluteTarget, absoluteErr := filepath.Abs(target)
		if absoluteErr != nil {
			return nil, fmt.Errorf("解析目标路径 %q 失败: %w", target, absoluteErr)
		}
		cleanTarget := filepath.Clean(absoluteTarget)
		if _, exists := normalizedFiles[cleanTarget]; exists {
			return nil, fmt.Errorf("%w: %s", errAutoCodeDuplicateTarget, cleanTarget)
		}
		if _, classifyErr := layout.classify(cleanTarget); classifyErr != nil {
			return nil, classifyErr
		}
		normalizedFiles[cleanTarget] = content
		targets = append(targets, cleanTarget)
	}
	sort.Slice(targets, func(i, j int) bool {
		leftKind, _ := layout.classify(targets[i])
		rightKind, _ := layout.classify(targets[j])
		if leftKind != rightKind {
			return leftKind == autoCodeTaskBackend
		}
		return targets[i] < targets[j]
	})

	for index, target := range targets {
		kind, _ := layout.classify(target)
		content := normalizedFiles[target]
		file := autoCodeTaskFile{
			TargetPath: target,
			Kind:       kind,
			Mode:       0o666,
			AfterHash:  hashAutoCodeContent(content),
			StagedPath: filepath.Join(stagingDir, fmt.Sprintf("%06d", index)),
		}
		stat, statErr := os.Stat(target)
		switch {
		case statErr == nil:
			if !stat.Mode().IsRegular() {
				return nil, fmt.Errorf("自动代码目标不是普通文件: %s", target)
			}
			before, readErr := os.ReadFile(target)
			if readErr != nil {
				return nil, fmt.Errorf("读取自动代码目标 %s 失败: %w", target, readErr)
			}
			file.Existed = true
			file.Mode = stat.Mode().Perm()
			file.BeforeContent = before
			file.BeforeHash = hashAutoCodeContent(before)
		case errors.Is(statErr, fs.ErrNotExist):
		default:
			return nil, fmt.Errorf("检查自动代码目标 %s 失败: %w", target, statErr)
		}
		if writeErr := replaceAutoCodeFileAtomically(file.StagedPath, content, file.Mode); writeErr != nil {
			return nil, fmt.Errorf("写入自动代码 staging 失败: %w", writeErr)
		}
		task.files = append(task.files, file)
	}
	return task, nil
}

func commitAutoCodeFileTask(task *autoCodeFileTask, publish autoCodeFilePublisher, persist func() error) error {
	if task == nil || publish == nil || persist == nil {
		return errors.New("自动代码任务提交参数不能为空")
	}
	defer task.cleanup()

	backendApplied, err := task.apply(autoCodeTaskBackend, publish)
	if err != nil {
		return joinAutoCodeRollbackError(err, task.rollback(backendApplied))
	}
	if err = persist(); err != nil {
		return joinAutoCodeRollbackError(err, task.rollback(backendApplied))
	}
	frontendApplied, err := task.apply(autoCodeTaskFrontend, publish)
	if err != nil {
		return joinAutoCodeRollbackError(
			fmt.Errorf("数据库已提交，前端文件发布失败: %w", err),
			task.rollback(frontendApplied),
		)
	}
	return nil
}

func (task *autoCodeFileTask) apply(kind string, publish autoCodeFilePublisher) ([]int, error) {
	applied := make([]int, 0, len(task.files))
	for index := range task.files {
		file := &task.files[index]
		if file.Kind != kind {
			continue
		}
		published, err := publish(file)
		if err != nil {
			return applied, err
		}
		if published {
			applied = append(applied, index)
		}
	}
	return applied, nil
}

func publishPreparedAutoCodeFile(file *autoCodeTaskFile) (bool, error) {
	currentHash, exists, err := hashAutoCodeTarget(file.TargetPath)
	if err != nil {
		return false, err
	}
	if exists && currentHash == file.AfterHash {
		if removeErr := os.Remove(file.StagedPath); removeErr != nil && !errors.Is(removeErr, fs.ErrNotExist) {
			return false, fmt.Errorf("清理重复 staging 文件失败: %w", removeErr)
		}
		file.StagedPath = ""
		return false, nil
	}
	if exists != file.Existed || (file.Existed && currentHash != file.BeforeHash) {
		return false, fmt.Errorf("%w: %s", errAutoCodeFileConflict, file.TargetPath)
	}
	if err = os.MkdirAll(filepath.Dir(file.TargetPath), 0o755); err != nil {
		return false, fmt.Errorf("创建目标目录 %s 失败: %w", filepath.Dir(file.TargetPath), err)
	}
	if err = os.Rename(file.StagedPath, file.TargetPath); err != nil {
		return false, fmt.Errorf("原子发布 %s 失败: %w", file.TargetPath, err)
	}
	file.StagedPath = ""
	return true, nil
}

func (task *autoCodeFileTask) rollback(applied []int) error {
	for i := len(applied) - 1; i >= 0; i-- {
		index := applied[i]
		if index < 0 || index >= len(task.files) {
			return fmt.Errorf("自动代码回滚索引无效: %d", index)
		}
		file := &task.files[index]
		currentHash, exists, err := hashAutoCodeTarget(file.TargetPath)
		if err != nil {
			return err
		}
		if file.Existed && exists && currentHash == file.BeforeHash {
			continue
		}
		if !exists || currentHash != file.AfterHash {
			return fmt.Errorf("%w: %s", errAutoCodeFileConflict, file.TargetPath)
		}
		if !file.Existed {
			if err = os.Remove(file.TargetPath); err != nil && !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("删除新建文件 %s 失败: %w", file.TargetPath, err)
			}
			continue
		}
		if err = replaceAutoCodeFileAtomically(file.TargetPath, file.BeforeContent, file.Mode); err != nil {
			return fmt.Errorf("恢复自动代码文件 %s 失败: %w", file.TargetPath, err)
		}
	}
	return nil
}

func (task *autoCodeFileTask) cleanup() {
	if task != nil && task.stagingDir != "" {
		_ = os.RemoveAll(task.stagingDir)
		task.stagingDir = ""
	}
}

func joinAutoCodeRollbackError(cause, rollbackErr error) error {
	if rollbackErr == nil {
		return cause
	}
	return errors.Join(cause, fmt.Errorf("回滚自动代码文件失败: %w", rollbackErr))
}

func replaceAutoCodeFileAtomically(target string, content []byte, mode fs.FileMode) error {
	if mode == 0 {
		mode = 0o666
	}
	dir := filepath.Dir(target)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("创建目标目录 %s 失败: %w", dir, err)
	}
	tmp, err := os.CreateTemp(dir, ".autocode-*")
	if err != nil {
		return fmt.Errorf("创建临时文件失败: %w", err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if _, err = tmp.Write(content); err == nil {
		err = tmp.Sync()
	}
	if closeErr := tmp.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		return fmt.Errorf("写入临时文件失败: %w", err)
	}
	if err = os.Chmod(tmpName, mode.Perm()); err != nil {
		return fmt.Errorf("设置临时文件权限失败: %w", err)
	}
	if err = os.Rename(tmpName, target); err != nil {
		return fmt.Errorf("原子替换 %s 失败: %w", target, err)
	}
	return nil
}

func (l autoCodeTaskLayout) classify(target string) (string, error) {
	if isPathWithin(l.serverRoot, target) {
		return autoCodeTaskBackend, nil
	}
	if isPathWithin(l.webRoot, target) {
		return autoCodeTaskFrontend, nil
	}
	return "", fmt.Errorf("自动代码目标不在服务端或前端目录内: %s", target)
}

func pathWithin(root string, elems ...string) (string, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	joined := filepath.Join(append([]string{root}, elems...)...)
	joined, err = filepath.Abs(joined)
	if err != nil {
		return "", err
	}
	if !isPathWithin(root, joined) {
		return "", fmt.Errorf("路径越过根目录: %s", joined)
	}
	return filepath.Clean(joined), nil
}

func isPathWithin(root, target string) bool {
	rel, err := filepath.Rel(filepath.Clean(root), filepath.Clean(target))
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) && !filepath.IsAbs(rel)
}

func hashAutoCodeTarget(target string) (string, bool, error) {
	content, err := os.ReadFile(target)
	if errors.Is(err, fs.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("读取自动代码目标 %s 失败: %w", target, err)
	}
	return hashAutoCodeContent(content), true, nil
}

func hashAutoCodeContent(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}

package system

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func testAutoCodeTaskLayout(t *testing.T) autoCodeTaskLayout {
	t.Helper()
	layout, err := newAutoCodeTaskLayout(t.TempDir(), "server", filepath.Join("web", "src"))
	if err != nil {
		t.Fatalf("newAutoCodeTaskLayout() error = %v", err)
	}
	return layout
}

func writeAutoCodeTaskTarget(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll(%q): %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile(%q): %v", path, err)
	}
}

func readAutoCodeTaskTarget(t *testing.T, path string) string {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q): %v", path, err)
	}
	return string(content)
}

func prepareRequestFileTask(t *testing.T, frontendCount int) (autoCodeTaskLayout, *autoCodeFileTask, map[string]string) {
	t.Helper()
	layout := testAutoCodeTaskLayout(t)
	targets := map[string]string{
		"backend-a":  filepath.Join(layout.serverRoot, "api", "a.go"),
		"backend-b":  filepath.Join(layout.serverRoot, "api", "b.go"),
		"frontend-a": filepath.Join(layout.webRoot, "view", "a.vue"),
	}
	files := map[string][]byte{
		targets["backend-a"]:  []byte("backend-a-after"),
		targets["backend-b"]:  []byte("backend-b-after"),
		targets["frontend-a"]: []byte("frontend-a-after"),
	}
	if frontendCount > 1 {
		targets["frontend-b"] = filepath.Join(layout.webRoot, "view", "b.vue")
		files[targets["frontend-b"]] = []byte("frontend-b-after")
	}
	for name, target := range targets {
		writeAutoCodeTaskTarget(t, target, name+"-before")
	}
	task, err := prepareAutoCodeFileTask(layout, files)
	if err != nil {
		t.Fatalf("prepareAutoCodeFileTask() error = %v", err)
	}
	return layout, task, targets
}

func TestPrepareAutoCodeFileTaskStagesWithoutChangingTargets(t *testing.T) {
	layout, task, targets := prepareRequestFileTask(t, 1)
	t.Cleanup(task.cleanup)

	if !isPathWithin(layout.root, task.stagingDir) || isPathWithin(layout.serverRoot, task.stagingDir) || isPathWithin(layout.webRoot, task.stagingDir) {
		t.Fatalf("staging directory %q must be under root and outside watched roots", task.stagingDir)
	}
	if len(task.files) != 3 {
		t.Fatalf("task file count = %d, want 3", len(task.files))
	}
	if task.files[0].Kind != autoCodeTaskBackend || task.files[1].Kind != autoCodeTaskBackend || task.files[2].Kind != autoCodeTaskFrontend {
		t.Fatalf("task files are not ordered backend before frontend: %+v", task.files)
	}
	for _, file := range task.files {
		if _, err := os.Stat(file.StagedPath); err != nil {
			t.Fatalf("staged file %q missing: %v", file.StagedPath, err)
		}
	}
	if got := readAutoCodeTaskTarget(t, targets["backend-a"]); got != "backend-a-before" {
		t.Fatalf("backend target changed during prepare: %q", got)
	}
	if got := readAutoCodeTaskTarget(t, targets["frontend-a"]); got != "frontend-a-before" {
		t.Fatalf("frontend target changed during prepare: %q", got)
	}
}

func TestPrepareAutoCodeFileTaskRejectsEquivalentDuplicateTargets(t *testing.T) {
	layout := testAutoCodeTaskLayout(t)
	target := filepath.Join(layout.serverRoot, "api", "order.go")
	separator := string(filepath.Separator)
	equivalent := filepath.Join(layout.serverRoot, "api") + separator + "nested" + separator + ".." + separator + "order.go"

	_, err := prepareAutoCodeFileTask(layout, map[string][]byte{
		target:     []byte("first"),
		equivalent: []byte("second"),
	})
	if !errors.Is(err, errAutoCodeDuplicateTarget) {
		t.Fatalf("prepareAutoCodeFileTask() error = %v, want %v", err, errAutoCodeDuplicateTarget)
	}
}

func TestCommitAutoCodeFileTaskPublishesFrontendAfterDatabase(t *testing.T) {
	_, task, _ := prepareRequestFileTask(t, 1)
	stagingDir := task.stagingDir
	var events []string
	publish := func(file *autoCodeTaskFile) (bool, error) {
		events = append(events, file.Kind+":"+filepath.Base(file.TargetPath))
		return publishPreparedAutoCodeFile(file)
	}
	err := commitAutoCodeFileTask(task, publish, func() error {
		events = append(events, "database")
		return nil
	})
	if err != nil {
		t.Fatalf("commitAutoCodeFileTask() error = %v", err)
	}
	want := []string{"backend:a.go", "backend:b.go", "database", "frontend:a.vue"}
	if len(events) != len(want) {
		t.Fatalf("events = %v, want %v", events, want)
	}
	for index := range want {
		if events[index] != want[index] {
			t.Fatalf("events = %v, want %v", events, want)
		}
	}
	if _, err = os.Stat(stagingDir); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("staging directory error = %v, want not exist", err)
	}
}

func TestCommitAutoCodeFileTaskRestoresBackendWhenDatabaseFails(t *testing.T) {
	_, task, targets := prepareRequestFileTask(t, 1)
	stagingDir := task.stagingDir
	wantFailure := errors.New("database failed")
	err := commitAutoCodeFileTask(task, publishPreparedAutoCodeFile, func() error { return wantFailure })
	if !errors.Is(err, wantFailure) {
		t.Fatalf("commitAutoCodeFileTask() error = %v, want %v", err, wantFailure)
	}
	for _, name := range []string{"backend-a", "backend-b", "frontend-a"} {
		if got := readAutoCodeTaskTarget(t, targets[name]); got != name+"-before" {
			t.Fatalf("%s target = %q, want before", name, got)
		}
	}
	if _, err = os.Stat(stagingDir); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("staging directory error = %v, want not exist", err)
	}
}

func TestCommitAutoCodeFileTaskRollsBackPartialBackendPhase(t *testing.T) {
	_, task, targets := prepareRequestFileTask(t, 1)
	wantFailure := errors.New("second backend failed")
	backendCalls := 0
	publish := func(file *autoCodeTaskFile) (bool, error) {
		if file.Kind == autoCodeTaskBackend {
			backendCalls++
			if backendCalls == 2 {
				return false, wantFailure
			}
		}
		return publishPreparedAutoCodeFile(file)
	}

	err := commitAutoCodeFileTask(task, publish, func() error {
		t.Fatal("database must not run after a backend publish failure")
		return nil
	})
	if !errors.Is(err, wantFailure) {
		t.Fatalf("commitAutoCodeFileTask() error = %v, want %v", err, wantFailure)
	}
	for _, name := range []string{"backend-a", "backend-b", "frontend-a"} {
		if got := readAutoCodeTaskTarget(t, targets[name]); got != name+"-before" {
			t.Fatalf("%s target = %q, want before", name, got)
		}
	}
}

func TestCommitAutoCodeFileTaskRollsBackPartialFrontendPhase(t *testing.T) {
	_, task, targets := prepareRequestFileTask(t, 2)
	stagingDir := task.stagingDir
	wantFailure := errors.New("second frontend failed")
	frontendCalls := 0
	publish := func(file *autoCodeTaskFile) (bool, error) {
		if file.Kind == autoCodeTaskFrontend {
			frontendCalls++
			if frontendCalls == 2 {
				return false, wantFailure
			}
		}
		return publishPreparedAutoCodeFile(file)
	}
	err := commitAutoCodeFileTask(task, publish, func() error { return nil })
	if !errors.Is(err, wantFailure) {
		t.Fatalf("commitAutoCodeFileTask() error = %v, want %v", err, wantFailure)
	}
	for _, name := range []string{"backend-a", "backend-b"} {
		if got := readAutoCodeTaskTarget(t, targets[name]); got != name+"-after" {
			t.Fatalf("%s target = %q, want after", name, got)
		}
	}
	for _, name := range []string{"frontend-a", "frontend-b"} {
		if got := readAutoCodeTaskTarget(t, targets[name]); got != name+"-before" {
			t.Fatalf("%s target = %q, want before after frontend rollback", name, got)
		}
	}
	if _, err = os.Stat(stagingDir); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("staging directory error = %v, want not exist", err)
	}
}

func TestCommitAutoCodeFileTaskRejectsExternalModification(t *testing.T) {
	_, task, targets := prepareRequestFileTask(t, 1)
	writeAutoCodeTaskTarget(t, targets["backend-a"], "external-change")

	err := commitAutoCodeFileTask(task, publishPreparedAutoCodeFile, func() error {
		t.Fatal("database must not run after a file conflict")
		return nil
	})
	if !errors.Is(err, errAutoCodeFileConflict) {
		t.Fatalf("commitAutoCodeFileTask() error = %v, want %v", err, errAutoCodeFileConflict)
	}
	if got := readAutoCodeTaskTarget(t, targets["backend-a"]); got != "external-change" {
		t.Fatalf("conflicting target = %q, want external-change", got)
	}
}

func TestCommitAutoCodeFileTaskDoesNotRollBackMatchingExternalModification(t *testing.T) {
	_, task, targets := prepareRequestFileTask(t, 1)
	writeAutoCodeTaskTarget(t, targets["backend-a"], "backend-a-after")
	wantFailure := errors.New("database failed")

	err := commitAutoCodeFileTask(task, publishPreparedAutoCodeFile, func() error { return wantFailure })
	if !errors.Is(err, wantFailure) {
		t.Fatalf("commitAutoCodeFileTask() error = %v, want %v", err, wantFailure)
	}
	if got := readAutoCodeTaskTarget(t, targets["backend-a"]); got != "backend-a-after" {
		t.Fatalf("matching external target = %q, want backend-a-after", got)
	}
	if got := readAutoCodeTaskTarget(t, targets["backend-b"]); got != "backend-b-before" {
		t.Fatalf("task-published target = %q, want backend-b-before after rollback", got)
	}
}

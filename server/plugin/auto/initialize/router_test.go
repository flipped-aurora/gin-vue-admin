package initialize

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	pluginUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"github.com/gin-gonic/gin"
)

func TestInitRoutesRequireAuthentication(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testutil.InitNopLogger()
	engine := gin.New()
	InitializeRouter(engine)

	for _, path := range []string{"/autoCode/initMenu", "/autoCode/initAPI", "/autoCode/initDictionary"} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, path, nil)
			resp := httptest.NewRecorder()
			engine.ServeHTTP(resp, req)
			if resp.Code != http.StatusUnauthorized {
				t.Fatalf("POST %s status = %d, want %d", path, resp.Code, http.StatusUnauthorized)
			}
		})
	}
}

func TestApiRegistersAutoCodeInitEndpoints(t *testing.T) {
	testutil.NewMemoryDB(t, &model.SysApi{})
	Api(context.Background())
	apis, _, _ := pluginUtils.GetPluginData("auto")

	want := map[string]bool{
		"/autoCode/initMenu":       false,
		"/autoCode/initAPI":        false,
		"/autoCode/initDictionary": false,
	}
	for _, api := range apis {
		if _, ok := want[api.Path]; ok && api.Method == http.MethodPost {
			want[api.Path] = true
		}
	}
	for path, found := range want {
		if !found {
			t.Errorf("Api() did not register POST %s", path)
		}
	}
}

func TestInitializeRouterRegistersAutoCompatibilityPrefixes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	InitializeRouter(engine)

	foundAutoCode := false
	foundSkills := false
	for _, route := range engine.Routes() {
		if route.Path == "/autoCode/getDB" {
			foundAutoCode = true
		}
		if route.Path == "/skills/getTools" {
			foundSkills = true
		}
	}

	if !foundAutoCode {
		t.Fatalf("expected /autoCode/getDB to be registered by plugin router")
	}
	if !foundSkills {
		t.Fatalf("expected /skills/getTools to be registered by plugin router")
	}
}

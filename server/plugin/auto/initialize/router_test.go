package initialize

import (
	"testing"

	"github.com/gin-gonic/gin"
)

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

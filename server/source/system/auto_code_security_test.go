package system

import (
	"context"
	"testing"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func TestAutoCodeInitPoliciesAreAdministratorOnly(t *testing.T) {
	db := testutil.NewMemoryDBWithoutGlobal(t, &model.SysApi{}, &adapter.CasbinRule{})
	ctx := context.WithValue(context.Background(), "db", db)
	if _, err := (&initApi{}).InitializeData(ctx); err != nil {
		t.Fatalf("initApi.InitializeData() error = %v", err)
	}
	if _, err := (&initCasbin{}).InitializeData(ctx); err != nil {
		t.Fatalf("initCasbin.InitializeData() error = %v", err)
	}

	paths := []string{"/autoCode/initMenu", "/autoCode/initAPI", "/autoCode/initDictionary"}
	for _, path := range paths {
		var apiCount int64
		if err := db.Model(&model.SysApi{}).Where("path = ? AND method = ?", path, "POST").Count(&apiCount).Error; err != nil {
			t.Fatal(err)
		}
		if apiCount != 1 {
			t.Errorf("API initialization count for %s = %d, want 1", path, apiCount)
		}

		for _, role := range []string{"888", "8881", "9528"} {
			var policyCount int64
			query := adapter.CasbinRule{Ptype: "p", V0: role, V1: path, V2: "POST"}
			if err := db.Where(&query).Model(&adapter.CasbinRule{}).Count(&policyCount).Error; err != nil {
				t.Fatal(err)
			}
			want := int64(0)
			if role == "888" {
				want = 1
			}
			if policyCount != want {
				t.Errorf("policy count for role %s and %s = %d, want %d", role, path, policyCount, want)
			}
		}
	}
}

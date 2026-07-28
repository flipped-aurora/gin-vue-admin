package system

import (
	"context"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func TestUserServiceRegisterSnapshotsForcePasswordChangePolicy(t *testing.T) {
	tests := []struct {
		name  string
		force bool
	}{
		{name: "enabled", force: true},
		{name: "disabled", force: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testutil.NewMemoryDB(t, &model.SysUser{})
			setSecurityConfigCache(model.SysSecurityConfig{ForceNewUserChangePassword: tt.force})
			t.Cleanup(func() { setSecurityConfigCache(model.DefaultSecurityConfig()) })

			created, err := UserServiceApp.Register(context.Background(), model.SysUser{
				Username: "new-user",
				Password: "initial-password",
			})
			if err != nil {
				t.Fatal(err)
			}
			if created.MustChangePassword != tt.force {
				t.Fatalf("created.MustChangePassword = %v, want %v", created.MustChangePassword, tt.force)
			}

			var stored model.SysUser
			if err = db.First(&stored, created.ID).Error; err != nil {
				t.Fatal(err)
			}
			if stored.MustChangePassword != tt.force {
				t.Fatalf("stored.MustChangePassword = %v, want %v", stored.MustChangePassword, tt.force)
			}
		})
	}
}

func TestUserServiceChangePasswordClearsMustChangePassword(t *testing.T) {
	db := testutil.NewMemoryDB(t, &model.SysUser{})
	user := model.SysUser{
		Username:           "pending-user",
		Password:           utils.BcryptHash("old-password"),
		MustChangePassword: true,
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	err := UserServiceApp.ChangePassword(
		context.Background(),
		&model.SysUser{GVA_MODEL: user.GVA_MODEL, Password: "old-password"},
		"new-password",
	)
	if err != nil {
		t.Fatal(err)
	}

	var stored model.SysUser
	if err = db.First(&stored, user.ID).Error; err != nil {
		t.Fatal(err)
	}
	if stored.MustChangePassword {
		t.Fatal("MustChangePassword remained true after self-service password change")
	}
}

func TestUserServiceResetPasswordPreservesMustChangePassword(t *testing.T) {
	db := testutil.NewMemoryDB(t, &model.SysUser{})
	user := model.SysUser{
		Username:           "reset-user",
		Password:           utils.BcryptHash("old-password"),
		MustChangePassword: true,
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	if err := UserServiceApp.ResetPassword(context.Background(), user.ID, "reset-password"); err != nil {
		t.Fatal(err)
	}

	var stored model.SysUser
	if err := db.First(&stored, user.ID).Error; err != nil {
		t.Fatal(err)
	}
	if !stored.MustChangePassword {
		t.Fatal("ResetPassword changed MustChangePassword, want it preserved")
	}
}

package douyin

import (
	"fmt"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/config"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

func TestDouyin(t *testing.T) {
	dy := New()
	cfg := &config.Config{
		ClientKey:    "ttd2ffa4424dc3677001",
		ClientSecret: "0e212aca063d0b98e58147852bcb4ee13b5f2047",
		Cache:        util.NewMemCache(),
	}
	openAPI := dy.GetOpenAPI(cfg)
	mini := openAPI.GetMini()
	a, e := mini.GetSession("BDBfZ-CoattVcDG0bnIE2il2y_VkDX5ZY0CqdgHMiOnpwJV6x18_cmf7r1Adxt3i5u5oz4SPQUeMm2tAo3NHzLp6kUeaX-R4fTOVyHomqhVnLvG04mP2qjVaB8s")
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(a)
}

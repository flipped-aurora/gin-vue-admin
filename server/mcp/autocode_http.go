package mcpTool

import (
	"context"

	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

func fetchAutoCodePackages(ctx context.Context) ([]model.SysAutoCodePackage, error) {
	resp, err := postUpstream[map[string][]model.SysAutoCodePackage](ctx, "/autoCode/getPackage", map[string]any{})
	if err != nil {
		return nil, err
	}
	return resp.Data["pkgs"], nil
}

func fetchAutoCodeHistories(ctx context.Context) ([]model.SysAutoCodeHistory, error) {
	resp, err := postUpstream[pageResultData[[]model.SysAutoCodeHistory]](ctx, "/autoCode/getSysHistory", commonReq.PageInfo{
		Page:     1,
		PageSize: 10000,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}

func createAutoCodePackage(ctx context.Context, info *systemReq.SysAutoCodePackageCreate) error {
	_, err := postUpstream[map[string]any](ctx, "/autoCode/createPackage", info)
	return err
}

func createAutoCodeModule(ctx context.Context, info systemReq.AutoCode) error {
	_, err := postUpstream[map[string]any](ctx, "/autoCode/createTemp", info)
	return err
}

func deleteAutoCodePackage(ctx context.Context, id uint) error {
	_, err := postUpstream[map[string]any](ctx, "/autoCode/delPackage", commonReq.GetById{ID: int(id)})
	return err
}

func deleteAutoCodeHistory(ctx context.Context, id uint) error {
	_, err := postUpstream[map[string]any](ctx, "/autoCode/delSysHistory", commonReq.GetById{ID: int(id)})
	return err
}

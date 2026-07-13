package mcpTool

import (
	"context"
	"net/url"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func fetchDictionaryList(ctx context.Context, keyword string) ([]system.SysDictionary, error) {
	query := url.Values{}
	if keyword != "" {
		query.Set("name", keyword)
	}

	resp, err := getUpstream[[]system.SysDictionary](ctx, "/sysDictionary/getSysDictionaryList", query)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// fetchDictionaryListWithDetails 一次性拉取字典及其字典项明细(每个 SysDictionary 已预加载 SysDictionaryDetails),
// 供 query_dictionaries 消除逐条 export 的 N+1。
func fetchDictionaryListWithDetails(ctx context.Context, keyword string) ([]system.SysDictionary, error) {
	query := url.Values{}
	if keyword != "" {
		query.Set("name", keyword)
	}

	resp, err := getUpstream[[]system.SysDictionary](ctx, "/sysDictionary/getSysDictionaryListWithDetails", query)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func findDictionaryByType(ctx context.Context, dictType string) (*system.SysDictionary, error) {
	dictionaries, err := fetchDictionaryList(ctx, dictType)
	if err != nil {
		return nil, err
	}

	for _, dictionary := range dictionaries {
		if dictionary.Type == dictType {
			dict := dictionary
			return &dict, nil
		}
	}

	return nil, nil
}

// createDictionary 回传创建后的字典实体(含自增 ID),免去调用方二次按 type 回查
func createDictionary(ctx context.Context, dictionary system.SysDictionary) (*system.SysDictionary, error) {
	resp, err := postUpstream[system.SysDictionary](ctx, "/sysDictionary/createSysDictionary", dictionary)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func createDictionaryDetail(ctx context.Context, detail system.SysDictionaryDetail) error {
	_, err := postUpstream[map[string]any](ctx, "/sysDictionaryDetail/createSysDictionaryDetail", detail)
	return err
}

func enabledBoolPointer() *bool {
	return utils.Pointer(true)
}

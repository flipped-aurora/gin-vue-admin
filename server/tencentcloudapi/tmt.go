package tencentcloudapi

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

func Translate(str string) (string, error) {
	if str == "" {
		return str, nil
	}
	credential := getClient()
	client, err := tmt.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		return "", err
	}
	request := tmt.NewTextTranslateRequest()
	SourceText := str
	Source := "zh"
	Target := "en"
	var ProjectId int64
	ProjectId = 0
	request.SourceText = &SourceText
	request.Source = &Source
	request.Target = &Target
	request.ProjectId = &ProjectId
	response, err := client.TextTranslate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		go utils.Notify(err)
		return str, err
	}
	if err != nil {
		fmt.Printf("An API error has returned: %s", err)
		go utils.Notify(err)
		return str, err
	}
	fmt.Printf("%s\n", response.ToJsonString())
	return *response.Response.TargetText, nil
}

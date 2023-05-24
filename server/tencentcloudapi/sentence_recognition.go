package tencentcloudapi

import (
	"encoding/base64"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func SentenceRecognition(engSerViceType, voiceFormat, fileName string, sourceType uint64) (string, error) {
	credential := getClient()
	client, err := asr.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		return "", err
	}
	request := asr.NewSentenceRecognitionRequest()
	ProjectId := new(uint64)
	SubServiceType := new(uint64)
	EngSerViceType := new(string)
	SourceType := new(uint64)
	VoiceFormat := new(string)
	FileName := new(string)
	*ProjectId = 0
	*SubServiceType = 2
	*EngSerViceType = engSerViceType
	*SourceType = sourceType
	*VoiceFormat = voiceFormat
	*FileName = fileName
	request.ProjectId = ProjectId
	request.SubServiceType = SubServiceType
	request.EngSerViceType = EngSerViceType
	request.SourceType = SourceType
	request.VoiceFormat = VoiceFormat
	request.UsrAudioKey = FileName
	var dataBase64 string
	if sourceType == 0 {
		request.Url = FileName
	} else {
		oss := upload.NewOss()
		bytes, err := oss.GetFile(fileName)
		if err != nil {
			fmt.Println("oss.GetFile(fileName) error")
			go utils.Notify(err)
			return "", err
		}
		dataBase64 = base64.StdEncoding.EncodeToString(bytes)
		fmt.Println("remote len(bytes)", len(dataBase64))
		dataLen := new(int64)
		*dataLen = int64(len(bytes))
		request.DataLen = dataLen
		data := new(string)
		*data = dataBase64
		request.Data = data
	}

	response, err := client.SentenceRecognition(request)
	if err != nil {
		fmt.Printf("An API error has returned: %s", err)
		go utils.Notify(err)
		return "", err
	}
	fmt.Printf("%s\n", response.ToJsonString())
	return *response.Response.Result, nil
}

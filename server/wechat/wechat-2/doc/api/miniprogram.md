# 小程序

## 基础接口

TODO

## 内容安全

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html)

|               名称                | 请求方式 | URL                              | 是否已实现 | 使用方法                               |
| :-------------------------------: | -------- | :------------------------------- | ---------- | -------------------------------------- |
| 异步校验图片/音频 <sub>v1.0</sub> | POST     | /wxa/media_check_async           | YES        | (security *Security) MediaCheckAsyncV1 |
| 同步校验一张图片 <sub>v1.0</sub>  | POST     | /wxa/img_sec_check               | YES        | (security *Security) ImageCheckV1      |
|         异步校验图片/音频         | POST     | /wxa/media_check_async?version=2 | YES        | (security *Security) MediaCheckAsync   |
| 同步检查一段文本 <sub>v1.0</sub>  | POST     | /wxa/msg_sec_check               | YES        | (security *Security) MsgCheckV1        |
|         同步检查一段文本          | POST     | /wxa/msg_sec_check?version=2     | YES        | (security *Security) MsgCheck          |


## OCR

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html)

|      名称      | 请求方式 | URL                    | 是否已实现 | 使用方法 |
| :------------: | -------- | :--------------------- | ---------- | -------- |
|   银行卡识别   | POST     | /cv/ocr/bankcard       |            |          |
|  营业执照识别  | POST     | /cv/ocr/bizlicense     |            |          |
|   驾驶证识别   | POST     | /cv/ocr/drivinglicense |            |          |
|   身份证识别   | POST     | /cv/ocr/idcard         |            |          |
| 通用印刷体识别 | POST     | /cv/ocr/comm           |            |          |
|   行驶证识别   | POST     | /cv/ocr/driving        |            |          |


## 手机号

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html)

|        名称        | 请求方式 | URL                              | 是否已实现 | 使用方法                            |
| :----------------: | -------- | :------------------------------- | ---------- | ----------------------------------- |
| code换取用户手机号 | POST     | /wxa/business/getuserphonenumber | YES        | (business *Business) GetPhoneNumber |


## 安全风控

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/safety-control-capability/riskControl.getUserRiskRank.html)

|        名称        | 请求方式 | URL                  | 是否已实现 | 使用方法                                   |
| :----------------: | -------- | :------------------- | ---------- | ------------------------------------------ |
| 获取用户的安全等级 | POST     | /wxa/getuserriskrank | YES        | (riskControl *RiskControl) GetUserRiskRank |


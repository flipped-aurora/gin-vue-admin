# 微信公众号 API 列表

## 基础接口

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html)

|           名称            | 请求方式 | URL                        | 是否已实现 | 使用方法 |
| :-----------------------: | -------- | :------------------------- | ---------- | -------- |
|     获取 Access token     | GET      | /cgi-bin/token             | YES        |          |
|  获取微信服务器 IP 地址   | GET      | /cgi-bin/get_api_domain_ip | YES        |          |
| 获取微信 callback IP 地址 | GET      | /cgi-bin/getcallbackip     | YES        |          |
|     清理接口调用次数      | POST     | /cgi-bin/clear_quota       | YES        |          |

## 订阅通知

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html)

| 名称                 | 请求方式 | URL                                    | 是否已实现 | 使用方法                                     |
| -------------------- | -------- | -------------------------------------- | ---------- | -------------------------------------------- |
| 选用模板             | POST     | /wxaapi/newtmpl/addtemplate            | YES        | (tpl \*Subscribe) Add                        |
| 删除模板             | POST     | /wxaapi/newtmpl/deltemplate            | YES        | (tpl \*Subscribe) Delete                     |
| 获取公众号类目       | GET      | /wxaapi/newtmpl/getcategory            | YES        | (tpl \*Subscribe) GetCategory                |
| 获取模板中的关键词   | GET      | /wxaapi/newtmpl/getpubtemplatekeywords | YES        | (tpl \*Subscribe) GetPubTplKeyWordsByID      |
| 获取类目下的公共模板 | GET      | /wxaapi/newtmpl/getpubtemplatetitles   | YES        | (tpl \*Subscribe) GetPublicTemplateTitleList |
| 获取私有模板列表     | GET      | /wxaapi/newtmpl/gettemplate            | YES        | (tpl \*Subscribe) List()                     |
| 发送订阅通知         | POST     | /cgi-bin/message/subscribe/bizsend     | YES        | (tpl \*Subscribe) Send                       |

## 客服消息

### PC 客服能力

#### 客服管理

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html)

| 名称             | 请求方式  | URL                                    | 是否已实现 | 使用方法                         |
| ---------------- | --------- | -------------------------------------- | ---------- | -------------------------------- |
| 获取客服基本信息 | GET       | /cgi-bin/customservice/getkflist       | YES        | (csm \*Manager) List             |
| 添加客服帐号     | POST      | /customservice/kfaccount/add           | YES        | (csm \*Manager) Add              |
| 邀请绑定客服帐号 | POST      | /customservice/kfaccount/inviteworker  | YES        | (csm \*Manager) InviteBind       |
| 设置客服信息     | POST      | /customservice/kfaccount/update        | YES        | (csm \*Manager) Update           |
| 上传客服头像     | POST/FORM | /customservice/kfaccount/uploadheadimg | YES        | (csm \*Manager) UploadHeadImg    |
| 删除客服帐号     | POST      | /customservice/kfaccount/del           | YES        | (csm \*Manager) Delete           |
| 获取在线客服     | POST      | /cgi-bin/customservice/getonlinekflist | YES        | (csm \*Manager) OnlineList       |
| 下发客服输入状态 | POST      | /cgi-bin/message/custom/typing         | YES        | (csm \*Manager) SendTypingStatus |

#### 会话控制

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html)

| 名称               | 请求方式 | URL                                     | 是否已实现 | 使用方法 |
| ------------------ | -------- | --------------------------------------- | ---------- | -------- |
| 创建会话           | POST     | /customservice/kfsession/create         | NO         |          |
| 获取客户会话状态   | GET      | /customservice/kfsession/getsession     | NO         |          |
| 获取客服会话列表   | GET      | /customservice/kfsession/getsessionlist | NO         |          |
| 获取未接入会话列表 | POST     | /customservice/kfsession/getwaitcase    | NO         |          |

#### 获取聊天记录

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html)

| 名称         | 请求方式 | URL                                 | 是否已实现 | 使用方法 |
| ------------ | -------- | ----------------------------------- | ---------- | -------- |
| 获取聊天记录 | POST     | /customservice/msgrecord/getmsglist | NO         |          |

### 对话能力

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide.html)

#### 顾问管理

| 名称                           | 请求方式 | URL                                    | 是否已实现 | 使用方法 |
| ------------------------------ | -------- | -------------------------------------- | ---------- | -------- |
| 添加顾问                       | POST     | /cgi-bin/guide/addguideacct            | NO         |          |
| 获取顾问信息                   | POST     | /cgi-bin/guide/getguideacct            | NO         |          |
| 修改顾问信息                   | POST     | /cgi-bin/guide/updateguideacct         | NO         |          |
| 删除顾问                       | POST     | /cgi-bin/guide/delguideacct            | NO         |          |
| 获取服务号顾问列表             | POST     | /cgi-bin/guide/getguideacctlist        | NO         |          |
| 生成顾问二维码                 | POST     | /cgi-bin/guide/guidecreateqrcode       | NO         |          |
| 获取顾问聊天记录               | POST     | /cgi-bin/guide/getguidebuyerchatrecord | NO         |          |
| 设置快捷回复与关注自动回复     | POST     | /cgi-bin/guide/setguideconfig          | NO         |          |
| 获取快捷回复与关注自动回复     | POST     | /cgi-bin/guide/getguideconfig          | NO         |          |
| 设置敏感词与离线自动回复       | POST     | /cgi-bin/guide/setguideacctconfig      | NO         |          |
| 获取离线自动回复与敏感词       | POST     | /cgi-bin/guide/getguideacctconfig      | NO         |          |
| 允许微信用户复制小程序页面路径 | POST     | /cgi-bin/guide/pushshowwxapathmenu     | NO         |          |
| 新建顾问分组                   | POST     | /cgi-bin/guide/newguidegroup           | NO         |          |
| 获取顾问分组列表               | POST     | /cgi-bin/guide/getguidegrouplist       | NO         |          |
| 获取顾问分组信息               | POST     | /cgi-bin/guide/getgroupinfo            | NO         |          |
| 分组内添加顾问                 | POST     | /cgi-bin/guide/addguide2guidegroup     | NO         |          |
| 分组内删除顾问                 | POST     | /cgi-bin/guide/delguide2guidegroup     | NO         |          |
| 获取顾问所在分组               | POST     | /cgi-bin/guide/getgroupbyguide         | NO         |          |
| 删除指定顾问分组               | POST     | /cgi-bin/guide/delguidegroup           | NO         |          |

#### 客户管理

| 名称                     | 请求方式 | URL                                         | 是否已实现 | 使用方法 |
| ------------------------ | -------- | ------------------------------------------- | ---------- | -------- |
| 为顾问分配客户           | POST     | /cgi-bin/guide/addguidebuyerrelation        | NO         |          |
| 为顾问移除客户           | POST     | /cgi-bin/guide/delguidebuyerrelation        | NO         |          |
| 获取顾问的客户列表       | POST     | /cgi-bin/guide/getguidebuyerrelationlist    | NO         |          |
| 为客户更换顾问           | POST     | /cgi-bin/guide/rebindguideacctforbuyer      | NO         |          |
| 修改客户昵称             | POST     | /cgi-bin/guide/updateguidebuyerrelation     | NO         |          |
| 查询客户所属顾问         | POST     | /cgi-bin/guide/getguidebuyerrelationbybuyer | NO         |          |
| 查询指定顾问和客户的关系 | POST     | /cgi-bin/guide/getguidebuyerrelation        | NO         |          |

#### 标签管理

| 名称               | 请求方式 | URL                                    | 是否已实现 | 使用方法 |
| ------------------ | -------- | -------------------------------------- | ---------- | -------- |
| 新建标签类型       | POST     | /cgi-bin/guide/newguidetagoption       | NO         |          |
| 删除标签类型       | POST     | /cgi-bin/guide/delguidetagoption       | NO         |          |
| 为标签添加可选值   | POST     | /cgi-bin/guide/addguidetagoption       | NO         |          |
| 获取标签和可选值   | POST     | /cgi-bin/guide/getguidetagoption       | NO         |          |
| 为客户设置标签     | POST     | /cgi-bin/guide/addguidebuyertag        | NO         |          |
| 查询客户标签       | POST     | /cgi-bin/guide/getguidebuyertag        | NO         |          |
| 根据标签值筛选客户 | POST     | /cgi-bin/guide/queryguidebuyerbytag    | NO         |          |
| 删除客户标签       | POST     | /cgi-bin/guide/delguidebuyertag        | NO         |          |
| 设置自定义客户信息 | POST     | /cgi-bin/guide/addguidebuyerdisplaytag | NO         |          |
| 获取自定义客户信息 | POST     | /cgi-bin/guide/getguidebuyerdisplaytag | NO         |          |

#### 素材管理

| 名称               | 请求方式 | URL                                  | 是否已实现 | 使用方法 |
| ------------------ | -------- | ------------------------------------ | ---------- | -------- |
| 添加小程序卡片素材 | POST     | /cgi-bin/guide/setguidecardmaterial  | NO         |          |
| 查询小程序卡片素材 | POST     | /cgi-bin/guide/getguidecardmaterial  | NO         |          |
| 删除小程序卡片素材 | POST     | /cgi-bin/guide/delguidecardmaterial  | NO         |          |
| 添加图片素材       | POST     | /cgi-bin/guide/setguideimagematerial | NO         |          |
| 查询图片素材       | POST     | /cgi-bin/guide/getguideimagematerial | NO         |          |
| 删除图片素材       | POST     | /cgi-bin/guide/delguideimagematerial | NO         |          |
| 添加文字素材       | POST     | /cgi-bin/guide/setguidewordmaterial  | NO         |          |
| 查询文字素材       | POST     | /cgi-bin/guide/getguidewordmaterial  | NO         |          |
| 删除文字素材       | POST     | /cgi-bin/guide/delguidewordmaterial  | NO         |          |

#### 群发任务管理

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.addGuideMassendJob.html)

| 名称                 | 请求方式 | URL                                   | 是否已实现 | 使用方法 |
| -------------------- | -------- | ------------------------------------- | ---------- | -------- |
| 添加群发任务         | POST     | /cgi-bin/guide/addguidemassendjob     | NO         |          |
| 获取群发任务列表     | POST     | /cgi-bin/guide/getguidemassendjoblist | NO         |          |
| 获取指定群发任务信息 | POST     | /cgi-bin/guide/getguidemassendjob     | NO         |          |
| 修改群发任务         | POST     | /cgi-bin/guide/updateguidemassendjob  | NO         |          |
| 取消群发任务         | POST     | /cgi-bin/guide/cancelguidemassendjob  | NO         |          |

## 微信网页开发

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)

| 名称                                                                    | 请求方式 | URL                                                 | 是否已实现 | 使用方法                             |
| ----------------------------------------------------------------------- | -------- | --------------------------------------------------- | ---------- | ------------------------------------ |
| 获取跳转的 url 地址                                                     | GET      | https://open.weixin.qq.com/connect/oauth2/authorize | YES        | (oauth \*Oauth) GetRedirectURL       |
| 获取网页应用跳转的 url 地址                                             | GET      | https://open.weixin.qq.com/connect/qrconnect        | YES        | (oauth \*Oauth) GetWebAppRedirectURL |
| 通过网页授权的 code 换取 access_token(区别于 context 中的 access_token) | GET      | /sns/oauth2/access_token                            | YES        | (oauth \*Oauth) GetUserAccessToken   |
| 刷新 access_token                                                       | GET      | /sns/oauth2/refresh_token?                          | YES        | (oauth \*Oauth) RefreshAccessToken   |
| 检验 access_token 是否有效                                              | GET      | /sns/auth                                           | YES        | (oauth \*Oauth) CheckAccessToken(    |
| 拉取用户信息(需 scope 为 snsapi_userinfo)                               | GET      | /sns/userinfo                                       | YES        | (oauth \*Oauth) GetUserInfo          |
| 获取 jssdk 需要的配置参数                                               | GET      | /cgi-bin/ticket/getticket                           | YES        | (js \*Js) GetConfig                  |

## 素材管理

## 草稿箱

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html)

| 名称                        | 请求方式 | URL                                                          | 是否已实现 | 使用方法                      |
| --------------------------- | -------- | ------------------------------------------------------------ | ---------- | ----------------------------- |
| 新建草稿                    | POST     | /cgi-bin/draft/add                                           | YES        | (draft \*Draft) AddDraft      |
| 获取草稿                    | POST     | /cgi-bin/draft/get                                           | YES        | (draft \*Draft) GetDraft      |
| 删除草稿                    | POST     | /cgi-bin/draft/delete                                        | YES        | (draft \*Draft) DeleteDraft   |
| 修改草稿                    | POST     | /cgi-bin/draft/update                                        | YES        | (draft \*Draft) UpdateDraft   |
| 获取草稿总数                | GET      | /cgi-bin/draft/count                                         | YES        | (draft \*Draft) CountDraft    |
| 获取草稿列表                | POST     | /cgi-bin/draft/batchget                                      | YES        | (draft \*Draft) PaginateDraft |
| MP 端开关（仅内测期间使用） | POST     | /cgi-bin/draft/switch<br />/cgi-bin/draft/switch?checkonly=1 | NO         |                               |

## 发布能力

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Publish/Publish.html)

说明：「发表记录」包括群发和发布。

注意：该接口，只能处理 "发布" 相关的信息，无法操作和获取 "群发" 相关内容！！[官方回复](https://developers.weixin.qq.com/community/develop/doc/0002a4fb2109d8f7a91d421c556c00)

- 群发：主动推送给粉丝，历史消息可看，被搜一搜收录，可以限定部分的粉丝接收到。
- 发布：不会主动推给粉丝，历史消息列表看不到，但是是公开给所有人的文章。也不会占用群发的次数。每天可以发布多篇内容。可以用于自动回复、自定义菜单、页面模板和话题中，发布成功时会生成一个永久链接。

| 名称                           | 请求方式 | URL                             | 是否已实现 | 使用方法                                 |
| ------------------------------ | -------- | ------------------------------- | ---------- | ---------------------------------------- |
| 发布接口                       | POST     | /cgi-bin/freepublish/submit     | YES        | (freePublish \*FreePublish) Publish      |
| 发布状态轮询接口               | POST     | /cgi-bin/freepublish/get        | YES        | (freePublish \*FreePublish) SelectStatus |
| 事件推送发布结果               |          |                                 | YES        | EventPublishJobFinish                    |
| 删除发布                       | POST     | /cgi-bin/freepublish/delete     | YES        | (freePublish \*FreePublish) Delete       |
| 通过 article_id 获取已发布文章 | POST     | /cgi-bin/freepublish/getarticle | YES        | (freePublish \*FreePublish) First        |
| 获取成功发布列表               | POST     | /cgi-bin/freepublish/batchget   | YES        | (freePublish \*FreePublish) Paginate     |

## 图文消息留言管理

## 用户管理

| 名称                                       | 请求方式 | URL                                    | 是否已实现 | 使用方法                           |
| ------------------------------------------ | -------- | -------------------------------------- | ---------- | ---------------------------------- |
| 获取指定 OpenID 变化列表（公众号账号迁移） | POST     | /cgi-bin/changeopenid                  | YES        | (user \*User) ListChangeOpenIDs    |
| 获取所有用户 OpenID 列表（公众号账号迁移） |          |                                        | YES        | (user \*User) ListAllChangeOpenIDs |
| 获取用户基本信息                           | GET      | /cgi-bin/user/info                     | YES        | (user \*User) GetUserInfo          |
| 设置用户备注名                             | POST     | /cgi-bin/user/info/updateremark        | YES        | (user \*User) UpdateRemark         |
| 获取用户列表                               | GET      | /cgi-bin/user/get                      | YES        | (user \*User) ListUserOpenIDs      |
| 获取所有用户 OpenID 列表                   |          |                                        | YES        | (user \*User) ListAllUserOpenIDs   |
| 获取公众号的黑名单列表                     | POST     | /cgi-bin/tags/members/getblacklist     | YES        | (user \*User) GetBlackList         |
| 获取公众号的所有黑名单列表                 |          |                                        | YES        | (user \*User) GetAllBlackList      |
| 拉黑用户                                   | POST     | /cgi-bin/tags/members/batchblacklist   | YES        | (user \*User) BatchBlackList       |
| 取消拉黑用户                               | POST     | /cgi-bin/tags/members/batchunblacklist | YES        | (user \*User) BatchUnBlackList     |
| 创建标签                                   | POST     | /cgi-bin/tags/create                   | YES        | (user \*User) CreateTag            |
| 删除标签                                   | POST     | /cgi-bin/tags/delete                   | YES        | (user \*User) DeleteTag            |
| 编辑标签                                   | POST     | /cgi-bin/tags/update                   | YES        | (user \*User) UpdateTag            |
| 获取公众号已创建的标签                     | GET      | /cgi-bin/tags/get                      | YES        | (user \*User) GetTag               |
| 获取标签下粉丝列表                         | POST     | /cgi-bin/user/tag/get                  | YES        | (user \*User) OpenIDListByTag      |
| 批量为用户打标签                           | POST     | /cgi-bin/tags/members/batchtagging     | YES        | (user \*User) BatchTag             |
| 批量为用户取消标签                         | POST     | /cgi-bin/tags/members/batchuntagging   | YES        | (user \*User) BatchUntag           |
| 获取用户身上的标签列表                     | POST     | /cgi-bin/tags/getidlist                | YES        | (user \*User) UserTidList          |

## 账号管理

## 数据统计

## 微信卡券

## 微信门店

## 智能接口

## 微信设备功能

## 微信“一物一码”

## 微信发票

## 微信非税缴费

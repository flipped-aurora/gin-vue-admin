# 企业微信

host: https://qyapi.weixin.qq.com/

## 微信客服

[官方文档](https://work.weixin.qq.com/api/doc/90000/90135/94638)

### 客服账号管理

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94684)

|       名称       | 请求方式 | URL                         | 是否已实现 | 使用方法                             |贡献者       |
| :--------------: | -------- | :-------------------------- | ---------- | -------------------------------|------------|
|   添加客服帐号   | POST     | /cgi-bin/kf/account/add     | YES        | (r *Client) AccountAdd           | NICEXAI    |
|   删除客服帐号   | POST     | /cgi-bin/kf/account/del     | YES        | (r *Client) AccountDel           | NICEXAI    |
|   修改客服帐号   | POST     | /cgi-bin/kf/account/update  | YES        | (r *Client) AccountUpdate        | NICEXAI    |
| 获取客服帐号列表 | GET      | /cgi-bin/kf/account/list     | YES        | (r *Client) AccountList          | NICEXAI    |
| 获取客服帐号链接 | GET      | /cgi-bin/kf/add_contact_way  | YES        | (r *Client) AddContactWay        | NICEXAI    |

### 接待人员列表

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94693)

|       名称       | 请求方式 | URL                         | 是否已实现 | 使用方法                             |贡献者       |
| :--------------: | -------- | :-------------------------- | ---------- | -------------------------------|------------|
|   添加接待人员     | POST     | /cgi-bin/kf/servicer/add    | YES        | (r *Client) ReceptionistAdd     | NICEXAI    |
|   删除接待人员     | POST     | /cgi-bin/kf/servicer/del    | YES        | (r *Client) ReceptionistDel     | NICEXAI    |
| 获取接待人员列表    | GET      | /cgi-bin/kf/servicer/list   | YES        | (r *Client) ReceptionistList    | NICEXAI    |

### 会话分配与消息收发

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94694)

|       名称       | 请求方式 | URL                               | 是否已实现   | 使用方法                          |贡献者       |
| :--------------: | -------- | :-------------------------------| ---------- | ------------------------------- |------------|
|   获取会话状态     | POST     | /cgi-bin/kf/service_state/get   | YES        | (r *Client) ServiceStateGet     | NICEXAI    |
|   变更会话状态     | POST     | /cgi-bin/kf/service_state/trans | YES        | (r *Client) ServiceStateTrans   | NICEXAI    |
|   读取消息        | POST     | /cgi-bin/kf/sync_msg            | YES        | (r *Client) SyncMsg             | NICEXAI    |
|   发送消息        | POST     | /cgi-bin/kf/send_msg            | YES        | (r *Client) SendMsg             | NICEXAI    |
|   发送事件响应消息 | POST     | /cgi-bin/kf/send_msg_on_event   | YES        | (r *Client) SendMsgOnEvent      | NICEXAI    |

### 升级服务配置

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94702)

|       名称                | 请求方式  | URL                                               | 是否已实现 | 使用方法                            |贡献者       |
| :--------------:         | -------- | :-------------------------------------------------| ---------- | -------------------------------  |------------|
| 获取配置的专员与客户群       | POST     | /cgi-bin/kf/customer/get_upgrade_service_config   | YES        | (r *Client) UpgradeServiceConfig | NICEXAI    |
| 为客户升级为专员或客户群服务  | POST     | /cgi-bin/kf/customer/upgrade_service              | YES        | (r *Client) UpgradeService       | NICEXAI    |
| 为客户取消推荐             | POST     | /cgi-bin/kf/customer/cancel_upgrade_service       | YES        | (r *Client) UpgradeServiceCancel  | NICEXAI    |

### 其他基础信息获取

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/95148)

|       名称            | 请求方式  | URL                                     | 是否已实现   | 使用方法                            | 贡献者       |
| :--------------:     | -------- | :---------------------------------------| ---------- | -------------------------------   |------------|
| 获取客户基础信息        | POST     | /cgi-bin/kf/customer/batchget           | YES        | (r *Client) CustomerBatchGet      | NICEXAI    |
| 获取视频号绑定状态      | GET      |  /cgi-bin/kf/get_corp_qualification      | YES        | (r *Client) GetCorpQualification  | NICEXAI    |

### 客户联系
[官方文档](https://developer.work.weixin.qq.com/document/path/92132/92133/92228)

|            名称            | 请求方式  | URL                                                          | 是否已实现   | 使用方法                            | 贡献者      |
|:------------------------:| -------- |:-------------------------------------------------------------| ---------- | -------------------------------   |----------|
|       获取「联系客户统计」数据       | POST     | /cgi-bin/externalcontact/get_user_behavior_data              | YES        | (r *Client) GetUserBehaviorData      | MARKWANG |
| 获取「群聊数据统计」数据 (按群主聚合的方式)  | POST      | /cgi-bin/externalcontact/groupchat/statistic                 | YES        | (r *Client) GetGroupChatStat  | MARKWANG  |
| 获取「群聊数据统计」数据 (按自然日聚合的方式) | POST      | /cgi-bin/externalcontact/groupchat/statistic_group_by_day    | YES        | (r *Client) GetGroupChatStatByDay  | MARKWANG  |
|      配置客户联系「联系我」方式       | POST      | /cgi-bin/externalcontact/add_contact_way                     | YES        | (r *Client) AddContactWay  | MARKWANG  |
|     获取企业已配置的「联系我」方式      | POST      | /cgi-bin/externalcontact/get_contact_way                     | YES        | (r *Client) GetContactWay  | MARKWANG  |
|     更新企业已配置的「联系我」方式      | POST      | /cgi-bin/externalcontact/update_contact_way                  | YES        | (r *Client) UpdateContactWay  | MARKWANG  |
|     获取企业已配置的「联系我」列表      | POST      | /cgi-bin/externalcontact/list_contact_way                    | YES        | (r *Client) ListContactWay  | MARKWANG  |
|     删除企业已配置的「联系我」方式      | POST      | /cgi-bin/externalcontact/del_contact_way                     | YES        | (r *Client) DelContactWay  | MARKWANG  |
|          创建企业群发          | POST      | /cgi-bin/externalcontact/add_msg_template                    | YES        | (r *Client) AddMsgTemplate  | MARKWANG  |
|          获取群发记录列表          | POST      | /cgi-bin/externalcontact/get_groupmsg_list_v2                | YES        | (r *Client) GetGroupMsgListV2  | MARKWANG  |
|          获取群发成员发送任务列表          | POST      | /cgi-bin/externalcontact/get_groupmsg_task                   | YES        | (r *Client) GetGroupMsgTask  | MARKWANG  |
|          获取企业群发成员执行结果          | POST      | /cgi-bin/externalcontact/get_groupmsg_send_result            | YES        | (r *Client) GetGroupMsgSendResult  | MARKWANG  |
|          发送新客户欢迎语          | POST      | /cgi-bin/externalcontact/send_welcome_msg                    | YES        | (r *Client) SendWelcomeMsg  | MARKWANG  |
|          添加入群欢迎语素材          | POST      | /cgi-bin/externalcontact/group_welcome_template/add          | YES        | (r *Client) AddGroupWelcomeTemplate  | MARKWANG  |
|          编辑入群欢迎语素材          | POST      | /cgi-bin/externalcontact/group_welcome_template/edit         | YES        | (r *Client) EditGroupWelcomeTemplate  | MARKWANG  |
|          获取入群欢迎语素材          | POST      | /cgi-bin/externalcontact/group_welcome_template/get          | YES        | (r *Client) GetGroupWelcomeTemplate  | MARKWANG  |
|          删除入群欢迎语素材          | POST      | /cgi-bin/externalcontact/group_welcome_template/del          | YES        | (r *Client) DelGroupWelcomeTemplate  | MARKWANG  |

## 通讯录管理
[官方文档](https://developer.work.weixin.qq.com/document/path/90193)

### 部门管理

|    名称     | 请求方式 | URL                                     | 是否已实现   | 使用方法                            | 贡献者      |
|:---------:|------|:----------------------------------------| ---------- | -------------------------------   |----------|
| 获取子部门ID列表 | GET  | /cgi-bin/department/simplelist          | YES        | (r *Client) DepartmentSimpleList| MARKWANG |
|  获取部门成员   | GET | /cgi-bin/user/simplelist                | YES        | (r *Client) UserSimpleList  | MARKWANG  |
|  获取成员ID列表   | Post | /cgi-bin/user/list_id                | YES        | (r *Client) UserListId  | MARKWANG  |


## 素材管理
[官方文档](https://developer.work.weixin.qq.com/document/path/91054)

|    名称     | 请求方式 | URL                                     | 是否已实现   | 使用方法                            | 贡献者      |
|:---------:|------|:----------------------------------------| ---------- | -------------------------------   |----------|
| 上传图片 | POST  | /cgi-bin/media/uploadimg          | YES        | (r *Client) UploadImg| MARKWANG |

### 成员管理

| 名称     | 请求方式 | URL               | 是否已实现 | 使用方法            | 贡献者   |
| -------- | -------- | ----------------- | ---------- | ------------------- | -------- |
| 读取成员 | GET      | /cgi-bin/user/get | YES        | (r *Client) UserGet | chcthink |


## 群机器人

[官方文档](https://developer.work.weixin.qq.com/document/path/91770)

| 名称             | 请求方式 | URL                   | 是否已实现 | 使用方法                   | 贡献者   |
| ---------------- | -------- | --------------------- | ---------- | -------------------------- | -------- |
| 群机器人发送消息 | POST     | /cgi-bin/webhook/send | YES        | (r *Client) RobotBroadcast | chcthink |

## 应用管理
TODO

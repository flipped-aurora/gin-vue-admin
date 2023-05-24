// All Rights Reserved.
// *File ： WeWorkFinanceSdk_C.h
// @Brief：拉取企业聊天记录与媒体消息sdk头文件

#pragma once
//返回码	错误说明
//10000	参数错误，请求参数错误
//10001	网络错误，网络请求错误
//10002	数据解析失败
//10003	系统失败
//10004	密钥错误导致加密失败
//10005	fileid错误
//10006	解密失败
//10007 找不到消息加密版本的私钥，需要重新传入私钥对
//10008 解析encrypt_key出错
//10009 ip非法
//10010 数据过期
//10011	证书错误

typedef struct WeWorkFinanceSdk_t WeWorkFinanceSdk_t;

// 数据
typedef struct Slice_t {
    char* buf;
    int len;
} Slice_t;

typedef struct MediaData {
    char* outindexbuf;
    int out_len;
    char* data;    
    int data_len;
    int is_finish;
} MediaData_t;


#ifdef __cplusplus
extern "C" {
#endif    

    WeWorkFinanceSdk_t* NewSdk();
                               

	/**
	 * 初始化函数
	 * Return值=0表示该API调用成功
	 * 
	 * @param [in]  sdk			NewSdk返回的sdk指针
	 * @param [in]  corpid      调用企业的企业id，例如：wwd08c8exxxx5ab44d，可以在企业微信管理端--我的企业--企业信息查看
	 * @param [in]  secret		聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
	 *						
	 *
	 * @return 返回是否初始化成功
	 *      0   - 成功
	 *      !=0 - 失败
	 */
    int Init(WeWorkFinanceSdk_t* sdk, const char* corpid, const char* secret);

	/**
	 * 拉取聊天记录函数
	 * Return值=0表示该API调用成功
	 * 
	 *
	 * @param [in]  sdk				NewSdk返回的sdk指针
	 * @param [in]  seq				从指定的seq开始拉取消息，注意的是返回的消息从seq+1开始返回，seq为之前接口返回的最大seq值。首次使用请使用seq:0
	 * @param [in]  limit			一次拉取的消息条数，最大值1000条，超过1000条会返回错误
	 * @param [in]  proxy			使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
	 * @param [in]  passwd			代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
	 * @param [in]  timeout			超时时间，单位秒
	 * @param [out] chatDatas		返回本次拉取消息的数据，slice结构体.内容包括errcode/errmsg，以及每条消息内容。示例如下：

	 {"errcode":0,"errmsg":"ok","chatdata":[{"seq":196,"msgid":"CAQQ2fbb4QUY0On2rYSAgAMgip/yzgs=","publickey_ver":3,"encrypt_random_key":"ftJ+uz3n/z1DsxlkwxNgE+mL38H42/KCvN8T60gbbtPD+Rta1hKTuQPzUzO6Hzne97MgKs7FfdDxDck/v8cDT6gUVjA2tZ/M7euSD0L66opJ/IUeBtpAtvgVSD5qhlaQjvfKJc/zPMGNK2xCLFYqwmQBZXbNT7uA69Fflm512nZKW/piK2RKdYJhRyvQnA1ISxK097sp9WlEgDg250fM5tgwMjujdzr7ehK6gtVBUFldNSJS7ndtIf6aSBfaLktZgwHZ57ONewWq8GJe7WwQf1hwcDbCh7YMG8nsweEwhDfUz+u8rz9an+0lgrYMZFRHnmzjgmLwrR7B/32Qxqd79A==","encrypt_chat_msg":"898WSfGMnIeytTsea7Rc0WsOocs0bIAerF6de0v2cFwqo9uOxrW9wYe5rCjCHHH5bDrNvLxBE/xOoFfcwOTYX0HQxTJaH0ES9OHDZ61p8gcbfGdJKnq2UU4tAEgGb8H+Q9n8syRXIjaI3KuVCqGIi4QGHFmxWenPFfjF/vRuPd0EpzUNwmqfUxLBWLpGhv+dLnqiEOBW41Zdc0OO0St6E+JeIeHlRZAR+E13Isv9eS09xNbF0qQXWIyNUi+ucLr5VuZnPGXBrSfvwX8f0QebTwpy1tT2zvQiMM2MBugKH6NuMzzuvEsXeD+6+3VRqL"}]}

	 *
	 * @return 返回是否调用成功
	 *      0   - 成功
	 *      !=0 - 失败	
	 */		
    int GetChatData(WeWorkFinanceSdk_t* sdk, unsigned long long seq, unsigned int limit, const char *proxy,const char* passwd, int timeout,Slice_t* chatDatas);

	/**
     * @brief 解析密文.企业微信自有解密内容
     * @param [in]  encrypt_key, getchatdata返回的encrypt_random_key,使用企业自持对应版本秘钥RSA解密后的内容
     * @param [in]  encrypt_msg, getchatdata返回的encrypt_chat_msg
     * @param [out] msg, 解密的消息明文
	 * @return 返回是否调用成功
	 *      0   - 成功
	 *      !=0 - 失败
     */
    int DecryptData(const char* encrypt_key, const char* encrypt_msg, Slice_t* msg);

	/**
	 * 拉取媒体消息函数
	 * Return值=0表示该API调用成功
	 * 
	 *
	 * @param [in]  sdk				NewSdk返回的sdk指针
	 * @param [in]  sdkFileid		从GetChatData返回的聊天消息中，媒体消息包括的sdkfileid
	 * @param [in]  proxy			使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
	 * @param [in]  passwd			代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
	 * @param [in]  indexbuf		媒体消息分片拉取，需要填入每次拉取的索引信息。首次不需要填写，默认拉取512k，后续每次调用只需要将上次调用返回的outindexbuf填入即可。
	 * @param [in]  timeout			超时时间，单位秒
	 * @param [out] media_data		返回本次拉取的媒体数据.MediaData结构体.内容包括data(数据内容)/outindexbuf(下次索引)/is_finish(拉取完成标记)
	 
	 *
	 * @return 返回是否调用成功
	 *      0   - 成功
	 *      !=0 - 失败
	 */
	int GetMediaData(WeWorkFinanceSdk_t* sdk, const char* indexbuf,
                     const char* sdkFileid,const char *proxy,const char* passwd, int timeout, MediaData_t* media_data);

    /**
     * @brief 释放sdk，和NewSdk成对使用
     * @return 
     */
    void DestroySdk(WeWorkFinanceSdk_t* sdk);


    //--------------下面接口为了其他语言例如python等调用c接口，酌情使用--------------
    Slice_t* NewSlice();

    /**
     * @brief 释放slice，和NewSlice成对使用
     * @return 
     */
    void FreeSlice(Slice_t* slice);

    /**
     * @brief 为其他语言提供读取接口
     * @return 返回buf指针
     *     !=NULL - 成功
     *     NULL   - 失败
     */
    char* GetContentFromSlice(Slice_t* slice);
	int GetSliceLen(Slice_t* slice);

	// 媒体记录相关工具

    MediaData_t*  NewMediaData();

    void FreeMediaData(MediaData_t* media_data);

    char* GetOutIndexBuf(MediaData_t* media_data);
    char* GetData(MediaData_t* media_data);
	int GetIndexLen(MediaData_t* media_data);
	int GetDataLen(MediaData_t* media_data);
    int IsMediaDataFinish(MediaData_t* media_data);

#ifdef __cplusplus
}
#endif    

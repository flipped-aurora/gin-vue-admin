#include "WeWorkFinanceSdk_C.h"
#include <dlfcn.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string>
using std::string;

typedef WeWorkFinanceSdk_t* newsdk_t();
typedef int Init_t(WeWorkFinanceSdk_t*, const char*, const char*);
typedef void DestroySdk_t(WeWorkFinanceSdk_t*);

typedef int GetChatData_t(WeWorkFinanceSdk_t*, unsigned long long, unsigned int, const char*, const char*, int, Slice_t*);
typedef Slice_t* NewSlice_t();
typedef void FreeSlice_t(Slice_t*);

typedef int GetMediaData_t(WeWorkFinanceSdk_t*, const char*, const char*, const char*, const char*, int, MediaData_t*);
typedef int DecryptData_t(const char*, const char*, Slice_t*);
typedef MediaData_t* NewMediaData_t();
typedef void FreeMediaData_t(MediaData_t*);

int main(int argc, char* argv[])
{
    int ret = 0;
	//seq 表示该企业存档消息序号，该序号单调递增，拉取序号建议设置为上次拉取返回结果中最大序号。首次拉取时seq传0，sdk会返回有效期内最早的消息。
	//limit 表示本次拉取的最大消息条数，取值范围为1~1000
	//proxy与passwd为代理参数，如果运行sdk的环境不能直接访问外网，需要配置代理参数。sdk访问的域名是"https://qyapi.weixin.qq.com"。
	//建议先通过curl访问"https://qyapi.weixin.qq.com"，验证代理配置正确后，再传入sdk。
	//timeout 为拉取会话存档的超时时间，单位为秒，建议超时时间设置为5s。
	//sdkfileid 媒体文件id，从解密后的会话存档中得到
	//savefile 媒体文件保存路径
	//encrypt_key 拉取会话存档返回的encrypt_random_key，使用配置在企业微信管理台的rsa公钥对应的私钥解密后得到encrypt_key。
	//encrypt_chat_msg 拉取会话存档返回的encrypt_chat_msg
    if (argc < 2) {
        printf("./sdktools 1(chatmsg) 2(mediadata) 3(decryptdata)\n");
        printf("./sdktools 1 seq limit proxy passwd timeout\n");
        printf("./sdktools 2 fileid proxy passwd timeout savefile\n");
        printf("./sdktools 3 encrypt_key encrypt_chat_msg\n");
        return -1;
    }

    void* so_handle = dlopen("./libWeWorkFinanceSdk_C.so", RTLD_LAZY);
    if (!so_handle) {
        printf("load sdk so fail:%s\n", dlerror());
        return -1;
    }
    newsdk_t* newsdk_fn = (newsdk_t*)dlsym(so_handle, "NewSdk");
    WeWorkFinanceSdk_t* sdk = newsdk_fn();

	//使用sdk前需要初始化，初始化成功后的sdk可以一直使用。
	//如需并发调用sdk，建议每个线程持有一个sdk实例。
	//初始化时请填入自己企业的corpid与secrectkey。
    Init_t* init_fn = (Init_t*)dlsym(so_handle, "Init");
    DestroySdk_t* destroysdk_fn = (DestroySdk_t*)dlsym(so_handle, "DestroySdk");
    ret = init_fn(sdk, "wwd08c8e7c775ab44d", "zJ6k0naVVQ--gt9PUSSEvs03zW_nlDVmjLCTOTAfrew");
    if (ret != 0) {
        //sdk需要主动释放
        destroysdk_fn(sdk);
        printf("init sdk err ret:%d\n", ret);
        return -1;
    }

    int type = strtoul(argv[1], NULL, 10);
    if (type == 1) {
        //拉取会话存档
        uint64_t iSeq = strtoul(argv[2], NULL, 10);
        uint64_t iLimit = strtoul(argv[3], NULL, 10);
        uint64_t timeout = strtoul(argv[6], NULL, 10);
        
        NewSlice_t* newslice_fn = (NewSlice_t*)dlsym(so_handle, "NewSlice");
        FreeSlice_t* freeslice_fn = (FreeSlice_t*)dlsym(so_handle, "FreeSlice");

		//每次使用GetChatData拉取存档前需要调用NewSlice获取一个chatDatas，在使用完chatDatas中数据后，还需要调用FreeSlice释放。
        Slice_t* chatDatas = newslice_fn();
        GetChatData_t* getchatdata_fn = (GetChatData_t*)dlsym(so_handle, "GetChatData");
        ret = getchatdata_fn(sdk, iSeq, iLimit, argv[4], argv[5], timeout, chatDatas);
        if (ret != 0) {
            freeslice_fn(chatDatas);
            printf("GetChatData err ret:%d\n", ret);
            return -1;
        }
        printf("GetChatData len:%d data:%s\n", chatDatas->len, chatDatas->buf);
        freeslice_fn(chatDatas);
    } 
    else if (type == 2) {
		//拉取媒体文件
        std::string index;
        uint64_t timeout = strtoul(argv[5], NULL, 10);
        int isfinish = 0;

        GetMediaData_t* getmediadata_fn = (GetMediaData_t*)dlsym(so_handle, "GetMediaData");
        NewMediaData_t* newmediadata_fn = (NewMediaData_t*)dlsym(so_handle, "NewMediaData");
        FreeMediaData_t* freemediadata_fn = (FreeMediaData_t*)dlsym(so_handle, "FreeMediaData");

		//媒体文件每次拉取的最大size为512k，因此超过512k的文件需要分片拉取。若该文件未拉取完整，mediaData中的is_finish会返回0，同时mediaData中的outindexbuf会返回下次拉取需要传入GetMediaData的indexbuf。
		//indexbuf一般格式如右侧所示，”Range:bytes=524288-1048575“，表示这次拉取的是从524288到1048575的分片。单个文件首次拉取填写的indexbuf为空字符串，拉取后续分片时直接填入上次返回的indexbuf即可。
        while (isfinish == 0) {
            //每次使用GetMediaData拉取存档前需要调用NewMediaData获取一个mediaData，在使用完mediaData中数据后，还需要调用FreeMediaData释放。
            printf("index:%s\n", index.c_str());
            MediaData_t* mediaData = newmediadata_fn();
            ret = getmediadata_fn(sdk, index.c_str(), argv[2], argv[3], argv[4], timeout, mediaData);
            if (ret != 0) {
                //单个分片拉取失败建议重试拉取该分片，避免从头开始拉取。
                freemediadata_fn(mediaData);
                printf("GetMediaData err ret:%d\n", ret);
                return -1;
            }
            printf("content size:%d isfin:%d outindex:%s\n", mediaData->data_len, mediaData->is_finish, mediaData->outindexbuf);

			//大于512k的文件会分片拉取，此处需要使用追加写，避免后面的分片覆盖之前的数据。
            char file[200];
            snprintf(file, sizeof(file), "%s", argv[6]);
            FILE* fp = fopen(file, "ab+");
            printf("filename:%s \n", file);
            if (NULL == fp) {
                freemediadata_fn(mediaData);
                printf("open file err\n");
                return -1;
            }

            fwrite(mediaData->data, mediaData->data_len, 1, fp);
            fclose(fp);

            //获取下次拉取需要使用的indexbuf
            index.assign(string(mediaData->outindexbuf));
            isfinish = mediaData->is_finish;
            freemediadata_fn(mediaData);
        }
    } 
    else if (type == 3) {
		//解密会话存档内容
		//sdk不会要求用户传入rsa私钥，保证用户会话存档数据只有自己能够解密。
		//此处需要用户先用rsa私钥解密encrypt_random_key后，作为encrypt_key参数传入sdk来解密encrypt_chat_msg获取会话存档明文。
		//每次使用DecryptData解密会话存档前需要调用NewSlice获取一个Msgs，在使用完Msgs中数据后，还需要调用FreeSlice释放。
        NewSlice_t* newslice_fn = (NewSlice_t*)dlsym(so_handle, "NewSlice");
        FreeSlice_t* freeslice_fn = (FreeSlice_t*)dlsym(so_handle, "FreeSlice");

        Slice_t* Msgs = newslice_fn();
        // decryptdata api
        DecryptData_t* decryptdata_fn = (DecryptData_t*)dlsym(so_handle, "DecryptData");
        ret = decryptdata_fn(argv[2], argv[3], Msgs);
        printf("chatdata :%s ret :%d\n", Msgs->buf, ret);

        freeslice_fn(Msgs);
    }

    return ret;
}

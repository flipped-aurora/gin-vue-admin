import { reactive, ref, watchEffect } from "vue"
/**
 * 文件上传HOOKS
 * 只限h5,微信，qq使用。其它平台请自己修改实现。
 * @description 鉴于大家可能需要只是上传功能，可能界面需要自己定制，那么这个hooks函数将提升你的效率。你仅需要专注于界面设计，其它的交给本函数。
 * @author tmui|tmzdy|https://tmui.design
 * @version 1.0.0
 * 2023-10-7 22:51:00 by tmui
 * @copyright 本代码只允许跟随tmui组件库发布。如果需要其它地方合并使用，请保留此注释版权信息。
 */

//文件对象结构
export interface FILE_TYPE {
    url: string,//当前显示的图片地址，这是个本地临时地址。待上传前，成功后会替换服务器地址。
    status?: string,//上传状态文本
    progress?: number,//当前文件上传的进度
    uid?: string | number,//文件唯一标识id
    statusCode?: STATUS_CODE,//文件状态
    response?: any,//上传成功后的回调数据。
    name?: string,//文件名称
    size?: number,//文件大小字节单位
    FILE?: any,//文件对象。
    [key: string]: any
}

//文件上传的状态值
export enum STATUS_CODE {
    //待上传
    upload = 0,
    //上传中
    uploading = 1,
    //上传失败
    fail = 2,
    //上传成功
    success = 3,
    //超过大小限制
    max = 4,
}
export interface USE_UPLOAD_FILE_CONFIG_TYPE {
    /**media表示只允许图片或者视频选择，file表示允许任意文件，但只支持h5,微信平台  */
    uploadType:"media"|"file",
    maxCount: number,//一次选择文件最大数量。
    extension: string[],//文件选择的类型。
    type: "all" | "image" | "video" | "file" | undefined,
    /**只对h5 */
    sourceType: Array<'album' | 'camera'>,
    maxSize: number,//每一个文件上传的最大尺寸，默认为10mb
    hostUrl: string,//上传文件的服务器地址
    autoUpload: boolean,
    header: { [key: string]: any },//头部参数。
    formData: { [key: string]: any },//额外的表单数据。
    formName: string,
    code: number,//服务器响应码，如果不为此码，表示上传失败。
    maxDuration:number,//如果选择的类型是视频，可以定制此拍摄的最大时长。仅uploadType为media时有效
    sizeType:Array<'original'|'compressed'>,//仅对 mediaType 为 image 时有效，是否压缩所选文件,仅uploadType为media时有效
    camera:'back'|'front',//仅在 sourceType 为 camera 时生效，使用前置或后置摄像头,仅uploadType为media时有效
    mediaType:Array<'image'|'video'>,//注意当uploadType=media时，如果是微信，抖音，飞书这里可以为当前正常类型值；如果是其它平台：只取数组中第一个值，比如要选择图片设置为["image"],视频：["video"],仅uploadType为media时有效

}
export function getUid(length = 3) {
    return Number(Number(Math.random().toString().substr(3, length) + Date.now()).toString(8));
}
export const useUploadFile = (cfg:USE_UPLOAD_FILE_CONFIG_TYPE|null) => {
    const config = ref<USE_UPLOAD_FILE_CONFIG_TYPE>({
        uploadType:'file',
        extension: [],
        type: "all",
        sourceType: ['album', 'camera'],
        maxSize: 10 * 1024 * 1024,//每一个文件上传的最大尺寸，默认为10mb
        maxCount: 9,//一次选择文件最大数量。
		// 测试地址：https://mockapi.eolink.com/tNYKNA7ac71aa90bcbe83c5815871a5b419601e96a5524d/upload
        hostUrl: "",//上传文件的服务器地址
        autoUpload: true,
        header: Object,//头部参数。
        formData: {},//额外的表单数据。
        formName: 'file',
        code: 200,
        maxDuration:20,
        sizeType:['original', 'compressed'],
        camera:'back',
        mediaType:['image','video']
    })
    config.value = uni.$tm.u.deepObjectMerge(config.value, cfg || {});
    //文件列表。
    const list = ref<FILE_TYPE[]>([])
    //当前是否正在上传中，true表示上传中，false表示中断/或者未开始上传中。
    const uploading = ref(false);
    //当前上传的索引位置。
    const activeIndex = ref(0)
    //是否中断上传。
    let isStopUpload = false;
	//是否已经超过了限制数量
	let isLimitCount = ref(false)

    let uploadTask: UniApp.UploadTask|null = null;

    /**
     * 选择文件前执行
     * @param cfg 全局参数，可以修改
     * @return {USE_UPLOAD_FILE_CONFIG_TYPE|null} 需要返回配置，如果不想修改返回null即可。
     */
    let beforeChooseFile = (cfg: USE_UPLOAD_FILE_CONFIG_TYPE | null = null) => {

        return cfg
    }
    /**
     * 选择文件后执行
     * @param files 选择后的文件列表对象，
     * @return {FILE_TYPE[]} 必须要返回文件列表，可以在这里修改待上传的文件列表地址。
     */
    let chooseFileAfter = (files: FILE_TYPE[]) => {
        return files
    }

    /**
     * 所有文件上传结束时触发。
     * @param files 当前上传的文件列表
     */
    let uploadComplete = (files: FILE_TYPE[]) => {

    }

    /**
     * 某一个文件上传结束时触发，不管失败与否都会触发。
     * @param file 当前上传的文件
     * @param files 当前上传的文件列表
     */
    let complete = (file: FILE_TYPE, files: FILE_TYPE[]) => {

    }
    /**
     * 某一个文件上传失败时触发
     * @param file 当前上传的文件
     * @param files 当前上传的文件列表
     */
    let fail = (file: FILE_TYPE, files: FILE_TYPE[]) => {

    }

    /**
     * 某一个文件上传成功时触发
     * @param file 当前上传的文件
     * @param files 当前上传的文件列表
     * @returns {FILE_TYPE} 需要返回本文件
     */
    let success = (file: FILE_TYPE, files: FILE_TYPE[]) => {

    }
    /**
     * 某一个文件上传成功后触发
     * @param file 当前上传的文件
     * @param files 当前上传的文件列表
     * @returns {FILE_TYPE} 需要返回本文件，你可以在此修改文件的statusCode,status状态等，将会最终影响文件是否真的上传成功或者修改返回值等等。
     */
    let successAfter = (file: FILE_TYPE, files: FILE_TYPE[]) => {

        return file;
    }

    /**
     * 任意一个文件上传前触发
     * @param file 当前上传的文件
     * @param files 当前上传的文件列表
     * @returns {boolean} 如果返回true将继续上传，如果返回false将阻止本次上传并设置为该文件为上传失败，然后接着上传下一个文件。
     */
    let beforeFileStart = (file: FILE_TYPE, files: FILE_TYPE[]):Promise<boolean>|boolean => {

        return true;
    }

    /**
     * 开始上传前触发
     * @param cfg 当前上传文件的配置表
     * @returns {boolean} 如果返回false将中止本次整体的上传。
     */
    let beforeStart = (cfg:USE_UPLOAD_FILE_CONFIG_TYPE):Promise<boolean>|boolean => {

        return true;
    }
	/**
	 * 删除前触发
	 * @param file 当前上传文件
	 * @returns {boolean} 如果返回false将不允许删除
	 */
	let beforeRemove = (file: FILE_TYPE):Promise<boolean>|boolean => {
	    return true;
	}
	
	/**
	 * 任意事件变动触发，管是中断，上传结束，失败等都会触发。
	 * @param file 变动的文件，可能为null
	 */
	let change = (file: FILE_TYPE|null) => {
	   
	}
	
	



    /**
     * 选择文件
     */
    function choose(type:"file"|"media" = 'file'): Promise<FILE_TYPE[]> {
        if (list.value.length >= config.value.maxCount) {
            uni.showToast({ title: "超过上传数量", icon: "none", mask: true });
            return Promise.reject("超过上传数量");
        }
        //执行上传的勾子。
        let temconfig = beforeChooseFile(uni.$tm.u.deepClone(config.value));
        config.value = uni.$tm.u.deepObjectMerge(config.value, temconfig || {});
        let cfg = {
            count: config.value.maxCount,
            extension: config.value.extension,
            type: config.value.type,
            mediaType:config.value.mediaType,
            sourceType:config.value.sourceType,
            maxDuration:config.value.maxDuration,
            sizeType:config.value.sizeType,
            camera:config.value.camera,
        }
        if (typeof cfg.extension === 'undefined' || !cfg.extension || cfg.extension?.length === 0) {
            delete cfg.extension;
        }

        return new Promise((resolve, rejects) => {
            let apiName = uni?.chooseFile;
            let ischooseVideo = false;
            // #ifdef MP-WEIXIN || MP-QQ
            apiName = uni?.chooseMessageFile
            // #endif
            // #ifdef H5 
            apiName = uni?.chooseFile;
            // #endif
            
            if(config.value.uploadType == 'media'){
                
                // #ifdef MP-WEIXIN || MP-JD || MP-TOUTIAO || MP-LARK
                apiName = uni?.chooseMedia;
                // #endif
                // #ifndef MP-WEIXIN || MP-JD || MP-TOUTIAO || MP-LARK
                if(!config.value.mediaType[0]||config.value.mediaType[0]==='image'){
                    apiName = uni?.chooseImage;
                }else if(config.value.mediaType[0]==='video'){
                    apiName = uni?.chooseVideo;
                    ischooseVideo=true;
                }
                // #endif
                
                
            }
            
			console.log(apiName)
            if(!apiName){
                uni.showModal({
                    title:"警告",
                    content:"当前只支持微信，QQ，webPC或者H5平台，其它平台不支持文件上传",
                    showCancel:false,
                    confirm:"懂了",
                })
                return;
            }
            // @ts-ignore
            apiName({
                ...cfg,
                fail(result) {
                    uni.showToast({ title: typeof result == 'object' ? JSON.stringify(result) : result, icon: "none", mask: true });
                    rejects(typeof result == 'object' ? JSON.stringify(result) : result)
                },
                success(res) {
                    let temFiles = res?.tempFiles;
                    console.log(res)
                    let temlist: any[] = [];
                    if (Array.isArray(temFiles)) {
                        temFiles.forEach(ele => {
                            
                            temlist.push({
                                uid: getUid(4),
                                url: ele.path,
                                progress: 0,
                                statusCode: ele.size > config.value.maxSize ? STATUS_CODE.max : STATUS_CODE.upload,
                                status: ele.size > config.value.maxSize ? '超过大小' : '待上传',
                                response: null,
                                name: ele.name,
                                FILE: ele
                            })
                        });
                    } else {
                      
                        if(ischooseVideo){
                            temFiles = {
                                file:res.tempFile,
                                name:res.name,
                                size:res.size,
                                path:res.tempFilePath
                            }
                        }
                        
                        temlist.push({
                            uid: getUid(4),
                            url: temFiles.path,
                            progress: 0,
                            statusCode: temFiles.size > config.value.maxSize ? STATUS_CODE.max : STATUS_CODE.upload,
                            status: temFiles.size > config.value.maxSize ? '超过大小' : '待上传',
                            response: null,
                            name: temFiles.name,
                            FILE: ischooseVideo?temFiles.file:temFiles
                        })
                    }
                    let syuCount = Math.max(config.value.maxCount - list.value.length, 0)
                    temlist = [...temlist.slice(0, syuCount)]
                    temlist = chooseFileAfter(temlist);
                    list.value.push(...temlist)
                    if(config.value.autoUpload){
                        start();
                    }
                    resolve([...temlist])
                }
            })
        })
    }
    /**
     * 删除指定文件
     * @param file 可以是索引或者文件FILE_TYPE对象
     */
    async function remove(file: number | FILE_TYPE) {
        let removeIndex = null;
        if (typeof file === 'number') {
            removeIndex = file;
        } else {
            let uid = file.uid;
            let index = list.value.findIndex(ele => ele.uid === uid);
            if (index > -1) {
                removeIndex = index;

            }
        }
        if (typeof removeIndex === null) {
            uni.showToast({ title: "删除失败,无对应文件", icon: "none", mask: true });
            return;
        }
        if (typeof removeIndex === 'number') {
            if (!list.value[removeIndex]) {
                uni.showToast({ title: "删除失败,无对应文件", icon: "none", mask: true });
                return;
            }
			//是否允许删除
			let isDel = true;
			if(beforeRemove instanceof Promise){
				isDel = await beforeRemove(uni.$tm.u.deepClone(list.value[removeIndex]));
			}else{
				isDel = beforeRemove(uni.$tm.u.deepClone(list.value[removeIndex]));
			}
			if(isDel){
				if(uploading.value){
					uni.showToast({ title: "上传中，暂无法操作", icon: "none", mask: true });
					return;
				}
				list.value.splice(removeIndex, 1);
				change(uni.$tm.u.deepClone(list.value[removeIndex]))
			}else{
                uni.showToast({ title: "删除失败,不允许", icon: "none", mask: true });
                return;
            }
        }


    }
    /**
     * 添加已有（上传文件）
     * 通常是服务器返回时的数据，用来返选用。
     * @param files 需要添加的静态文件
     */
    function addFile(files: string | FILE_TYPE | FILE_TYPE[] | string[] | Array<string | FILE_TYPE>) {
        let temList: any[] = [];
        if (typeof files === 'string') {
            temList.push({
                uid: getUid(4),
                url: files,
                progress: 100,
                statusCode: STATUS_CODE.success,
                status: '已上传',
                response: JSON.stringify({ code: 0, msg: "上传成功", data: files }),
                name: files,
                FILE: null
            })

        } else if (typeof files === 'object') {
            if (Array.isArray(files)) {
                files.forEach(el => {
                    if (typeof el === 'string') {
                        temList.push({
                            uid: getUid(4),
                            url: el,
                            progress: 100,
                            statusCode: STATUS_CODE.success,
                            status: '已上传',
                            response: JSON.stringify({ code: 0, msg: "上传成功", data: el }),
                            name: el,
                            FILE: null
                        })
                    } else if (typeof el === 'object') {
						const indexNow = list.value.findIndex(el=>el.uid===files?.uid)
						if(indexNow==-1){
							temList.push({
							    uid: getUid(4),
							    progress: 100,
							    statusCode: STATUS_CODE.success,
							    status: '已上传',
							    response: JSON.stringify(el?.response ?? { code: 0, msg: "上传成功", data: el }),
							    name: files,
							    FILE: null,
							    ...files
							})
						}
                        
                    }
                })
            } else {
				const indexNow = list.value.findIndex(el=>el.uid===files?.uid)
				if(indexNow==-1){
					temList.push({
					    uid: getUid(4),
					    progress: 100,
					    statusCode: STATUS_CODE.success,
					    status: '已上传',
					    response: JSON.stringify(files?.response ?? { code: 0, msg: "上传成功", data: files }),
					    name: files,
					    FILE: null,
					    ...files
					})
				}
                
            }
        }

        list.value.push(...temList);
    }

    /**
     * 停止当前的上传
     */
    function stop() {
        isStopUpload = true;
        uploading.value = false;
        uploadTask?.abort();
		change(null)
    }

    /**
     * 清空文件
     * @param type {"fail"|"success"|"all" } 清除文件类型
     */
    function clearFile(type:"fail"|"success"|"all" = "all"){
        if(uploading.value){
            uni.showToast({ title: "任务进行中", icon: "none", mask: true });
            return;
        }
        list.value = [];
		change(null)
    }

    /**
     * 获取文件
     * @param type {"fail"|"success"|"all" } 获取当前所有文件时的筛选条件
     */
    function getFile(type:"fail"|"success"|"all" = "success"){
        let temp = uni.$tm.u.deepClone(list.value);
        let temlist:FILE_TYPE[] = [];
        temp.forEach(el=>{
            if(type == 'all'){
                temlist.push(el) 
            }else if(type == 'fail'){
                if(el.statusCode == STATUS_CODE.fail || el.statusCode == STATUS_CODE.max){
                    temlist.push(el) 
                }
            }else if(type == 'success'){
                if(el.statusCode == STATUS_CODE.success){
                    temlist.push(el) 
                }
            }
        })

        return temlist;
    }

    async function start() {
        if (list.value.length <= 0) {
            uni.showToast({ title: "没有文件可上传", icon: "none", mask: true });
            return;
        }
        // await beforeStart(uni.$tm.u.deepClone(config.value));
        let isupload = true
        if(beforeStart instanceof Promise ){
            isupload = await beforeStart(uni.$tm.u.deepClone(config.value));
        }else{
            // @ts-ignore
            isupload = beforeStart(uni.$tm.u.deepClone(config.value));
        }
        if (uploading.value || !isupload) return;
        uploading.value = true;
        activeIndex.value = 0;
        isStopUpload = false;
		change(null)
        _upload();
    }

    async function _upload() {
        if (isStopUpload) return;
        let nowItem = reactive(list.value[activeIndex.value])
        //上结束
        if (!nowItem || typeof nowItem === 'undefined') {
            uploading.value = false;
            uploadComplete(uni.$tm.u.deepClone(list.value))
			change(null)
            return;
        }


        //成功，超过大小，上传中需要跳过上传直接下一个文件。
        if(nowItem.statusCode == STATUS_CODE.success || nowItem.statusCode == STATUS_CODE.max || nowItem.statusCode == STATUS_CODE.uploading){
            activeIndex.value+=1;
			change(uni.$tm.u.deepClone(nowItem))
            _upload();
            return;
        }

        //是否允许本次上传。
        let isupload = true
        // 执行上传前项目的勾子。
        if(beforeFileStart instanceof Promise ){
            isupload = await beforeFileStart(uni.$tm.u.deepClone(nowItem),uni.$tm.u.deepClone(list.value));
        }else{
            // @ts-ignore
            isupload = beforeFileStart(uni.$tm.u.deepClone(nowItem),uni.$tm.u.deepClone(list.value));
        }
        //跳过本次上传。
        if(!isupload){
            nowItem.status = "请更换文件";
            nowItem.statusCode = STATUS_CODE.fail;
            activeIndex.value+=1;
			change(uni.$tm.u.deepClone(nowItem))
            _upload();
            return;
        }

        //正式进入上传阶段。
        nowItem.status = "上传中.."
        nowItem.statusCode = STATUS_CODE.uploading;

        
        uploadTask = uni.uploadFile({
            url:config.value.hostUrl,
            name:config.value.formName||"file",
            filePath:nowItem.url,
            header:config.value.header||{},
            formData:{name:nowItem.name,...config.value.formData},
            fail(result) {
                nowItem.status = "上传失败";
                nowItem.statusCode = STATUS_CODE.fail;
                uni.showToast({ title: "上传失败", icon: "none", mask: true });
                fail(uni.$tm.u.deepClone(nowItem),uni.$tm.u.deepClone(list.value));
                activeIndex.value+=1;
            },
            success(result) {
                if(result.statusCode !==config.value.code){
                    nowItem.status = "上传失败";
                    nowItem.statusCode = STATUS_CODE.fail;
                    uni.showToast({ title: result.errMsg, icon: "none", mask: true });
                    fail(uni.$tm.u.deepClone(nowItem),uni.$tm.u.deepClone(list.value));
                    activeIndex.value+=1;
                    return;
                }
                let temp = uni.$tm.u.deepClone(nowItem);
                temp.status = "上传成功";
                temp.statusCode = STATUS_CODE.success;
                temp.progress = 100;
                temp.response = result.data;
                let chuliitem = successAfter(temp,uni.$tm.u.deepClone(list.value))
                for(let key in chuliitem){
                    nowItem[key] = chuliitem[key];
                }
                success(uni.$tm.u.deepClone(nowItem),uni.$tm.u.deepClone(list.value))
                activeIndex.value+=1;

            },
            complete(result) {
                complete(uni.$tm.u.deepClone(nowItem),uni.$tm.u.deepClone(list.value));
				change(uni.$tm.u.deepClone(nowItem))
				if (isStopUpload) return
                _upload();
            },
        })

        uploadTask.onProgressUpdate((result)=>{
            nowItem.progress = result.progress;
            nowItem.statusCode = STATUS_CODE.uploading;
            nowItem.status = "进度:"+parseInt(result.progress.toFixed(1))+"%";
        })

    }
	/** 添加事件函数 */
	function addEventListener(
	type:'change'|'beforeRemove'|'beforeStart'|'beforeFileStart'|'successAfter'|'success'|'fail'|'complete'|'uploadComplete'|'chooseFileAfter',
	Fun:any
	){
		if(type == 'change'){
			change = Fun;
		}else if(type == 'beforeRemove'){
			beforeRemove = Fun;
		}else if(type == 'beforeStart'){
			beforeStart = Fun;
		}else if(type == 'beforeFileStart'){
			beforeFileStart = Fun;
		}else if(type == 'successAfter'){
			successAfter = Fun;
		}else if(type == 'success'){
			success = Fun;
		}else if(type == 'fail'){
			fail = Fun;
		}else if(type == 'complete'){
			complete = Fun;
		}else if(type == 'uploadComplete'){
			uploadComplete = Fun;
		}else if(type == 'chooseFileAfter'){
			chooseFileAfter = Fun;
		}
	}

	
	
	watchEffect(()=>{
		if(list.value.length>=config.value.maxCount){
			isLimitCount.value = true;
		}else{
			isLimitCount.value = false;
		}
	})

    return {
        config,
        choose,
        files: list,
        remove,
        addFile,
        start,
        stop,
        clearFile,
        getFile,
        uploading,
        activeIndex,
		isLimitCount,
		addEventListener
    }
}
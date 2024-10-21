
//文件上传的状态值
export enum statusCode {
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
//文件对象结构
export interface file {
	url: string,//当前显示的图片地址，这是个本地临时地址。待上传前，成功后会替换服务器地址。
	status?: string,//上传状态文本
	progress?: number,//当前文件上传的进度
	uid?: string | number,//文件唯一标识id
	statusCode?: statusCode,//文件状态
	response?: any,//上传成功后的回调数据。
	name?: string,//文件名称
	[propName: string]: any
}

//上传class对象的配置。
export interface fileConfig {
	maxSize?: number,//每一个文件上传的最大尺寸，默认为10mb
	maxFile?: number,//一次选择文件最大数量。
	fileType?: Array<string>,//文件选择的类型。
	hostUrl?: string,//上传文件的服务器地址
	fileList?: Array<file>,//已上传的文件列表。
	autoUpload?: Boolean,
	header?: Object,//头部参数。
	formData?: Object,//额外的表单数据。
	formName?: string,
	statusCode?: number,//服务返回成功的状态码标志，默认200表示成功。
}
export function getUid(length = 3) {
	return Number(Number(Math.random().toString().substr(3, length) + Date.now()).toString(8));
}
/**
 * 上传文件。
 * 作者：tmzdy
 * 时间：2022年4月27日
 * 联系：zhongjihan@sina.com
 * @method {Function} beforeChooesefile -- 选择图片上传前执行的勾子。
 * @method {Function} chooesefile -- 选择图片上传，弹出层进行选择文件 。
 * @method {Function} chooesefileAfter -- 选择图片、视频文件成功后触发。返回选择后的文件。此时还未加入待上传列表，会返回 一个文件列表用于过滤。需要在函数中再返回 一个文件 列表，
 * @method {Function} chooesefileSuccess -- 文件已经加入到了待上传文件列表,需要返回过滤的文件列表
 * @method {Function} beforeAddfile -- 动态加入预计加入文件前执行的勾子，返回true才会正式加入到待上传列表中。
 * @method {Function} addFile -- 动态加入预上传的文件。
 * @method {Function} progress -- 进度。
 * @method {Function} fail -- 失败。服务器出现非200时触发。
 * @method {Function} beforeSuccess -- 服务器返回成功，立即执行的勾子。如果此时返回false,则表示文件 上传失败。
 * @method {Function} success -- 成功。
 * @method {Function} complete -- 单个文件完成。不管上传成功与失败否，都会触发此函数。
 * @method {Function} uploadComplete -- 所有文件上传完时触发
 * @method {Function} beforeStart -- 开始上传。前的校验勾子。如果返回false,将阻止上传
 * @method {Function} start -- 开始上传。
 * @method {Function} stop -- 停止上传。中止上传时触发的函数。
 */
export class uploadfile {
	//文件列表。
	filelist: Array<file> = [];
	isStop = false;
	index = 0;
	config: fileConfig = {};
	uploadobj: UniNamespace.UploadTask | null = null;
	constructor(config: fileConfig) {
		let cf: fileConfig = { maxSize: 10 * 1024 * 1024, maxFile: 9, fileType: ['album', 'camera'], fileList: [], autoUpload: true, header: {}, formData: {}, formName: 'file' }
		cf = { ...cf, ...arguments[0] ?? {} };
		//配置{name: 'file', // 上传时的文件key名。默认file,header: {}, // 上传的头部参数。}
		this.config = cf;
		this.addFile(cf.fileList);
		delete this.config.fileList;
	}
	async beforeChooesefile() {
		return true;
	}
	async chooesefileAfter(fileList: Array<file>) {

		return fileList;
	}
	async chooesefileSuccess(fileList: Array<file>) {

		return fileList;
	}
	delete(item: file) {
		let index = this.filelist.findIndex(el => el.uid == item.uid);
		if (index > -1) {
			let p = [...this.filelist]
			p.splice(index, 1)
			this.filelist = [...p];
		}

		return this.filelist;
	}
	async clear() {
		/** 清清前要选暂停所有正在上传的文件 */
		this.stop();
		this.filelist = [];
	}
	setFileStatus(item: file) {
		let index = this.filelist.findIndex(el => el.uid == item.uid);
		if (index > -1) {
			let p = [...this.filelist]
			p.splice(index, 1, item)
			this.filelist = [...p];
		}
	}
	/**
	 * 成功后返回选择后的图片列表。
	 */
	async chooesefile(): Promise<Array<file>> {
		let t = this;

		return new Promise(async (rs, rj) => {
			let isready = await t.beforeChooesefile();
			if (!isready) {
				rs([])
				return;
			}
			uni.chooseImage({
				count: t.config.maxFile,
				sourceType: t.config.fileType,
				fail: (e) => {
					rj("取消选择");
				},
				success: async (res) => {

					if (res.tempFilePaths.length == 0) {
						rj("未选择")
						return;
					}

					let imgarray = res.tempFilePaths;
					let fielist = res.tempFiles;
					let jgsk: Array<file> = [];
					//0待上传，1上传中，2上传失败，3上传成功。4超过大小限制
					imgarray.forEach((item: string, index: number) => {
						let isMaxsize = fielist[index].size > t.config.maxSize ? true : false;
						jgsk.push({
							url: item,
							status: isMaxsize ? '超过大小' : "待上传",
							progress: isMaxsize ? 100 : 0,
							uid: getUid(),
							statusCode: isMaxsize ? statusCode.max : statusCode.upload,
							response: null,
							name: fielist[index].name ?? ""
						})
					})

					let isreadyChoose = await t.chooesefileAfter(jgsk);

					if (!Array.isArray(isreadyChoose) || typeof isreadyChoose != 'object') {
						rj("chooesefileAfter:函数过滤，没有返回文件列表。")
						return;
					}

					t.filelist.push(...isreadyChoose)

					t.chooesefileSuccess(isreadyChoose);
					rs(isreadyChoose)
					if (t.config.autoUpload) {
						setTimeout(function () {
							t.start();
						}, 500);
					}


				}
			})
		})
	}
	async chooseMPH5weixinFile() {
		let t = this;
		return new Promise((rs, rj) => {
			var fs = uni.chooseFile;
			// #ifdef MP-WEIXIN || MP-QQ
			fs = uni.chooseMessageFile;
			// #endif
			var config = {
				count: t.config.maxfile,
				type: t.config.type,
				extension: t.config.extension,
			}
			if (!t.config.extension || !Array.isArray(t.config.extension) || t.config.extension?.length == 0) {
				delete config.extension
			}
			fs({
				...config,
				fail: (e) => {
					console.error(e);
					uni.$tm.toast("已取消选择");
					rj(e);
				},
				success: (res) => {
					if (res.tempFiles.length == 0) {
						uni.$tm.toast("未选择");
						return;
					}
					let fielist = res.tempFiles;
					let jgsk = [];
					//0待上传，1上传中，2上传失败，3上传成功。4超过大小限制
					fielist.forEach((item, index) => {
						let isMaxsize = fielist[index].size > t.config.maxsize ? true : false;
						let ftype = item.name || ""
						if (ftype) {
							ftype = ftype.substr(ftype.lastIndexOf(".") + 1).toLocaleLowerCase();
						}
						jgsk.push({
							url: item.path,
							name: item.name || '默认文件名称',
							type: ftype,
							status: isMaxsize ? '超过大小' : "待上传",
							progress: isMaxsize ? 100 : 0,
							fileId: guid(),
							statusCode: isMaxsize ? 4 : 0,
							data: null,//上传成功后的回调数据。
						})
					})
					t.filelist.push(...jgsk)

					t.selected(t.filelist);
					if (t.config.isAuto) {
						t.start();
					}

					rs(t.filelist)
				}
			})

		})
	}
	setConfig(config: fileConfig) {
		this.config = { ...this.config, ...config ?? {} }
	}
	/**
	 * 动态加入文件
	 * @param {Object} filelist
	 */
	addFile(filelist: Array<file> = []) {
		if (typeof filelist !== 'object' && !Array.isArray(filelist)) return;
		let total_uid = new Set(this.filelist.map(e => e.uid))
		let total_url = new Set(this.filelist.map(e => e.url))
		let cfilelist = filelist.map(el => {
			return {
				...el,
				status: el?.status ?? "待上传",
				statusCode: el?.statusCode ?? statusCode.upload,
				uid: el?.uid ?? getUid(),
				progress: el?.progress ?? 0,
				name: el?.name ?? "",
				response: el?.response ?? null,
				url: el?.url ?? ""
			}
		})
		let filterFIle = cfilelist.filter(item => !total_uid.has(item.uid) && !total_url.has(item.url))
		this.filelist.push(...filterFIle)

	}
	beforeSuccess(item: file) {
		return Promise.resolve(true);
	}
	beforeStart(item: file) {
		return Promise.resolve(true);
	}
	// 进度。
	progress(item: file,index:number) { }
	// 失败
	fail(item: file) { }
	// 成功
	success(item: file, fileList: Array<file>) { }
	// 完成。
	complete(filelist: file) { }
	uploadComplete(filelist: Array<file>) { }
	awaitTime() {
		return new Promise((resolve) => {
			setTimeout(() => {
				resolve(true)
			}, 20)
		})
	}
	// 开始上传。
	async start() {

		if (this.filelist.length <= 0) {
			console.error("未选择图片,已取消上传")
			return;
		}
		let t = this;
		// t重新开始上传从头开始。
		this.index = 0;
		this.isStop = false;
		async function startupload() {
			if (t.isStop) return;

			let item = t.filelist[t.index];

			if (!item || typeof item === 'undefined') {
				// 文件不存在。直接结束。
				t.uploadComplete(t.filelist);
				return;
			}

			let canbleStart = await t.beforeStart(item)
			if (!canbleStart) {
				item.statusCode = statusCode.fail;
				item.status = "不允许上传"
				t.filelist.splice(t.index, 1, item)
				t.index++;
				t.setFileStatus(item)
				t.fail(item)
				t.complete(item);
				startupload();
				return;
			}

			if (item.statusCode == 3 || item.statusCode == 1 || item.statusCode == 4 || item.statusCode == 2) {
				// 直接跳过。至下一个文件。
				t.index++;
				startupload();
				return;
			}

			item.statusCode = statusCode.uploading;
			item.status = "上传中..."
			t.setFileStatus(item)
			const upObj = t.uploadobj = uni.uploadFile({
				url: String(t.config.hostUrl),
				name: t.config?.formName ?? 'file',
				header: t.config?.header ?? {},
				filePath: item.url,
				formData: { name: item.name, ...t.config.formData },
				success: async (res) => {
					if (t.isStop) return
					item.response = res.data;
					let isOksuccess = await t.beforeSuccess(item);
					const statusCode_reonese = t.config?.statusCode ?? 200
					if (res.statusCode != statusCode_reonese || !isOksuccess) {
						item.statusCode = statusCode.fail;
						item.status = "上传失败";
						t.fail(item)
						t.setFileStatus(item)
						t.index++;
						return;
					}

					// 上传成功。
					item.statusCode = statusCode.success;
					item.status = "上传成功";
					t.setFileStatus(item)
					t.success(item, t.filelist)
					t.index++;
				},
				fail: (res) => {
					if (t.isStop) return
					item.statusCode = statusCode.fail;
					item.status = "上传失败";
					t.setFileStatus(item)
					t.fail(item)
					t.index++;
				},
				complete: async (res) => {
					if (t.isStop) return
					await t.awaitTime();
					t.complete(item);
					// 直接下一个文件。
					startupload();
				}
			})
			if (upObj) {
				let item = t.filelist[t.index];
				upObj.onProgressUpdate(async (res) => {
					if (t.isStop) return
					item.progress = res.progress;
					item.statusCode = statusCode.uploading;
					item.status = "...";
					t.setFileStatus(item)
					t.progress(item,t.index)

				})
			}

		}
		await startupload();
	}
	// 停止上传
	stop() {
		this.isStop = true;
		if (this.uploadobj != null) {
			this.uploadobj.abort()
		}
	}

}

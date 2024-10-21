import { fetchNet } from './tool/lib/fetch';
import fontJson from './tool/tmicon/fontJson';
import { setDomDarkOrWhite } from './tool/theme/util';
import preview, * as util from './tool/function/util';
import { language, languageByGlobal } from "./tool/lib/language"
import { share } from "./tool/lib/share"
import { App, nextTick} from "vue"
import PageJsonInit from "../pages.json"
import { useTmRouterBefore } from "./tool/router/index"
import tmuiconfigdefault from "./tool/lib/tmuiconfigDefault"
import { pagesType, tabBarType, pagesCustomType } from './interface';
import * as Pinia from 'pinia';
let pages: Array<pagesType> = [];
if (typeof PageJsonInit?.pages == 'undefined') {
	PageJsonInit.pages = [];
}
PageJsonInit.pages.forEach((el: any) => {
	let customType: pagesCustomType = <pagesCustomType>(el?.style?.navigationStyle ?? "default");
	let bg = (el.style?.navigationBarBackgroundColor ?? PageJsonInit?.globalStyle?.navigationBarBackgroundColor ?? '#FFFFFF') || '#FFFFFF'
	let txtColor = (el.style?.navigationBarTextStyle ?? PageJsonInit?.globalStyle?.navigationBarTextStyle ?? 'black') || 'black'
	pages.push({
		path: el.path,
		custom: customType,
		navigationBarBackgroundColor: bg,
		navigationBarTextStyle: txtColor
	})
})
if (Array.isArray(PageJsonInit?.subPackages ?? null)) {
	PageJsonInit?.subPackages.forEach((el: any) => {
		let rootPath = el.root;
		el.pages.forEach((el2: any) => {
			let elany: any = el2;
			let bg = (el2.style?.navigationBarBackgroundColor ?? PageJsonInit?.globalStyle?.navigationBarBackgroundColor ?? '#FFFFFF') || '#FFFFFF'
			let txtColor = (el2.style?.navigationBarTextStyle ?? PageJsonInit?.globalStyle?.navigationBarTextStyle ?? 'black') || 'black'
			pages.push({
				path: rootPath + "/" + elany.path,
				custom: elany?.style?.navigationStyle ?? "default",
				navigationBarBackgroundColor: bg,
				navigationBarTextStyle: txtColor
			})
		})
	})
}
let pagers: any = PageJsonInit;
let tabBar: tabBarType = pagers?.tabBar ?? {
	color: "",
	selectedColor: "",
	borderStyle: "",
	backgroundColor: "",
	list: []
}

// custom icon
let cusutomIconList = [];
// #ifdef APP
cusutomIconList = fontJson;
// #endif
let $tm = {
	tabBar: tabBar,
	pages: pages,
	globalNavStyle: (PageJsonInit?.globalStyle.navigationStyle ?? ""),
	isOpenDarkModel: (PageJsonInit?.globalStyle?.navigationBarBackgroundColor ?? "").indexOf("@") > -1,
	isColor: (color: string) => {
		const reg1 = /^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/;
		const reg2 = /^(rgb|RGB|rgba|RGBA)/;
		return reg1.test(color) || reg2.test(color);
	},
	/**tmui3.0工具函数 */
	u: { ...util, preview },
	/**tmui3.0国际化语言辅助函数 */
	language: language,
	fetch: fetchNet,
	tmicon: [
		{
			font: "tmicon",
			prefix: "tmicon-",
			fontJson: cusutomIconList
		}
	],
	config: tmuiconfigdefault as Tmui.tmuiConfig
};

export default {
	/**
	 * tmui3.0
	 * @param app Vue
	 * @param options tmui3.0配置
	 */
	install: (app: App, options: Tmui.tmuiConfig = {}) => {
		uni.addInterceptor('navigateTo', {
			invoke(result) {
				nextTick(() => {
					linsInko({
						path: result.url,
						context: null,
						openType: 'navigateTo'
					})
				})
			},
			success(result) {

			}
		})
		uni.addInterceptor('redirectTo', {
			success(result) {
				let pages = getCurrentPages().pop()
				let path = pages?.route ?? "";
				let msg = result.errMsg ?? "";
				let opentype = msg.split(":")[0] ?? "";
				linsInko({
					path: path,
					context: null,
					openType: opentype
				})
			}
		})
		uni.addInterceptor('reLaunch', {
			success(result) {
				let pages = getCurrentPages().pop()
				let path = pages?.route ?? "";
				let msg = result.errMsg ?? "";
				let opentype = msg.split(":")[0] ?? "";
				// 这里的链接是去往的链接
				linsInko({
					path: path,
					context: null,
					openType: opentype
				})
			}
		})
		uni.addInterceptor('navigateBack', {
			invoke(result) {
				nextTick(() => {
					let pages = getCurrentPages().pop()
					let path = pages?.route ?? "";
					let msg = result.errMsg ?? "";
					let opentype = msg.split(":")[0] ?? "";
					// 这里返回的链接是返回前的链接，并非返回后的链接。
					//这里在h5端不需要监测，因为有全局监测实现。
					// #ifndef H5
					linsInko({
						path: path,
						context: null,
						openType: "navigateBack"
					})
					// #endif
				})
			},
			success(result) {

			}
		})
		// #ifdef H5
		window.addEventListener('popstate', (ev: any) => {
			linsInko({
				path: ev?.state?.forward ?? "",
				context: null,
				openType: "navigateBack"
			})
		})
		// #endif
		//路由拦截
		function linsInko(obj: any) {
			// #ifdef H5
			setDomDarkOrWhite();
			// #endif
			obj.path = obj.path[0] == "/" ? obj.path.substr(1) : obj.path
			// useTmRouterBefore(obj)
			options.router?.useTmRouterBefore ? options.router?.useTmRouterBefore(obj) : useTmRouterBefore(obj)
		}

		options = util.deepObjectMerge($tm.config, options)

		const pinia = app.config.globalProperties.$pinia || null
		const tmPiniaPlugin = (context: Pinia.PiniaPluginContext) => {
			if (context.store.$id === 'tmpinia') {
				context.store.tmuiConfig = options
				context.store.$state.tmuiConfig = options
			}
		};
		if (pinia) {
			pinia.use(tmPiniaPlugin)
		} else {
			const pinia = Pinia.createPinia()
			pinia.use(tmPiniaPlugin)
			app.use(pinia)
		}

		// #ifndef APP-NVUE
		app.use(languageByGlobal())
		// #endif
		let appconfig = {};
		// #ifdef MP

		if (!$tm.config.shareDisable) {
			const { onShareAppMessage, onShareTimeline } = share()
			appconfig = { ...appconfig, onShareAppMessage, onShareTimeline }

		}
		// #endif

		app.mixin({
			...appconfig,
		})


		/**合并插件的全局配置 */
		$tm = {
			...$tm,
			config: options
		}

		/**对外暴露 */
		uni.$tm = $tm;
		// #ifdef APP-VUE
		uni.setStorageSync("$tm", JSON.stringify($tm.config.theme));
		// #endif

		/**app应用上下文的暴露 */
		app.config.globalProperties.tm = $tm;
	}
};

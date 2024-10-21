<template>
	<!-- #ifdef APP-NVUE -->
	<view class="flex relative"
	:style="
	[
		appConfig.theme ? { background: transparent ? 'rgba(0,0,0,0)' : appConfig.theme } : '',
		 _bgStyle
	 ]"
	>
		<slot name="default"></slot>
	</view>
	<!-- #endif -->
	<!-- #ifndef APP-NVUE -->
	<view
		class="flex relative app flex-1"
		:style="
		[
			appConfig.theme ? { background: transparent ? '' : appConfig.theme } : '',
			 _props.bgImg ? { backgroundImage: `url(${_props.bgImg})` } : '',
			 _bgStyle
		 ]"
	>
		<slot name="default"></slot>
	</view>
	<!-- #endif -->
</template>
<script lang="ts" setup>
/**
 * 应用容器
 * @description 这是所有创建应用的最外层基础容器。
 */
import { provide,computed, onMounted, ref, watch, PropType } from 'vue'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import { custom_props} from '../../tool/lib/minxs'
import { onLoad} from '@dcloudio/uni-app'
import useTheme from '../../tool/useFun/useTheme'
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'
const store = useTmpiniaStore()


//路由守卫---------------------------------

// end-----------------------------------------
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	/**
	 * 是否透明背景
	 */
	transprent: {
		type: [Boolean, String],
		default: false
	},
	/**
	 * 是否透明背景,等同transprent,因单词拼写错误，现在写一个正确的。
	 */
	transparent: {
		type: [Boolean, String],
		default: false
	},
	//整体的主题色调同全局色一样。
	/**暂时不可用 */
	theme: {
		type: String,
		default: 'grey-5'
	},
	bgImg: {
		type: String,
		default: ''
	},
	/** 背景层div的样式 */
	bgStyle: {
		type: String,
		default: ''
	},
	//应用的背景颜色。
	color: {
		type: String,
		default: 'grey-4'
	},
	/**自定义暗黑的背景，你也可以通过全局配置 */
	darkColor: {
		type: String,
		default: '#050505'
	},
	/**是否模糊背景，暂时不可用 */
	blur: {
		type: [Boolean, String],
		default: false
	},
	/**这是一个淘汰的属性，请在pages.json中配置即可，会自动读取。而不需要在这里设置 */
	navbar: {
		type: Object as PropType<{ background: string; fontColor: '#ffffff' | '#000000' }>,
		default: () => {
			return {
				background: '',
				fontColor: ''
			}
		}
	},
	/**是否自动修改系统自带的navbar的主题。 */
	navbarDarkAuto:{
		type: Boolean,
		default: true
	}
})
// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore)
const isSetThemeOk = ref(false)

const {dark,isNvue,customCSSStyle,customClass,parentClass,transparent,_props,proxy,blur,
	round,margin,padding,theme
} = useTheme(computed(()=>props),tmcfg);
//计算主题
const tmcomputed = theme()

let timids: any = NaN
let flag = false

//本页面是否是tabar切换页面。
let isTabbarPage = false
let nowPage = getCurrentPages().pop()
const _bgStyle = computed(()=>props.bgStyle)
/**整体的配置 */
let winSize = useWindowInfo();
let appConfig = ref({
	width: winSize.width,
	height: winSize.height,
	theme: tmcomputed.value.backgroundColor,
	bgImg: props.bgImg,
	dark: dark.value
})

onLoad((opts: any) => {
	try {
		store.tmuiConfig.router?.useTmRouterAfter({
			path: nowPage?.route ?? '',
			opts: opts,
			context: proxy ?? null
		})
	} catch (error) {}
})
onMounted(() => {
	_onInit()
})



function _onInit() {
	/**
	 * 读取pages.json中的bar页面列表。
	 */
	let barLit = uni.$tm.tabBar?.list ?? []
	for (let i = 0; i < barLit.length; i++) {
		if (nowPage?.route == barLit[i].pagePath) {
			isTabbarPage = true
			break
		}
	}

	if (store.tmuiConfig.autoDark) {
		// #ifdef H5
		if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
			osChangeTheme('dark')
		} else {
			osChangeTheme('light')
		}
		// #endif

		// #ifndef H5
		osChangeTheme(uni.getSystemInfoSync().osTheme)
		// #endif
	} else {
		setAppStyle()
	}
}

watch([() => tmcfg.value.color, dark], () => {
	isSetThemeOk.value = false
	setAppStyle()
})
watch([() => props.color], () => {
	appConfig.value.theme = tmcomputed.value.backgroundColor
})


function setAppStyle() {
	if (dark.value) {
		appConfig.value.theme = store.tmuiConfig?.themeConfig?.dark?.bodyColor || props.darkColor
	} else {
		appConfig.value.theme = tmcomputed.value.backgroundColor
	}
	// #ifdef MP-WEIXIN || MP-BAIDU || MP-QQ || MP-KUAISHOU || MP-LARK
	uni.setBackgroundColor({
		backgroundColor: appConfig.value.theme,
		backgroundColorBottom: appConfig.value.theme,
		backgroundColorTop: appConfig.value.theme
	}).catch((error) => {})
	// #endif

	// #ifdef APP-NVUE ||  APP-VUE
	if (plus?.webview?.currentWebview()?.setStyle) {
		plus.webview.currentWebview().setStyle({
			background: appConfig.value.theme,
			backgroundColorTop: appConfig.value.theme,
			backgroundColorBottom: appConfig.value.theme,
			userSelect: true,
			webviewBGTransparent: true
		})
	}

	// #endif

	// #ifdef H5
	document.body.style.background = appConfig.value.theme || ''
	localStorage.setItem(
		'tmuiNavStyle',
		JSON.stringify({
			navbarBackground: dark.value ? appConfig.value.theme : props.navbar.background,
			navbarFontColor: dark.value ? '#ffffff' : props.navbar.fontColor
		})
	)
	// #endif

	if (dark.value) {
		// #ifndef MP-ALIPAY
		try {
			if (!uni.$tm.isOpenDarkModel&&props.navbarDarkAuto) {
				uni.setNavigationBarColor({
					backgroundColor: '#000000',
					frontColor: '#ffffff'
				}).catch((error) => {})
			}
		} catch (error) {}

		// #endif

		// #ifdef APP
		plus.navigator.setStatusBarStyle('light')
		// #endif

		/**如果当前页面有tabbar则需要设定原生的tabbar */
		if (isTabbarPage) {
			if (uni.$tm.tabBar.selectedColor.indexOf('@') == -1) {
				uni.setTabBarStyle({
					backgroundColor: '#141415',
					borderStyle: 'black',
					color: '#ffffff',
					selectedColor: uni.$tm.tabBar.selectedColor || tmcomputed.value.textColor
				}).catch((error) => {})
			}
		}
	} else {
		// #ifndef MP-ALIPAY
		try {
			let nowPageConfigs = uni.$tm.pages.filter((el) => el.path == nowPage?.route)
			if (nowPageConfigs.length > 0 && !props.navbar.background) {
				if (nowPageConfigs[0].navigationBarTextStyle.indexOf('@') > -1) return
				let tcolor = nowPageConfigs[0].navigationBarTextStyle
				tcolor = tcolor.toLocaleLowerCase()
				tcolor = tcolor == 'black' ? '#000000' : tcolor
				tcolor = tcolor == 'white' ? '#ffffff' : tcolor
				if(props.navbarDarkAuto){
					uni.setNavigationBarColor({
						backgroundColor: nowPageConfigs[0].navigationBarBackgroundColor,
						frontColor: tcolor
					}).catch((error) => {})
				}
				
				uni.setStorageSync(
					'tmuiNavStyle',
					JSON.stringify({
						navbarBackground: nowPageConfigs[0].navigationBarBackgroundColor,
						navbarFontColor: tcolor
					})
				)
			} else if (!uni.$tm.isOpenDarkModel) {
				if(props.navbarDarkAuto){
					uni.setNavigationBarColor({
						backgroundColor: props.navbar.background,
						frontColor: props.navbar.fontColor
					}).catch((error) => {})
				}
				
				uni.setStorageSync(
					'tmuiNavStyle',
					JSON.stringify({
						navbarBackground: props.navbar.background,
						navbarFontColor: props.navbar.fontColor
					})
				)
			}
		} catch (error) {}
		// #endif
		// #ifdef APP
		plus.navigator.setStatusBarStyle('dark')
		// #endif
		try {
			if (isTabbarPage && !uni.$tm.isOpenDarkModel) {
				uni.setTabBarStyle({
					backgroundColor: uni.$tm.tabBar.backgroundColor || props.navbar.background,
					borderStyle: uni.$tm.tabBar.borderStyle || 'white',
					color: uni.$tm.tabBar.color || props.navbar.fontColor,
					selectedColor: uni.$tm.tabBar.selectedColor || tmcomputed.value.textColor
				}).catch((error) => {})
			}
		} catch (error) {}
	}
	isSetThemeOk.value = true
}

/**
 * 设定主题
 * @params colorName 主题名称
 */
function setTheme(colorName: string) {
	store.setTmVuetifyTheme(colorName)
}

/**
 * 设定暗黑
 * @param dark boolean 空表示自动处理和切换暗黑效果
 */
function setDark(darks?: boolean) {
	let maindark = !dark.value
	if (typeof darks !== 'undefined' && typeof darks == 'boolean') {
		maindark = darks
	}
	appConfig.value.dark = maindark
	store.setTmVuetifyDark(maindark)

	setAppStyle()
}

/** 监听当前系统的主题变化。 */
try {
	uni.onThemeChange((ev) => {
		osChangeTheme(ev.theme)
	})
} catch (error) {
	console.warn('没有主题切换功能：', error)
}
/**自动跟随系统设定暗黑主题。 */
function osChangeTheme(ev: 'light' | 'dark' | string | undefined) {
	if (!store.tmuiConfig.autoDark) return
	if (ev === 'light') {
		setDark(false)
	} else if (ev === 'dark') {
		setDark(true)
	}
}


//向下子组件传递，相关参数，可代替store使用。
provide(
	'appTextColor',
	computed(() => tmcomputed.value.textColor)
)
//各个组件之间的间隙。
provide('custom_space_size', [0, 0])

//向ref外层公开本组件的特定方法。
defineExpose({
	setTheme: setTheme,
	setDark: setDark
})
</script>

<style scoped>
.app {
	/* #ifndef APP-NVUE */
	background-attachment: fixed;
	background-position: 0px 0px;
	background-repeat: no-repeat;
	background-size: 100% 100%;
	
	/* #endif */
	flex-direction: column;
}
</style>

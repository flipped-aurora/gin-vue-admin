//配置国际化标准语言。
import en from '../../locale/en.json'
import zhHans from '../../locale/zh-Hans.json'
import { createI18n, I18nOptions,   } from 'vue-i18n'
import { initVueI18n} from '@dcloudio/uni-i18n'
export const language= function(key:string){
	const messages:any = {
		en,
		'zh-Hans':zhHans
	}
	let i18nConfig:I18nOptions | undefined = {
	  locale: uni.getLocale(),// 获取已设置的语言
	  messages
	}
	// #ifndef APP-NVUE
	const i18n = createI18n(i18nConfig)
	return i18n.global.t(key)
	// #endif
	// #ifdef APP-NVUE
	const { t } = initVueI18n(messages)
	return t(key);
	// #endif
	
}
//保留原始调用方法。
export const languageByGlobal= function(){
	const messages  = {
		en,
		'zh-Hans':zhHans
	}
	let i18nConfig = {
	  locale: uni.getLocale(),// 获取已设置的语言
	  messages
	}
	const i18n = createI18n(i18nConfig)
	return i18n
}
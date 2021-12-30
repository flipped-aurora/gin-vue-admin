import { createI18n as _createI18n } from 'vue-i18n'
import elementEnLocale from 'element-plus/es/locale/lang/en' // english lang
import elementZhLocale from 'element-plus/es/locale/lang/zh-cn'// chinese lang
import elementArLocale from 'element-plus/es/locale/lang/ar'// arabic lang
import enLocale from './locales/en.json'
import zhLocale from './locales/zh.json'
import arLocale from './locales/ar.json'
import Cookies from 'js-cookie'

// const messageImports = import.meta.glob('./locales/*.json')
// console.log(messageImports)

// export const SUPPORT_LOCALES = ['en', 'zh', 'ar']

const messages = {
  en: {
    ...enLocale,
    ...elementEnLocale
  },
  zh: {
    ...zhLocale,
    ...elementZhLocale
  },
  ar: {
    ...arLocale,
    ...elementArLocale
  }
}

export function createI18n() {
  return _createI18n({
    legacy: false,
    globalInjection: true,
    // locale: 'en',
    // set locale
    // options: en | zh | ar
    locale: Cookies.get('language') || 'en', // get selected language from cookies
    fallbackLocale: 'en',
    // set locale messages
    messages
    // messages: {
    //   enLang,
    //   zhLang,
    //   arLang
    // }

  })
}

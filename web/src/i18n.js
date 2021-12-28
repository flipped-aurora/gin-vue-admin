import { createI18n as _createI18n } from 'vue-i18n'
import en from './locales/en.json'
import cn from './locales/cn.json'
import ar from './locales/ar.json'

// const messageImports = import.meta.glob('./locales/*.json')
// console.log(messageImports)

export const SUPPORT_LOCALES = ['en', 'cn', 'ar']

export function createI18n() {
  return _createI18n({
    legacy: false,
    globalInjection: true,
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
      en,
      cn,
      ar
    }
  })
}
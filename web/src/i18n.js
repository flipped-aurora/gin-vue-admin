import { createI18n } from 'vue-i18n'
import enLocale from './locales/en.json'
import zhLocale from './locales/zh.json'
import arLocale from './locales/ar.json'
import Cookies from 'js-cookie'

const messages = {
  en: {
    ...enLocale,
  },
  zh: {
    ...zhLocale,
  },
  ar: {
    ...arLocale,
  }
}

export default createI18n({
  legacy: false, // you must set `false`, to use Composition API
  globalInjection: true,
  locale: process.env.VUE_APP_I18N_LOCALE || Cookies.get('language') || 'en', // get selected language from cookies
  fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || 'en',
  // messages: loadLocaleMessages()
  messages
})

/*
function loadLocaleMessages () {
  const locales = require.context('./locales', true, /[A-Za-z0-9-_,\s]+\.json$/i)
  const messages = {}
  locales.keys().forEach(key => {
    const matched = key.match(/([A-Za-z0-9-_]+)\./i)
    if (matched && matched.length > 1) {
      const locale = matched[1]
      messages[locale] = locales(key).default
    }
  })
  return messages
}

export default createI18n({
  locale: process.env.VUE_APP_I18N_LOCALE || 'en',
  fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || 'en',
  messages: loadLocaleMessages()
})
*/

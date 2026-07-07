import { fmtTitle } from '@/utils/fmtRouterTitle'
import config from '@/core/config'
import i18n from '@/i18n'
export default function getPageTitle(pageTitle, route) {
  if (pageTitle) {
    const title = i18n.global.t(fmtTitle(pageTitle, route))
    return `${title} - ${config.appName}`
  }
  return `${config.appName}`
}

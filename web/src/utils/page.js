import config from '@/core/config'
export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${config.appName}`
  }
  return `${config.appName}`
}

import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage


export const formatBoolean = (bool) => {
  const { t } = useI18n() // added by mohamed hassan to support multilanguage
  if (bool !== null) {
    return bool ? t('general.yes') : t('general.no')
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}

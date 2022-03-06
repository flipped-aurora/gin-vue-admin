import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import i18n from '@/i18n' // added by mohamed hassan to multilangauge

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? i18n.t('general.yes') : i18n.t('general.no')
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

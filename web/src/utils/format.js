import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
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

export const filterDataSource = (dataSource, value) => {
  const rowLabel = dataSource && dataSource.find(item => item.value === value)
  return rowLabel?.label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}

const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'
export const ReturnArrImg = (arr) => {
  const imgArr = []
  if (arr instanceof Array) { // 如果是数组类型
    for (const arrKey in arr) {
      if (arr[arrKey].slice(0, 4) !== 'http') {
        imgArr.push(path + arr[arrKey])
      } else {
        imgArr.push(arr[arrKey])
      }
    }
  } else { // 如果不是数组类型
    if (arr.slice(0, 4) !== 'http') {
      imgArr.push(path + arr)
    } else {
      imgArr.push(arr)
    }
  }
  return imgArr
}

export const onDownloadFile = (url) => {
  window.open(path + url)
}
const colorToHex = u=>{
  let e = u.replace("#", "").match(/../g);
  for (let t = 0; t < 3; t++)
    e[t] = parseInt(e[t], 16);
  return e
}

const  hexToColor = (u,e,t)=>{
  let a = [u.toString(16), e.toString(16), t.toString(16)];
  for (let n = 0; n < 3; n++)
    a[n].length === 1 && (a[n] = `0${a[n]}`);
  return `#${a.join("")}`
}
const  generateAllColors = (u,e)=> {
  let t = colorToHex(u);
  for (let a = 0; a < 3; a++)
    t[a] = Math.floor((255 - t[a]) * e + t[a]);
  return hexToColor(t[0], t[1], t[2])
}



function addOpacityToColor(u, opacity) {
  let t = colorToHex(u);
  return `rgba(${t[0]}, ${t[1]}, ${ t[2]}, ${opacity})`;
}


export const setBodyPrimaryColor = (  primaryColor ) =>{
  document.documentElement.style.setProperty('--el-color-primary', primaryColor)
  for (let times = 1; times <= 2; times++) {
    document.documentElement.style.setProperty(`--el-color-primary-dark-${times}`,  generateAllColors(primaryColor, times / 10))
  }
  for (let times = 1; times <= 10; times++) {
    document.documentElement.style.setProperty(`--el-color-primary-light-${times}`,  generateAllColors(primaryColor, times / 10))
  }
  document.documentElement.style.setProperty(`--el-menu-hover-bg-color`,  addOpacityToColor(primaryColor, 0.1))
}

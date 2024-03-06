function converResDataToPrintData(responseData, info, source) {
  const currentDate = new Date() // 获取当前时间
  const year = currentDate.getFullYear().toString() // 获取年份
  const month = (currentDate.getMonth() + 1).toString().padStart(2, '0') // 获取月份（注意要加上1）
  const day = currentDate.getDate().toString().padStart(2, '0') // 获取天数
  const hours = currentDate.getHours().toString().padStart(2, '0') // 获取小时（两位）
  const minutes = currentDate.getMinutes().toString().padStart(2, '0') // 获取分钟（两位）
  const formattedDate = `${year}-${month}-${day}` // 组合成指定格式的字符串

  const printData = []

  responseData.map((item) => {
    const printItem = {}
    printItem.hall = info.hall + '号厅' || ''
    printItem.today = formattedDate + ' ' + info.playTime
    printItem.filmName = info.name
    printItem.seat = item[0] + '排' + item[1] + '座'
    printItem.type = source
    printItem.datetime = formattedDate + ' ' + hours + ':' + minutes
    printItem.qrcodeUrl = getRandomString()
    printItem.orderNo = getOrderNo()
    printItem.time = info.playTime
    printItem.price = parseFloat(info.price).toFixed(2) + '元'
    printItem.seller = '黄岩星驰国际影城'
    printItem.sellerFee = '0.00元'
    printItem.charge = '0.00'

    printData.push(printItem)
  })

  return printData
}

function getRandomString() {
  const letters = 'QWERTYUIOPASDFGHJKLZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789';
  let randomLetters = ''
  for (let i = 0; i < 90; i++) {
    const randomIndex = Math.floor(Math.random() * letters.length)
    // eslint-disable-next-line no-const-assign
    randomLetters += letters.charAt(randomIndex)
  }
  return randomLetters
}

function getOrderNo() {
  // 生成3位随机值
  const lettersOrder = '0123456789'
  let randomOrder = ''
  for (let i = 0; i < 15; i++) {
    const randomIndex = Math.floor(Math.random() * lettersOrder.length);
    randomOrder += lettersOrder.charAt(randomIndex)
  }
  // 拼接时间戳和随机值
  return randomOrder
}

export default converResDataToPrintData;

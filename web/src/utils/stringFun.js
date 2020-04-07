export const toUpperCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toUpperCase())
    } else {
        return ""
    }
}

// 驼峰转换下划线
export const toSQLLine = (str) => {
    if (str=="ID") return "ID"
    return str.replace(/([A-Z])/g,"_$1").toLowerCase();
  }
export const toUpperCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toUpperCase())
    } else {
        return ""
    }
}
import { subtract } from '../util';
import { colorKeys } from './default';
export const color = {
  isValidColor(color: string): boolean {
    const hexRegExp = /^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$/;
    const rgbRegExp = /^rgb\((\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5])\)$/;
    const rgbaRegExp = /^rgba\((\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(0?\.\d|1(\.0)?)\)$/;
    return hexRegExp.test(color) || rgbRegExp.test(color) || rgbaRegExp.test(color);
  },
  convertColorToRGBA(color: string): [string, number[]] | null {

    if (!this.isValidColor(color)) {
      color = colorKeys.get(color)
    }
    // 匹配十六进制颜色值，格式为 #RRGGBB 或 #RGB
    const hexRegExp = /^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$/;
    // 匹配 RGB 颜色值，格式为 rgb(R, G, B)
    const rgbRegExp = /^rgb\((\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5])\)$/;
    // 匹配 RGBA 颜色值，格式为 rgba(R, G, B, A)
    const rgbaRegExp = /^rgba\((\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(\d{1,2}|1\d{1,2}|2[0-4]\d|25[0-5]),\s*(0?\.\d|1(\.0)?)\)$/;

    let red = 0, green = 0, blue = 0, alpha = 1;
    let rgbaString = '';
    let rgbaArray: number[] = [];

    if (hexRegExp.test(color)) {
      // 将 #RRGGBB 或 #RGB 格式的颜色值转换为 RGBA 格式和对应的 RGBA 数组
      const hexValue = color.substr(1);
      const hexLength = hexValue.length;
      red = parseInt(hexLength === 3 ? hexValue.charAt(0) + hexValue.charAt(0) : hexValue.substr(0, 2), 16);
      green = parseInt(hexLength === 3 ? hexValue.charAt(1) + hexValue.charAt(1) : hexValue.substr(2, 2), 16);
      blue = parseInt(hexLength === 3 ? hexValue.charAt(2) + hexValue.charAt(2) : hexValue.substr(4, 2), 16);
      rgbaString = `rgba(${red}, ${green}, ${blue}, ${alpha})`;
      rgbaArray = [red, green, blue, alpha];
    } else if (rgbRegExp.test(color)) {
      // 将 rgb(R, G, B) 格式的颜色值转换为 RGBA 格式和对应的 RGBA 数组
      const match = rgbRegExp.exec(color);
      if (match) {
        red = parseInt(match[1], 10);
        green = parseInt(match[2], 10);
        blue = parseInt(match[3], 10);
        rgbaString = `rgba(${red}, ${green}, ${blue}, ${alpha})`;
        rgbaArray = [red, green, blue, alpha];
      }
    } else if (rgbaRegExp.test(color)) {
      // 将 rgba(R, G, B, A)格式的颜色值转换为 RGBA 格式和对应的 RGBA 数组
      const match = rgbRegExp.exec(color);
      if (match) {
        red = parseInt(match[1], 10);
        green = parseInt(match[2], 10);
        blue = parseInt(match[3], 10);
        alpha = parseInt(match[4], 10);
        rgbaString = `rgba(${red}, ${green}, ${blue}, ${alpha})`;
        rgbaArray = [red, green, blue, alpha];
      }
    }
    return [rgbaString, rgbaArray]
  },
  /**一个rgba转换到另一个rgba,通过一个比率值0-1 */
  rgbaToRgba(rgba1: [number, number, number, number], rgba2: [number, number, number, number], blv: number = 0) {
    let result = subtract(rgba1, rgba2);
    if (result[0] <= 0) {
      result[0] = rgba2[0] - rgba1[0]
    }
    if (result[1] <= 0) {

      result[1] = rgba2[1] - rgba1[1]
    }
    if (result[2] <= 0) {
      result[2] = rgba2[2] - rgba1[2]
    }
    if (result[3] <= 0) {
      result[3] = rgba2[3] - rgba1[3]
    }

    return result
  }
}
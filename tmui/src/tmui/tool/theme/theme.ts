
/**
 * 主题工具
 * @author tmzdy tmui3.0
 * @description 主题样式生成工具
 * @copyright tmzdy|tmui|https://tmui.design
 */
import { colortool } from './colortool';
import { ssrRef } from '@dcloudio/uni-app'
import { cssStyleConfig, cssstyle, colorThemeType, cssDirection, linearDirection, linearDeep, linearDirectionType } from '../lib/interface';
//导入用户自定义的主题色值。
// import { theme } from '../../../theme/index';

let localTheme = {};
// #ifdef APP
try {
	localTheme = JSON.parse(uni.getStorageSync("$tmTheme"))
} catch (e) {
	//TODO handle the exception
}
// #endif
let theme = uni?.$tm?.config?.theme ? { ...uni.$tm.config.theme } : localTheme;

var colors: Array<colorThemeType> = [];
var colorObj: any = {
	red: '#FE1C00',
	pink: '#CA145D',
	purple: '#A61BC3',
	'deep-purple': '#6A0E81',
	indigo: '#652DF4',
	blue: '#0163FF',
	'light-blue': '#0889FF',
	cyan: '#11CDE8',
	teal: '#00998a',
	green: '#5DBD1F',
	'light-green': '#83D54A',
	lime: '#D4ED00',
	yellow: '#FFC400',
	amber: '#FFFB01',
	orange: '#FEA600',
	'deep-orange': '#FE5C00',
	brown: '#795548',
	'blue-grey': '#607D8B',
	grey: '#9E9E9E',
	black: '#000000',
	white: '#FFFFFF',
	primary: '#0163FF',
	'grey-5': '#fafafa',
	'grey-4': '#f5f5f5',
	'grey-3': '#eeeeee',
	'grey-2': '#e0e0e0',
	'grey-1': '#bdbdbd',
	'grey-darken-1': '#757575',
	'grey-darken-2': '#616161',
	'grey-darken-3': '#404044',
	'grey-darken-4': '#202022',
	'grey-darken-5': '#111112',
	'grey-darken-6': '#0A0A0B',
	...theme
};
for (const key in colorObj) {
	if (Object.prototype.hasOwnProperty.call(colorObj, key)) {
		const element: string = String(colorObj[key]);
		if (isCssColor(element)) {
			let rgba = colortool.cssToRgba(element);
			colors.push({
				name: key,
				value: element,
				hsva: colortool.rgbaToHsva(colortool.cssToRgba(element)),
				rgba: colortool.cssToRgba(element),
				hsla: colortool.rgbaToHsla(rgba),
				csscolor: `rgba(${rgba.r},${rgba.g},${rgba.b},${rgba.a})`
			});
		}
	}
}

function isCssColor(color: string) {
	const reg1 = /^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/;
	const reg2 = /^(rgb|RGB|rgba|RGBA)/;
	return reg1.test(color) || reg2.test(color);
}
function getColor(colorName: string) {
	let isHand: number = colors.findIndex(function (el, index) {
		return el.name == colorName;
	});
	if (isHand == -1) {
		colorName = "primary";
		isHand = colors.findIndex(function (el, index) {
			return el.name == colorName;
		});
		console.warn('主题中不存在相关名称的主题。');
	}


	return colors[isHand];
}
class themeColors {
	colors: Array<colorThemeType> = [];
	constructor(c: Array<colorThemeType> = colors) {
		this.colors = c;
	}
	public hasColors(colorName: string = "") {
		let isHand: Array<colorThemeType> = this.colors.filter(function (el, index) {
			return el.name == colorName;
		});
		return isHand.length > 0;
	}
	public add(colorName: string = "", value: string = "") {
		let isHand: Array<colorThemeType> = this.colors.filter(function (el, index) {
			return el.name == colorName;
		});
		if (isHand.length > 0) {
			// console.error('已存在相关颜色名称!!!');
			return this.colors;
		}
		if (!value) {
			console.error('颜色值必填!!!');
			return this.colors;
		}
		let rgba = colortool.cssToRgba(value);
		let color: colorThemeType = {
			csscolor: "",
			hsva: { h: 0, s: 0, v: 0, a: 0 },
			hsla: { h: 0, s: 0, l: 0, a: 0 },
			rgba: { r: 0, g: 0, b: 0, a: 0 },
			name: colorName, value: value
		};
		color.csscolor = colortool.rgbaToCss(rgba);
		color.hsva = colortool.rgbaToHsva(rgba);
		color.rgba = rgba;
		color.hsla = colortool.rgbaToHsla(rgba);
		this.colors.push(color);

		return this.colors;
	}
	public del(colorName: string) {
		let isHand: number = this.colors.findIndex(function (el, index) {
			return el.name == colorName;
		});
		if (isHand == -1) {
			console.error('删除失败，主题中不存在相关名称的主题。');
			return;
		}
		this.colors.splice(isHand, 1);
	}
	public getColor(colorName: string): colorThemeType {
		let isHand: number = this.colors.findIndex(function (el, index) {
			return el.name == colorName;
		});
		if (isHand == -1) {
			colorName = "primary";
			isHand = this.colors.findIndex(function (el, index) {
				return el.name == colorName;
			});
			console.error('主题中不存在相关名称的主题。');
		}


		return this.colors[isHand];
	}
	/**
	 * 计算主题
	 * @author tmui3.0|tmzdy
	 * @param config 样式的细化
	 * @returns cssstyle 返回一个计算好的主题系。
	 */
	public getTheme(config: cssStyleConfig = { colorname: 'primary', dark: false }): cssstyle {
		if (!config['colorname']) {
			console.error('颜色名称必填');
			config.colorname = 'primary';
		}
		let index = this.colors.findIndex(el => el.name == config.colorname);
		if (index == -1) {
			console.error('主题不存在，默认为primary');
			config.colorname = 'primary';
		}

		//当前颜色对象。
		let nowColor = { ...this.colors[index] };
		config.borderWidth = isNaN(parseInt(String(config['borderWidth']))) ? 0 : config['borderWidth'] ?? 0;
		config.borderStyle = config['borderStyle'] ? config['borderStyle'] : 'solid';
		config.borderColor = config['borderColor'] || '';
		config.borderDirection = config['borderDirection'] || cssDirection.all;
		config.linearDirection = config['linearDirection'] || linearDirection.none;
		config.linearDeep = config['linearDeep'] || linearDeep.light;
		config.shadow = isNaN(parseInt(String(config['shadow']))) ? 6 : config['shadow'];
		config.round = isNaN(parseInt(String(config['round']))) ? 4 : config['round'];
		config.opaticy = isNaN(parseInt(String(config['opaticy']))) ? 1 : config['opaticy'];
		config.outlined = typeof config['outlined'] == 'boolean' ? config['outlined'] : false;
		config.text = typeof config['text'] == 'boolean' ? config['text'] : false;
		config.blur = typeof config['blur'] == 'boolean' ? config['blur'] : false;
		// 确定一个颜色值是明亮的还是深色的,
		//以方便确定文本颜色是亮还是深来区别,否则颜色看不清.
		function isDarkColorFun(r: number, g: number, b: number) {
			const yiq = (r * 2126 + g * 7152 + b * 722) / 10000;
			return yiq < 180;
		}
		/**是否是黑色 */
		let isBlack = false;
		/**是否是白色 */
		let isWhite = false;
		/**黑或者白 */
		let isBlackAndWhite = false;
		/**是否是灰色 */
		let isGrey = false
		/**该颜色在人眼中属于深，还是浅，以适配文本色 */
		let isDarkColor = false;

		isDarkColor = isDarkColorFun(nowColor.rgba.r, nowColor.rgba.g, nowColor.rgba.b)
		//黑
		if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0 && nowColor.hsla.l == 0) {
			isBlack = true;
		}
		//白
		if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0 && nowColor.hsla.l == 100) {
			isWhite = true;
		}
		//灰
		if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0 && nowColor.hsla.l < 100) {
			isGrey = true;
		}
		//黑或者白
		if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0) {
			isBlackAndWhite = true;
		}

		let css: cssstyle = {};
		css.color = nowColor.value;
		css.config = { ...config };
		css.isBlackAndWhite = isBlackAndWhite;
		css.gradientColor = []
		css.colorname = config.colorname;
		let borderhsl = { ...nowColor.hsla };
		let borderDir = "all";
		css.borderCss = {};

		//背景颜色。
		let bghsl = { ...nowColor.hsla };
		/**非黑非白,h,s不变，只要降10%的亮度即可。 */
		if (config.dark && !isBlackAndWhite) {
			bghsl.l = 40;
		}
		if (config.blur) {
			bghsl.a = 0.85
		}
		css.backgroundColor = colortool.rgbaToCss(colortool.hslaToRgba({ ...bghsl }));

		if (isBlackAndWhite && config.dark) {
			css.backgroundColor = colortool.rgbaToCss(colortool.hslaToRgba({ ...bghsl, h: 240, s: 3, l: 8 }));
			css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...borderhsl, h: 240, s: 3, l: 12 }));
		}
		if (isWhite && !config.dark) {
			css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...borderhsl, l: 90 }));
		}
		if (isBlack && !config.dark) {
			css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...borderhsl, l: 12 }));
		}
		css.backgroundColorCss = { 'background-color': css.backgroundColor }
		//文字颜色。
		let txcolor = { ...nowColor.hsla };
		//当亮度小于（含）50需要降低文本颜色的亮度，即加深。，否则加亮，即变浅色。
		if (config.dark) {
			txcolor.l = 95;
		} else {
			if (isDarkColor) {
				txcolor.l = 95;
			} else {
				if (isGrey) {
					txcolor.l = 10;
				} else {
					txcolor.l = 20;
				}
			}
		}

		//外边框轮廓时
		//outlined
		if (config.outlined) {
			txcolor.l = nowColor.hsla.l;
			if (config.dark) {
				txcolor.l = 55;
			} else {
				if (nowColor.hsla.h != 0 && nowColor.hsla.s != 0 && !isDarkColorFun(nowColor.rgba.r, nowColor.rgba.g, nowColor.rgba.b)) {
					txcolor.l = 20;
				}
			}
			if ((isBlack || isWhite) && config.dark) {
				txcolor.l = 100
			}
			config.borderWidth = config['borderWidth'] || 2;
			let n_hsl = { h: nowColor.hsla.h, s: nowColor.hsla.s, l: 0, a: 0 };
			let o_bgcss = colortool.rgbaToCss(colortool.hslaToRgba(n_hsl));
			css.backgroundColor = o_bgcss;
			css.backgroundColorCss = { 'background-color': o_bgcss }
			css.textColor = colortool.rgbaToCss(colortool.hslaToRgba(txcolor));
		}

		//text
		if (config.text) {
			txcolor.l = nowColor.hsla.l;
			if (isGrey) {
				txcolor.l = 15;
			} else {
				// txcolor.l = 55;
				if (nowColor.hsla.h != 0 && nowColor.hsla.s != 0 && !isDarkColorFun(nowColor.rgba.r, nowColor.rgba.g, nowColor.rgba.b)) {
					txcolor.l = 20;
				}

			}
			if (config.dark) {
				txcolor.l = 60;
				if (!isBlackAndWhite) {
					txcolor.s = 100;
				}
			}
			if (isBlack) {
				txcolor.l = 90
			}
			if (isWhite) {
				txcolor.l = 15
			}
			if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0 && config.dark) {
				txcolor.l = 90;
			}


			css.textColor = colortool.rgbaToCss(colortool.hslaToRgba(txcolor));
			css.border = css.textColor;
			let o_now_bgColor = nowColor.csscolor;
			let n_hsl = { h: nowColor.hsla.h, s: nowColor.hsla.s, l: 96, a: nowColor.hsla.a };
			if (config.dark) {
				if (nowColor.hsla.h != 0 && nowColor.hsla.s != 0) {
					n_hsl.l = 12;
					n_hsl.s = 35;
				} else {
					n_hsl.l = 12;
					n_hsl.s = 0;
				}

			}
			if (config.blur) {
				n_hsl.a = 0.85
			}
			o_now_bgColor = colortool.rgbaToCss(colortool.hslaToRgba(n_hsl));

			css.backgroundColor = o_now_bgColor;
			css.backgroundColorCss = { 'background-color': o_now_bgColor }
		}

		//shadow
		if (config.shadow) {
			let n_hsl = { h: nowColor.hsla.h, s: 100, l: 50, a: 0.2 };
			if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0) {
				//黑白要反转。
				n_hsl = { h: 0, s: 0, l: 20, a: 0.07 };
			}
			let o_bgcss = colortool.rgbaToCss(colortool.hslaToRgba(n_hsl));
			css.shadowColor = {
				boxShadow: `0rpx ${config.shadow * 2.5}rpx ${config.shadow * 6}rpx ${o_bgcss}`
			}
		}



		//处理渐变色
		if (config.linearDirection) {

			let liner_color_1 = { h: 0, s: 0, l: 0, a: nowColor.hsla.a };
			let liner_color_2 = { h: 0, s: 0, l: 0, a: nowColor.hsla.a };
			let dir_str = linearDirection[config.linearDirection];
			// 增减控制参数。
			let addling = 0;
			if (nowColor.hsla.h < 180 && nowColor.hsla.h > 0) {
				addling = 20
			} else {
				addling = -37
			}


			//先计算渐变的亮色系。
			// 先算白或者黑
			// 如果是白
			if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0 && nowColor.hsla.l == 100) {
				//白。
				if (config.linearDeep == 'light') {
					liner_color_1.l = 80;
					liner_color_2.l = 20;
				} else {
					liner_color_1.l = 50;
					liner_color_2.l = 40;
				}
			} else if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0 && nowColor.hsla.l == 0) {
				//黑。
				if (config.linearDeep == 'light') {
					liner_color_1.l = 40;
					liner_color_2.l = 10;
				} else {
					liner_color_1.l = 30;
					liner_color_2.l = 0;
				}

			} else {
				liner_color_2.h = nowColor.hsla.h;
				liner_color_2.s = nowColor.hsla.s;

				liner_color_1.h = nowColor.hsla.h;
				liner_color_1.s = nowColor.hsla.s;
				if (config.linearDeep == 'light') {
					liner_color_1.h = liner_color_1.h;//色相需要往前偏移加强色系
					liner_color_1.s = 90;//饱和度需要加强
					liner_color_1.l = 70;
					liner_color_2.l = 44;

				} else if (config.linearDeep == 'dark') {

					liner_color_2.s = 90;
					liner_color_2.l = 26;

					liner_color_1.s = 90;
					liner_color_1.l = 50;
				} else if (config.linearDeep == 'accent') {
					liner_color_1.h -= 0;//色相需要往前偏移加强色系
					liner_color_1.s = 90;//饱和度需要加强
					liner_color_1.l = 54;

					liner_color_2.h -= addling;//偏移30度的色相搭配色进行渐变
					liner_color_2.s = 90;//饱和度需要加强
					liner_color_2.l = 54;
				}

			}
			if (config.dark) {
				liner_color_1.l = 40
				liner_color_2.l = 40
				txcolor.l = 90;
			}
			// 背景颜色取中间。
			let color_t_1 = colortool.rgbaToCss(colortool.hslaToRgba(liner_color_1));
			let color_t_2 = colortool.rgbaToCss(colortool.hslaToRgba(liner_color_2));
			if (!config.text && !config.outlined) {
				css.backgroundColorCss = { 'background-image': `linear-gradient(${dir_str},${color_t_1},${color_t_2})` }
				let newBgcolor = {
					h: (liner_color_1.h + liner_color_2.h) / 2,
					s: (liner_color_1.s + liner_color_2.s) / 2,
					l: (liner_color_1.l + liner_color_2.l) / 2,
					a: (liner_color_1.a + liner_color_2.a) / 2
				}
				let newBgcolorRgb = colortool.hslaToRgba(newBgcolor)

				if (!config.dark) {
					if (!isDarkColorFun(newBgcolorRgb.r, newBgcolorRgb.g, newBgcolorRgb.b) && nowColor.hsla.h != 0 && nowColor.hsla.s != 0) {
						txcolor.l = 20;
					}
				}
				css.backgroundColor = colortool.rgbaToCss(colortool.hslaToRgba(newBgcolor));
				css.gradientColor = [color_t_1, color_t_2]
				css.linearDirectionStr = dir_str;
			}

		}

		if (config.dark == true) {
			// css.cardcolor = '#0A0A0B'; //项目
			// css.inputcolor = '#111112';//输入框，表单等
			// css.bodycolor = 'rgba(5,5,5, 1.0)';//背景
			// css.disablecolor = 'rgba(30, 30, 30, 1.0)';//禁用的项目或者表单
			// css.textDisableColor = 'rgba(100, 100, 100, 1.0)';//文本禁用色.
			css = { ...css, ...uni.$tm.config?.themeConfig?.dark ?? {} }
		}

		css.textColor = colortool.rgbaToCss(colortool.hslaToRgba(txcolor));
		if (config.dark) {

			if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0) {
				css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...nowColor.hsla, l: 12 }));
			} else {
				css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...nowColor.hsla, l: bghsl.l + 10 }));
			}
		} else {

			if (nowColor.hsla.h == 0 && nowColor.hsla.s == 0) {
				css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...nowColor.hsla, l: 90 }));
			} else {
				// text时,使用浅色线条,outlined时与颜色相同
				if ((config.text && config.outlined)) {
					css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...nowColor.hsla, l: 90 }));
				} else if (!config.text && config.outlined) {
					css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...txcolor }));
				} else if (!config.text && !config.outlined && config.borderWidth > 0) {
					css.border = colortool.rgbaToCss(colortool.hslaToRgba({ ...nowColor.hsla, l: bghsl.l - 3 }));
				}

			}
			css.border = config.borderColor || css.border
		}

		//设置边线样式。
		let bcss = `${config.borderWidth}rpx ${config.borderStyle} ${css.border}`;
		if (config.borderDirection == 'all') {
			css.borderCss[`border`] = bcss;
		} else if (config.borderDirection == 'x' || config.borderDirection == "leftright") {
			css.borderCss[`border-left`] = bcss;
			css.borderCss[`border-right`] = bcss;
		} else if (config.borderDirection == 'y' || config.borderDirection == "topbottom") {
			css.borderCss[`border-top`] = bcss;
			css.borderCss[`border-bottom`] = bcss;
		} else if (config.borderDirection == 'bottomleft') {
			css.borderCss[`border-left`] = bcss;
			css.borderCss[`border-bottom`] = bcss;
		} else if (config.borderDirection == 'bottomright') {
			css.borderCss[`border-right`] = bcss;
			css.borderCss[`border-bottom`] = bcss;
		} else if (config.borderDirection == 'topleft') {
			css.borderCss[`border-left`] = bcss;
			css.borderCss[`border-top`] = bcss;
		} else if (config.borderDirection == 'topright') {
			css.borderCss[`border-right`] = bcss;
			css.borderCss[`border-top`] = bcss;
		} else {
			let str = '-' + config.borderDirection;
			css.borderCss[`border${str}`] = bcss;
		}

		return css;
	}
}

export default {
	isCssColor,
	themeColors,
	getColor
};

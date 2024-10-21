import { ComputedRef, Ref, computed, reactive, ref, watchEffect, getCurrentInstance, nextTick } from "vue";
import { valToMarginAr, valToRoundStrClass } from "../function/util";
import { cssstyle, tmVuetify, cssDirectionType, linearDirectionType, linearDeep } from '../lib/interface'
import colors from '../theme/theme';

/**
 * 计算tmui主题属性
 */
export default (props : ComputedRef<any>, store : ComputedRef<any>) => {
	let dark = ref(false);
	let isNvue = ref(false)
	// #ifdef APP-NVUE
	isNvue.value = true
	// #endif
	let customCSSStyle = ref({})
	let parentClass = ref("")
	let transparent = ref(false)
	let blur = ref(false)
	let customClass = ref("")
	let margin = ref<number[]>([])
	let padding = ref<number[]>([])
	let round = ref("")
	let theme = ref(computedTheme(props.value, dark.value, store.value))
	let customThemeConfig = {};
	watchEffect(() => {
		const followDark = props.value.followDark;
		let lsdark = props.value.dark;
		if (followDark) {
			lsdark = store.value.dark;
		}

		dark.value = lsdark;


		if (store.value.os == 'android' && isNvue.value) {
			blur.value = false
		} else {
			blur.value = props.value?.blur ?? false;
		}


		// 计算样式。
		customCSSStyle.value = computedStyle(props.value?._style)
		parentClass.value = props.value?.parenClass || props.value?.parentClass || "";
		transparent.value = (props.value?.transprent ?? props.value?.transparent) || false;
		customClass.value = computedClas(props.value?._class ?? "");
		margin.value = valToMarginAr(props.value?.margin ?? null)
		padding.value = valToMarginAr(props.value?.padding ?? null)
		round.value = valToRoundStrClass(props.value?.round ?? "")
		theme.value = computedTheme({ ...props.value, ...customThemeConfig }, dark.value, store.value)

		
	})



	return {
		dark,
		isNvue,
		round,
		padding,
		margin,
		customCSSStyle,
		theme: (config : any = {}) => {
			for(let key in config){
				/**如果自定的数据为null，表示为自动，不采用。如果非null就采用自定义的配置。 */
				if(config[key]!==null){
					// @ts-ignore
					customThemeConfig[key] = config[key]
				}
			}
			return theme;
		},
		customClass,
		parentClass,
		transparent,
		_props: props,
		proxy: getCurrentInstance()?.proxy ?? null,
		blur
	}

}



function computedStyle(_style : any) {
	if (typeof _style == 'string') {
		let p = _style.split(";");
		let k = p.map(el => {
			el = el.replace(";", "");
			let node : any = {};
			let idx = el.split(":");
			node[idx[0]] = idx[1];
			return node;
		})
		let kl = {};
		k.forEach(el => {
			kl = { ...kl, ...el }
		})
		return kl;
	}
	if (typeof _style == 'object' && !Array.isArray(_style)) {
		return _style;
	}
	if (typeof _style == 'object' && Array.isArray(_style)) {
		let kl = {};
		_style.forEach(el => {
			kl = { ...kl, ...el }
		})
		return kl;
	}
	return {};
}

function computedClas(_class : string | string[]) : string {
	if (typeof _class == 'string') {
		return _class
	}
	if (Array.isArray(_class)) {
		return _class.join(' ');
	}
	return '';
}


function computedTheme(props : any, dark : boolean, store : any) : cssstyle {
	const color = props.color;
	const border = props.border;
	const shadow = props.shadow;
	const round = props.round;
	const outlined = props.outlined;
	const text = props.text;
	const borderStyle = props.borderStyle;
	const borderDirection = props.borderDirection;
	const linear = props.linear;
	const linearDeep = props.linearDeep;
	const blur = props.blur;
	var borderColor = props?.borderColor ?? '';
	var theme = new colors.themeColors(store.colorList);
	if (colors.isCssColor(color) && !theme.hasColors(color)) {
		// console.error('不支持自定义组件上的颜色值，请在theme/theme.js中添加自定义的颜色值为主题。当前已切换为primary主题。');
		theme = new colors.themeColors(theme.add(color, color));
	}
	let defaultColorName = color || 'primary';

	if (props?.followTheme == true && store.color) {
		defaultColorName = store.color;
		borderColor = ""
	}

	let c = theme.getTheme({
		colorname: defaultColorName,
		dark,
		borderWidth: border,
		shadow: parseInt(String(shadow)),
		round: parseInt(String(round)),
		outlined,
		text,
		borderStyle: borderStyle,
		borderDirection: <cssDirectionType>borderDirection,
		linearDirection: <linearDirectionType>linear,
		linearDeep: <linearDeep>linearDeep,
		blur,
		borderColor: borderColor
	});

	return c;

};
<template>
	<view class="flex flex-row" :style="[{  height: `${props.height}rpx`,width: `${props.width}rpx`,}]">
		<tmSheet text class="flex-1" :transprent="props.circular" :followTheme="false" _class="flex flex-row flex-1"
			:color="props.bgColor" :margin="[0, 0]" :padding="[0, 0]">
			<view @click="setStep('-')" @longpress="longpressEvent('-')" @touchend="endlongpressEvent('-')"
				:class="[!props.circular ? `` : `round-${10}`, 'overflow', _disabled || isJianDisabled ? 'opacity-5' : '']">
				<tmSheet :followTheme="props.followTheme" :round="props.circular ? 10 : props.round"
					:linear="props.linear" :linear-deep="props.linearDeep" _class="flex-center" :color="props.color"
					:margin="[0, 0]" :padding="[0, 0]" :height="height" :width="height">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-minus"></tm-icon>
				</tmSheet>
			</view>

			<input class="mx-4" :disabled="props.disabledInput || props.disabled" @blur="inputVal" v-model="setVal"
				auto-blur :style="[
					{
						height: `${props.height}rpx`,
						textAlign: 'center',
						fontSize: props.fontSize + 'rpx',
						width: `${props.width - (props.height-4)*2}rpx`,
						color: tmcomputed.textColor,
						
					}
				]" type="digit" />

			<view @click="setStep('+')" @longpress="longpressEvent('+')" @touchend="endlongpressEvent('+')"
				:class="[!props.circular ? `` : `round-${10}`, 'overflow', _disabled || isAddDisabled ? 'opacity-5' : '']">
				<tmSheet :followTheme="props.followTheme" :round="props.circular ? 10 : props.round"
					:linear="props.linear" :linear-deep="props.linearDeep" :_class="'flex-center'" :color="props.color"
					:margin="[0, 0]" :padding="[0, 0]" :height="height" :width="height">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-plus"></tm-icon>
				</tmSheet>
			</view>
		</tmSheet>
	</view>
</template>
<script lang="ts" setup>
	/**
 * 步进器
 * @description 可以根据所需要的步骤进行增加和减少。
 */
	import { computed, Ref, ref, nextTick, watch } from 'vue'
	import { custom_props, computedTheme, computedDark } from '../../tool/lib/minxs'
	import { cssstyle, tmVuetify } from '../../tool/lib/interface'
	import tmSheet from '../tm-sheet/tm-sheet.vue'
	import tmIcon from '../tm-icon/tm-icon.vue'
	import { useTmpiniaStore } from '../../tool/lib/tmpinia'
	import { isNumber } from '@/tmui/tool/function/util'
	const store = useTmpiniaStore()
	const props = defineProps({
		...custom_props,
		//是否跟随全局主题的变换而变换
		followTheme: {
			type: [Boolean, String],
			default: true
		},
		/**
		 * 宽是input宽
		 */
		width: {
			type: [Number],
			default: 210
		},
		height: {
			type: [Number],
			default: 52
		},
		disabled: {
			type: Boolean,
			default: false
		},
		//禁用输入功能
		disabledInput: {
			type: [Boolean],
			default: false
		},
		black: {
			type: [Boolean, String],
			default: null
		},

		// 步幅，默认1
		step: {
			type: Number,
			default: 1
		},
		//固定小数点位数，0表示整数
		fixed: {
			type: Number,
			default: 0
		},
		//按钮的主题
		color: {
			type: String,
			default: 'grey-4' //grey-2
		},
		bgColor: {
			type: String,
			default: 'grey-4'
		},
		linear: {
			type: String,
			default: ''
		},
		linearDeep: {
			type: String,
			default: 'light'
		},
		round: {
			type: [String, Number],
			default: 2
		},
		fontSize: {
			type: [String, Number],
			default: 28
		},
		//圆形按钮。
		circular: {
			type: [Boolean],
			default: false
		},
		max: {
			type: [Number],
			default: 999
		},
		min: {
			type: [Number],
			default: 0
		},
		//按钮增加或者 减少前执行，返回 fase取消当前操作。
		beforeEnter: {
			type: [Function, Boolean],
			default: true
		},
		modelValue: {
			type: Number,
			default: null
		},
		defaultValue: {
			type: Number,
			default: null
		}
	})
	const emits = defineEmits(['update:modelValue', 'change'])
	const setVal : Ref<number> = ref(props.defaultValue ?? '0')
	const _min = computed(() => Number(props.min))
	const _max = computed(() => Number(props.max))
	const _step = computed(() => Number(props.step))
	const _fixed = computed(() => Number(props.fixed))

	// 设置响应式全局组件库配置表。
	const tmcfg = computed<tmVuetify>(() => store.tmStore)
	//是否暗黑模式。
	const isDark = computed(() => computedDark(props, tmcfg.value))
	//计算主题
	const tmcomputed = computed<cssstyle>(() => computedTheme({ ...props, color: props.bgColor, text: true }, isDark.value, tmcfg.value))
	const _disabled = computed(() => props.disabled)
	let timeid : any = NaN
	let timeid2 : any = NaN
	let timeid3 : any = NaN

	const isJianDisabled = computed(() => {
		if (setVal.value <= _min.value) return true
		return false
	})

	const isAddDisabled = computed(() => {
		if (setVal.value >= _max.value) return true
		return false
	})
	watch(
		() => props.modelValue,
		(newValue, oldValue) => {
			clearTimeout(timeid3)
			timeid3 = setTimeout(function () {
				let vs = forMart(props.modelValue)
				// 多此一局的做法是在微信端值为"",0,"0"之间的变化是无法被watch监测到，也就触发不了更新函数。
				if (setVal.value !== vs) {
					setVal.value = null;
					nextTick(() => {
						setVal.value = vs;
					})
				}
			}, 20);

		}
	)

	async function setStep(ty : string) {
		if (props.disabled) return
		if (typeof props.beforeEnter === 'function') {
			uni.showLoading({
				title: '...',
				mask: true
			})
			let p = await props.beforeEnter(ty)
			if (typeof p === 'function') {
				p = await p(ty)
			}
			uni.hideLoading()
			if (!p) return false
		}

		let val = setVal.value
		if (ty == "+") {
			val = _handleIncrement(setVal.value)
		} else if (ty == "-") {
			val = _handleDecrement(setVal.value)
		}
		nextTick(function () {
			jianchData(val)
		})
	}

	function inputVal(e : any) {
		var val = e.detail.value
		clearTimeout(timeid2)
		timeid2 = setTimeout(function () {
			let rs = forMart(val);
			jianchData(rs)

		}, 50)
	}

	function jianchData(val : number) {
		let vs = val.toFixed(_fixed.value)
		let vsr = _clampValue(Number(vs))
		if (setVal.value !== vsr) {
			setVal.value = null;
			nextTick(() => {
				setVal.value = vsr;
			})
		}
		emits('update:modelValue', vsr)
		emits('change', vsr)
	}
	function longpressEvent(ty : string) {
		if (props.disabled) return
		clearInterval(timeid)
		timeid = setInterval(async function () {
			await setStep(ty)
		}, 250)
	}


	/**加 */
	function _handleIncrement(value : number) {
		const newValue = Math.min(value + Number(_step.value), _max.value);
		if (_getDecimalPlaces(value)) {
			return _clampValue(parseFloat(newValue.toFixed(_fixed.value)));
		}
		return _clampValue(newValue);
	}
	/**减 */
	function _handleDecrement(value : number) {
		const newValue = Math.max(value - Number(_step.value), _min.value);
		if (_getDecimalPlaces(value)) {
			return _clampValue(parseFloat(newValue.toFixed(_fixed.value)));
		}
		return _clampValue(newValue);
	}


	function _isInRange(value : number) : boolean {
		return value >= _min.value && value <= _max.value;
	}
	function _clampValue(value : number) : number {
		return Math.min(Math.max(value, _min.value), _max.value);
	}

	//返回正确的小数点值。
	function _getDecimalPlaces(value : number | string) : number {
		const match = value.toString().match(/(?:\.(\d+))?(?:[eE]([+-]?\d+))?$/);
		if (!match) {
			return 0;
		}
		return Math.max(0, (match[1] ? match[1].length : 0) - (match[2] ? +match[2] : 0));
	}

	function endlongpressEvent(ty : string) {
		clearInterval(timeid)
	}

	function forMart(val : string | number) : number {
		val = parseFloat(typeof val === 'number' ? String(val) : val)
		if (!isNaN(val) && _isInRange(val)) {

			if (_getDecimalPlaces(val)) {
				val = Number(val.toFixed(_fixed.value))
			}
			return _clampValue(val);
		} else {
			let value = _clampValue(isNaN(val) ? 0 : val);
			return _clampValue(value);
		}
	}
</script>
<style scoped>
	.inputInit {}
</style>
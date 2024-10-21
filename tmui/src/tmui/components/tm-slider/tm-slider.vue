<template>
	<view class="flex" :class="[props.direction == 'vertical' ? 'flex-row' : 'flex-col', _disabled ? 'opacity-6' : '']">
		<view
			class="relative flex flex-col"
			:class="[props.direction == 'horizontal' ? 'flex-col-center-start' : 'flex-row-center-center']"
			:style="[{ width: _sliderBarCssWidth + 'rpx', height: _sliderBarCssHeight + 'rpx' }]"
		>
			<tm-sheet
				noLevel
				:round="10"
				:color="props.bgColor"
				:height="props.direction == 'horizontal' ? props.height : _sliderBarCssHeight"
				:width="props.direction == 'horizontal' ? _sliderBarCssWidth : props.height"
				:margin="[0, 0]"
				:padding="[0, 0]"
			></tm-sheet>
			<slider-bar
				:followTheme="props.followTheme"
				:class="[props.direction == 'horizontal' ? 'flex-col-center-start' : 'flex-row-center-center']"
				:direction="props.direction"
				:color="props.color"
				:size="props.height"
				:x="_barLet"
				:width="_barWidth"
			></slider-bar>
			<slider-mask
				:followTheme="props.followTheme"
				:color="props.color"
				v-if="props.showLabel"
				:size="props.buttonSize"
				:step="props.step"
				:min="props.min"
				:max="props.max"
				:width="_sliderBarCssWidth"
				:height="_sliderBarCssHeight"
				:direction="props.direction"
			></slider-mask>
			<slider-button
				:maxLeft="buttonStaticsMaxLeft"
				:followTheme="props.followTheme"
				:direction="props.direction"
				ref="btn0"
				:color="props.color"
				:x="btnPos[0].x"
				@movestart="butnMoveStart($event, 0)"
				@moveing="butnMove($event, 0)"
				@moveend="butnMoveEnd($event, 0)"
				:size="props.buttonSize"
			></slider-button>
			<slider-button
				:maxLeft="buttonStaticsMaxLeft"
				:followTheme="props.followTheme"
				:direction="props.direction"
				v-if="isDablue"
				ref="btn1"
				:color="props.color"
				:x="btnPos[1].x"
				@movestart="butnMoveStart($event, 1)"
				@moveing="butnMove($event, 1)"
				@moveend="butnMoveEnd($event, 1)"
				:size="props.buttonSize"
			></slider-button>
		</view>
		<view v-if="props.showLabel" :class="[props.direction == 'vertical' ? 'flex-col' : 'flex-row']">
			<slider-label
				:size="props.buttonSize"
				:step="props.step"
				:min="props.min"
				:max="props.max"
				:width="_sliderBarCssWidth"
				:height="_sliderBarCssHeight"
				:direction="props.direction"
			></slider-label>
		</view>
		<view
			v-if="showDetail"
			class="flex absolute"
			:class="[props.direction == 'horizontal' ? 'flex-col flex-col-start-center' : ' flex-row ']"
			:style="[
				props.direction == 'horizontal'
					? { width: props.width + props.buttonSize + 'rpx' }
					: { height: props.width + props.buttonSize + 'rpx' }
			]"
		>
			<view
				:class="[isNvue ? 'fixed' : 'absolute ', props.direction == 'horizontal' ? (isNvue ? 't-0' : 'b-0') : 't-0 ']"
				class="mb-0 flex flex-col flex-col-bottom-center"
				:style="[
					!isNvue
						? props.direction == 'horizontal'
							? {
									transform: `translateX(${btnPos[BtnIndex].x}px)`,
									left: -(100 - props.buttonSize + 24) / 2 + 'rpx'
							  }
							: {
									transform: `translateY(${btnPos[BtnIndex].x}px)`,
									top: '-70rpx',
									left: -(100 - props.buttonSize + 24) / 2 + 'rpx'
							  }
						: { left: `${nvueDetailPos.left}px`, top: `${nvueDetailPos.bottom}px` }
				]"
			>
				<tm-sheet :followDark="false" :dark="false" _class="flex-center" color="grey-darken-5" :border="2" :margin="[0, 0]" :padding="[10, 6]" :width="100" :round="3">
					<tm-text :label="_value"></tm-text>
				</tm-sheet>
				<tm-icon :followDark="false" :dark="false" color="grey-darken-5" _class="t--10" :font-size="32" name="tmicon-sort-down-nogap-copy"></tm-icon>
			</view>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 滑块
 * description 注意，如果想开启单滑块（默认），default-value=0单个值即可，如果想双滑块，值等于数组比如:=[0,5]
 */
import { computed, ref, Ref, toRaw, getCurrentInstance, watchEffect, watch, inject,PropType } from 'vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
import sliderBar from './slider-bar.vue'
import sliderButton from './slider-button.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import TmIcon from '../tm-icon/tm-icon.vue'
import sliderLabel from './slider-label.vue'
import sliderMask from './slider-mask.vue'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const btn0 = ref<InstanceType<typeof sliderButton> | null>(null)
const btn1 = ref<InstanceType<typeof sliderButton> | null>(null)

const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['update:modelValue', 'change'])
const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	//竖向时这个是高度。
	width: {
		type: Number,
		default: 500
	},
	// 条的高度
	//竖向时，这个是宽度。
	height: {
		type: Number,
		default: 8
	},
	//按钮的大小。
	buttonSize: {
		type: Number,
		default: 52
	},
	/**
	 * 方向
	 * horizontal:水平,
	 * vertical:竖向。
	 */
	direction: {
		type: String,
		default: 'horizontal'
	},
	color: {
		type: String,
		default: 'primary'
	},
	bgColor: {
		type: String,
		default: 'grey-3'
	},
	max: {
		type: Number,
		default: 100
	},
	min: {
		type: Number,
		default: 0
	},
	//默认的值一定要和v-model相同的类型，
	defaultValue: {
		type: [Array, Number] as PropType<number[]|number>,
		default: 0
	},

	//双向绑定数据可以使用v-model,记住格式一定要和defaultValue相同。
	modelValue: {
		type: [Array, Number],
		default: 0
	},
	//格式为自己的标签显示 ，是按钮上方的显示 标签。
	formart: {
		type: Function,
		default: () => {
			return (val: number | string) => {
				return val
			}
		}
	},
	//只有这个为true,下面的step才会显示标签。
	showLabel: {
		type: Boolean,
		default: false
	},
	//需要显示的步骤标签。
	step: {
		type: Number,
		default: 5
	},
	disabled: {
		type: Boolean,
		default: false
	}
})

interface btnMovetype {
	x: number
	y: number
}
const width = computed(() => uni.upx2px(props.width))
const isDarg = ref(false)
const _disabled = computed(() => props.disabled)
const btnPos: Ref<Array<btnMovetype>> = ref([
	{
		x: 0,
		y: 0
	},
	{
		x: 0,
		y: 0
	}
])
let _x = 0
let _y = 0
let buttonStaticsLeft = 0
let buttonStaticsMaxLeft = width.value
const isNvue = ref(false)
const nvueDetailPos = ref({ left: 0, bottom: 0 })
// #ifdef APP-NVUE
isNvue.value = true
// #endif

const _sliderBarCssWidth = computed(() => {
	if (props.direction == 'horizontal') return props.width + props.buttonSize
	return props.buttonSize
})
const _sliderBarCssHeight = computed(() => {
	if (props.direction == 'horizontal') return props.buttonSize
	return props.width + props.buttonSize
})

//当前操作的是哪个按钮。
const BtnIndex = ref(0)
//显示上访值。
const showDetail = ref(false)
const _valueMax = computed(() => {
	return props.max - props.min
})
const _barWidth = computed(() => {
	return Math.abs(btnPos.value[0].x - btnPos.value[1].x)
})
const _barLet = computed(() => {
	return Math.min(Math.abs(btnPos.value[0].x), Math.abs(btnPos.value[1].x))
})

const _value = ref(0)
//是双滑块，还是单滑块。
const isDablue = ref(Array.isArray(props.defaultValue)?true:false)

zhuanghuaValue()
watchEffect(() => {
	let val = Math.ceil((Math.abs(btnPos.value[BtnIndex.value].x) / uni.upx2px(props.width)) * _valueMax.value + props.min)
	if (typeof props.formart === 'function') {
		let p = props.formart(val)
		if (typeof p === 'function') {
			p = p(val)
		}
		val = p
	}
	_value.value = val
})

const _blackValue = getValue()
watch(
	() => props.modelValue,
	() => {
		zhuanghuaValue(props.modelValue)
	}
)
/**转换值 */
function zhuanghuaValue(defaultValue:number|number[] = props.defaultValue) {
	let maxWidth  = uni.upx2px(props.width);
	if (!Array.isArray(defaultValue)) {
		let vsp = Number(defaultValue)
		vsp = Math.min(props.max,Math.max(props.min,vsp))
		btnPos.value[0].x =  mapToPercentage(vsp) * maxWidth
	} else {
		let vsp_0 = Number(defaultValue[0])
		vsp_0 = Math.max(props.min,Math.min(props.max,vsp_0))
		btnPos.value[0].x = mapToPercentage(vsp_0) * maxWidth
		let vsp_1 = Number(defaultValue[1])
		vsp_1 = Math.min(Math.max(vsp_1,props.min),props.max);
		btnPos.value[1].x = mapToPercentage(vsp_1) * maxWidth
	}
}

function mapToPercentage(value:number) {
  var minValue = props.min;
  var maxValue = props.max;
  var percentage = (value - minValue) / (maxValue - minValue);
  return percentage;
}

function getValue() {
	if (!isDablue.value) {
		return Math.ceil((Math.abs(btnPos.value[0].x) / uni.upx2px(props.width)) * _valueMax.value + props.min)
	} else {
		return [
			Math.ceil((Math.abs(btnPos.value[0].x) / uni.upx2px(props.width)) * _valueMax.value + props.min),
			Math.ceil((Math.abs(btnPos.value[1].x) / uni.upx2px(props.width)) * _valueMax.value + props.min)
		]
	}
}

function butnMoveStart(e: btnMovetype, index: number) {
	if (props.disabled) return
	isDarg.value = true

	if (props.direction == 'horizontal') {
		_x = e.x - btnPos.value[index].x
	} else {
		_x = e.y - btnPos.value[index].x
	}

	BtnIndex.value = index
}
function butnMove(e: btnMovetype, index: number) {
	if (props.disabled) return
	if (!isDarg.value) return
	let left = e.x - _x
	if (props.direction != 'horizontal') {
		left = e.y - _x
	}
	if (left < buttonStaticsLeft) {
		left = buttonStaticsLeft
	} else if (left > buttonStaticsMaxLeft) {
		left = buttonStaticsMaxLeft
	}

	btnPos.value[index].x = left
	showDetail.value = true
	getDomRectBound()
}
function butnMoveEnd(e: btnMovetype, index: number) {
	if (props.disabled) return
	isDarg.value = false
	showDetail.value = false
	emits('update:modelValue', getValue())
	emits('change', getValue())
	pushFormItem()
}

function getDomRectBound() {
	// #ifdef APP-NVUE
	dom.getComponentRect(proxy?.$refs['btn' + BtnIndex.value], function (res) {
		if (res?.size) {
			// { ...res.size }
			// console.log(res.size)
			const { left, top } = res.size
			nvueDetailPos.value = {
				left: left - (uni.upx2px(100) - uni.upx2px(props.buttonSize)) / 2,
				bottom: top - 45
			}
		}
	})
	// #endif
}

/** -----------form专有------------ */
//父级方法。
const rulesObj = inject(
	'tmFormItemRules',
	computed<Array<rulesItem>>(() => {
		return [
			{
				message: '请选择',
				required: false,
				validator: false
			}
		]
	})
)
//父级方法。
let parentFormItem: any = proxy?.$parent
while (parentFormItem) {
	if (parentFormItem?.tmFormComnameFormItem == 'tmFormComnameFormItem' || !parentFormItem) {
		break
	} else {
		parentFormItem = parentFormItem?.$parent ?? undefined
	}
}
const validate = (rules: Array<rulesItem>) => {
	rules = rules.map((el) => {
		if (typeof el.validator === 'function' && el.required === true) {
			return el
		} else if (typeof el.validator === 'boolean' && el.required === true) {
			return {
				...el,
				validator: (val: string | number | Array<string | number>) => {
					if (Array.isArray(val)) {
						return val.reduce((a, b) => Number(a) + Number(b)) == 0
					} else {
						return Number(val) == 0
					}
				}
			}
		} else {
			return {
				...el,
				validator: (val: string | number) => {
					return true
				}
			}
		}
	})
	let rules_filter: Array<rulesItem> = rules.filter((el) => {
		return typeof el.validator === 'function' && el.required === true
	})
	let _valueSlider = getValue()
	let rules_fun: Array<Promise<rulesItem>> = rules_filter.map((el) => {
		return new Promise(async (res, rej) => {
			if (typeof el.validator === 'function') {
				let vr = await el.validator(_valueSlider)
				if (vr) {
					res({
						message: String(el.message),
						validator: true
					})
				} else {
					rej({
						message: el.message,
						validator: false
					})
				}
			} else {
				res({
					message: el.message,
					validator: true
				})
			}
		})
	})
	return Promise.all(rules_fun)
}

async function pushFormItem(isCheckVail = true) {
	if (parentFormItem) {
		if (isCheckVail) {
			let _valueSlider = getValue()
			validate(toRaw(rulesObj.value))
				.then((ev) => {
					parentFormItem.pushCom({
						value: _valueSlider,
						isRequiredError: false, //true,错误，false正常 检验状态
						componentsName: 'tm-slider', //表单组件类型。
						message: ev.length == 0 ? '' : ev[0].message //检验信息提示语。
					})
				})
				.catch((er) => {
					parentFormItem.pushCom({
						value: _valueSlider,
						isRequiredError: true, //true,错误，false正常 检验状态
						componentsName: 'tm-slider', //表单组件类型。
						message: er.message //检验信息提示语。
					})
				})
		}
	}
}
pushFormItem()

const tmFormFun = inject(
	'tmFormFun',
	computed(() => '')
)
watch(tmFormFun, () => {
	if (tmFormFun.value == 'reset') {
		emits('update:modelValue', _blackValue)
		pushFormItem(false)
	}
})

/** -----------end------------ */
</script>
<style scoped></style>

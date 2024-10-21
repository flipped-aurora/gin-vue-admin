<template>
	<view class="flex flex-row flex-row-center-start" @click.stop="">
		<view v-for="(item, index) in _count" :key="item" :class="[`pr-${gutter}`, props.disabled ? 'opacity-6' : '']">
			<view :class="[ani && index == _start - 1 ? 'rateOn' : '']">
				<tm-icon
					ref="nvueElAni"
					:follow-dark="false"
					:color="item <= _start ? _color : 'grey-2'"
					@click="startClick(item)"
					:font-size="props.size"
					:name="props.icon"
				></tm-icon>
			</view>
		</view>

		<slot><tm-text :dark="isDark" v-if="showLabel" :color="_color" :label="_label"></tm-text></slot>
	</view>
</template>
<script lang="ts" setup>
import { computed, ref, watch, getCurrentInstance, inject, toRaw, nextTick } from 'vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
import { custom_props, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import TmText from '../tm-text/tm-text.vue'
// #ifdef APP-PLUS-NVUE
const Binding = uni.requireNativePlugin('bindingx')
const dom = uni.requireNativePlugin('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
const store = useTmpiniaStore()
const emits = defineEmits(['click', 'change', 'update:modelValue'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	count: {
		type: Number,
		default: 5
	},
	modelValue: {
		type: Number,
		default: 0
	},
	defaultValue: {
		type: Number,
		default: 0
	},
	//只读式，样式无变化，可以触发点击事件，但同样无法切换数值。
	readonly: {
		type: Boolean,
		default: false
	},
	//禁用后无法点击和切换值
	disabled: {
		type: Boolean,
		default: false
	},
	icon: {
		type: String,
		default: 'tmicon-collection-fill'
	},
	size: {
		type: Number,
		default: 42
	},
	//可以是数据也可以是单独主题名称。如果出现多个主题名称时。颜色将会对应count出现。比如1星是红，2星是蓝。依此类推。
	color: {
		type: [Array, String],
		default: 'orange'
	},
	//图标之间的间距
	gutter: {
		type: Number,
		default: 16
	},
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	//暗黑
	dark: {
		type: [Boolean, String],
		default: false
	},
	//是否跟随主题全局切换暗黑模式。
	followDark: {
		type: [Boolean, String],
		default: true
	},
	//需要展示在右边的分值，默认为空即显示星数值。如果提供了其它值，就不显示默认的，比如9.6分
	label: {
		type: String,
		default: ''
	},
	showLabel: {
		type: Boolean,
		default: false
	}
})
let timid: any = NaN
const _count = computed(() => props.count)
const _start = ref(props.defaultValue)
// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore)
const isDark = computed(() => computedDark(props, tmcfg.value))
const _color = computed(() => {
	if (props.followTheme && tmcfg.value.color) return tmcfg.value.color
	if (typeof props.color == 'string') return props.color
	if (Array.isArray(props.color)) {
		if (props.color[_start.value - 1]) {
			return props.color[_start.value - 1]
		}
		return props.color[props.color.length - 1]
	}
	return 'grey-2'
})
const ani = ref(false)
const _label = computed(() => {
	if (props.label != '') return props.label
	return _start.value + '.0'
})

watch(
	() => props.modelValue,
	() => {
		let valueStart = props.modelValue >= _count.value ? _count.value : props.modelValue
		_start.value = valueStart <= 0 ? 0 : valueStart
	}
)
function startClick(index: number) {
	if (props.disabled) return
	if (props.readonly) {
		emits('click', index - 1)
		return
	}
	_start.value = index
	emits('change', _start.value)
	emits('update:modelValue', _start.value)
	emits('click', index - 1)
	ani.value = false
	clearTimeout(timid)
	timid = setTimeout(() => {
		ani.value = true
		// #ifdef APP-NVUE
		animationPlayNvue()
		// #endif
	}, 50)
}

function animationPlayNvue(s = 0.8, isend = false) {
	var el = proxy?.$refs.nvueElAni[_start.value - 1]
	if (!el) return
	animation.transition(
		el,
		{
			styles: {
				transform: `scale(${s},${s})`,
				transformOrigin: 'center center'
			},
			duration: 300, //ms
			timingFunction: 'ease',
			delay: 0 //ms
		},
		() => {
			if (isend) return
			if (s == 0.8) {
				animationPlayNvue(1.04, false)
			} else {
				animationPlayNvue(1, true)
			}
		}
	)
}
</script>
<style scoped>
/* #ifndef APP-NVUE */
.rateOn {
	transition-timing-function: ease-in;
	transition-property: transform;
	transition-duration: 0.3s;
	transition-delay: 0s;
	animation: scaleAni 0.3s ease-in;
}
/* #endif */
/* #ifndef APP-NVUE */
@keyframes scaleAni {
	0% {
		transform: scale(0.8, 0.8);
	}

	50% {
		transform: scale(1.1, 1.1);
	}
	100% {
		transform: scale(1, 1);
	}
}

/* #endif */
</style>

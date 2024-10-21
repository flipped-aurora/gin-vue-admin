<template>
	<view
		class="flex flex-shrink"
		
		:class="[_disabled ? props.disabledClass : '', tmCheckedBoxDir == 'row' ? 'flex-row' : '', 
		tmCheckedBoxDir == 'customCol' ? 'flex-1 flex-col' : '']"
	>
		<view @click.stop="hanlerClick" class="" :class="[
			(tmCheckedBoxDir == 'customCol'&&props.custom)||tmCheckedBoxDir == 'row' ? 'flex flex-1 flex-row flex-row-center-start' : ''
			]">
			<tm-sheet
				parenClass="flex-shrink"
				class="flex-shrink"
				v-if="props.custom"
				:eventPenetrationEnabled="true"
				:linear="props.linear"
				:linearDeep="props.linearDeep"
				:followTheme="props.followTheme"
				:followDark="props.followDark"
				:dark="props.dark"
				:shadow="props.shadow"
				:userInteractionEnabled="false"
				:width="_is_radio ? props.size : 0"
				:height="props.size"
				:text="!_checked && !props.outlined"
				:border="props.border"
				:border-style="props.borderStyle"
				:transprent="props.transprent"
				:outlined="props.outlined"
				:padding="_is_radio ? [0, 0] : [16, 10]"
				:margin="_is_radio ? props.margin : [8, 8]"
				:color="_disabled ? 'grey-2' : props.color"
				:round="props.round"
				_class="flex-row flex-row-center-center"
				_style="transition:background-color 0.24s"
			>
				<view
					v-if="_checked && _is_radio"
					:style="{
						width: props.size - props.border * 2 + 'rpx',
						height: props.size - props.border * 2 + 'rpx'
					}"
					class="flex flex-row flex-row-center-center relative t--0"
				>
					<tm-icon :font-size="props.size * 0.6" :line-height="-1" :name="props.icon"></tm-icon>
				</view>
				
				<tm-text v-if="!_is_radio" :line-height="1" :font-size="props.fontSize" :label="props.label"></tm-text>
			</tm-sheet>
			<slot name="default" :checked="{ checked: _checked }">
				<tm-text
					:userInteractionEnabled="false"
					class="flex-1 flex-wrap"
					
					v-if="_is_radio"
					:font-size="props.fontSize"
					:label="props.label"
				></tm-text>
			</slot>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 单选框
 * @description 单选框不可单独使用，必须配合tm-radio-group，放置在tm-radio-group中使用。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import { custom_props } from '../../tool/lib/minxs'
import { ref, computed, watch, inject, getCurrentInstance, nextTick, PropType } from 'vue'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['update:modelValue', 'change', 'click'])
const props = defineProps({
	...custom_props,
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	size: {
		type: Number,
		default: 42
	},
	outlined: {
		type: Boolean,
		default: false
	},
	transprent: {
		type: Boolean,
		default: false
	},
	//为false时将隐藏所有内容，只显示插槽内容，但点击插槽还是会触发选选择状态。
	custom: {
		type: Boolean,
		default: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	round: {
		type: Number,
		default: 24
	},
	border: {
		type: Number,
		default: 2
	},
	//选项值，选中后返回的值。
	value: {
		type: [String, Number, Boolean,null],
		default: ''
	},
	/**
	 * v-model双向绑定，如果选中后以数组形式给出value值。
	 */
	modelValue: {
		type: [String, Number, Boolean,null],
		default: ''
	},
	label: {
		type: [String, Number],
		default: ''
	},
	//默认是否选中状态。不受上方的modelValue控制，直接选中。
	defaultChecked: {
		type: [Boolean],
		default: false
	},
	//选中前的勾子。返回false将阻止选中。也可以返回 Promise异步
	beforeChecked: {
		type: [Function, String, Boolean],
		default: () => {
			return false
		}
	},
	disabled: {
		type: Boolean,
		default: false
	},
	fontSize: {
		type: Number,
		default: 26
	},
	/**
	 * 自定义选中的图标名称
	 */
	icon: {
		type: String,
		default: 'tmicon-check'
	},
	disabledClass: {
		type: String,
		default: 'opacity-5'
	},
	margin: {
		type: Array as PropType<number[]>,
		default: () => [16, 8]
	}
})
const is_nvue = ref(false)
// #ifdef APP-NVUE
is_nvue.value = true
// #endif
const _checked = ref(props.defaultChecked ?? false)
const _groupCheckedVal = inject(
	'tmRadioBoxVal',
	computed(() => '')
)
const tmCheckedBoxDisabled = inject(
	'tmRadioBoxDisabled',
	computed(() => false)
)
const _is_radio = inject(
	'tmRadioBoxModel',
	computed(() => false)
)
const tmCheckedBoxDir = inject(
	'tmCheckedBoxDir',
	computed(() => 'row')
)
let timed: any = NaN
const _disabled = computed(() => props.disabled || tmCheckedBoxDisabled.value)
function vailChecked(val?: string | number | boolean) {
	let checked_val = false
	let val_self = typeof val === 'undefined' ? _groupCheckedVal.value : val
	if (props.modelValue === props.value && typeof props.value !== 'undefined' && props.value !== '' && props.modelValue !== '') {
		checked_val = true
	}
	if (props.value === val_self && val_self !== '' && props.value !== '') {
		checked_val = true
	}
	return checked_val
}
if (vailChecked()) {
	_checked.value = true
	emits('update:modelValue', props.value)
}
async function hanlerClick() {
	emits('click')
	// console.log(new Date().getSeconds(),'-----1')
	if (_disabled.value || _checked.value) {
		return
	}

	if (typeof props.beforeChecked === 'function') {
		uni.showLoading({
			title: '...',
			mask: true
		})
		let p = await props.beforeChecked(props.modelValue,props.value)
		if (typeof p === 'function') {
			p = await p(props.modelValue,props.value)
		}
		uni.hideLoading()
		if (!p) return
	}
	_checked.value = true
	if (parent) {
		parent.addKey(props.value)
	}
	emits('update:modelValue', props.value)
	// console.log(new Date().getSeconds(),'-----2')

	nextTick(() => {
		emits('change', _checked.value)
	})
}
watch(
	[() => props.modelValue, () => props.value, () => _groupCheckedVal.value],
	() => {
		_checked.value = vailChecked()
	},
	{ deep: true }
)
//父级方法。
let parent: any = proxy?.$parent

while (parent) {
	if (parent?.checkBoxkeyId == 'tmRadioBoxGroup' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
if (parent) {
	parent.pushKey(props.value)
}
</script>

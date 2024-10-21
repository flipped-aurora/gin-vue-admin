<template>
	<view
		class="flex"
		:class="[_disabled ? props.disabledClass : '', tmCheckedBoxDir == 'row' ? 'flex-row' : '', 
		tmCheckedBoxDir == 'customCol' ? 'flex-1 flex-col' : ''
		]"
	>
		<view @click.stop="hanlerClick" class="" :class="[
			(tmCheckedBoxDir == 'customCol'&&props.custom)||tmCheckedBoxDir == 'row' ? 'flex flex-1 flex-row flex-row-center-start' : ''
			]">
			
			<tm-sheet
				parenClass="flex-shrink"
				class="flex-shrink"
				v-if="props.custom"
				:eventPenetrationEnabled="true"
				:userInteractionEnabled="false"
				:linear="props.linear"
				:linearDeep="props.linearDeep"
				:followTheme="props.followTheme"
				:followDark="props.followDark"
				:dark="props.dark"
				:shadow="props.shadow"
				:width="props.size"
				:height="props.size"
				:text="((!props.indeterminate && !_checked) || _disabled) && !props.outlined"
				:border="props.border"
				:borderStyle="props.borderStyle"
				:transprent="props.transprent"
				:padding="[0, 0]"
				:margin="props.margin"
				:color="_disabled ? 'white' : props.color"
				:round="props.round"
				_class="flex-row flex-row-center-center"
				:outlined="props.outlined"
				_style="transition:background-color 0.24s"
			>
				<template v-if="!props.closeAni">
					<tm-translate :duration="100" v-if="_checked && !props.indeterminate" name="zoom" style="line-height: 1">
						<tm-icon :font-size="props.size * 0.6" :name="props.icon"></tm-icon>
					</tm-translate>
					<tm-translate v-if="props.indeterminate" :duration="100" name="zoom" style="line-height: 1">
						<tm-icon :font-size="props.size * 0.6" name="tmicon-minus"></tm-icon>
					</tm-translate>
				</template>
				<template v-if="props.closeAni">
					<tm-icon v-if="_checked && !props.indeterminate" :font-size="props.size * 0.6" :name="props.icon"></tm-icon>
					<tm-icon v-if="props.indeterminate" :font-size="props.size * 0.6" name="tmicon-minus"></tm-icon>
				</template>
			</tm-sheet>
			<slot name="default" :checked="{ checked: _checked }" style="display: flex;flex-direction: column;">
				<tm-text :userInteractionEnabled="false" class="flex-1 flex-wrap" :font-size="props.fontSize" :label="props.label"></tm-text>
			</slot>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 复选框
 * @description 复选框可以单独使用，也可以配合复选框组tm-checkbox-group使用。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import tmCheckboxGroup from '../tm-checkbox-group/tm-checkbox-group.vue'
import { custom_props } from '../../tool/lib/minxs'
import {
	ref,
	computed,
	watch,
	inject,
	getCurrentInstance,
	PropType,
	ComponentInternalInstance,
	ComputedRef,
	onMounted,
	onUnmounted,
	onBeforeUnmount
} from 'vue'
const CheckboxGropup = ref<InstanceType<typeof tmCheckboxGroup> | null>(null)
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['update:modelValue', 'change', 'click'])
const props = defineProps({
	...custom_props,
	/**
	 * 是否跟随全局主题的变换而变换
	 */
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	size: {
		type: Number,
		default: 42
	},
	//为false时将隐藏所有内容，只显示插槽内容，但点击插槽还是会触发选选择状态。
	custom: {
		type: Boolean,
		default: true
	},
	outlined: {
		type: Boolean,
		default: false
	},
	transprent: {
		type: Boolean,
		default: false
	},
	round: {
		type: Number,
		default: 2
	},
	border: {
		type: Number,
		default: 2
	},
	//选项值，选中后返回的值。
	value: {
		type: [String, Number, Boolean],
		default: true
	},
	/**
	 * v-model双向绑定，如果选中后以数组形式给出value值。
	 */
	modelValue: {
		type: [String, Number, Boolean],
		default: false
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
		default: 28
	},
	//半选中状态。
	indeterminate: {
		type: [Boolean, String],
		default: false
	},
	//是否关闭动画 ，对于大批量的数据时，建议关闭动画。
	closeAni: {
		type: [Boolean, String],
		default: true
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
let timed: any = NaN
const _checked = ref(props.defaultChecked ?? false)
const _groupCheckedVal: ComputedRef<Array<string | number | boolean>> = inject(
	'tmCheckedBoxVal',
	computed(() => [])
)
const tmCheckedBoxDisabled = inject(
	'tmCheckedBoxDisabled',
	computed(() => false)
)
const tmCheckedBoxMax = inject(
	'tmCheckedBoxMax',
	computed(() => false)
)
const tmCheckedBoxDir = inject(
	'tmCheckedBoxDir',
	computed(() => 'row')
)
const _disabled = computed(() => props.disabled || tmCheckedBoxDisabled.value)

function vailChecked(val?: Array<string | number | boolean>) {
	let checked_val = false
	let val_self: Array<string | number | boolean> = typeof val === 'undefined' ? _groupCheckedVal.value : val

	if (props.modelValue === props.value && typeof props.value !== 'undefined' && props.value !== '' && props.modelValue !== '') {
		checked_val = true
	}
	// let index = val_self.findIndex((el) => el === props.value);
	if (val_self.includes(props.value)) {
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
	if (_disabled.value) {
		return
	}
	if (tmCheckedBoxMax.value && !_checked.value) {
		uni.showToast({ title: '超最大选择', icon: 'error' })
		return
	}
	if (typeof props.beforeChecked === 'function') {
		uni.showLoading({ title: '...', mask: true })
		let p = await props.beforeChecked(props.modelValue,props.value)
		if (typeof p === 'function') {
			p = await p(props.modelValue,props.value)
		}
		uni.hideLoading()
		if (!p) return
	}
	_checked.value = !_checked.value
	if (_checked.value) {
		emits('update:modelValue', props.value)
		if (parent) {
			parent.addKey(props.value)
		}
	} else {
		emits('update:modelValue', false)
		if (parent) {
			parent.delKey(props.value)
		}
	}
	emits('change', _checked.value)
}
watch(
	[() => props.modelValue, () => props.value, () => _groupCheckedVal.value],
	() => {
		_checked.value = vailChecked()
	},
	{ deep: true }
)

const _blackValue = _groupCheckedVal.value
//父级方法。
let parent: any = proxy?.$parent
while (parent) {
	if (parent?.checkBoxkeyId == 'tmCheckBoxGroup' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
if (parent) {
	parent.pushKey(props.value)
}

defineExpose({ hanlerClick })
</script>

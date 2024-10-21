<template>
	<view
		:class="['relative flex-1', 'flex flex-col flex-col-center-center overflow', _cureent < _countCurrent ? 'flex-1' : '']"
		:style="{ width: _width }"
	>
		<!-- 线 -->
		<view class="flex-1 flex-row flex-row-center-between absolute" :style="{ width: _width, top: nodeSize / 2 + 'rpx' }">
			<view :style="{ width: _widthLine }">
				<tm-divider :border="2" v-if="_cureent > 0" :color="_color" :followDark="props.followDark" :dark="props.dark" :margin="[0, 0]">
				</tm-divider>
			</view>
			<view :style="{ width: _widthLine }">
				<tm-divider
					:border="2"
					v-if="_cureent < _countCurrent"
					:color="_color"
					:followDark="props.followDark"
					:dark="props.dark"
					:margin="[0, 0]"
				>
				</tm-divider>
			</view>
		</view>
		<!-- 内容 区域。 -->
		<view @click.stop="stepClick" class="flex flex-col relative" style="justify-content: flex-start; align-items: center">
			<tm-sheet
				:eventPenetrationEnabled="true"
				no-level
				:followTheme="props.followTheme"
				:followDark="props.followDark"
				:dark="props.dark"
				:shadow="props.shadow"
				:outlined="props.outlined"
				:borderStyle="props.borderStyle"
				:borderDirection="props.borderDirection"
				:text="_isCheck ? false : !status"
				:linearDeep="props.linearDeep"
				:linear="props.linear"
				:color="_color"
				:round="24"
				_class="flex-center"
				:margin="[0, 0]"
				:padding="[0, 0]"
				:width="nodeSize"
				:height="nodeSize"
			>
				<tm-text
					v-if="!_icon && !status && _typeModel != 'dot'"
					:userInteractionEnabled="false"
					:font-size="22"
					:label="_cureent + 1"
				></tm-text>
				<tm-icon :font-size="22" v-if="status && _isActive && _typeModel != 'dot'" :name="status.icon"></tm-icon>
				<tm-icon :font-size="22" v-if="!status && _icon && _typeModel != 'dot'" :name="_icon"> </tm-icon>
			</tm-sheet>
			<view :eventPenetrationEnabled="true" style="width: 120rpx" class="flex flex-col flex-col-center-center mt-12">
				<slot name="default">
					<tm-text
						v-if="props.title"
						:userInteractionEnabled="false"
						:color="_isCheck ? _activeColor : ''"
						_class="text-overflow-2"
						:font-size="24"
						:label="props.title"
					>
					</tm-text>
					<tm-text
						v-if="props.label"
						:userInteractionEnabled="false"
						:color="_isCheck ? _activeColor : ''"
						_class="text-overflow-2 opacity-5"
						:font-size="22"
						:label="props.label"
					></tm-text>
				</slot>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 步骤条项目组件
 * @description 必须放置在tm-steps中使用，不可单独使用。
 * @slot default 默认插槽是底部标题的插槽。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import { computed, getCurrentInstance, inject, onUnmounted, ref } from 'vue'
import { custom_props } from '../../tool/lib/minxs'
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	transprent: {
		type: Boolean,
		default: false
	},
	color: {
		type: String,
		default: ''
	},
	activeColor: {
		type: String,
		default: ''
	},
	title: {
		type: String,
		default: ''
	},
	label: {
		type: String,
		default: ''
	},
	icon: {
		type: String,
		default: ''
	},
	//圆点的大小。
	size: {
		type: Number,
		default: 32
	},
	width: {
		type: Number,
		default: 150
	}
})
//父级方法。
let parent: any = proxy?.$parent

while (parent) {
	if (parent?.compoenentName == 'tmSteps' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
const _isNvue = ref(false)
// #ifdef APP-NVUE
_isNvue.value = true
// #endif

//本节点的位置步骤。
const _cureent = ref(parent?.pushKey(props.width) ?? 0)
onUnmounted(() => {
	if (parent) {
		_cureent.value = parent?.delKey(props.width)
	}
})
//总步骤数量
const _countCurrent = inject(
	'tmStepsCountCureent',
	computed(() => 1)
)
// 当前被激活的步骤
const _tmStepsCureent = inject(
	'tmStepsCureent',
	computed(() => -1)
)
const tmStepsCountActiveColor = inject(
	'tmStepsCountActiveColor',
	computed(() => 'primary')
)
const tmStepsCountColor = inject(
	'tmStepsCountColor',
	computed(() => 'grey-3')
)

const _activeColor = computed(() => {
	if (props.activeColor) return props.activeColor
	return tmStepsCountActiveColor.value
})

const _typeModel = computed(() => parent.$props.type)
const nodeSize = computed(() => (_typeModel.value != 'dot' ? props.size : 20))

const status_obj = {
	wait: {
		color: _activeColor.value,
		icon: 'tmicon-clock-fill'
	},
	process: {
		color: 'grey-2',
		icon: 'tmicon-loading'
	},
	finish: {
		color: 'green',
		icon: 'tmicon-check'
	},
	error: {
		color: 'red',
		icon: 'tmicon-times'
	}
}
const status = computed(() => {
	if (!_isActive.value) return null
	if (!status_obj.hasOwnProperty(parent.$props.status)) return null
	return status_obj[String(parent.$props.status)]
})
const _isActive = computed(() => _cureent.value === _tmStepsCureent.value)
const _isCheck = computed(() => _cureent.value < _tmStepsCureent.value)
const _color = computed(() => {
	if (status.value && _isActive.value) {
		return status.value.color
	}
	if (_isCheck.value) return _activeColor.value
	if (props.color) return props.color
	return tmStepsCountColor.value
})

const _icon = computed(() => {
	return props.icon
})

const showLine = computed(() => {
	return parent.$props.showLine
})

const _width = computed(() => {
	return Math.ceil(uni.upx2px(props.width)) + 'px'
})
const _widthLine = computed(() => {
	let _itemwidth = Math.ceil(uni.upx2px(props.width - nodeSize.value - 16) / 2) + 'px'
	return _itemwidth
})

async function stepClick() {
	if (!parent.$props.changeable) return
	if (typeof parent.$props.beforeStepChange === 'function') {
		uni.showLoading({ title: '...', mask: true })
		let p = await parent.$props.beforeStepChange()
		if (typeof p === 'function') {
			p = await p()
		}
		uni.hideLoading()
		if (!p) return
	}
	parent?.steplick(_cureent.value)
}
</script>

<style></style>

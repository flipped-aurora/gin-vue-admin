<template>
	<tm-sheet
		:color="props.color"
		:followTheme="props.followTheme"
		:dark="props.dark"
		:followDark="props.followDark"
		:round="props.round"
		:shadow="props.shadow"
		:outlined="props.outlined"
		:border="props.border"
		:borderStyle="props.borderStyle"
		:borderDirection="props.borderDirection"
		:text="props.text"
		:transprent="props.transprent"
		:linear="props.linear"
		:linearDeep="props.linearDeep"
		:width="props.width"
		:height="props.height"
		:margin="props.margin"
		:padding="props.padding"
		ref="tmDescriptions"
	>
		<slot name="title">
			<tm-text v-if="props.title" _class="pb-16" :label="props.title" :font-size="30"></tm-text>
		</slot>

		<view
			class="tmDescriptions"
			:class="[_cellWidth == 'auto' ? '' : 'flex flex-wrap flex-row-top-start']"
			:style="[{ flexDirection: _cellWidth == 'auto' ? 'column' : 'row' }]"
		>
			<slot>
				<block v-for="(item, index) in _dataList" :key="index">
					<tm-descriptions-item :label="item.label" :value="item.value"></tm-descriptions-item>
				</block>
			</slot>
		</view>
	</tm-sheet>
</template>

<script lang="ts" setup>
/**
 * 描述
 * @description 主要用于详细字段的陈述，可用于详情，列表一些描述性展示 。
 * @template title 标题，default 默认数据内容插槽
 */
import tmDescriptionsItem from '../tm-descriptions-item/tm-descriptions-item.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import { getCurrentInstance, computed, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const tmDescriptions = ref<InstanceType<typeof tmSheet> | null>(null)
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	shadow: {
		type: [Number, String],
		default: 0
	},
	round: {
		type: [Number, String],
		default: 0
	},
	border: {
		type: [Number],
		default: 0
	},
	margin: {
		type: Array,
		default: () => [0, 0]
	},
	padding: {
		type: Array,
		default: () => [16, 16]
	},
	transprent: {
		type: [Boolean, String],
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	//需要展示的列数。
	column: {
		type: [Number, String],
		default: 2
	},
	//数据。如果有。
	//这个是快捷方法，如果提供了，那么插槽中不需要填写tm-descriptions-item
	data: {
		type: [Array],
		default: () => []
	},
	//字段映射。
	keyMap: {
		type: [Object],
		default: { key: 'label', value: 'value' }
	},
	title: {
		type: String,
		default: ''
	},
	//定标签为等宽
	labelWidth: {
		type: [String, Number],
		default: ''
	},
	//定标签为等宽
	width: {
		type: [Number],
		default: 0
	}
})
const _dataList = computed(() => {
	return props.data.map((el: any) => {
		return {
			label: el[props.keyMap.key],
			value: el[props.keyMap.value]
		}
	})
})
const _cellWidth = ref('0px')

function nvueGetRect() {
	// #ifdef APP-PLUS-NVUE
	if (Number(props.column) <= 1) {
		_cellWidth.value = 'auto'
		return
	}
	try {
		nextTick(function () {
			dom?.getComponentRect(tmDescriptions.value, function (res: any) {
				if (!res?.size) return
				_cellWidth.value = res.size.width / Number(props.column) - uni.upx2px(20) + 'px'
				if (res.size.width == 0) {
					nvueGetRect()
				}
			})
		})
	} catch (e) {
		//TODO handle the exception
	}
	// #endif

	// #ifdef MP
	if (props.column <= 1) {
		_cellWidth.value = 'auto'
		return
	}
	uni.createSelectorQuery()
		.in(proxy)
		.select('.tmDescriptions')
		.boundingClientRect()
		.exec(function (res) {
			_cellWidth.value = res[0].width / Number(props.column) + 'px'

			if (res[0].width == 0) {
				nvueGetRect()
			}
		})

	// #endif
	// #ifdef APP-VUE || H5
	if (Number(props.column) <= 1) {
		_cellWidth.value = 'auto'
		return
	}
	_cellWidth.value = 100 / Number(props.column) + '%'
	// #endif
}
onMounted(() => {
	if (_cellWidth.value == '0px') {
		nvueGetRect()
	}
})
provide(
	'tmDescriptionsItem',
	computed(() => _cellWidth.value)
)
provide(
	'tmDescriptionsLabelWidth',
	computed(() => props.labelWidth)
)
</script>

<style></style>

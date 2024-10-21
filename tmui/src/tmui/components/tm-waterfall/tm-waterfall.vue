<template>
	<view>
		<view
			class="flex flex-col flex-col-top-start flex-between relative overflow"
			:style="[
				{
					width: _containerWidth + 'rpx',
					height: _containerHeight.maxHeight + props.bottomHeight + 'px'
				}
			]"
		>
			<slot></slot>
		</view>
		<!-- 虚拟加载占位符。 -->
		<tm-sheet
			v-if="_totalNum.length != _list.length && props.isLoadPlaceholder"
			_class="flex flex-center"
			:margin="[0, 0]"
			:padding="[0, 0]"
			:transprent="props.isLoadPlaceholderTransprent"
		>
			<tm-icon name="tmicon-shuaxin" spin></tm-icon>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
/**
 * 瀑布流
 * @description 瀑布流,只能放置tm-waterfall-item组件不可放置其它组件。
 * @example <tm-waterfall><tm-waterfall-item ></tm-waterfall-item></tm-waterfall>
 */
import { computed, nextTick, provide, Ref, ref, watch } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { itemParenSG } from './interface'
const props = defineProps({
	/**组件整体宽度 */
	width: {
		type: Number,
		default: 750
	},
	/**元素间的间距 */
	gutter: {
		type: Number,
		default: 12
	},
	/**底部高度 */
	bottomHeight: {
		type: Number,
		default: 50
	},
	/**是否开启虚拟加载占位符 */
	isLoadPlaceholder: {
		type: Boolean,
		default: true
	},
	/**虚拟加载占位符背景是否透明 */
	isLoadPlaceholderTransprent: {
		type: Boolean,
		default: false
	}
})

//容器的宽度
const _containerWidth = computed(() => props.width)
//项目的宽度
const _itemRealWidth = computed(() => {
	return (_containerWidth.value - props.gutter) / 2
})
//组件标识
const parentNameId = 'tmWaterfallId'
//这是首次进行缓存的数据。后续列表如果更新，针对id进行队列更新。
const _cacheList: Ref<Array<itemParenSG>> = ref([])
//这是按队列排序的左右数据主要是用来计算最大高度的。
const _totalSort: Ref<Array<Array<itemParenSG>>> = ref([[], []])
//排序的一列表数据
const _list: Ref<Array<itemParenSG>> = ref([])
const _totalNum: Ref<Array<number>> = ref([])
// //当前列表的高度,这是第一次的缓存高度。
//当前列表的排序后的列表高度

const _containerHeight = computed(() => {
	let gu = uni.upx2px(props.gutter)
	let lh = _totalSort.value[0].map((el) => el.height)
	let l_height = lh.length == 0 ? 0 : lh.reduce((a, b) => a + b + gu)
	let rh = _totalSort.value[1].map((el) => el.height)
	let r_height = rh.length == 0 ? 0 : rh.reduce((a, b) => a + b + gu)
	return {
		left: l_height,
		right: r_height,
		maxHeight: Math.max(r_height, l_height),
		minHeight: Math.min(r_height, l_height)
	}
})

provide(
	'tmWaterFallItemRealWidth',
	computed(() => uni.upx2px(_itemRealWidth.value))
)
function sumTotal(id: number) {
	_totalNum.value.push(id)
}
async function pushKey(n: itemParenSG) {
	let index = _cacheList.value.findIndex((el) => el.id == n.id)
	let item = n
	if (index > -1) {
		//已存在的不在更新高度，只作更新数据。
		_cacheList.value[index] = item
		return item
	} else {
		_cacheList.value.push(item)
		return countPushSort(item)
	}
}
function clear() {
	_cacheList.value = []
	_list.value = []
	_totalNum.value = []
	_totalSort.value = [[], []]
}

//排序数据。
function countPushSort(item: itemParenSG) {
	//当前数据是
	let dir = _containerHeight.value.left > _containerHeight.value.right ? 1 : 0
	let bottom = _totalSort.value[dir][_totalSort.value[dir].length - 1]?.bottom ?? 0
	item.top = bottom + uni.upx2px(props.gutter)
	item.bottom = item.top + item.height
	item.left = dir == 0 ? 0 : uni.upx2px(_itemRealWidth.value) + uni.upx2px(props.gutter)
	let index = _list.value.findIndex((el) => el.id == item.id)
	_totalSort.value[dir].push(item)
	if (index > -1) {
		_list.value[index] = item
	} else {
		_list.value.push(item)
	}
	return item
}

defineExpose({ parentNameId: parentNameId, pushKey, sumTotal, clear })
</script>

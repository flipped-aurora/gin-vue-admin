<template>
	<view class="flex flex-col">
		<scroll-view
			:scroll-into-view="viewNode"
			:refresher-triggered="Loading"
			@scrolltolower="pullEnd"
			:scroll-with-animation="true"
			@scroll="scroll"
			:scroll-y="true"
			:style="[
				{
					height: rootHeight + 'px',
					'overflow-anchor': 'auto',
					width: `${props.width}rpx`
				}
			]"
		>
			<!-- 顶部锚点 -->
			<view id="top"></view>
			<tm-sheet :height="40" unit="px" v-if="Loading && pullType == 'top'" :margin="[0, 0]" _class="flex flex-col flex-col-center-center">
				<tm-icon :color="props.color" :font-size="24" v-if="status == 'loading'" spin name="tmicon-shuaxin"></tm-icon>
				<view @click="reset('pullTop')" v-if="status == 'error'" class="flex flex-row flex-center">
					<tm-icon :userInteractionEnabled="false" color="red" :font-size="24" name="tmicon-times-circle-fill"></tm-icon>
					<tm-text :userInteractionEnabled="false" color="red" :font-size="24" _class="pl-16" label="加载失败,点我重试"></tm-text>
				</view>
			</tm-sheet>
			<view class="flex flex-col relative" :style="[{ height: totalHeight + 'px' }]">
				<view class="absolute l-0 t-0 flex flex-col" :style="{ height: `${props.topHeight}rpx`, width: `${props.width}rpx` }">
					<slot name="top"></slot>
				</view>
				<view class="absolute l-0 t-0 flex flex-col" :style="{ transform: `translateY(${offsetY}px)`, width: `${props.width}rpx` }">
					<slot name="default" :data="visibleItems"></slot>
				</view>
			</view>
			<tm-sheet :height="40" unit="px" v-if="Loading && pullType == 'bottom'" :margin="[0, 0]" _class="flex flex-col flex-col-center-center">
				<tm-icon :color="props.color" :font-size="24" v-if="status == 'loading'" spin name="tmicon-shuaxin"></tm-icon>
				<view @click="reset('pullEnd')" v-if="status == 'error'" class="flex flex-row flex-center">
					<tm-icon :userInteractionEnabled="false" color="red" :font-size="24" name="tmicon-times-circle-fill"></tm-icon>
					<tm-text :userInteractionEnabled="false" color="red" :font-size="24" _class="pl-16" label="加载失败,点我重试"></tm-text>
				</view>
			</tm-sheet>
			<!-- 底部锚点 -->
			<view id="bottom"></view>
		</scroll-view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 虚拟列表
 * @description 用来展示长列表数据使用，采用虚拟数据展示，只展示视窗内数据，其它不展示。因此上万条列表数据，也可以轻松展示。
 * @example 下面演示了加载1900张图片的示例
 * <tm-virtual-list :width="750" :height="800" :data="imglist" :itemHeight="160">
    <template v-slot:default="{data}">
      <tm-image-group >
        <tm-image :padding="[5,5]" preview :width="750" :height="150" :src="item" 
        v-for="(item,index) in data" :key="index"></tm-image>
      </tm-image-group>
    </template>
  </tm-virtual-list>
  -----js生成1900张图片链接。-----
  const imglist = ref([])
  for(let i=0;i<1900;i++){
    imglist.value.push('https://picsum.photos/200/300?id='+i)
  }
 */
import { ref, computed, Ref, PropType, onMounted, ComputedRef,watch } from 'vue'
import { scrollDetailFace, statusType } from './interface'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
const emits = defineEmits(['pullEnd', 'pullStart', 'status'])
const props = defineProps({
	width: {
		type: Number,
		default: 300
	},
	height: {
		type: Number,
		default: 500
	},
	topHeight: {
		type: Number,
		default: 0
	},
	//预估项目的高度，必填。
	itemHeight: {
		type: Number,
		default: 0,
		required: true
	},
	data: {
		type: Array as PropType<Array<any>>,
		default: () => []
	},
	//触底结束或者下拉刷新前执行，可以返回Promise<boolean>,真，结束触底操作，假触底加载中
	load: {
		type: [Function, Boolean],
		default: () => true
	},
	//首次加载渲染时，是否触发加载数据事件。
	firstLoad: {
		type: Boolean,
		default: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	scrollViewInTo:{
		type:String as PropType<'top'|'bottom'|''>,
		default:""
	}
})
const rowHeight = uni.upx2px(props.itemHeight)
const outTopHeight = uni.upx2px(props.topHeight)
const rootHeight = uni.upx2px(props.height)
const scrollTop = ref(0)
const renderAhead = 2
const rowCount = computed(() => props.data.length)
const childPositions = computed(() => {
	const results = [0]
	for (let i = 1; i < rowCount.value; i++) {
		results.push(results[i - 1] + rowHeight)
	}
	return results
})
const totalHeight = computed(() => {
	return (rowCount.value ? childPositions.value[rowCount.value - 1] + rowHeight : 0) + outTopHeight
})
const firstVisibleNode = computed(() => findStartNode())
const startNode = computed(() => Math.max(0, firstVisibleNode.value - renderAhead))
const lastVisibleNode = computed(() => findEndNode())
const endNode = computed(() => Math.min(rowCount.value - 1, lastVisibleNode.value + renderAhead))
const visibleNodeCount = computed(() => endNode.value - startNode.value + 1)
const offsetY = computed(() => childPositions.value[startNode.value] + outTopHeight)
const visibleItems: ComputedRef<Array<any>> = computed(() => {
	return props.data.slice(startNode.value, startNode.value + visibleNodeCount.value)
})
const viewNode = ref("")
watch(()=>props.scrollViewInTo,()=>{
	viewNode.value = props.scrollViewInTo
})

const Loading = ref(false)
const pullType = ref('')
/**
 * never 从未加载过。
 */
const status: Ref<statusType> = ref('never')

function findStartNode() {
	let startRange = 0
	let endRange = rowCount.value ? rowCount.value - 1 : rowCount.value

	while (endRange !== startRange) {
		const middle = Math.floor((endRange - startRange) / 2 + startRange)

		if (childPositions.value[middle] <= scrollTop.value && childPositions.value[middle + 1] > scrollTop.value) {
			return middle
		}
		if (middle === startRange) {
			// edge case - start and end range are consecutive
			return endRange
		}
		if (childPositions.value[middle] <= scrollTop.value) {
			startRange = middle
		} else {
			endRange = middle
		}
	}
	return rowCount.value
}
function findEndNode() {
	let endNode
	for (endNode = firstVisibleNode.value; endNode < rowCount.value; endNode++) {
		if (childPositions.value[endNode] > childPositions.value[firstVisibleNode.value] + rootHeight) {
			return endNode
		}
	}
	return endNode
}

function scroll(e: any) {
	let detail: scrollDetailFace = e.detail
	scrollTop.value = detail.scrollTop
	if (Math.ceil(scrollTop.value) < -80) {
		pullStart('pullStart')
	}
	if (Math.ceil(scrollTop.value) >= 0 && status.value == 'error') {
		Loading.value = false
	}
}

const pullStart = async (e: any) => {
	emits('pullStart')
	if (typeof props.load === 'function') {
		if (Loading.value) return
		pullType.value = 'top'
		Loading.value = true
		status.value = 'loading'
		let p = await props.load('top')
		if (typeof p === 'function') {
			p = await p('top')
		}
		if (!p) {
			status.value = 'error'
			return
		}
		Loading.value = false
		status.value = 'success'
	}
}

const pullEnd = async (e: any) => {
	emits('pullEnd')
	if (typeof props.load === 'function') {
		if (Loading.value) return
		pullType.value = 'bottom'
		Loading.value = true
		status.value = 'loading'
		let p = await props.load('bottom')
		if (typeof p === 'function') {
			p = await p('bottom')
		}
		if (!p) {
			status.value = 'error'
			return
		}
		Loading.value = false
		status.value = 'success'
	}
}

const reset = function (e: string) {
	Loading.value = false
	if (e == 'pullEnd') {
		pullEnd('bottom')
	} else {
		pullStart('top')
	}
}

onMounted(() => {
	if (props.firstLoad) {
		pullStart('top')
	}
})
</script>

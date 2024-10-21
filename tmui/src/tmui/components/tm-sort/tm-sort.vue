<template>
	<view id="box" ref="box" :style="{ width: props.width + 'rpx', height: props.height + 'rpx' }">
		<!-- <view 
        class="flex flex-row-top-start flex-row flex-wrap"
        :style="{ width: props.width + 'rpx', height: props.height + 'rpx', background: 'red' }">
            <view v-for="(item, index) in _data" :key="index"
                :style="{ width: `${item.width}px`, height: `${item.height}px`, background: 'yellow' }">
                {{ item._id }}
            </view>
        </view> -->
		<movable-area :style="{ width: props.width + 'rpx', height: props.height + 'rpx' }">
			<movable-view
				:class="[moveIndex === item._index ? 'opacity-6 zIndex-1' : 'zIndex-0']"
				:animation="false"
				@touchstart="drageStart"
				@touchmove="onTouchMove($event, index)"
				@touchcancel="dragEnd($event, index)"
				@touchend="dragEnd($event, index)"
				@change="eleChange($event, index)"
				:x="item._x"
				:y="item._y"
				:style="{
					width: `${item.width}px`,
					height: `${item.height}px`,
					background: 'white'
				}"
				class=""
				:data-index="item._index"
				direction="all"
				v-for="(item, index) in _data"
				:key="index"
			>
				<view :class="[onmoveIndex === index ? 'onSelecteItemStyle ' : '', 'onSelecteItem ']">
					<tm-sheet
						:border="0"
						:margin="[0]"
						:padding="[0]"
						:width="item.width"
						:height="item.height"
						unit="px"
						:color="onmoveIndex === index ? 'grey-2' : props.color"
						_class="flex flex-col flex-col-center-center"
					>
						<tm-icon v-if="item.icon" :font-size="props.iconSize" :name="item?.icon ?? ''"></tm-icon>
						<tm-text _class="text-align-center" :font-size="props.fontSize" :label="item?.text ?? ''"></tm-text>
					</tm-sheet>
				</view>
			</movable-view>
		</movable-area>
	</view>
</template>
<script setup lang="ts">
/**
 * 宫格排序，
 * 在nvue上：
 * 截止3.6.18 uni sdk 的movable有bug,具体来说是动态修改宽和高时，可移动范围会跑上面去。
 * 理论上是支持nvue，只有当uni修复此bug才可使用，否则，请使用vue页面。
 *
 */
import customProps from './customProps'
import type { cellOptsType } from './interface'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { ref, Ref, nextTick, onMounted, watch, getCurrentInstance } from 'vue'
const props = defineProps(customProps)
const emits = defineEmits(['start', 'end'])
const boxHeight = ref(0)
const proxy = getCurrentInstance()?.proxy ?? null
let cellWidth = 0
let cellHeight = 0
let maxRow = 0
const _data: Ref<Array<cellOptsType>> = ref([])
const movexy = { x: 0, y: 0 }
//当前拖动的项目
const moveIndex = ref(-1)
//即将替换的项目
const onmoveIndex = ref(-1)
let tid: any = NaN
let boxOpts = { left: 0, top: 0 }
const onMoveStauts = ref(false)
let testdata: Array<any> = props.list
onMounted(() => {
	watchOpts()
	setTimeout(() => {
		uni.$tm.u.quereyDom(proxy, '#box').then((ev: any) => {
			if (ev) {
				boxOpts.left = ev.left
				boxOpts.top = ev.top
			}
		})
	}, 50)
})
watch(
	() => props.list,
	() => {
		testdata = uni.$tm.u.deepClone(props.list)
		watchOpts()
	},
	{ deep: true }
)
function watchOpts() {
	clearTimeout(tid)
	tid = setTimeout(() => {
		initOpts()
	}, 100)
}
function initOpts(isOld: boolean = false) {
	let col = props.column
	if (!isOld) {
		testdata = testdata.map((el, index) => {
			el['_id'] = index
			return el
		})
	}
	let data = uni.$tm.u.deepClone(testdata)

	let d: Array<Array<cellOptsType>> = uni.$tm.u.splitData(data, col)
	let row = d.length
	maxRow = row
	// uni sdk3.6.18有bug，如果自动计算高度会导致拖动区域异常。
	// boxHeight.value = uni.upx2px(row * props.itemHeight)
	cellWidth = uni.upx2px(props.width / col)
	cellHeight = uni.upx2px(props.itemHeight)
	d = d.map((el, index) => {
		el = el.map((ele, index2) => {
			ele._index = index * col + index2
			ele._id = ele['_id']
			if (!isOld) {
				ele._x = index2 * cellWidth
				ele._y = index * cellHeight
			} else {
				ele._x = index2 * cellWidth - 0.5
				ele._y = index * cellHeight - 0.5
			}
			ele._newx = index2 * cellWidth
			ele._newy = index * cellHeight
			ele.width = cellWidth
			ele.height = cellHeight
			return ele
		})
		return el
	})

	_data.value = []
	let dtestdata: any = []
	d.forEach((el) => {
		dtestdata.push(...el)
	})
	_data.value = [...dtestdata]
	if (isOld) {
		setTimeout(() => {
			_data.value = _data.value.map((ele) => {
				ele._x = ele._newx
				ele._y = ele._newy
				return ele
			})
		}, 50)
	}
}

function eleChange(ev: any, index: number) {
	if (!onMoveStauts.value) return
	let { x, y } = ev.detail
	movexy.x = x
	movexy.y = y
	moveIndex.value = index
}
function onTouchMove(ev: any, index: number) {
	if (!onMoveStauts.value) return
	clearTimeout(tid)
	tid = setTimeout(() => {
		onmoveIndex.value = getNowXyindex(ev, index)
		// #ifndef H5
		try {
			uni.vibrateShort({})
		} catch (error) {}
		// #endif
		if (!onMoveStauts.value) {
			onmoveIndex.value = -1
		}
	}, 150)
}
function drageStart() {
	onMoveStauts.value = true
	emits('start', testdata)
}
function dragEnd(ev: any, index: number) {
	onMoveStauts.value = false
	let olddata = uni.$tm.u.deepClone(testdata)
	let data = uni.$tm.u.deepClone(testdata)
	const item = data[index]
	let newinex = getNowXyindex(ev, index)
	const [removed] = data.splice(newinex, 1, item)
	data.splice(index, 1, removed)
	testdata = [...data]
	initOpts(true)
	moveIndex.value = -1
	onmoveIndex.value = -1
	emits('end', testdata, olddata)
}
function getNowXyindex(ev: any, nowIndex: number = 0) {
	let { clientX, clientY, pageX, pageY } = ev.changedTouches[0]
	let cx = clientX ?? pageX
	let cy = clientY ?? pageY
	let newinex = 0
	cx = cx <= boxOpts.left ? boxOpts.left : cx
	cy = cy <= boxOpts.top ? boxOpts.top : cy
	let x = cx - boxOpts.left
	let y = cy - boxOpts.top
	let nx = Math.floor(x / cellWidth)
	let ny = Math.floor(y / cellHeight)
	if (ny >= maxRow - 1) {
		ny = maxRow - 1
	}
	if (nx >= props.column - 1) {
		nx = props.column - 1
	}
	newinex = ny * props.column + nx
	if (newinex >= testdata.length) {
		newinex = testdata.length - 1
	}
	return newinex
}
</script>
<style scoped>
.onSelecteItem {
	transition-property: transform;
	transition-duration: 0.3s;
	transition-timing-function: ease;
}

.onSelecteItemStyle {
	transform: scale(0.9);
}
</style>

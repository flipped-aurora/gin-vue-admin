<template>
	<view ref="tm-dragList" class="tm-dragList flex flex-col flex-col-center-center">
		<view :style="{ height: h * list.length + 'px', width: w + 'px' }" class="relative flex flex-col" :class="[disabled ? 'opacity-4' : '']">
			<view
				class="overflow flex flex-col"
				:class="['absolute', 'tm-dragList-item', 'flex-between']"
				v-for="(item, index) in listData"
				:key="index"
				:style="[
					{
						top: `${item.top}px`,
						height: h + 'px',
						width: w + 'px',
						zIndex: nowMove_index == index ? 5 : 0
					},
					isNvue
						? { 'transition-delay': '0.1s' }
						: {
								'transition-duration': nowMove_index == index || endDrage ? '0s' : '0.25s'
						  }
				]"
			>
				<tm-sheet
					hover-class="opacity-6"
					:border="1"
					border-direction="bottom"
					:color="nowMove_index == index ? 'grey-3' : 'white'"
					_class="flex-1 flex flex-row flex-between"
					class="flex-1"
					:margin="[0, 0]"
					:padding="[0, 0]"
					@click="onClick(index)"
				>
					<view class="flex flex-row flex-row-center-start pl-12" :style="{ height: h - 1 + 'px' }">
						<view v-if="item['icon']" class="flex-shrink fulled-height flex-center">
							<tm-icon :color="item['color']" :name="item['icon']" :fontSize="40"></tm-icon>
						</view>
						<tm-text _class=" pl-24" :font-size="30" :label="item[props.rangKey]"></tm-text>
					</view>
					<view
						:style="{ height: h - 1 + 'px', width: '100rpx' }"
						@touchstart.stop.prevent="m_start($event, index)"
						@longpress="m_start_longpress(index)"
						@mousedown="m_start($event, index)"
						@touchmove.stop.prevent="m_move($event, index)"
						@mousemove.stop.prevent="m_move($event, index)"
						@touchend="m_end($event, index)"
						@mouseup="m_end($event, index)"
						class="flex-shrink flex flex-row flex-row-center-center opacity-3"
					>
						<tm-icon :userInteractionEnabled="false" name="tmicon-menu"></tm-icon>
					</view>
				</tm-sheet>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 拖放排序列表
 * @description 在h5和pc端，点击右边区域即可触发拖动。在MP和APP端需要长按右边区域触发拖动排序。
 * @description 已知hack:在nvue端，因zIndex不起作用，导致前端组件移动时，可能会被后面渲染的列表覆盖，但不影响使用，只影响些许美观。
 */
import { Ref, ref, nextTick, toRaw, watch, getCurrentInstance, PropType } from 'vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import { itemList } from './interface'
const props = defineProps({
	disabled: {
		type: [String, Boolean],
		default: false
	},
	width: {
		type: Number,
		default: 700
	},
	height: {
		type: Number,
		default: 90
	},
	list: {
		type: Array as PropType<Array<itemList>>,
		default: () => {
			return []
		}
	},
	rangKey: {
		type: String,
		default: 'text'
	},
	rightIcon: {
		type: String,
		default: 'icon-menu'
	},
	bgColor: {
		type: String,
		default: 'white'
	}
})
const emits = defineEmits(['change','click'])
const w = ref(0)
const h = ref(0)
const totalH = ref(0)
let y = 0
const new_index: Ref<number> = ref(NaN) //即将被替换的索引（实质性被替换）
const nowMove_index: Ref<number> = ref(NaN) //现在正在移动的索引
const listData = ref<itemList[]>([])
let new_item = [] //虚拟列表，内部排列好，但未在页面中渲染。
const endDrage = ref(false)
const isNvue = ref(false)
// #ifdef APP-NVUE
isNvue.value = true
// #endif

watch(
	() => props.list,
	() => {
		jishunTopData()
	
	},
	{
		deep: true,
		immediate: true
	}
)
function jishunTopData() {
	let nowList = uni.$tm.u.deepClone(props.list)
	nextTick(function () {
		listData.value = []
		w.value = uni.upx2px(props.width) || 700
		h.value = uni.upx2px(props.height) || 120
		totalH.value = h.value * nowList.length
		for (let i = 0; i < nowList.length; i++) {
			let p: itemList = nowList[i]
			p['top'] = i * h.value
			p['i'] = i
			p['__id'] = uni.$tm.u.getUid()
			// testList.push(p);
		}
		// console.log(nowList)
		listData.value = [...nowList]
		new_item = [...nowList]
	})
}
function onClick(index:number){
	emits('click',index)
}
function m_start_longpress(index: number) {
	endDrage.value = false
	nowMove_index.value = index
	// #ifdef APP-NVUE
	uni.vibrateShort({
		success: function () {}
	})
	// #endif
}

function m_start(event: Event, index: number) {
	event.preventDefault()
	event.stopPropagation()
	if (props.disabled) return
	// #ifndef APP-NVUE
	endDrage.value = false
	nowMove_index.value = index
	// #endif

	new_item = [...listData.value]
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]
		y = touch.pageY - event.currentTarget.offsetTop - listData.value[index].top
	} else {
		y = event.pageY - event.currentTarget.offsetTop - listData.value[index].top
	}

	// #ifdef APP-NVUE || MP
	// uni.vibrateShort()
	// #endif
}
function m_move(event: Event, index: number) {
	if (props.disabled) return
	event.preventDefault()
	event.stopPropagation()

	if (isNaN(nowMove_index.value)) return
	//当前元素的top位置。
	let ch = 0
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]
		ch = touch.pageY - y
	} else {
		ch = event.pageY - y
	}

	listData.value.splice(index, 1, {
		...listData.value[index],
		top: ch
	})
	const currenit_index = index
	const currentSort = listData.value[currenit_index].i
	const currenit_id = listData.value[currenit_index].__id

	// 计算当前移动的index.
	let moveIndex = Math.round(ch / h.value)

	moveIndex = moveIndex < 0 ? 0 : moveIndex
	moveIndex = moveIndex > listData.value.length - 1 ? listData.value.length - 1 : moveIndex

	moveIndex = Math.abs(moveIndex)

	index = moveIndex
	let elList = [...listData.value]
	for (let i = 0; i < elList.length; i++) {
		if (currentSort < moveIndex) {
			if (elList[i].i > currentSort && elList[i].i <= moveIndex) {
				elList[i].i -= 1
			}
		} else if (currentSort > moveIndex) {
			if (elList[i].i < currentSort && elList[i].i >= moveIndex) {
				elList[i].i += 1
			}
		}
	}
	elList[currenit_index].i = moveIndex
	elList = elList.map((im) => {
		if (im.__id != currenit_id) {
			im.top = im.i * h.value
		}

		return im
	})
	listData.value = elList
	new_index.value = moveIndex
}
function m_end(event: Event, index: number) {
	if (props.disabled) return
	event.preventDefault()
	event.stopPropagation()
	nowMove_index.value = NaN
	endDrage.value = true
	if (isNaN(new_index.value)) return
	let elList = [...listData.value]
	elList = elList.map((im) => {
		im.top = im.i * h.value
		return im
	})
	elList.sort((a, b) => a.i - b.i)
	listData.value = [...elList]
	moveChange()
}

function moveChange() {
	if (props.disabled) return
	//change后修改的数据 。
	emits('change', toRaw(listData.value))
}
</script>

<style scoped>
.tm-dragList-item {
	transition-delay: 0s;
	transition-property: top, left, transform;
	transition-timing-function: ease;
	transition-duration: 0.15s;
}
</style>

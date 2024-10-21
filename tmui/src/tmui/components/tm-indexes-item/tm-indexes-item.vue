<template>
	<view :id="'id_' + readId">
		<slot name="title">
			<tm-sheet
				no-level
				@click="emits('title-click')"
				:transprent="props.transprent"
				color="grey-3"
				v-if="_title !== ''"
				:height="50"
				_class="flex flex-row flex-row-center-start"
				:margin="[props.margin[0], 0]"
				:padding="[props.padding[0], 0]"
			>
				<tm-text :userInteractionEnabled="false" :font-size="24" _class="text-weight-b opacity-6" :label="_title"></tm-text>
			</tm-sheet>
		</slot>
		<tm-sheet
			@click="itemClick"
			color="white"
			:border="props.showBorder ? 1 : 0"
			:transprent="props.transprent"
			borderDirection="bottom"
			:margin="props.margin"
			:padding="props.padding"
		>
			<view :eventPenetrationEnabled="true" class="flex-1">
				<slot></slot>
			</view>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
/**
 * 列表索引项目
 * @description 索引列表项目，内部只能放置在tm-indexes组件中。
 */
import { computed, PropType, ref, getCurrentInstance, onUnmounted, onMounted, onUpdated, nextTick } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['click', 'title-click'])
const props = defineProps({
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 0]
	},
	height: {
		type: Number,
		default: 100
	},
	//如果title不为空，此组件就会当分类标题来显示，并隐藏内容。
	title: {
		type: [String, Number],
		default: ''
	},
	// 右边快捷显示的标题，只能是一个字符。
	//注意如果title为空时，这个才会显示在右边。不空此字段无任何意义。
	navTitle: {
		type: [String, Number],
		default: ''
	},
	transprent: {
		type: Boolean,
		default: false
	},
	showBorder: {
		type: Boolean,
		default: true
	}
})

const _title = computed(() => props.title)
const _navTitle = computed(() => props.navTitle)
const readId = uni.$tm.u.getUid(1)
let _itemHeight = 0
//父级方法。
let parent: any = proxy.$parent

while (parent) {
	if (parent?.compentNameId == 'tmIndexesId' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
if (parent) {
	parent.pushKey(_itemHeight, readId, String(_title.value), String(_navTitle.value))
}
onUnmounted(() => parent.delKey(_itemHeight, readId))
onMounted(() => getRect())
onUpdated(() => getRect())
function getRect() {
	let id = '#id_' + readId
	// #ifdef MP
	uni.createSelectorQuery()
		.in(proxy)
		.select(id)
		.boundingClientRect()
		.exec((ret) => {
			_itemHeight = ret[0].height
			parent.pushKey(ret[0].height, readId, String(_title.value), String(_navTitle.value))
		})
	// #endif
}
function itemClick() {
	emits('click')
}
</script>

<style></style>

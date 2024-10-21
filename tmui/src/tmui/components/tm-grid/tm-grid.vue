<template>
	<tm-sheet
		:round="props.round"
		:width="props.width"
		:transprent="props.transprent"
		:color="props.color"
		:margin="[0, 0]"
		:padding="[0, 0]"
		_class="flex flex-row flex-row-top-start"
		contStyle="flex-wrap:wrap;"
	>
		<slot></slot>
	</tm-sheet>
</template>
<script lang="ts" setup>
/**
 * 宫格
 * @description 注意，它内部只能放置tm-grid-item，且不能嵌套tm-grid
 * @slot 默认插槽 放置tm-grid-item
 * @example 例子：
    <tm-grid >
        <tm-grid-item v-for="(item,index) in 11" :key="index">
            <tm-icon  :fontSize="64" name="https://gw.alicdn.com/imgextra/i4/O1CN01XCiY1B1px9ivHkDGm_!!6000000005426-2-tps-183-144.png"></tm-icon>
            <tm-text :fontSize="22" label="苹果"></tm-text>
        </tm-grid-item>
    </tm-grid>
 */
import { computed, provide, ref, Ref, inject } from 'vue'
import { custom_props } from '../../tool/lib/minxs'
import tmSheet from '../tm-sheet/tm-sheet.vue'
const props = defineProps({
	...custom_props,
	round: {
		type: Number,
		default: 2
	},
	width: {
		type: Number,
		default: 750
	},
	//一行放置几个。
	col: {
		type: Number,
		default: 5
	},
	//是否显示边线。
	showBorder: {
		type: Boolean,
		default: false
	},
	// 背景
	color: {
		type: String,
		default: 'white'
	},
	//是否透明背景
	transprent: {
		type: Boolean,
		default: false
	}
})
interface arrayid {
	id: string | number
	type: string | number
}

let _cachList: Ref<Array<arrayid>> = ref([])
const _colWidth = computed(() => {
	return Math.ceil(props.width / props.col - 1)
})
provide('tmGridItemWidth', _colWidth.value + (props.showBorder ? 1 : 0))
provide(
	'tmGridshowBorder',
	computed(() => props.showBorder)
)
provide(
	'tmGridshowCachList',
	computed(() => _cachList.value)
)

function pushKey(e: arrayid) {
	let index = _cachList.value.findIndex((el) => el.id == e.id)
	if (index == -1) {
		_cachList.value.push(e)
	} else {
		_cachList.value.splice(index, 1, e)
	}

	setIndexType()
}

function delKey(e: arrayid) {
	let index = _cachList.value.findIndex((el) => el.id == e.id)
	setIndexType()
}

function setIndexType() {
	// type
	/**
	 * 1,all,
	 * 2,bottom,right
	 * 3,left,bottom,right
	 * 4,top,left,bottom
	 */
	let totallen = _cachList.value.length
	_cachList.value = _cachList.value.map((el, index) => {
		let aIndex = index + 1
		if (aIndex <= props.col) {
			el.type = 4
			if ((aIndex == totallen && totallen == 1) || aIndex == 1) {
				el.type = 1
			}
		} else {
			if (aIndex % props.col == 1) {
				el.type = 3
			} else {
				el.type = 2
			}
		}
		return el
	})
}
defineExpose({
	pushKey,
	delKey,
	keyName: 'tmGrid'
})
</script>

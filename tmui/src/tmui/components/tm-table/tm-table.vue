<template>
	<!-- #ifdef APP-NVUE -->
	<table></table>
	<!-- #endif -->
	<!-- #ifndef APP-NVUE -->

	<view class="tableBox" :style="[{ width: props.width + 'rpx' }, props.height ? { height: props.height + 'rpx' } : '']">
		<tm-row v-if="props.showHeader" :transprent="true" :width="totalTableWidth" :column="_data.header.length">
			<tm-col
				_class="flex flex-row flex-row-center-center"
				:text="item.opts?.light ?? false"
				:height="props.headerHeight"
				:color="item.opts?.color ?? 'primary'"
				:transprent="false"
				v-for="(item, index) in _data.header"
				:key="index"
				@click="headerClick(item, index)"
			>
				<view
					:style="{
						width:
							(item.opts?.sort ?? false ? totalTableWidth / _data.header.length - 60 : totalTableWidth / _data.header.length - 60) +
							'rpx'
					}"
				>
					<tm-text
						_class="text-overflow-2 text-align-center"
						:line-height="0"
						:font-size="item.opts?.fontSize ?? 0"
						:color="item.opts?.fontColor ?? ''"
						:label="item.name"
					></tm-text>
				</view>
				<tm-icon v-if="item.opts?.sort ?? false" name="tmicon-sort" :font-size="26"></tm-icon>
			</tm-col>
		</tm-row>

		<block v-for="(item, index) in _maxrows" :key="item">
			<view :style="{ width: totalTableWidth + 'rpx' }">
				<tm-divider color="grey-5" v-if="props.showBottomBorder && !_stripe" :margin="[0, 0]"></tm-divider>
			</view>
			<tm-row
				:color="_stripe ? (index % 2 ? 'grey-5' : 'white') : 'white'"
				:width="totalTableWidth"
				:column="_data.header.length"
				:transprent="false"
			>
				<tm-col
					:text="getOptsCellStyle(index2, index)?.light ?? false"
					:height="props.cellHeight"
					v-for="(item2, index2) in _data.header.length"
					:key="item2"
					:color="getOptsCellStyle(index2, index)?.color ?? ''"
					:transprent="_stripe"
					@click="cellClick(index2, index)"
					_class="flex flex-row flex-row-center-center"
				>
					<view class="flex-1" v-if="((_rows[index2] ?? [])[index].opts?.type ?? 'text') == 'button'">
						<tm-button
							:font-size="getOptsCellStyle(index2, index)?.fontSize ?? 26"
							:color="getOptsCellStyle(index2, index)?.fontColor ?? 'primary'"
							:label="(_rows[index2] ?? [])[index].value ?? '-'"
							:margin="[10, 10]"
							:height="props.cellHeight - 20"
							size="mini"
							block
						></tm-button>
					</view>
					<tm-text
						v-else
						:font-size="getOptsCellStyle(index2, index).fontSize"
						:line-height="0"
						:color="getOptsCellStyle(index2, index).fontColor"
						:label="(_rows[index2] ?? [])[index].value ?? '-'"
					></tm-text>
				</tm-col>
			</tm-row>
		</block>
	</view>

	<!-- #endif -->
</template>

<script setup lang="ts">
// #ifdef APP-NVUE
import table from './table.vue'
// #endif
// #ifndef APP-NVUE
import type { headresItem, cellItem, dataTypeArray } from './interface'
import { tableDataType, tabaleCellStyleType, defaultCellStyle, tabaleCellData, headerType } from './newInterface'
import { computed, nextTick, onMounted, PropType, Ref, ref, toRaw, watchEffect, watch } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmButton from '../tm-button/tm-button.vue'
import tmRow from '../tm-row/tm-row.vue'
import tmCol from '../tm-col/tm-col.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import { cutomProps } from './cutomProps'
const emits = defineEmits(['rowClick'])
const props = defineProps(cutomProps)
const _data: Ref<tableDataType> = ref(props.tableData)
const sortType = ref('none')
let tid: any = NaN
const _stripe = computed(() => props.stripe)
const totalTableWidth = computed(() => {
	let d = props.cellWidth * _data.value.header.length

	if (d <= props.width) d = props.width
	return d
})
const _rows: Ref<Array<Array<tabaleCellData>>> = ref([])
let _rows_back: Array<Array<tabaleCellData>> = []
const _maxrows = ref(0)

onMounted(() => {
	_rows.value = chuliRows(_data.value.data)
	_rows_back = uni.$tm.u.deepClone(_rows.value)
})
watch(
	[() => props.tableData],
	() => {
		clearTimeout(tid)
		tid = setTimeout(() => {
			let cdatas: any = uni.$tm.u.deepClone(props.tableData)
			_data.value = { ...cdatas }
			_rows.value = chuliRows(_data.value.data)
			_rows_back = uni.$tm.u.deepClone(_rows.value)
		}, 150)
	},
	{ deep: true }
)

function chuliRows(bigdata: Array<{ [key: string]: any }>) {
	let d: Array<Array<tabaleCellData>> = []
	let dlen: number[] = []
	let dstyle: Array<Array<{ [key: string]: tabaleCellStyleType }>> = []
	if (!bigdata) {
		bigdata = []
	}
	_data.value.data = bigdata.map((el) => {
		let ptps: { [key: string]: tabaleCellStyleType } = el['opts'] ?? {}
		_data.value.header.forEach((ielem) => {
			if (ptps[ielem.field]) {
				ptps[ielem.field] = { ...(ptps[ielem.field] ?? {}) }
			} else {
				ptps[ielem.field] = {}
			}
		})
		el.opts = ptps
		return el
	})

	_data.value.header.forEach((el) => {
		let pd: Array<tabaleCellData> = []
		pd = _data.value.data.map((ele: { [key: string]: any }) => {
			let isasync = el?.opts?.asyncStyleCell ?? false

			return {
				value: ele[el.field] ?? '-',
				opts: {
					...defaultCellStyle,
					...{
						color: isasync === true ? el.opts?.color ?? 'white' : 'white',
						fontColor: isasync === true ? el.opts?.fontColor ?? 'black' : '',
						light: isasync === true ? el.opts?.light ?? false : false
					},
					...ele['opts'][el.field]
				}
			}
		})

		d.push(pd)
		dlen.push(pd.length)
	})
	_maxrows.value = Math.max(...(dlen.length ? dlen : [0, 0]))
	return d
}
function getOptsCellStyle(index1: number, index2: number) {
	let d = (_rows.value[index1] ?? [])[index2].opts ?? { ...defaultCellStyle }
	return d
}
function cellClick(index1: number, index2: number) {
	let dp = (_rows.value[index1] ?? [])[index2].value ?? ''
	/**row,col,value */
	emits('rowClick', index2, index1, dp)
}
function sort(data: Array<Array<tabaleCellData>>, index: number, type = 'none', callback: Function) {
	uni.showLoading({
		title: '...',
		mask: true
	})
	let d = data[index]
	// 排序
	if (type == 'none') {
		data = uni.$tm.u.deepClone(_rows_back)
	}

	//降序
	if (type == 'desc') {
		let dbiaoji = new Array()
		for (let i = 0; i < d.length; i++) {
			dbiaoji.push(i)
		}
		d = d.map((el, iof) => {
			el['__ids'] = iof
			return el
		})
		d.sort((a, b) => Number(a.value) - Number(b.value))
		let pd: Array<Array<tabaleCellData>> = []
		data.forEach((element, index2) => {
			let p: Array<tabaleCellData> = []
			if (index2 !== index) {
				d.forEach((el, index3) => {
					let nm = el['__ids']
					p.push(element[nm])
				})
			} else {
				p = d
			}
			pd.push(p)
		})

		data = pd
	}
	//升序。
	if (type == 'asce') {
		let pd: Array<Array<tabaleCellData>> = []
		data.forEach((element, index2) => {
			let p: Array<tabaleCellData> = []
			p = [...element.reverse()]
			pd.push(p)
		})

		data = pd
	}

	nextTick(() => {
		if (callback) {
			callback(data)
		}
		uni.hideLoading()
	})
}
function headerClick(item: headerType, index: number) {
	if (item.opts?.sort ?? false) {
		let d = uni.$tm.u.deepClone(_rows.value)
		if (sortType.value == 'none') {
			sortType.value = 'desc'
		} else if (sortType.value == 'desc') {
			sortType.value = 'asce'
		} else if (sortType.value == 'asce') {
			sortType.value = 'none'
		}
		sort(d, index, sortType.value, (ds: Array<Array<tabaleCellData>>) => {
			_rows.value = [...ds]
		})
	}
}

// #endif
</script>

<style scoped>
.tableBox {
	/* #ifndef APP-NVUE */
	overflow: auto;
	/* #endif */
}
</style>

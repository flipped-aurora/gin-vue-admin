<template>
	<list :style="{ width: `${defaultProps.width}rpx`, height: `${defaultProps.height}rpx` }">
		<header
			v-if="_showHeader"
			class="tableBox flex flex-row"
			:style="{
				width: `${defaultProps.width}rpx`,
				height: `${defaultProps.headerHeight}rpx`
			}"
		>
			<scroller ref="scrolleHead" :scrollable="false" class="tableRight flex flex-row flex-1" scroll-direction="horizontal">
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
							:userInteractionEnabled="false"
							:style="{
								width:
									(item.opts?.sort ?? false
										? totalTableWidth / _data.header.length - 60
										: totalTableWidth / _data.header.length - 60) + 'rpx'
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
						<tm-icon :userInteractionEnabled="false" v-if="item.opts?.sort ?? false" name="tmicon-sort" :font-size="26"></tm-icon>
					</tm-col>
				</tm-row>
			</scroller>
		</header>
		<cell class="tableBox" :style="{ width: `${defaultProps.width}rpx` }">
			<scroller ref="scrollerRight" class="tableRight" scroll-direction="horizontal" :style="{ width: `${defaultProps.width}rpx` }">
				<view v-for="(item, index) in _maxrows" :key="item" :style="{ width: totalTableWidth + 'rpx' }">
					<view>
						<tm-divider color="grey-5" v-if="props.showBottomBorder && !_stripe" :margin="[0, 0]"></tm-divider>
					</view>

					<tm-row
						:color="_stripe ? (index % 2 ? 'grey-5' : 'white') : 'white'"
						:width="totalTableWidth"
						:column="_data.header.length"
						:transprent="false"
					>
						<tm-col
							:col="1"
							:text="getOptsCellStyle(index2, index)?.light ?? false"
							:height="props.cellHeight"
							v-for="(item2, index2) in _data.header.length"
							:key="item2"
							:color="getOptsCellStyle(index2, index)?.color ?? ''"
							:transprent="_stripe"
							@click="cellClick(index2, index)"
							_class="flex flex-row flex-row-center-center"
						>
							<view
								:userInteractionEnabled="false"
								class="flex-1"
								v-if="((_rows[index2] ?? [])[index].opts?.type ?? 'text') == 'button'"
							>
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
								:userInteractionEnabled="false"
								v-else
								:font-size="getOptsCellStyle(index2, index).fontSize"
								:line-height="0"
								:color="getOptsCellStyle(index2, index).fontColor"
								:label="(_rows[index2] ?? [])[index].value ?? '-'"
							></tm-text>
						</tm-col>
					</tm-row>
				</view>
			</scroller>
		</cell>
	</list>
</template>

<script setup lang="ts">
import { computed, ref, Ref, getCurrentInstance, onMounted, nextTick, onBeforeUnmount, PropType, watch } from 'vue'
import type { headresItem, cellItem, dataTypeArray, FixedItemType } from '../interface'
import { tableDataType, tabaleCellStyleType, defaultCellStyle, tabaleCellData, headerType } from './newInterface'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmButton from '../tm-button/tm-button.vue'
import tmRow from '../tm-row/tm-row.vue'
import tmCol from '../tm-col/tm-col.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import { cutomProps } from './cutomProps'
const Binding = uni.requireNativePlugin('bindingx')
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({ ...cutomProps })

let bindingx_token: any = null
const defaultProps = computed(() => {
	return {
		width: props.width,
		height: props.height,
		cellHeight: props.cellHeight,
		headerHeight: props.headerHeight,
		showFixed: props.showFixed
	}
})
const _fixedKey = ref('')
const _showBottomBorder = computed(() => props.showBottomBorder)
const _showHeader = computed(() => props.showHeader)

const emits = defineEmits(['rowClick'])

const _data: Ref<tableDataType> = ref({
	fields: {
		columns: []
	},
	header: [],
	data: []
})
const sortType = ref('none')
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
	// _rows.value = chuliRows(_data.value.data)
	// _rows_back = uni.$tm.u.deepClone(_rows.value)
	setTimeout(function () {
		bindingXscroll()
	}, 200)
})
watch(
	[() => props.tableData],
	() => {
		_data.value = { ...props.tableData }
		_rows.value = chuliRows(uni.$tm.u.deepClone(props.tableData.data))
		_rows_back = uni.$tm.u.deepClone(_rows.value)
	},
	{ deep: true, immediate: true }
)

function chuliRows(
	bigdata: Array<{
		[key: string]: any
	}>
) {
	let d: Array<Array<tabaleCellData>> = []
	let dlen: number[] = []
	let dstyle: Array<
		Array<{
			[key: string]: tabaleCellStyleType
		}>
	> = []
	if (!bigdata) {
		bigdata = []
	}
	_data.value.data = bigdata.map((el) => {
		let ptps: {
			[key: string]: tabaleCellStyleType
		} = el['opts'] ?? {}
		_data.value.header.forEach((ielem) => {
			if (ptps[ielem.field]) {
				ptps[ielem.field] = { ...(ptps[ielem.field] ?? {}) }
			} else {
				ptps[ielem.field] = {}
			}
		})
		el['opts'] = ptps
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

onBeforeUnmount(() => unbindXscroll())

function bindingXscroll() {
	let scroller_right = proxy.$refs.scrollerRight
	let scroller_right_head = proxy.$refs.scrolleHead
	nextTick(() => {
		scroller_right = scroller_right.ref
		scroller_right_head = scroller_right_head.ref
		Binding.bind(
			{
				eventType: 'scroll',
				anchor: scroller_right,
				props: [
					{
						element: scroller_right_head,
						property: 'scroll.contentOffsetX',
						expression: 'x+0'
					}
				]
			},
			(e) => {
				bindingx_token = e.token
			}
		)
	})
}

function unbindXscroll() {
	if (bindingx_token) {
		Binding.unbind({
			token: bindingx_token,
			eventType: 'scroll'
		})
	}
}

function getEl(el: HTMLElement) {
	if (typeof el === 'string' || typeof el === 'number') return el
	if (WXEnvironment) {
		return el.ref
	} else {
		return el instanceof HTMLElement ? el : el.$el
	}
}
</script>

<style></style>

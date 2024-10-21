<template>
	<view>
		<scroll-view
			v-if="showHeader"
			:enable-flex="isNvue"
			:class="[isNvue ? 'flex-row flex' : 'tableHeader relative']"
			:scroll-x="true"
			:scroll-with-animation="false"
			show-scrollbar
			:scroll-y="defaultProps.height > 0"
			:style="[
				defaultProps.height ? { height: `${defaultProps.height}${defaultProps.unit}` } : '',
				{ width: `${defaultProps.width}${defaultProps.unit}` }
			]"
		>
			<!-- #ifndef APP-NVUE -->
			<!--============ S 头部区域 ============= -->
			<view class="flex-1 flex flex-row flex-nowrap absolute zIndex-6" :style="{ height: `${defaultProps.headerHeight}${defaultProps.unit}` }">
				<tm-sheet v-bind="item.headerProps" v-for="(item, index) in columnsList" :key="index" @click="headerClick(item.key, item.sort)">
					<view :style="`width: ${item._headerProps.width}${item.headerProps.unit}`" class="flex flex-row flex-1" :class="[item.align]">
						<view @click.stop="" class="flex-1 flex flex-center" style="width: 0px">
							<tm-text
								@click="headerClick(item.key, item.sort)"
								_style="line-height:normal;"
								:font-size="item._headerProps.fontSize"
								:unit="item.headerProps.unit"
								_class="text-weight-b text-align-center"
								:label="item.title"
							>
							</tm-text>
						</view>
						<view @click.stop="" v-if="item.sort" class="flex flex-col flex-col-center-center">
							<tm-icon
								:line-height="item._headerProps.fontSize * 0.5"
								@click="headerClick(item.key, item.sort)"
								:_class="sortKeys[item.key] == 'asc' || sortKeys[item.key] == 'none' || !sortKeys[item.key] ? '' : 'opacity-6'"
								:font-size="item._headerProps.fontSize * 0.7"
								:unit="item.headerProps.unit"
								name="tmicon-sort-up"
							></tm-icon>
							<tm-icon
								:line-height="item._headerProps.fontSize * 0.5"
								@click="headerClick(item.key, item.sort)"
								:_class="sortKeys[item.key] == 'desc' || sortKeys[item.key] == 'none' || !sortKeys[item.key] ? '' : 'opacity-6'"
								:font-size="item._headerProps.fontSize * 0.7"
								:unit="item.headerProps.unit"
								name="tmicon-sort-down"
							></tm-icon>
						</view>
					</view>
				</tm-sheet>
			</view>
			<!--============ E 头部区域 ============= -->
			<!--============ S 内容区域 ============= -->
			<view :style="{ paddingTop: `${defaultProps.headerHeight}${defaultProps.unit}` }">
				<view class="flex flex-row flex-nowrap" v-for="(cell, cellIndex) in data" :key="cellIndex" :margin="[0, 0]">
					<tm-sheet v-bind="col.cellProps" v-for="(col, colIndex) in columnsList" :key="cellIndex + colIndex">
						<template v-if="col.slot">
							<slot :name="col.slot" :record="cell" :index="cellIndex" :rowIndex="colIndex"></slot>
						</template>
						<tm-text
							v-else
							:fontSize="col._cellProps.fontSize"
							:unit="defaultProps.unit"
							:_class="col.ellipsis ? 'text-overflow-1' : ''"
							:label="cell[col.key] ?? ''"
						/>
					</tm-sheet>
				</view>
			</view>

			<!--============ E 内容区域 ============= -->
			<!-- #endif -->
			<!-- #ifdef APP-NVUE -->
			<!--============ S 头部区域 ============= -->
			<view class="flex-1 flex flex-row flex-nowrap absolute zIndex-6" :style="{ height: `${defaultProps.headerHeight}${defaultProps.unit}` }">
				<tm-sheet v-bind="item.headerProps" v-for="(item, index) in columnsList" :key="index" @click="headerClick(item.key, item.sort)">
					<view
						:userInteractionEnabled="false"
						:style="{
							width: `${item._headerProps.width}${item.headerProps.unit}`,
							height: defaultProps.headerHeight - 6 + 'rpx'
						}"
						class="flex flex-row-center-center flex-row"
						:class="[item.align]"
					>
						<view @click.stop="" class="flex-1 flex" style="width: 0px">
							<tm-text
								@click="headerClick(item.key, item.sort)"
								_style="line-height:normal;"
								:font-size="item._headerProps.fontSize"
								:unit="item.headerProps.unit"
								_class="text-weight-b text-align-center"
								:label="item.title"
							/>
						</view>
						<view @click.stop="" v-if="item.sort" class="flex flex-col flex-col-center-center">
							<tm-icon
								:line-height="item._headerProps.fontSize * 0.5"
								@click="headerClick(item.key, item.sort)"
								:_class="sortKeys[item.key] == 'asc' || sortKeys[item.key] == 'none' || !sortKeys[item.key] ? '' : 'opacity-6'"
								:font-size="item._headerProps.fontSize * 0.7"
								:unit="item.headerProps.unit"
								name="tmicon-sort-up"
							></tm-icon>
							<tm-icon
								:line-height="item._headerProps.fontSize * 0.5"
								@click="headerClick(item.key, item.sort)"
								:_class="sortKeys[item.key] == 'desc' || sortKeys[item.key] == 'none' || !sortKeys[item.key] ? '' : 'opacity-6'"
								:font-size="item._headerProps.fontSize * 0.7"
								:unit="item.headerProps.unit"
								name="tmicon-sort-down"
							></tm-icon>
						</view>
					</view>
				</tm-sheet>
			</view>
			<!--============ E 头部区域 ============= -->
			<!--============ S 内容区域 ============= -->
			<view :style="{ paddingTop: `${defaultProps.headerHeight}${defaultProps.unit}` }">
				<view class="flex flex-row flex-nowrap" v-for="(cell, cellIndex) in data" :key="cellIndex" :margin="[0, 0]">
					<tm-sheet v-bind="col.cellProps" v-for="(col, colIndex) in columnsList" :key="cellIndex + colIndex">
						<template v-if="col.slot">
							<slot :name="col.slot" :record="cell" :index="cellIndex" :rowIndex="colIndex"></slot>
						</template>
						<tm-text
							v-else
							:fontSize="col._cellProps.fontSize"
							:unit="defaultProps.unit"
							:_class="col.ellipsis ? 'text-overflow-1' : ''"
							:label="cell[col.key] ?? ''"
						/>
					</tm-sheet>
				</view>
			</view>
			<!--============ E 内容区域 ============= -->
			<!-- #endif -->
		</scroll-view>
	</view>
</template>
<script lang="ts" setup>
import { computed, nextTick, onMounted, PropType, ref, toRaw, watch } from 'vue'
import { DataTableColumn } from './interface'
import TmSheet from '../tm-sheet/tm-sheet.vue'
import TmText from '../tm-text/tm-text.vue'
import TmIcon from '../tm-icon/tm-icon.vue'
const emits = defineEmits(['sorter'])
const props = defineProps({
	// 是否显示头部
	showHeader: {
		type: Boolean,
		default: true
	},
	columns: {
		type: Array as PropType<Array<DataTableColumn>>,
		default: () => []
	},
	// 表格数据
	data: {
		type: Object || Array,
		default: () => []
	},
	//整体宽度
	width: {
		type: Number
	},
	// 内容高度
	height: {
		type: Number
	},
	//单元格的高度。
	cellHeight: {
		type: Number
	},
	//头部的高度。
	headerHeight: {
		type: Number
	},
	// 计量单位
	unit: {
		type: String,
		default: 'rpx'
	},
	// 显示下划线
	showBottomBorder: {
		type: Boolean,
		default: true
	}
})

const isNvue = ref(false),
	scrollDong = ref(''),
	scrollIndex = ref(0),
	tableLeft = ref(0),
	columnsList = ref([]),
	sortKeys: any = ref({})
// #ifdef APP-NVUE
isNvue.value = true
// #endif

//计算基础属性
const defaultProps = computed(() => {
	if (props.unit == 'px') {
		return {
			width: props.width ?? 750,
			height: props.height ?? 0,
			cellHeight: props.cellHeight ?? 0,
			headerHeight: props.headerHeight ?? 44,
			unit: props.unit
		}
	}
	return {
		width: props.width ?? 750,
		height: props.height ?? 0,
		cellHeight: props.cellHeight ?? 0,
		headerHeight: props.headerHeight ?? 88,
		unit: props.unit ?? 'rpx'
	}
})
// 计算头部内容区域
//===========S 渲染表格 =========//
onMounted(() => {
	nextTick(() => setColData())
})
watch(
	[() => props.data, () => props.columns],
	() => {
		setColData()
	},
	{ deep: true }
)
const _columns = computed(() => {
	return JSON.parse(JSON.stringify(props.columns))
})
const _tableData = computed(() => {
	return props.data
})
function setColData() {
	columnsList.value = []
	//======= S 计算宽度 =========//
	let colWidthNum = 0,
		colWidth = 0,
		colNum = _columns.value.length,
		cellWidth = 0
	_columns.value?.forEach((item: DataTableColumn) => {
		if (item.width) {
			colWidthNum += 1
			colWidth += item.width
		}
	})
	if (defaultProps.value.width > colWidth) {
		cellWidth = parseFloat(String((defaultProps.value.width - colWidth) / (colNum - colWidthNum))).toFixed(2) - 0
	}
	//======= E 计算宽度 =========//
	_columns.value?.forEach((col: DataTableColumn, index: number) => {
		const _minWidth = col.minWidth ?? 0,
			_width = col.width ?? 0
		let _parseWidth = _width
		if (_width === 0) {
			_parseWidth = cellWidth > _minWidth ? cellWidth : _minWidth
		}

		//======= S 头部tSheet props =========//
		if (col?.headerProps) {
			col.headerProps.border = props.showBottomBorder ? 1 : 0
			col.headerProps.color = col.headerProps.color ?? col.bgColor
			col.headerProps.text = col.headerProps.text ?? col.light
			col.headerProps._class = col.headerProps._class ?? `flex-col-center-${col.align || 'center'}`
			col.headerProps.margin = col.headerProps.margin ?? [0, 0]
			col.headerProps.padding = col.headerProps.padding ?? [10, 6]
			col.headerProps.height =
				col.headerProps.height ??
				defaultProps.value.headerHeight -
					(defaultProps.value.unit == 'rpx' ? col.headerProps.padding[1] * 4 : uni.upx2px(col.headerProps.padding[1] * 4))
			col.headerProps.width =
				col.headerProps.width ??
				(_parseWidth == 0 ? 44 : _parseWidth) -
					(defaultProps.value.unit == 'rpx' ? col.headerProps.padding[0] * 2 : uni.upx2px(col.headerProps.padding[0] * 2))
			col.headerProps.unit = col.headerProps.unit ?? defaultProps.value.unit
			col.headerProps.borderDirection = 'bottom'
		} else {
			col.headerProps = {
				border: props.showBottomBorder ? 1 : 0,
				borderDirection: 'bottom',
				color: col.bgColor,
				text: col.light,
				_class: `flex-col-center-${col.align || 'center'}`,
				height: defaultProps.value.headerHeight - (defaultProps.value.unit == 'rpx' ? 24 : uni.upx2px(24)),
				width: _parseWidth - (defaultProps.value.unit == 'rpx' ? 40 : uni.upx2px(40)),
				margin: [0, 0],
				padding: [10, 6],
				unit: defaultProps.value.unit
			}
		}
		col._headerProps = {
			width:
				col.headerProps.width - (col.headerProps.unit == 'rpx' ? col.headerProps.padding[0] * 2 : uni.upx2px(col.headerProps.padding[0] * 2)),
			height:
				col.headerProps.height -
				(col.headerProps.unit == 'rpx' ? col.headerProps.padding[1] * 2 : uni.upx2px(col.headerProps.padding[1] * 2)),
			fontSize: col.headerFontSize ?? (col.headerProps.unit == 'rpx' ? 26 : 14)
		}
		//======= E 头部tSheet props =========//
		//======= S 内容tSheet props =========//
		if (col?.cellProps) {
			col.cellProps.border = props.showBottomBorder ? 1 : 0
			col.cellProps.color = col.cellProps.color ?? col.cellColor
			col.cellProps.text = col.cellProps.text ?? col.light
			col.cellProps._class = col.cellProps._class ?? `flex-col-center-${col.align || 'center'}`
			col.cellProps.margin = col.cellProps.margin ?? [0, 0]
			col.cellProps.padding = col.cellProps.padding ?? [10, 6]
			col.cellProps.height = col.cellProps.height ?? defaultProps.value.cellHeight
			col.cellProps.width = col.cellProps.width ?? (_parseWidth == 0 ? 44 : _parseWidth)
			col.cellProps.unit = col.cellProps.unit ?? defaultProps.value.unit
			col.cellProps.borderDirection = 'bottom'
		} else {
			col.cellProps = {
				border: props.showBottomBorder ? 1 : 0,
				borderDirection: 'bottom',
				color: col.cellColor,
				text: col.light,
				_class: `flex-col-center-${col.align || 'center'}`,
				height:
					defaultProps.value.cellHeight > 0 ? defaultProps.value.cellHeight - (defaultProps.value.unit == 'rpx' ? 24 : uni.upx2px(24)) : 0,
				width: (_parseWidth == 0 ? 44 : _parseWidth) - (defaultProps.value.unit == 'rpx' ? 40 : uni.upx2px(40)),
				margin: [0, 0],
				padding: [10, 6],
				unit: defaultProps.value.unit
			}
		}
		col._cellProps = {
			width: col.cellProps.width - (col.cellProps.unit == 'rpx' ? col.cellProps.padding[0] * 2 : uni.upx2px(col.cellProps.padding[0] * 2)),
			height:
				col.cellProps.height > 0
					? col.cellProps.height - (col.cellProps.unit == 'rpx' ? col.cellProps.padding[1] * 2 : uni.upx2px(col.cellProps.padding[1] * 2))
					: 0,
			fontSize: col.fontSize ?? (col.cellProps.unit == 'rpx' ? 24 : 13)
		}
		//======= E 内容tSheet props =========//
		col.ellipsis = col.ellipsis ?? false
		columnsList.value.push(col)
	})
}
//===========E 渲染表格 =========//
// 头部滚动
function headerScroll(e: any, index: number) {
	if (scrollDong.value != 't') return
	nextTick(() => {
		tableLeft.value = e.detail.scrollLeft
	})
}
function touchStartScroll(index: number) {
	scrollIndex.value = index
	scrollDong.value = 't'
}
// 排序点击事件
function headerClick(key: string, sort: boolean) {
	if (sort) {
		if (!sortKeys.value[key] || sortKeys.value[key] == 'none') {
			sortKeys.value[key] = 'desc'
			dataSort(key, 'desc')
		} else if (sortKeys.value[key] == 'desc') {
			sortKeys.value[key] = 'asc'
			dataSort(key, 'asc')
		} else if (sortKeys.value[key] == 'asc') {
			sortKeys.value[key] = 'none'
			dataSort(key, 'none')
		}
		emits('sorter', sortKeys.value)
	}
}

// 数据排序
function dataSort(key: string, order: string = 'none') {
	uni.showLoading({
		title: '...',
		mask: true
	})
	// 排序
	if (order == 'none' || key === '') {
		setColData()
		uni.hideLoading()
		return
	}

	//降序
	if (order == 'desc') {
		props.data.sort((a, b) => {
			if (typeof a[key] === 'number' && typeof b[key] === 'number') {
				return a[key] - b[key]
			} else if (typeof a[key] === 'string' && typeof b[key] === 'string') {
				return a[key].localeCompare(b[key])
			}
		})
	}
	//升序。
	if (order == 'asc') {
		props.data.sort((b, a) => {
			if (typeof a[key] === 'number' && typeof b[key] === 'number') {
				return a[key] - b[key]
			} else if (typeof a[key] === 'string' && typeof b[key] === 'string') {
				return a[key].localeCompare(b[key])
			}
		})
	}
	nextTick(() => {
		uni.hideLoading()
	})
}
</script>
<style scoped>
.flex-nowrap {
	flex-wrap: nowrap;
}
</style>

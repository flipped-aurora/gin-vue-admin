<template>
	<view class="flex flex-row">
		<pickerPanelVue
			:immediateChange="props.immediateChange"
			:followTheme="props.followTheme"
			@end="emits('end')"
			@start="emits('start')"
			:dataKey="props.mapKey || props.dataKey"
			:mapKey="props.mapKey || props.dataKey"
			@change="pickerChange($event, index)"
			:col="_colIndex[index]"
			v-for="(item, index) in _data"
			:data="item"
			:key="index"
			:height="props.height"
			class="flex-1"
		></pickerPanelVue>
	</view>
</template>
<script lang="ts" setup>
/**
 * 级联选择(滚选)
 * @description 嵌入在页面的级联选择(滚选)
 */
import { computed, PropType, watchEffect, ref, toRaw, onMounted, nextTick, watch, Ref } from 'vue'
import { columnsItem } from './interface'
import pickerPanelVue from './picker-panel.vue'
/**
 * 事件说明
 * change 每一项滑动时会触发事件，并返回父索引levelIndex：当前第几列，子项目索引itemindex：当前列中的第几项被选中。
 * v-model:双向绑定返回当前选中的索引数组。
 * v-model:modelStr 单向输出索引字符串，与v-model区别一个显示索引，一个显示text文本，这个用处主要是用来绑定显示文本用，但与后台交互时可能只需索引。所以它只用来展示用比较方便。
 * start和end分别 为开始选择和结束选择触发的事件，目前在cli 8001结尾的版本（至2022/5/1此事件不会触发系hbx的bug）
 */
const emits = defineEmits(['change', 'update:modelValue', 'update:modelStr', 'end', 'start'])
const props = defineProps({
	followTheme: {
		type: [Boolean],
		default: true
	},
	height: {
		type: Number,
		default: 450
	},
	//可v-model,每一列选中的索引值
	modelValue: {
		type: Array as PropType<Array<number>>,
		default: () => []
	},
	/**
	 * 注意：这里是单向输出显示的value值，而不是modelValue的index索引值。
	 * 这里主要是为了方便表单上页面的显示。如果真要保存到数据库，你应该保存modelValue的值。
	 */
	modelStr: {
		type: [String],
		default: ''
	},
	//默认选中的索引值。
	defaultValue: {
		type: Array as PropType<Array<number>>,
		default: () => []
	},
	columns: {
		type: Array as PropType<Array<columnsItem>>,
		default: () => [],
		required: true
	},
	//当columns项目中的data数据为对象时的key取值字段。
	dataKey: {
		type: String,
		default: 'text'
	},
	//当columns项目中的data数据为对象时的key取值字段。兼容上方dataKey,因为微信dataKey与本字段重名，无法设置。
	mapKey: {
		type: String,
		default: 'text'
	},
	//当前改变index项时，改变时执行的函数。如果返回false，将会阻止本次改变,可以是Promise
	//提供了即将改变的数据和将要改变到目标的数据
	//结构 为 from:{itemindex,levelIndex,data},to:{itemindex,levelIndex,data}。
	beforeChange: {
		type: [Boolean, Function],
		default: () => false
	},
	immediateChange: {
		type: Boolean,
		default: false
	}
})
const _colIndex = ref([...props.defaultValue])
const _data: Ref<Array<Array<columnsItem>>> = ref([])
const _modelStr = computed(() => {
	let str: Array<string | number> = []
	_data.value.forEach((el, index) => {
		let item = el[_colIndex.value[index]]
		if (typeof item == 'undefined') return
		str.push(item[props.mapKey || props.dataKey] ?? '')
	})
	return str.join('/')
})

watch(
	() => _colIndex.value,
	() => {
		nextTick(() => {
			emits('update:modelStr', _modelStr.value)
		})
	},
	{ deep: true }
)
//循环获取子级
function getIndexLoop(defaultindex = 0, data: Array<columnsItem>): Array<Array<columnsItem>> {
	let ds: Array<Array<columnsItem>> = []
	if (data.length == 0) return []
	if (typeof _colIndex.value[defaultindex] == 'undefined') {
		_colIndex.value.push(0)
	}
	let nowData = data[_colIndex.value[defaultindex]]
	if (!nowData) {
		_colIndex.value[defaultindex] = 0
		nowData = data[_colIndex.value[defaultindex]]
	}
	if (nowData && nowData?.children && Array.isArray(nowData?.children) && nowData?.children?.length > 0) {
		ds.push(data)
		let dy = getIndexLoop(defaultindex + 1, nowData?.children)
		ds = [...ds, ...dy]
	} else {
		if (data?.length > 0 && Array.isArray(data) && data) {
			ds.push(data)
		}
	}
	return ds
}
_data.value = getIndexLoop(0, props.columns)

watch(
	() => props.columns,
	() => {
		_data.value = getIndexLoop(0, props.columns)
	},
	{ deep: true }
)
watch(
	() => props.modelValue,
	() => {
		_colIndex.value = props.modelValue
		_data.value = getIndexLoop(0, props.columns)
	},
	{ deep: true }
)
//itemindex 子级索引
//levelIndex 父级索引
async function pickerChange(itemindex: number, levelIndex: number) {
	let isActive = true
	let toItem = _data.value[levelIndex][itemindex]
	const params = {
		from: {
			itemindex: _colIndex.value[levelIndex],
			levelIndex: levelIndex,
			data: _data.value[levelIndex][_colIndex.value[levelIndex]]
		},
		to: { itemindex: itemindex, levelIndex: levelIndex, data: toItem }
	}

	_colIndex.value.splice(levelIndex, 1, itemindex)
	if (typeof props.beforeChange === 'function') {
		uni.showLoading({ title: '...', mask: true })
		let p = await props.beforeChange(params)
		if (typeof p === 'function') {
			p = await p(params)
		}
		if (!p) {
			isActive = false
			nextTick(() => {
				_colIndex.value.splice(levelIndex, 1, params.from.itemindex)
			})
			uni.hideLoading()
		}
	}
	if (toItem?.disabled == true) {
		isActive = false
		nextTick(() => {
			_colIndex.value.splice(levelIndex, 1, params.from.itemindex)
		})
	}

	if (isActive) {
		_data.value = getIndexLoop(0, props.columns)
		emits('change', levelIndex, itemindex)
		emits('update:modelValue', toRaw(_colIndex.value))
	}
}

nextTick(() => {
	emits('update:modelValue', toRaw(_colIndex.value))
	emits('update:modelStr', _modelStr.value || props.modelStr)
})
</script>

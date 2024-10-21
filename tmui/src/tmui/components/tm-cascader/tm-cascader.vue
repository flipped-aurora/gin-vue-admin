<template>
	<tm-sheet :color="props.color" :height="props.height" :padding="[0, 0]">
		<tm-sheet :color="props.color" text :round="3" :border="0" _class="flex flex-row overflow" :height="68" :margin="[0, 0]" :padding="[24, 0]">
			<view style="height: 68rpx" class="flex flex-row flex-row-center-start">
				<view v-for="(item, index) in _idArrays_str" :key="index" class="flex flex-row flex-row-center-start">
					<view>
						<tm-text
							:followTheme="_activeIndex == index ? props.followTheme : false"
							:fontSize="28"
							_class="text-overflow-1"
							:color="_activeIndex == index ? props.activeColor : ''"
							:label="item"
							@click="tabClick(index)"
						></tm-text>
					</view>
					<view class="mx-10" v-if="index < _idArrays.length - 1">
						<tm-text :followTheme="props.followTheme" label="/" :fontSize="24" :color="props.activeColor"> </tm-text>
					</view>
				</view>
				<tm-text
					:followTheme="_activeIndex == _nowLevel ? props.followTheme : false"
					:color="_activeIndex == _nowLevel ? props.activeColor : ''"
					v-if="_nowLevel == _idArrays.length"
					:fontSize="28"
					_class="ml-10"
					label=">请选择"
					@click="tabClick(_nowLevel)"
				></tm-text>
			</view>
		</tm-sheet>

		<view style="height: 16rpx"></view>
		<view v-if="slotTabHeigth" class="flex flex-row flex-row-center-start">
			<slot></slot>
		</view>
		<view style="height: 16rpx"></view>
		<BaseCascader
			:followTheme="props.followTheme"
			:color="props.activeColor"
			:height="props.height - 120 - slotTabHeigth"
			:data="listData"
			:beforeCellClick="props.beforeCellClick"
			@cell-click="emits('cell-click', $event)"
		></BaseCascader>
	</tm-sheet>
</template>
<script lang="ts" setup>
/**
 * 级联选择(点选)
 * @description 级联选择有很多用处，不管是地址选择，城市或者自定数据关联选择，都非常有用处。
 */
import { computed, provide, ref, toRaw, watch, Ref, nextTick, isRef, isProxy, PropType, onMounted } from 'vue'
import BaseCascader from './base-cascader.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import { childrenData } from './interface'
import { getNodeRouterData, getNodeRouterCustomStr, queryNodeIsParent } from './util'
/**
 * 事件说明。
 * @update:modelValue 为v-model
 * @tab-click 点击切换tab触发。
 * @cell-click 点击项目时触发
 * @change 最后选中值时触发，结果同v-model
 */
const emits = defineEmits(['update:modelValue', 'tab-click', 'cell-click', 'change'])
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	/**
	 * 导入的数据
	 */
	data: {
		type: Array as PropType<Array<childrenData>>,
		default: () => [],
		required: true
	},
	/**
	 * 默认选中的数据
	 */
	defaultValue: {
		type: Array as PropType<Array<string>>,
		default: () => []
	},
	/**
	 * 双向绑定输入数据
	 */
	modelValue: {
		type: Array,
		default: () => []
	},
	height: {
		type: Number,
		default: 650
	},
	//激活状态下的颜色。
	activeColor: {
		type: String,
		default: 'primary'
	},
	// 背景主题
	color: {
		type: String,
		default: 'white'
	},
	//点击tab切换之前执行的勾子函数。可以是promise.返回假，即阻切换。真切换正常。方便动态加载数据。
	beforeTabClick: {
		type: [Function, Boolean],
		default: () => {
			return false
		}
	},
	//点击列表中项目时再自动切换到下一项时之前执行的勾子函数，方便动态加载数据。
	beforeCellClick: {
		type: [Function, Boolean],
		default: () => {
			return false
		}
	},
	//介于tab和下面选项中间的插槽区域。如果想自定内容在这这间，可以设置高度，将会显示 。
	slotTabHeigth: {
		type: Number,
		default: 0
	}
})
let listData: Ref<Array<childrenData>> = ref(props.data)
const tmCascaderName = 'tmCascader'
//保存的选中对象数据
let save_value_obj: Ref<Array<childrenData>> = ref([])
//当前选中的id数组。
const _value: Ref<string | number> = ref('')
const _idArrays: Ref<Array<string | number>> = ref([])
const _idArrays_str = ref<string[]>([])
const _nowActiveIndex = ref<number[]>([])

//当前的层级
const _activeIndex = ref(0)
//下一步的层级。
const _nowLevel = ref(0)

//查找节点路径。
// 判断是父节点还是子节点，如果是父或者不存在的节点，不作操作。
function initNode(once: boolean = false) {
	let listdata = uni.$tm.u.deepClone(props.data)
	let moisref = uni.$tm.u.deepClone(props.modelValue)
	if (once) {
		moisref = uni.$tm.u.deepClone(props.defaultValue)
	}
	let ls_ids = Array.isArray(moisref) ? moisref : []
	_value.value = ls_ids.length == 0 ? '' : ls_ids[ls_ids.length - 1]

	if (_value.value == '' || typeof _value.value == 'undefined') return
	let isParent = queryNodeIsParent(listdata, String(_value.value), 'id')

	//父节点不能选中。
	if (isParent) return
	let parentStr = getNodeRouterCustomStr(listdata, ls_ids, 'text')

	_idArrays.value = moisref
	if (parentStr.length > 0) {
		_idArrays_str.value = parentStr.map((el) => el.text)
	}
	_activeIndex.value = _idArrays.value.length ? _idArrays.value.length - 1 : 0
	_nowLevel.value = _activeIndex.value
}

onMounted(() => {
	//初始化
	nextTick(() => initNode(true))
})

provide(
	'tmCascaderValue',
	computed(() => _idArrays.value)
)
provide(
	'tmCascaderShowIndex',
	computed(() => _activeIndex.value)
)

function pushValue(key: childrenData, level: number, id: string | number, keyIndex: number) {
	if (_idArrays.value.length < level) {
		_idArrays.value.push(id)
		_idArrays_str.value.push(key.text)
		_nowActiveIndex.value.push(keyIndex)
	} else {
		_idArrays.value[level] = id
		_idArrays_str.value.splice(level, 1, key.text)
		_nowActiveIndex.value[level] = keyIndex

		nextTick(() => {
			_idArrays.value = _idArrays.value.slice(0, level + 1)
			_idArrays_str.value = _idArrays_str.value.slice(0, level + 1)
			_nowActiveIndex.value = _nowActiveIndex.value.slice(0, level + 1)
		})
	}
}

async function addActiveIndex(level: number, childrenIndex: number, childrenData: childrenData) {
	if (typeof props.beforeCellClick === 'function') {
		uni.showLoading({ title: '...' })
		let p = await props.beforeCellClick(_nowActiveIndex.value, childrenIndex, childrenData)
		if (typeof p === 'function') {
			p = await props.beforeCellClick(_nowActiveIndex.value, childrenIndex, childrenData)
		}
		uni.hideLoading()
		if (!p) return
	}
	_activeIndex.value = level
	_nowLevel.value = level
}

function endSelected() {
	emits('update:modelValue', _idArrays.value)
	nextTick(() => {
		emits('change', _idArrays.value)
	})
}
watch(
	() => props.modelValue,
	() => {
		if (props.modelValue.length == 0) {
			_value.value = ''
			save_value_obj.value = []
			_idArrays.value = []
			_idArrays_str.value = []
			_activeIndex.value = 0
			_nowLevel.value = 0
			return
		}
		initNode()
	},
	{ deep: true }
)
// watch(
//   () => props.data,
//   () => {
//     save_value_obj.value = [];
//     _activeIndex.value = 0;
//     _nowLevel.value = 0;
//     _idArrays.value = [];
// 	_idArrays_str.value = [];
//   },
//   { deep: true }
// );
/**
 * 返回当前选中的数据对象数组。
 */
function getValueObject() {
	let pr = uni.$tm.u.deepClone(_idArrays.value)
	let [lastId] = pr.reverse()
	let ar = getNodeRouterData(uni.$tm.u.deepClone(props.data), String(lastId), [], 'id')
	return ar
}
/**
 * 返回当前选中的数据字符串路径
 */
function getValueStr(): Array<string> {
	let ar: Array<childrenData> = getValueObject()
	let str = ar.map((el) => el.text)
	return str
}

async function tabClick(index: number) {
	if (typeof props.beforeTabClick === 'function') {
		uni.showLoading({ title: '...' })
		let p = await props.beforeTabClick()
		if (typeof p === 'function') {
			p = await p()
		}
		uni.hideLoading()
		if (!p) return
	}
	_activeIndex.value = index
	emits('tab-click', index)
}

function reFresh(data: Array<childrenData> = []) {
	listData.value = []
	if (!Array.isArray(data)) {
		console.error('需要提供数组格式')
		return
	}
	listData.value = data.length > 0 ? data : props.data
}
defineExpose({
	pushValue,
	addActiveIndex,
	tmCascaderName,
	getValueStr,
	getValueObject,
	endSelected,
	reFresh
})
</script>

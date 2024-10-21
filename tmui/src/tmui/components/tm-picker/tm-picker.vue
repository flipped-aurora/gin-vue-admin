<template>
	<view @click="showCity = !props.disabled ? !showCity : false">
		<!-- #ifdef APP-NVUE -->
		<view :eventPenetrationEnabled="true">
			<slot></slot>
		</view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<slot></slot>
		<!-- #endif -->
		<tm-drawer :zIndex="props.zIndex" :inContent="props.inContent" :disabbleScroll="true" :round="props.round"
			ref="drawer" :height="dHeight" :closeable="true" :overlayClick="aniover" @open="drawerOpen" @cancel="cancel"
			@ok="confirm" :show="showCity" @close="closeDrawer" title="请选择" ok-text="确认">
			<tm-picker-view v-if="showCity" :dataKey="props.mapKey || props.dataKey"
				:map-key="props.mapKey || props.dataKey" :height="dHeight - 230" @end="aniover = true"
				@start="aniover = false" :modelValue="_colIndex" @update:modelValue="_colIndex = $event"
				@update:model-str="_colStr = $event" :model-str="_colStr" :default-value="_colIndex"
				:beforeChange="props.beforeChange" :immediateChange="props.immediateChange" :columns="_data">
			</tm-picker-view>
			<tm-button label="确认选择" block :margin="[32, 12]" :color="props.color" :linear="props.linear"
				:linear-deep="props.linearDeep" @click="confirm" :round="props.btnRound">
			</tm-button>
			<view :style="{ height: sysinfo.bottomSafe + 'px' }"></view>
		</tm-drawer>
	</view>
</template>
<script lang="ts" setup>
	/**
 * 级联选择（弹层）
 * @description 这是弹出式级联
 */
	import { PropType, Ref, ref, watchEffect, onMounted, getCurrentInstance, computed, toRaw, inject, nextTick, watch } from 'vue'
	import { custom_props } from '../../tool/lib/minxs'
	import tmDrawer from '../tm-drawer/tm-drawer.vue'
	import { columnsItem } from '../tm-picker-view/interface'
	import tmPickerView from '../tm-picker-view/tm-picker-view.vue'
	import TmSheet from '../tm-sheet/tm-sheet.vue'
	import tmText from '../tm-text/tm-text.vue'
	import tmButton from '../tm-button/tm-button.vue'
	import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

	const drawer = ref<InstanceType<typeof tmDrawer> | null>(null)
	const proxy = getCurrentInstance()?.proxy ?? null

	/**
	 * 事件说明：
	 * v-model:show 双向绑定显示和隐藏选择器
	 * v-model 双向绑定以selectedModel为模式的值。一般用来给数据库后台用的。
	 * v-model:model-str 单向输出地区名称，
	 * 一般用来绑定在Input组件或者其它组件上用来展示当前选择的地址名称，
	 * 但我们服务器可能有的是id或者index，所以你可以v-model绑定数据，用此v-model:model-str来显示数据。
	 * confirm 点击了确认按钮时触发，返回当前选择数据。
	 * cancel 取消或者点击遮罩关闭了选择器都将触发。
	 */
	const emits = defineEmits(['update:show', 'update:modelValue', 'update:modelStr', 'confirm', 'cancel', 'close', 'open'])
	const props = defineProps({
		...custom_props,

		//可v-model,每一列选中的索引值
		modelValue: {
			type: Array as PropType<Array<number | string>>,
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
			type: Array as PropType<Array<number | string>>,
			default: () => []
		},
		//赋值和选值方式
		//name:名称模式赋值和选择
		//id:id模式赋值和选择
		//index:索引模式赋值和选择
		selectedModel: {
			type: String,
			default: 'index'
		},
		//数据。
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
		//v-model:show来双向绑定显示和隐藏选择器。
		show: {
			type: [Boolean],
			default: false
		},
		color: {
			type: String,
			default: 'primary'
		},
		linear: {
			type: String,
			default: ''
		},
		linearDeep: {
			type: String,
			default: 'light'
		},
		btnRound: {
			type: Number,
			default: 0
		},
		round: {
			type: Number,
			default: 12
		},
		height: {
			type: Number,
			default: 700
		},
		immediateChange: {
			type: Boolean,
			default: false
		},
		/** 是否嵌入弹层，开启后将在它的父组件内执行弹层。 */
		inContent: {
			type: Boolean,
			default: false
		},
		/**禁用时，通过插槽点击时，不会触发显示本组件，适合表单 */
		disabled: {
			type: Boolean,
			default: false
		},
		zIndex: {
			type: [Number, String],
			default: 999
		},
		//弹出的动画时间单位ms.
		duration: {
			type: Number,
			default: 300
		}
	})
	const showCity = ref(true)
	const _colIndex : Ref<Array<number>> = ref([])
	const _data = computed(() => props.columns)
	const _colStr = ref(props.modelStr)
	const aniover = ref(true)
	const sysinfo = useWindowInfo()

	let tmid : any = NaN
	watchEffect(() => {
		showCity.value = props.show
	})

	watch(
		[() => props.columns, () => props.modelValue],
		() => {
			clearTimeout(tmid)
			tmid = setTimeout(function () {
				getIndexBymodel(_data.value, props.selectedModel, 0, props.modelValue)
				defaultModerStrGet()
			}, 500)
		},
		{ deep: true }
	)

	function closeDrawer(e : boolean) {
		showCity.value = e
		emits('update:show', showCity.value)
		getIndexBymodel(_data.value, props.selectedModel, 0, props.modelValue)
		emits('close')
	}

	//弹层打开时触发。
	function drawerOpen() {
		emits('open')
	}
	onMounted(() => {
		
		// getIndexBymodel(_data.value, props.selectedModel, 0, props.defaultValue)
		if (props.defaultValue.length > 0) {
			getIndexBymodel(_data.value, props.selectedModel, 0, props.defaultValue)


			defaultModerStrGet()
		}
	})
	//点击确认了地区。
	function confirm() {
		// if (!aniover.value) return
		if (_colIndex.value.length == 0) {
			getIndexBymodelByEmptyOnConfirm(_data.value, props.selectedModel, 0, props.defaultValue)
		}
		setVal()
		nextTick(() => {
			emits('confirm', toRaw(_colIndex.value))
			drawer.value?.close()
		})
	}
	function cancel() {
		if (!aniover.value) return
		emits('cancel')
	}
	function setVal() {
		let val = []
		if (props.selectedModel == 'name') {
			val = _colStr.value.split('/') ?? []
		} else if (props.selectedModel == 'id') {
			val = getRouterId(_data.value, 0)
		} else {
			val = [..._colIndex.value]
		}
		emits('update:modelValue', val)
		emits('update:modelStr', _colStr.value)
	}
	function defaultModerStrGet() {
		clearTimeout(tmid)
		tmid = setTimeout(function () {

			if (!_colStr.value && _colIndex.value.length > 0) {
				let text = getRouterText(_data.value, 0)
				let str = text.join('/')
				emits('update:modelStr', str)

			}
			if (_colIndex.value.length == 0) {
				emits('update:modelStr', '')
			}
		}, 100)
	}
	//模拟模型来返回index值
	function getIndexBymodel(vdata : Array<columnsItem> = [], model = 'name', parentIndex : number = 0, value : Array<number | string> = []) : Array<number> {
		if (value.length == 0) {
			_colIndex.value = []
			return []
		}
		let p_colIndex = [..._colIndex.value]
		_colIndex.value = []


		if (model == 'name') {
			let item = vdata.filter((el) => value[parentIndex] == el['text'])
			if (item.length == 0) {
				//    如果不存在,不再默认选中第一个,2022年10月14日修改
				// item = vdata[0];
				// if (item) {
				//     value[parentIndex] = item['text'];
				//     p_colIndex[parentIndex] = 0
				//     if (item['children']) {
				//         getIndexBymodel(item['children'], model, parentIndex + 1, value);
				//     }
				// }
			} else {
				item = item[0]
				if (item) {
					p_colIndex[parentIndex] = vdata.findIndex((el) => el['text'] == item['text'])
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			}
		} else if (model == 'id') {
			let item = vdata.filter((el) => value[parentIndex] == el['id'])

			if (item.length == 0) {
				//    如果不存在,不再默认选中第一个,2022年10月14日修改
				// item = vdata[0];
				// if (item) {
				//     value[parentIndex] = item['id'];
				//     p_colIndex[parentIndex] = 0
				//     if (item['children']) {
				//         getIndexBymodel(item['children'], model, parentIndex + 1, value);
				//     }
				// }
			} else {
				item = item[0]
				if (item) {
					p_colIndex[parentIndex] = vdata.findIndex((el) => el['id'] == item['id'])
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			}
		} else {
			p_colIndex = [...value]
		}
		_colIndex.value = [...p_colIndex]

		return _colIndex.value
	}
	// 和上方一样的功能,区别是,当什么都不选的时候,点了确认,就要默认选中第一行数据.
	function getIndexBymodelByEmptyOnConfirm(
		vdata : Array<columnsItem> = [],
		model = 'name',
		parentIndex : number = 0,
		value : Array<number | string> = []
	) : Array<number> {
		let p_colIndex = [..._colIndex.value]
		_colIndex.value = []
		if (model == 'name') {
			let item = vdata.filter((el) => value[parentIndex] == el[props.mapKey])
			if (item.length == 0) {
				//    如果不存在,不再默认选中第一个,2022年10月14日修改
				item = vdata[0]
				if (item) {
					value[parentIndex] = item[props.mapKey]
					p_colIndex[parentIndex] = 0
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			} else {
				item = item[0]
				if (item) {
					p_colIndex[parentIndex] = vdata.findIndex((el) => el[props.mapKey] == item[props.mapKey])
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			}
		} else if (model == 'id') {
			let item = vdata.filter((el) => value[parentIndex] == el['id'])
			if (item.length == 0) {
				//    如果不存在,不再默认选中第一个,2022年10月14日修改
				item = vdata[0]
				if (item) {
					value[parentIndex] = item['id']
					p_colIndex[parentIndex] = 0
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			} else {
				item = item[0]
				if (item) {
					p_colIndex[parentIndex] = vdata.findIndex((el) => el['id'] == item['id'])
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			}
		} else {
			p_colIndex = [...value]
		}
		_colIndex.value = [...p_colIndex]
		return _colIndex.value
	}
	//返回 一个节点从父到子的路径id组。
	function getRouterId(list = [], parentIndex = 0) : Array<string | number> {
		let p : Array<string | number> = []
		for (let i = 0; i < list.length; i++) {
			if (i == _colIndex.value[parentIndex]) {
				p.push(list[i]['id'])
				if (typeof _colIndex.value[parentIndex] != 'undefined') {
					let c = getRouterId(list[i]['children'], parentIndex + 1)
					p = [...p, ...c]
				}
				break
			}
		}
		return p
	}
	//返回 一个节点从父到子的路径text组。
	function getRouterText(list : Array<childrenData> = [], parentIndex = 0) : Array<string | number> {
		let p : Array<string | number> = []
		for (let i = 0; i < list.length; i++) {
			if (i == _colIndex.value[parentIndex]) {

				p.push(list[i][props.mapKey || "text"])
				if (typeof _colIndex.value[parentIndex] != 'undefined') {
					let c = getRouterText(list[i]['children'], parentIndex + 1)
					p = [...p, ...c]
				}
				break
			}
		}
		return p
	}
	const dHeight = computed(() => {
		return props.height + sysinfo.bottomSafe + 80
	})
</script>
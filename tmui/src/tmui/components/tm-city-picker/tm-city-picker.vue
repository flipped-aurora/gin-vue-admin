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
		<tm-drawer :disabbleScroll="true" :round="props.round" ref="drawer" :height="820" :closable="true"
			:overlayClick="aniover" @open="drawerOpen" @cancel="cancel" @ok="confirm" :show="showCity"
			@update:show="closeDrawer" title="请选择地区" ok-text="确认" :duration="props.duration">
			<tm-picker-view v-if="showCity" ref="picker" :height="590" @end="aniover = true" @start="aniover = false"
				:value="_colIndex" @update:modelValue="_colIndex = $event" @update:model-str="_colStr = $event"
				:model-str="_colStr" :default-value="_colIndex" :columns="_data"
				:ok-color="props.color"></tm-picker-view>
			<tm-button label="确认选择" block :margin="[32, 12]" :color="props.color" :linear="props.linear"
				:linear-deep="props.linearDeep" @click="confirm" :round="props.btnRound"></tm-button>
			<view :style="{ height: sysinfo.bottomSafe + 'px' }"></view>
		</tm-drawer>
	</view>
</template>
<script lang="ts" setup>
	/**
 * 地区选择器
 * @description 这是弹出式，滑动地址选择器，另还一个按步选择地区的组件，见：tm-city-cascader
 * @example <tm-city-picker v-model:show="show" v-model="status" v-model:model-str="statusw"></tm-city-picker>
 */
	import { PropType, inject, Ref, ref, watchEffect, getCurrentInstance, watch, toRaw, computed, nextTick, onMounted } from 'vue'
	import { custom_props } from '../../tool/lib/minxs'
	import tmDrawer from '../tm-drawer/tm-drawer.vue'
	import { childrenData } from './interface'
	import { provinceData } from '../../tool/static/province'
	import { cityData } from '../../tool/static/city'
	import { areaData } from '../../tool/static/area'
	import tmPickerView from '../tm-picker-view/tm-picker-view.vue'
	import TmSheet from '../tm-sheet/tm-sheet.vue'
	import tmText from '../tm-text/tm-text.vue'
	import tmButton from '../tm-button/tm-button.vue'
	const proxy = getCurrentInstance()?.proxy ?? null
	const drawer = ref<InstanceType<typeof tmDrawer> | null>(null)
	const picker = ref<InstanceType<typeof tmPickerView> | null>(null)
	import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

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
	const emits = defineEmits(['update:show', 'update:modelValue', 'update:modelStr', 'confirm', 'cancel'])
	const props = defineProps({
		...custom_props,
		followTheme: {
			type: [Boolean, String],
			default: true
		},
		//v-model以selectedModel为索引的值结果。
		modelValue: {
			type: Array as PropType<Array<string | number>>,
			default: () => []
		},
		//单向输出地区名称以/分割。不管selectedModel是以哪种索引选项，此处始终以地区名称输出显示。
		//可以使用v-model:modelStr，外部不可更改。
		modelStr: {
			type: String,
			default: ''
		},
		//v-model:show来双向绑定显示和隐藏选择器。
		show: {
			type: [Boolean],
			default: false
		},
		//赋值和选值方式
		//name:名称模式赋值和选择
		//id:地区id模式赋值和选择
		//index:索引模式赋值和选择
		selectedModel: {
			type: String,
			default: 'id'
		},
		/**
		 * 城市选择的级别
		 * province:省级别。
		 * city:省，市
		 * area:省，市，县/区.
		 */
		cityLevel: {
			type: String,
			default: 'area'
		},
		// 手动赋值城市数据
		city: {
			type: Array as PropType<Array<childrenData>>,
			default: () => []
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
			default: 3
		},
		round: {
			type: Number,
			default: 12
		},
		/**禁用时，通过插槽点击时，不会触发显示本组件，适合表单 */
		disabled: {
			type: Boolean,
			default: false
		},
		//弹出的动画时间单位ms.
		duration: {
			type: Number,
			default: 300
		}
	})
	const showCity = ref(true)
	const _cityData = computed(() => props.city)
	const _colIndex : Ref<Array<number>> = ref([])
	const _data = ref(chiliFormatCity_area())
	let tmid : any = NaN
	const _colStr = ref('')
	const aniover = ref(true)
	const sysinfo = useWindowInfo()

	watchEffect(() => {
		showCity.value = props.show
	})
	watch(
		[() => props.city, () => props.modelValue],
		() => {
			_data.value = chiliFormatCity_area()
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
		getIndexBymodel(_data.value, 'index', 0, toRaw(_colIndex.value))
	}
	// onMounted(()=>defaultModerStrGet())
	function drawerOpen() {
		getIndexBymodel(_data.value, props.selectedModel, 0, props.modelValue)
	}
	//点击确认了地区。
	function confirm() {
		if (!aniover.value) return
		setVal()
		nextTick(() => {
			emits('confirm', props.modelValue)
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
			if (_colIndex.value.length > 0) {
				let text = getRouterText(_data.value, 0)
				emits('update:modelStr', text.join('/'))
			}
		}, 200)
	}
	//模拟模型来返回index值
	function getIndexBymodel(
		vdata : Array<childrenData> = [],
		model = 'name',
		parentIndex = 0,
		value : Array<number | string> = []
	) : Array<number | string> {
		if (value.length == 0) return []
		if (model == 'name') {
			let item : any = vdata.filter((el) => value[parentIndex] == el['text'])
			if (item.length == 0) {
				item = vdata[0]
				if (item) {
					value[parentIndex] = item['text']
					_colIndex.value[parentIndex] = 0
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			} else {
				item = item[0]
				if (item) {
					_colIndex.value[parentIndex] = vdata.findIndex((el) => el['text'] == item['text'])
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			}
		} else if (model == 'id') {
			let item : any = vdata.filter((el) => value[parentIndex] == el['id'])
			if (item.length == 0) {
				item = vdata[0]
				if (item) {
					value[parentIndex] = item['id']
					_colIndex.value[parentIndex] = 0
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			} else {
				item = item[0]
				if (item) {
					_colIndex.value[parentIndex] = vdata.findIndex((el) => el['id'] == item['id'])
					if (item['children']) {
						getIndexBymodel(item['children'], model, parentIndex + 1, value)
					}
				}
			}
		}

		return _colIndex.value
	}
	//返回 一个节点从父到子的路径id组。
	function getRouterId(list : Array<childrenData> = [], parentIndex = 0, value : Array<number | string> = []) : Array<string | number> {
		let p : Array<string | number> = []
		const _defalutValue = value.length == 0 ? _colIndex.value : value
		for (let i = 0; i < list.length; i++) {
			if (i == _defalutValue[parentIndex]) {
				p.push(list[i]['id'])
				if (typeof _defalutValue[parentIndex] != 'undefined') {
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
				p.push(list[i]['text'])
				if (typeof _colIndex.value[parentIndex] != 'undefined') {
					let c = getRouterText(list[i]['children'], parentIndex + 1)
					p = [...p, ...c]
				}
				break
			}
		}
		return p
	}
	//格式化数据格式。
	function chiliFormatCity_area() {
		let list : Array<childrenData> = []

		if (_cityData.value.length > 0) {
			return _cityData.value
		}

		provinceData.forEach((item : childrenData, index : number) => {
			list.push({
				id: item.value ?? '',
				text: String(item.label),
				children: []
			})
		})
		if (props.cityLevel == 'province') return list
		cityData.forEach((item : childrenData, index : number) => {
			item.forEach((citem : childrenData, cindex : number) => {
				list[index]?.children.push({
					id: citem.value ?? '',
					text: citem.label ?? '',
					children: []
				})
			})
		})
		if (props.cityLevel == 'city') return list
		list.forEach((item, index) => {
			item.children.forEach((citem, cindex : number) => {
				areaData[index][cindex].forEach((jitem) => {
					list[index]?.children[cindex]?.children.push({
						id: jitem.value ?? '',
						text: jitem.label ?? ''
					})
				})
			})
		})

		return list
	}

	defineExpose({
		getList: chiliFormatCity_area,
		getIndexs: getIndexBymodel,
		getRouterId
	})
</script>
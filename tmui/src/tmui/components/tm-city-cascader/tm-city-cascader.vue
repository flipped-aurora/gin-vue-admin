<template>
	<view>
		<tm-cascader
			:value="_value"
			@update:modelValue="_value = $event"
			@change="changData"
			:default-value="_def"
			:color="props.color"
			:active-color="props.activeColor"
			:height="props.height"
			:slotTabHeigth="88"
			:data="cityAlldata"
		></tm-cascader>
	</view>
</template>
<script lang="ts" setup>
/**
 * 地区级联选择器
 * @description 这是一个城市级联选择器。基于tm-cascader级联选择器
 */
import { Ref, ref, PropType, isProxy } from 'vue'
import { childrenData } from './interface'
import { provinceData } from '../../tool/static/province'
import { cityData } from '../../tool/static/city'
import { areaData } from '../../tool/static/area'
import tmCascader from '../tm-cascader/tm-cascader.vue'
const emits = defineEmits(['update:modelValue'])
const props = defineProps({
	/**
	 * 赋值方式，
	 * id:城市id为返选赋值。
	 * name:以城市名称作为返选和赋值。
	 */
	selectionModel: {
		type: String,
		default: 'name'
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
	/**
	 * 热门城市选择。
	 */
	// hotCity: {
	//   type: Array,
	//   default: () => [
	//     { name: "南昌市", data: ["江西省", "南昌市", "红谷滩区"] },
	//     { name: "杭州市", data: ["浙江省", "杭州市", "余杭区"] },
	//   ],
	// },
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
		type: Array as PropType<Array<string>>,
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
	}
})
const cityAlldata = ref(chiliFormatCity_area())
let _value = ref([])
let _def: Ref<Array<string>> = ref(props.defaultValue)
function changData(e: any) {
	emits('update:modelValue', e)
}

//格式化数据格式。
function chiliFormatCity_area() {
	let list: Array<childrenData> = []
	provinceData.forEach((item: childrenData, index: number) => {
		list.push({
			id: String(props.selectionModel == 'id' ? item.value : item.label),
			text: String(item.label),
			children: []
		})
	})
	if (props.cityLevel == 'province') return list
	cityData.forEach((item: childrenData, index: number) => {
		item.forEach((citem: childrenData, cindex: number) => {
			list[index].children.push({
				id: props.selectionModel == 'id' ? citem.value : citem.label,
				text: citem.label,
				children: []
			})
		})
	})
	if (props.cityLevel == 'city') return list
	list.forEach((item, index) => {
		item.children.forEach((citem, cindex: number) => {
			areaData[index][cindex].forEach((jitem) => {
				list[index].children[cindex].children.push({
					id: props.selectionModel == 'id' ? jitem.value : jitem.label,
					text: jitem.label
				})
			})
		})
	})
	return list
}
</script>

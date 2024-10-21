<template>
	<view class="flex flex-row flex-row-center-start">
		<TmCheckbox
			:followTheme="props.followTheme"
			closeAni
			v-if="showCheckbox"
			:color="nodeData['color']"
			:indeterminate="_indeterminate"
			@change="checkboxChange"
			:disabled="nodeData['disabled']"
			v-model="_checked"
			:value="nodeData[props.fieldNames.id]"
			ref="checkboxRef"
		></TmCheckbox>
	</view>
</template>
<script lang="ts" setup>
import TmCheckbox from '../tm-checkbox/tm-checkbox.vue'
import { ref, getCurrentInstance, inject, computed, toRaw, watch, watchEffect, Ref } from 'vue'
import { baseNodeData } from './interface'
const checkboxRef = ref<InstanceType<typeof TmCheckbox> | null>(null)
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	/**
	 * 数据
	 * @description 生成树结构的数据。结构必须要有id字段。当然可以通过field-names来映射，如果你的唯一标识字段不是Id的话。
	 */
	data: {
		type: Object,
		default: () => {},
		required: true
	},
	/**
	 * 数据结构字段的映射表。
	 */
	fieldNames: {
		type: Object,
		default: () => {
			return {
				id: 'id',
				text: 'text'
			}
		}
	}
})
const emits = defineEmits(['change'])
//父级方法。
let parent: any = proxy?.$parent
while (parent) {
	if (parent?.TreeParaentName == 'tmTree' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
const TreePareantSelectedIds = inject('TreePareantSelectedIds', ref([]))
const color = parent?.$props.color ?? 'primary'
const nodeData = computed(
	() =>
		<baseNodeData>{
			icon: props.data['icon'] ?? '', //节点图标。
			color: props.data['color'] || color, //节点颜色主题
			disabled: props.data['disabled'] ?? false, //节点是否禁用
			text: props.data[props.fieldNames.text], //节点标题
			id: props.data[props.fieldNames.id] //节点标识
		}
)

const showCheckbox = inject(
	'TreeNodeCheckable',
	computed(() => true)
)
//当前是否选中Id;
const _checked: Ref<string | number> = ref('')
//当前是否半选中状态。
const _indeterminate = ref(false)
//是否允许多选。
const _multiple = parent.$props.multiple ?? true

function setChecked() {
	let sletctedIdArray = TreePareantSelectedIds.value.filter((el: string | number) => el == nodeData.value.id)
	if (sletctedIdArray.length >= 0) {
		_checked.value = sletctedIdArray[0]
	} else {
		_checked.value = ''
	}
	emits('change', _checked.value ? true : false)
}
watchEffect(() => setChecked())

function checkboxChange(e: boolean) {
	if (!_multiple && TreePareantSelectedIds.value.length >= 1) {
		parent?.checkAll(false)
	}
	if (e) {
		parent?.onSelected([nodeData.value.id])
	} else {
		parent?.onUnSelected([nodeData.value.id])
	}

	parent?.onCheck({ checked: e, data: toRaw(props.data) })
}

function setStatus() {
	if (nodeData.value['disabled'] === true) {
		return
	}
	checkboxRef.value?.hanlerClick()
}
defineExpose({ setStatus })
</script>

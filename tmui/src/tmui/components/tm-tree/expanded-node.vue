<template>
	<view>
		<view class="flex flex-row flex-row-center-start">
			<TmIcon
				@click="expandHandler(_item[props.fieldNames.id])"
				v-if="_item['children']"
				:font-size="26"
				:name="_checked ? expandedIconOpenIcon : expandedIconCloseIcon"
			></TmIcon>
			<slot></slot>
		</view>
		<view v-if="_item['children'] && _checked">
			<base-node-vue
				@text-click="emits('textClick', $event)"
				:followTheme="props.followTheme"
				:fieldNames="props.fieldNames"
				:data="_item['children']"
			></base-node-vue>
		</view>
	</view>
</template>
<script lang="ts" setup>
import { inject, computed, getCurrentInstance, ref, toRefs, toRaw, Ref, watch, watchEffect } from 'vue'
import TmIcon from '../tm-icon/tm-icon.vue'
import baseNodeVue from './base-node.vue'

import { baseNodeData } from './interface'
const emits = defineEmits(['textClick'])

const { proxy } = getCurrentInstance()
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
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
//父级方法。
let parent = proxy.$parent

while (parent) {
	if (parent?.TreeParaentName == 'tmTree' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
const TreePareantSelectedExpandedId = inject('TreePareantSelectedExpandedId', ref([]))
const expandedIconOpenIcon = parent?.$props.expandedIconOpen ?? 'tmicon-sort-down'
const expandedIconCloseIcon = parent?.$props.expandedIconClose ?? 'tmicon-caret-right'
const _item = computed(() => <baseNodeData>props.data)
const _checked: Ref<string | number> = ref('')
function expandHandler() {
	_checked.value = _checked.value ? '' : _item.value[props.fieldNames.id]
	if (_checked.value) {
		parent?.onExpand({ data: toRaw(_item.value), expand: true })
	} else {
		parent?.onUnExpand({ data: toRaw(_item.value), expand: false })
	}
}
function isCheckedOpen(): void {
	let sAr = TreePareantSelectedExpandedId.value.filter((el: string | number) => el == _item.value[props.fieldNames.id])
	_checked.value = sAr[0] ?? ''
}
watchEffect(() => isCheckedOpen())

function setStatus() {
	expandHandler(_item.value[props.fieldNames.id])
}

defineExpose({ setStatus, filedId: _item.value[props.fieldNames.id] })
</script>

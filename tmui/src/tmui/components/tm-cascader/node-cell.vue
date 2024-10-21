<template>
	<view @click.stop="nodeClick" class="flex px-24 py-24 border-b-1 flex-col-full" :class="[_value['disabled'] ? 'opacity-5' : '']">
		<view :userInteractionEnabled="false" class="flex flex-row flex-between">
			<tm-text :followTheme="isSelected ? props.followTheme : false" :color="isSelected ? props.color : ''" :label="_value.text"></tm-text>
			<view :font-size="30" class="flex flex-row flex-row-center-end">
				<tm-icon
					:followTheme="isSelected ? props.followTheme : false"
					:font-size="28"
					v-if="isSelected"
					:color="isSelected ? props.color : ''"
					name="tmicon-check"
				>
				</tm-icon>
				<tm-icon :font-size="22" v-if="!isSelected && _value['children']" name="tmicon-angle-right" color="grey-1"></tm-icon>
			</view>
		</view>
	</view>
</template>
<script lang="ts" setup>
import { computed, ref, Ref, watch, getCurrentInstance, inject, toRaw, watchEffect, nextTick, PropType } from 'vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import { childrenData } from './interface'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['click'])
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	/**
	 * 导入的数据
	 */
	data: {
		type: Object as PropType<childrenData>,
		default: () => {},
		required: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	level: {
		type: Number,
		default: 0
	},
	keyIndex: {
		type: Number,
		default: 0
	}
})
//父级方法。
let parent: any = proxy?.$parent

while (parent) {
	if (parent?.tmCascaderName == 'tmCascader' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? null
	}
}
const ParentActivedLs = inject(
	'tmCascaderValue',
	computed(() => [])
)
const _value = computed((): childrenData => props.data)
const _activeId: Ref<string | number> = ref('')
function queryNode() {
	let xd = ParentActivedLs.value.filter((el) => el == _value.value.id)
	_activeId.value = xd.length > 0 ? xd[0] : ''
}
watchEffect(() => queryNode())
const isSelected = computed(() => _activeId.value == _value.value.id)
function nodeClick() {
	if (_value.value['disabled']) return
	parent?.pushValue(_value.value, props.level, _value.value.id, props.keyIndex)
	if (typeof _value.value?.children === 'object' && Array.isArray(_value.value?.children) && _value.value?.children.length >= 0) {
		//进入下一步页面操作。

		parent?.addActiveIndex(props.level + 1, props.keyIndex, props.data)
	} else {
		//说明已经到了最后一层，发布确认id请求。
		nextTick(() => {
			parent?.endSelected()
		})
	}

	emits('click')
}
</script>

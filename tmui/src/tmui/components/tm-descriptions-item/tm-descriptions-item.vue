<template>
	<view v-if="_width != '0px' && _width != '0%'" :style="[_width ? { width: _width } : '']" class="flex flex-row flex-row-top-start py-6">
		<view :style="[labelWidth ? { width: labelWidth + 'rpx' } : '']" class="flex-shrink">
			<tm-text _class="opacity-7" :font-size="_fontSize" :label="props.label"></tm-text>
		</view>
		<view class="flex-1 wrap" :style="{ width: '0px' }">
			<slot><tm-text :color="props.color" _class="px-16" :font-size="_fontSize" :label="props.value"></tm-text></slot>
		</view>
	</view>
</template>

<script lang="ts" setup>
import tmText from '../tm-text/tm-text.vue'
import { computed, inject } from 'vue'
const _width: string = inject('tmDescriptionsItem', '')
const labelWidth: string = inject('tmDescriptionsLabelWidth', '')
const props = defineProps({
	label: {
		type: String,
		default: ''
	},
	value: {
		type: String,
		default: ''
	},
	color: {
		type: String,
		default: ''
	},
	fontSize: {
		type: Number,
		default: 23
	}
})
const _fontSize = computed(() => props.fontSize)
</script>

<style></style>

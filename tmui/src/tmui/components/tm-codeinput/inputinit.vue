<template>
	<view ref="dombg">
		<tm-sheet
			:followTheme="props.followTheme"
			:text="false"
			:userInteractionEnabled="false"
			:width="6"
			:color="props.color"
			:height="_size / 2"
			:margin="[0, 0]"
			:padding="[0, 0]"
		>
		</tm-sheet>
	</view>
</template>
<script lang="ts" setup>
import { computed, onMounted, onUnmounted, getCurrentInstance, nextTick } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
// #ifdef APP-PLUS-NVUE
const animation = uni.requireNativePlugin('animation')
// #endif
const { proxy } = getCurrentInstance()
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	size: {
		type: Number,
		default: 50
	}
})
const _size = computed(() => props.size)

onMounted(() => {
	// #ifdef APP-PLUS-NVUE
	nextTick(function () {
		setTimeout(function () {
			spinNvueAni()
		}, 50)
	})
	// #endif
})

function spinNvueAni(opacity = 0) {
	let icon = proxy?.$refs?.dombg
	if (!icon) return
	animation.transition(
		icon,
		{
			styles: {
				opacity: opacity
			},
			duration: 600, //ms
			timingFunction: 'linear',
			delay: 0 //ms
		},
		() => {
			nextTick(function () {
				spinNvueAni(opacity == 0 ? 1 : 0)
			})
		}
	)
}
</script>

<template>
	<view class="flex flex-col relative flex-center">
		<view v-if="!simple" class="flex flex-row relative flex-center" :class="[props.disabled ? 'opacity-5' : '']">
			<tm-sheet
				noLevel
				hover-class=" opacity-5"
				@click="btnClick(item)"
				:width="70"
				:height="70"
				:padding="[0, 0]"
				:round="props.round"
				:shadow="item.page == '...' ? 0 : props.shadow"
				:followTheme="computedCurrent == item.page ? props.followTheme : false"
				:followDark="props.followDark"
				:dark="props.dark"
				:text="props.text"
				:linear="computedCurrent == item.page ? props.linear : ''"
				:linearDeep="props.linearDeep"
				_class="flex-center"
				:transprent="item.page == '...'"
				:color="computedCurrent == item.page ? props.color : props.btnColor"
				:margin="[8, 24]"
				v-for="(item, index) in pageList"
				:key="index"
				><tm-text :userInteractionEnabled="false" _style="line-height:70rpx;cursor: pointer;" :label="String(item.page)"></tm-text
			></tm-sheet>
		</view>
		<view v-if="simple" class="flex flex-row relative flex-center" :class="[props.disabled ? 'opacity-5' : '']">
			<tm-sheet
				noLevel
				hover-class=" opacity-5"
				@click="pnbtn('prev')"
				:width="140"
				:height="70"
				:padding="[0, 0]"
				:round="props.round"
				:shadow="props.shadow"
				:followTheme="false"
				:followDark="props.followDark"
				:dark="props.dark"
				:text="props.text"
				:_class="['flex-center', computedCurrent == 1 ? 'opacity-3' : '']"
				:color="props.btnColor"
				:margin="[10, 24]"
			>
				<tm-icon
					:userInteractionEnabled="false"
					:dark="props.dark"
					color="grey-darken-2"
					:fontSize="24"
					:followTheme="false"
					name="tmicon-angle-left"
				></tm-icon>
			</tm-sheet>
			<tm-text :dark="props.dark" _style="line-height:70rpx;" _class="px-40" :label="`${computedCurrent}/${pages}`"></tm-text>
			<tm-sheet
				noLevel
				hover-class=" opacity-5"
				@click="pnbtn('next')"
				:width="140"
				:height="70"
				:padding="[0, 0]"
				:round="props.round"
				:shadow="props.shadow"
				:followTheme="false"
				:followDark="props.followDark"
				:dark="props.dark"
				:text="props.text"
				:_class="['flex-center', computedCurrent == pages ? 'opacity-3' : '']"
				:color="props.btnColor"
				:margin="[10, 24]"
			>
				<tm-icon
					:userInteractionEnabled="false"
					:dark="props.dark"
					color="grey-darken-2"
					:fontSize="24"
					:followTheme="false"
					name="tmicon-angle-right"
				></tm-icon>
			</tm-sheet>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 分页
 * @description 分页按钮组
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { getCurrentInstance, computed, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch } from 'vue'
const emits = defineEmits(['update:current', 'change'])
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: false
	},
	total: {
		type: Number,
		default: 0,
		required: true
	},
	current: {
		type: Number,
		default: 0
	},
	defaultCurrent: {
		type: Number,
		default: 1
	},
	pageSize: {
		type: Number,
		default: 10
	},
	disabled: {
		type: Boolean,
		default: false
	},
	//简洁模式
	simple: {
		type: Boolean,
		default: false
	},
	//当前容器内的按钮数量，超过将会显示劣略号。
	btnSize: {
		type: Number,
		default: 5
	},
	//当出现省略号时，左右两边出现的个数。
	btnSizeNum: {
		type: Number,
		default: 1
	},
	//主按钮色
	color: {
		type: String,
		default: 'primary'
	},
	//未选中的按钮色
	btnColor: {
		type: String,
		default: 'white'
	},

	followDark: {
		type: [Boolean, String],
		default: true
	},
	//暗黑
	dark: {
		type: [Boolean, String],
		default: false
	},
	round: {
		type: [Number],
		default: 2
	},
	shadow: {
		type: [Number],
		default: 2
	},
	text: {
		type: [Boolean, String],
		default: false
	},
	linear: {
		type: [String],
		default: 'top' //left:右->左，right:左->右。top:下->上，bottom:上->下。
	},
	// 渐变的亮浅
	linearDeep: {
		type: [String],
		default: 'light' //light,dark,accent亮系渐变和深色渐变。
	}
})
const p_current = ref(props.defaultCurrent || 1)
const _pageSize = ref(props.pageSize || 10)
const computedCurrent = computed(() => p_current.value)
const computedPageSize = computed(() => _pageSize.value)
const pages = computed(() => Math.ceil(props.total / computedPageSize.value))
const pageList = computed(() => {
	const pageList = []

	if (pages.value < props.btnSize + props.btnSizeNum * 2) {
		for (let i = 1; i <= pages.value; i++) {
			pageList.push({
				key: i,
				page: i
			})
		}
	} else {
		let left = 1
		let right = pages.value
		let hasLeftEllipsis = false
		let hasRightEllipsis = false

		if (computedCurrent.value > 2 + props.btnSizeNum) {
			hasLeftEllipsis = true
			left = Math.min(computedCurrent.value - props.btnSizeNum, pages.value - 2 * props.btnSizeNum)
		}
		if (computedCurrent.value < pages.value - (props.btnSizeNum + 1)) {
			hasRightEllipsis = true
			right = Math.max(computedCurrent.value + props.btnSizeNum, 2 * props.btnSizeNum + 1)
		}

		if (hasLeftEllipsis) {
			pageList.push({
				key: 1,
				page: 1
			})
			pageList.push({
				key: 'left-more',
				page: '...',
				step: -(props.btnSizeNum * 2 + 1)
			})
		}

		for (let i = left; i <= right; i++) {
			pageList.push({
				key: i,
				page: i
			})
		}

		if (hasRightEllipsis) {
			pageList.push({
				key: 'right-more',
				page: '...',
				step: props.btnSizeNum * 2 + 1
			})
			pageList.push({
				key: pages.value,
				page: pages.value
			})
		}
	}

	return pageList
})
function btnClick(item) {
	if (props.disabled) return
	if (item.page != '...') {
		p_current.value = parseInt(item.page)
		emits('update:current', p_current.value)
		emits('change', p_current.value)
	} else {
		p_current.value = p_current.value + item.step
		emits('update:current', p_current.value)
		emits('change', p_current.value)
	}
}
watch(
	() => props.current,
	() => {
		p_current.value = props.current
	}
)
function pnbtn(type) {
	if (props.disabled) return
	if (type == 'prev') {
		let p = p_current.value - 1
		if (p <= 1) p = 1
		p_current.value = p
		emits('update:current', p_current.value)
		emits('change', p_current.value)
	} else if (type == 'next') {
		let p = p_current.value + 1
		if (p >= pages.value) p = pages.value
		p_current.value = p
		emits('update:current', p_current.value)
		emits('change', p_current.value)
	}
}
</script>

<style>
@import url(./ani.css);
</style>

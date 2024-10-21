<template>
	<view @click="emits('click', $event)" class="flex relative" :class="[props.status ? 'flex-row flex-row-center-center mx-8' : '']">
		<view v-if="!props.status" eventPenetrationEnabled="true">
			<slot></slot>
		</view>
		<view
			eventPenetrationEnabled="true"
			v-if="show"
			:class="[
				(_dot || _count || _icon) && !props.status ? 'absolute flex-top-start-end r-0' : '',
				props.top ? `t-${String(props.top)}` : '',
				props.right ? `r-${String(props.right)}` : ''
			]"
			:style="{ zIndex: 10 }"
		>
			<tm-sheet
				:color="props.color"
				:_class="[customClass, 'flex-center flex-col']"
				:_style="[customCSSStyle, { flexShrink: 1 }]"
				:followTheme="props.followTheme"
				:dark="props.dark"
				:round="props.round"
				:shadow="props.shadow"
				:outlined="props.outlined"
				:border="props.border"
				:borderStyle="props.borderStyle"
				:borderDirection="props.borderDirection"
				:text="props.text"
				:transprent="props.transprent"
				:linear="props.linear"
				:linearDeep="props.linearDeep"
				:width="size.w"
				:height="size.h"
				:margin="props.margin"
				:padding="props.padding"
			>
				<tm-text
					color="white"
					:font-size="props.fontSize"
					:_class="size.h == 0 ? 'py-3 px-6' : ''"
					v-if="_count > 0 && !istext"
					:label="_count > props.maxCount ? props.maxCount + '+' : _count"
				></tm-text>
				<tm-text
					color="white"
					:font-size="props.fontSize"
					:_class="size.h == 0 ? 'py-3 px-6' : ''"
					v-if="_count && istext"
					:label="_count"
				></tm-text>
				<tm-icon color="white" :font-size="props.fontSize" :name="_icon" v-if="_icon"></tm-icon>
			</tm-sheet>
		</view>
		<tm-text eventPenetrationEnabled="true" :font-size="props.fontSize" _class="ml-10" v-if="props.status" :label="props.label"></tm-text>
	</view>
</template>

<script lang="ts" setup>
/**
   * 徽标
   * @description 可单独使用，也可使用插槽。
   * @example <view class="flex-row flex-wrap pa-32" >
      <tm-badge dot color="red">
        <tm-avatar :round="0" label="A"></tm-avatar>
      </tm-badge>
      <tm-badge :count="5"  color="red">
        <tm-avatar label="A"></tm-avatar>
      </tm-badge>
      <tm-badge dot color="red">
        <tm-icon color="primary" name="tmicon-clock-fill"></tm-icon>
      </tm-badge>
      <tm-badge :count="999"  color="red">
        <tm-avatar label="A"></tm-avatar>
      </tm-badge>
      <tm-badge :count="1000"  color="red">
        <tm-avatar label="A"></tm-avatar>
      </tm-badge>
      <tm-badge count="HOT"  color="red">
        <tm-avatar label="A"></tm-avatar>
      </tm-badge>
      <tm-badge icon="tmicon-collection-fill" color="red">
        <tm-avatar label="A"></tm-avatar>
      </tm-badge>
      <tm-badge status dot  label="情况不秒呀." color="green"></tm-badge>
    </view>
   */
import { computed, PropType } from 'vue'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
const emits = defineEmits(['click'])
const props = defineProps({
	...custom_props,
	round: {
		type: [Number],
		default: 6
	},
	border: {
		type: [Number],
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	transprent: {
		type: [Boolean],
		default: false
	},
	label: {
		type: String,
		default: ''
	},
	fontSize: {
		type: Number,
		default: 22
	},
	//为真时，隐藏插槽数据，展现状态文本模式。
	status: {
		type: [Boolean],
		default: false
	},
	dot: {
		type: [Boolean],
		default: false
	},
	icon: {
		type: [String],
		default: ''
	},
	//如果count为数字时，显示数字角标，如果为string是显示文本角标。
	count: {
		type: [Number, String],
		default: 0
	},
	maxCount: {
		type: [Number],
		default: 999
	},
	top: {
		type: [Number],
		default: 0
	},
	right: {
		type: [Number],
		default: 0
	}
})
//自定义样式：
const customCSSStyle = computed(() => computedStyle(props))
//自定类
const customClass = computed(() => computedClass(props))
//让出多少蹑以显示角标。
const istext = computed(() => {
	return isNaN(parseInt(String(props.count)))
})
const show = computed(() => {
	if (!props.dot && !props.icon && !props.count) return false
	return true
})
const size = computed(() => {
	if (props.status || props.dot) {
		return {
			w: 12,
			h: 12,
			pr: 6,
			t: 3
		}
	}
	if (props.icon) {
		let p = props.fontSize * 1.6
		return {
			w: p,
			h: p,
			pr: 12,
			t: 10
		}
	}
	if (isNaN(parseInt(String(props.count)))) {
		return {
			w: 0,
			h: 0,
			pr: 10,
			t: 10
		}
	}
	if (props.count < 10) {
		return {
			w: 30,
			h: 30,
			pr: 12,
			t: 10
		}
	}
	if (props.count >= 10) {
		return {
			w: 0,
			h: 0,
			pr: 10,
			t: 10
		}
	}
	return {
		w: 0,
		h: 0,
		pr: 0,
		t: 0
	}
})
const _icon = computed(() => props.icon)
const _dot = computed(() => props.dot)
const _count = computed(() => props.count)
</script>

<style scoped></style>

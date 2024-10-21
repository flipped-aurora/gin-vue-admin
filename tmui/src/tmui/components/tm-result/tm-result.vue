<template>
	<view @click="onClick" class="flex flex-col flex-col-center-center pa-32">
		<tm-translate :eventPenetrationEnabled="true" name="zoom" :delay="300">
			<tm-sheet
				:dark="props.dark"
				:followTheme="false"
				:followDark="props.followDark"
				_class="flex-center flex-row rounded"
				:width="props.size"
				:height="props.size"
				:round="25"
				:color="icon_color"
				:text="props.text"
			>
				<tm-icon _style="line-height:normal" :dark="props.dark" :followDark="props.followDark" :fontSize="80" :name="icon_rp"></tm-icon>
			</tm-sheet>
		</tm-translate>
		<view :eventPenetrationEnabled="true" class="flex flex-col flex-center pb-10">
			<tm-text :dark="props.dark" :followDark="props.followDark" _class="text-weight-b" :fontSize="34" :label="icon_title"></tm-text>
		</view>
		<view :eventPenetrationEnabled="true" class="flex flex-col flex-center pb-24">
			<tm-text :dark="props.dark" :followDark="props.followDark" _class="opacity-6 " :fontSize="24" :label="icon_subtitle"></tm-text>
		</view>
		<view :eventPenetrationEnabled="true" v-if="props.showBtn" hover-class="opacity-6" class="flex flex-col flex-center">
			<slot name="default">
				<tm-sheet
					:padding="[0, 0]"
					@click="emits('click', $event)"
					:height="80"
					:dark="dark"
					:followTheme="props.followTheme"
					:followDark="props.followDark"
					:shadow="3"
					linear="right"
					:color="icon_color"
					_style="cursor: pointer;"
					:round="4"
					:width="420"
					_class="flex-center"
					:margin="[0, 32]"
				>
					<tm-text
						:userInteractionEnabled="false"
						:dark="props.dark"
						_class="text-size-n"
						:followDark="props.followDark"
						:label="btnText"
					></tm-text>
				</tm-sheet>
			</slot>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 结果页
 * @description 主要用来显示页面当前状态，比如空，出错等。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import { computed } from 'vue'
//click为按钮事件，resultClick为整个组件点击事件。
const emits = defineEmits(['click', 'resultClick'])
const props = defineProps({
	// empty,error,success,warning,lock,network
	status: {
		type: String,
		default: 'empty'
	},
	icon: {
		type: String,
		default: ''
	},
	color: {
		type: String,
		default: 'primary'
	},
	title: {
		type: String,
		default: ''
	},
	subTitle: {
		type: String,
		default: ''
	},
	btnText: {
		type: String,
		default: '确认'
	},
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
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
	//是否显示底部的操作按钮。
	showBtn: {
		type: Boolean,
		default: true
	},
	//是否禁用整个组件的点击事件。
	clickDisabled: {
		type: Boolean,
		default: true
	},
	text: {
		type: Boolean,
		default: true
	},
	size: {
		type: Number,
		default: 140
	}
})
const statusData = {
	empty: {
		icon: 'tmicon-shiliangzhinengduixiang-',
		title: '数据空',
		subTitle: '下拉刷新',
		color: 'primary'
	},
	error: {
		icon: 'tmicon-times',
		title: '错误',
		subTitle: '请根据指示操作',
		color: 'red'
	},
	success: {
		icon: 'tmicon-check',
		title: '操作正确',
		subTitle: '请稍候',
		color: 'green'
	},
	warning: {
		icon: 'tmicon-exclamation-circle',
		title: '异常',
		subTitle: '请根据指示操作',
		color: 'orange'
	},
	lock: {
		icon: 'tmicon-ios-unlock',
		title: '授权提醒',
		subTitle: '需要你的授权进一步操作。',
		color: 'blue'
	},
	network: {
		icon: 'tmicon-wifi-off',
		title: '网络错误',
		subTitle: '请关注你的网络情况',
		color: 'blue'
	}
}
const icon_rp = computed(() => {
	if (props.icon) return props.icon
	if (!props.status) return ''
	return statusData[props.status].icon || ''
})
const icon_title = computed(() => {
	if (props.title) return props.title
	if (!props.status) return ''
	return statusData[props.status].title || ''
})
const icon_subtitle = computed(() => {
	if (props.subTitle) return props.subTitle
	if (!props.status) return ''
	return statusData[props.status].subTitle || ''
})
const icon_color = computed(() => {
	if (props.color) return props.color
	if (!props.status) return ''
	return statusData[props.status].color || ''
})
const onClick = (e: Event) => {
	if (props.clickDisabled) return
	emits('resultClick', e)
}
</script>

<style scoped></style>

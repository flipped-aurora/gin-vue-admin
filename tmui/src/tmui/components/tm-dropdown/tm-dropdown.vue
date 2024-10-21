<template>
	<view class="flex flex-row relative">
		<view
			v-if="show"
			@click.stop="closeDromenu"
			@touchmove.stop=""
			class="l-0 t-0 fixed zIndex-9"
			:style="[
				{
					width: windowWidth + 'px',
					height: windowHeight + 'px',
					background: 'rgba(0,0,0,0)'
				}
			]"
		>
		</view>
		<view
			class="flex flex-col"
			:class="[
				props.position == 'tc' ? 'popover-tc' : '',
				props.position == 'tl' ? 'popover-tl' : '',
				props.position == 'tr' ? 'popover-tr' : '',
				props.position == 'bc' ? 'popover-bc' : '',
				props.position == 'bl' ? 'popover-bl' : '',
				props.position == 'br' ? 'popover-br' : ''
			]"
		>
			<view @click.stop="openDromenu" class="relative zIndex-1 flex flex-row">
				<view :eventPenetrationEnabled="false" :userInteractionEnabled="false" ref="popver" class="flex flex-row"><slot></slot></view>
			</view>
			<view
				v-if="show"
				style="z-index: 400"
				:class="[
					isNvue ? 'fixed' : 'absolute',
					props.position == 'tc' || props.position == 'tl' || props.position == 'tr' ? 'popover-tcc' : '',
					props.position == 'bc' || props.position == 'bl' || props.position == 'br' ? 'popover-bcc' : ''
				]"
				:style="[
					isNvue && props.position == 'tc'
						? {
								top: domNvuePosCss.top - domNvueContentCss.height + 'px',
								left: domNvuePosCss.left + domNvuePosCss.width / 2 - domNvueContentCss.width / 2 + 'px'
						  }
						: '',
					isNvue && props.position == 'tl'
						? {
								top: domNvuePosCss.top - domNvueContentCss.height + 'px',
								left: domNvuePosCss.left + 'px'
						  }
						: '',
					isNvue && props.position == 'tr'
						? {
								top: domNvuePosCss.top - domNvueContentCss.height + 'px',
								left: domNvuePosCss.right - domNvueContentCss.width + 'px'
						  }
						: '',
					isNvue && props.position == 'bc'
						? {
								top: domNvuePosCss.bottom + 'px',
								left: domNvuePosCss.left + domNvuePosCss.width / 2 - domNvueContentCss.width / 2 + 'px'
						  }
						: '',
					isNvue && props.position == 'bl'
						? {
								top: domNvuePosCss.bottom + 'px',
								left: domNvuePosCss.left + 'px'
						  }
						: '',
					isNvue && props.position == 'br'
						? {
								top: domNvuePosCss.bottom + 'px',
								left: domNvuePosCss.right - domNvueContentCss.width + 'px'
						  }
						: ''
				]"
				ref="content"
			>
				<tm-translate ref="aniDom" reverse :name="tarnslateName" :duration="180" :autoPlay="!isNvue">
					<view
						class="flex flex-col"
						:class="[
							props.position == 'tc' ? 'flex-col-center-center' : '',
							props.position == 'tl' ? 'flex-col-top-start' : '',
							props.position == 'tr' ? 'flex-col-bottom-end' : '',
							props.position == 'bc' ? 'flex-col-center-center' : '',
							props.position == 'bl' ? 'flex-col-top-start' : '',
							props.position == 'br' ? 'flex-col-bottom-end' : ''
						]"
						:style="[props.width ? { width: props.width + 'rpx' } : '']"
					>
						<!-- #ifndef APP-NVUE -->
						<tm-sheet
							:color="props.color"
							:_class="props._class"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="0"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
							:text="props.text"
							:transprent="props.transprent"
							:linear="props.linear"
							:linearDeep="props.linearDeep"
							v-if="props.position == 'bc' || props.position == 'bl' || props.position == 'br'"
							:_style="[
								{ zIndex: 1 },
								props.position == 'bc' ? { transform: ' rotate(45deg) translateY(8rpx) translateX(8rpx)' } : {},
								props.position == 'bl' ? { transform: ' rotate(45deg) translateY(-12rpx) translateX(30rpx)' } : {},
								props.position == 'br'
									? {
											transform: ' rotate(45deg) translateY(20rpx) translateX(-16rpx)',
											transformOrigin: '0rpx 50%'
									  }
									: {}
							]"
							:margin="[0, 0]"
							:padding="[0, 0]"
							:width="20"
							:height="20"
						>
						</tm-sheet>
						<tm-sheet
							:color="props.color"
							:_class="props._class"
							:_style="[{ zIndex: 2, position: 'relative', overflow: 'hidden' }]"
							:padding="[0, 0]"
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
							:width="props.width"
							:margin="[0, 0]"
						>
							<view style="height: 24rpx"></view>
							<view
								hover-class="opacity-7"
								@click.stop="onclick(index, item)"
								v-for="(item, index) in listData"
								:key="index"
								class="flex-1 flex px-24 py-16"
							>
								<view class="flex-row flex pb-12 flex-row-center-between">
									<view class="flex flex-row flex-row-center-start flex-1">
										<tm-icon
											:color="item.iconColor"
											:dark="props.dark"
											:fontSize="36"
											_class="pr-12"
											v-if="item.icon"
											:name="item.icon"
										></tm-icon>
										<tm-text :fontSize="32" :label="item.text"></tm-text>
									</view>
									<tm-icon :font-size="22" _class="pl-24" name="tmicon-angle-right" v-if="props.haveArrow"></tm-icon>
								</view>
							</view>
							<view style="height: 8rpx"></view>
						</tm-sheet>
						<tm-sheet
							:color="props.color"
							:_class="props._class"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="0"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
							:text="props.text"
							:transprent="props.transprent"
							:linear="props.linear"
							:linearDeep="props.linearDeep"
							v-if="props.position == 'tc' || props.position == 'tl' || props.position == 'tr'"
							:_style="[
								{ zIndex: 1 },
								props.position == 'tc' ? { transform: ' rotate(45deg) translateY(-8rpx) translateX(-8rpx)' } : {},
								props.position == 'tl' ? { transform: ' rotate(45deg) translateY(-26rpx) translateX(10rpx)' } : {},
								props.position == 'tr'
									? {
											transform: ' rotate(45deg) translateY(0rpx) translateX(-24rpx)',
											transformOrigin: '0rpx 50%'
									  }
									: {}
							]"
							:margin="[0, 0]"
							:padding="[0, 0]"
							:width="20"
							:height="20"
						>
						</tm-sheet>

						<!-- #endif -->
						<!-- #ifdef APP-NVUE -->
						<tm-sheet
							:color="props.color"
							:_class="props._class"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="0"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
							:text="props.text"
							:transprent="props.transprent"
							:linear="props.linear"
							:linearDeep="props.linearDeep"
							v-if="props.position == 'bc' || props.position == 'bl' || props.position == 'br'"
							:_style="[
								{ zIndex: 1 },
								props.position == 'bc' ? { transform: ' rotate(45deg)', bottom: '-10rpx', marginRight: '0rpx' } : {},
								props.position == 'bl' ? { transform: ' rotate(45deg)', bottom: '-12rpx', marginLeft: '12rpx' } : {},
								props.position == 'br'
									? {
											transform: ' rotate(45deg)',
											bottom: '-12rpx',
											marginRight: '12rpx'
									  }
									: {}
							]"
							:margin="[0, 0]"
							:padding="[0, 0]"
							:width="20"
							:height="20"
						>
						</tm-sheet>
						<tm-sheet
							:color="props.color"
							:_class="props._class"
							:padding="[0, 0]"
							:_style="[{ zIndex: 2, position: 'relative' }]"
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
							:width="props.width"
							:margin="[0, 0]"
						>
							<view style="height: 24rpx"></view>
							<view
								hover-class="opacity-7"
								@click.stop="onclick(index, item)"
								v-for="(item, index) in listData"
								:key="index"
								class="flex-1 flex px-24 py-16"
							>
								<view class="flex-row flex pb-12 flex-row-center-between">
									<view class="flex flex-row flex-row-center-start flex-1">
										<tm-icon
											:color="item.iconColor"
											:dark="props.dark"
											:fontSize="36"
											_class="pr-12"
											v-if="item.icon"
											:name="item.icon"
										></tm-icon>
										<tm-text :fontSize="32" :label="item.text"></tm-text>
									</view>
									<tm-icon :font-size="22" _class="pl-24" v-if="props.haveArrow" name="tmicon-angle-right"></tm-icon>
								</view>
							</view>
							<view style="height: 8rpx"></view>
						</tm-sheet>
						<tm-sheet
							:color="props.color"
							:_class="props._class"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="0"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
							:text="props.text"
							:transprent="props.transprent"
							:linear="props.linear"
							:linearDeep="props.linearDeep"
							v-if="props.position == 'tc' || props.position == 'tl' || props.position == 'tr'"
							:_style="[
								{ zIndex: 1 },
								props.position == 'tc' ? { transform: ' rotate(45deg)', top: '-12rpx' } : {},
								props.position == 'tl' ? { transform: ' rotate(45deg)', top: '-12rpx', marginLeft: '12rpx' } : {},
								props.position == 'tr' ? { transform: ' rotate(45deg)', top: '-12rpx', marginRight: '12rpx' } : {}
							]"
							:margin="[0, 0]"
							:padding="[0, 0]"
							:width="20"
							:height="20"
						>
						</tm-sheet>
						<!-- #endif -->
					</view>
				</tm-translate>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 下拉选项
 * @description 用来制作下拉菜单，选项等。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import { itemList } from './interface'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedDark, computedTheme } from '../../tool/lib/minxs'
import { getCurrentInstance, computed, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch, PropType } from 'vue'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const sysinfo = inject(
	'tmuiSysInfo',
	computed(() => {
		return {
			bottom: 0,
			height: 750,
			width: uni.upx2px(750),
			top: 0,
			isCustomHeader: false,
			sysinfo: null
		}
	})
)
const emits = defineEmits(['click'])
const proxy = getCurrentInstance()?.proxy ?? null
const aniDom = ref<InstanceType<typeof tmTranslate> | null>(null)
const props = defineProps({
	...custom_props,
	border: {
		type: [Number],
		default: 0
	},
	round: {
		type: [Number],
		default: 3
	},
	transprent: {
		type: [Boolean],
		default: false
	},
	color: {
		type: String,
		default: 'grey-darken-4'
	},
	width: {
		type: Number,
		default: 0
	},
	position: {
		type: String,
		default: 'bc' //tl,tc,tr,bc,bl,br,上左，中，右。下左，中，右。
	},
	list: {
		type: Array as PropType<Array<itemList>>,
		default: () => [],
		required: true,
		validator: (val) => {
			return typeof val === 'object' && Array.isArray(val)
		}
	},
	rangKey: {
		type: String,
		default: 'text'
	},
	disabled: {
		type: Boolean,
		default: false
	},
	/**是否显示箭头 */
	haveArrow: {
		type: Boolean,
		default: true
	}
})
const windowWidth = computed(() => sysinfo.value.width)
const windowHeight = computed(() => sysinfo.value.height)
let isNvue = ref(false)
// #ifdef APP-PLUS-NVUE
isNvue.value = true
// #endif

let timeid = ref(uni.$tm.u.getUid(5))
let show = ref(false)
let domNvuePosCss = ref({
	left: 0,
	right: 0,
	bottom: 0,
	top: 0,
	width: 0,
	height: 0
})
let domNvueContentCss = ref({
	left: 0,
	right: 0,
	bottom: 0,
	top: 0,
	width: 0,
	height: 0
})
const listData = computed(() => {
	let list = props.list.map((item) => {
		let el: itemList = {
			text: '',
			icon: '',
			iconColor: ''
		}
		if (typeof item === 'string' || typeof item === 'number') {
			el.text = item
		} else {
			el.text = item[props.rangKey]
			el = { ...el, ...item }
		}
		return el
	})
	return list
})
const tarnslateName = computed(() => {
	if (props.position == 'bc' || props.position == 'bl' || props.position == 'br') return 'up'
	return 'down'
})
function nvueDomPos() {
	// #ifdef APP-PLUS-NVUE
	try {
		nextTick(function () {
			dom.getComponentRect(proxy.$refs.popver, function (res) {
				domNvuePosCss.value = { ...res.size }
			})
			dom.getComponentRect(proxy.$refs.content, function (res) {
				if (res?.size) {
					domNvueContentCss.value = { ...res.size }
				}
			})
		})
	} catch (e) {
		//TODO handle the exception
	}
	// #endif
}
onUpdated(() => {
	// #ifdef APP-PLUS-NVUE
	if (domNvuePosCss.value.width == 0 || !domNvueContentCss.value.height) {
		nvueDomPos()
	}
	// #endif
})
onMounted(() => nvueDomPos())
watch(
	() => show.value,
	() => {
		// #ifdef APP-PLUS-NVUE
		clearTimeout(timeid.value)
		if (show.value == true) {
			timeid.value = setTimeout(function () {
				aniDom.value?.play()
			}, 80)
		}
		// #endif
	}
)
function openDromenu() {
	uni.$tm.u.throttle(() => (show.value = true), 200)
}
function closeDromenu() {
	uni.$tm.u.debounce(() => (show.value = false), 250)
}
function onclick(index: number, data: itemList) {
	emits('click', { index: index, data: data })
	show.value = false
}
</script>

<style scoped>
/* #ifdef APP-PLUS-NVUE */
.popover-tcc {
	transform: translateY(-15rpx);
}
.popover-bcc {
	transform: translateY(15rpx);
}
.popover-tr {
	display: flex;
	justify-content: flex-start;
	align-items: flex-start;
}
/* #endif */
/* #ifndef APP-PLUS-NVUE */
.popover-tc {
	display: flex;
	justify-content: flex-end;
	align-items: center;
}
.popover-tcc {
	transform: translateY(-35rpx);
}
.popover-bcc {
	transform: translateY(35rpx);
}
.popover-tl {
	display: flex;
	justify-content: flex-end;
	align-items: flex-start;
}
.popover-tr {
	display: flex;
	justify-content: flex-end;
	align-items: flex-end;
}

.popover-bc {
	display: flex;
	justify-content: flex-start;
	align-items: center;
}
.popover-bl {
	display: flex;
	justify-content: flex-start;
	align-items: flex-start;
}
.popover-br {
	display: flex;
	justify-content: flex-start;
	align-items: flex-end;
}
/* #endif */
</style>

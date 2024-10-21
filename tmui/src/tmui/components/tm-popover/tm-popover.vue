<template>
	<view class="flex flex-row relative" @click.stop="">
		<view
			v-if="show"
			@click.stop="show = false"
			class="l-0 t-0 fixed zIndex-9"
			:style="[
				{
					width: windowWidth + 'px',
					height: windowHeight + 'px',
					background: 'rgba(0,0,0,0)'
				}
			]"
		></view>
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
			<view @click="show = true" class="relative zIndex-1 flex flex-row">
				<view :userInteractionEnabled="false" :eventPenetrationEnabled="true" ref="popver" class="flex flex-row"><slot></slot></view>
			</view>
			<view
				v-if="show"
				class="zIndex-10"
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
				<tm-translate ref="aniDom" reverse :name="tarnslateName" :duration="120" :autoPlay="!isNvue">
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
						<!-- #ifndef APP-PLUS-NVUE -->
						<tm-sheet
							:text="props.text"
							:color="props.color"
							_class="flex-col flex-col flex-col-center-center"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="0"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
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
						></tm-sheet>
						<tm-sheet
							@click="show = false"
							:width="props.width"
							:text="props.text"
							:color="props.color"
							:_class="props._class"
							:_style="[{ zIndex: 2, position: 'relative' }]"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="props.round"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
							:transprent="props.transprent"
							:linear="props.linear"
							:linearDeep="props.linearDeep"
							:margin="[0, 0]"
							:padding="[0, 12]"
						>
							<slot name="label">
								<tm-text
									:font-size="24"
									_style="line-height:normal;"
									:_class="[props.width ? '' : 'nowrap  ', 'px-16']"
									:label="props.label"
								></tm-text>
							</slot>
						</tm-sheet>
						<tm-sheet
							:text="props.text"
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
						></tm-sheet>

						<!-- #endif -->
						<!-- #ifdef APP-NVUE -->
						<tm-sheet
							:text="props.text"
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
						></tm-sheet>
						<tm-sheet
							@click="show = false"
							:width="props.width"
							:text="props.text"
							:color="props.color"
							_class="flex-col flex-col flex-col-center-center"
							:_style="[{ zIndex: 2, position: 'relative' }]"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:round="props.round"
							:shadow="props.shadow"
							:outlined="props.outlined"
							:border="props.border"
							:borderStyle="props.borderStyle"
							:borderDirection="props.borderDirection"
							:transprent="props.transprent"
							:linear="props.linear"
							:linearDeep="props.linearDeep"
							:margin="[0, 0]"
							:padding="[16, 12]"
						>
							<view :eventPenetrationEnabled="true">
								<slot name="label">
									<tm-text
										:font-size="24"
										_style="line-height:normal;"
										:_class="[props.width ? '' : 'nowrap  ', 'px-16']"
										:label="props.label"
									></tm-text>
								</slot>
							</view>
						</tm-sheet>
						<tm-sheet
							:text="props.text"
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
						></tm-sheet>
						<!-- #endif -->
					</view>
				</tm-translate>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 汽泡卡片
 * @description 用来提示，帮助展示信息等。
 */

import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import { custom_props } from '../../tool/lib/minxs'
import { getCurrentInstance, computed, ref, PropType, inject, onUpdated, onMounted, onUnmounted, nextTick, watch } from 'vue'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const aniDom = ref<InstanceType<typeof tmTranslate> | null>(null)
const props = defineProps({
	...custom_props,
	shadow: {
		type: Number,
		default: 0
	},
	border: {
		type: [Number, String],
		default: 0
	},
	round: {
		type: [Number, String],
		default: 3
	},
	transprent: {
		type: [Boolean, String],
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	width: {
		type: Number,
		default: 0
	},
	/**tl,tc,tr,bc,bl,br,上左，中，右。下左，中，右。 */
	position: {
		type: String as PropType<'tl' | 'tc' | 'tr' | 'bc' | 'bl' | 'br'>,
		default: 'tc'
	},
	label: {
		type: String,
		default: '提示内容'
	},
	/**
	 * 默认是否显示。
	 */
	defaultShow:{
		type:Boolean,
		default:false
	}
})
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

const tarnslateName = computed(() => {
	if (props.position == 'bc' || props.position == 'bl' || props.position == 'br') return 'up'
	return 'down'
})
function nvueDomPos() {
	
	try {
		nextTick(function () {
			dom.getComponentRect(proxy?.$refs.popver, function (res) {
				domNvuePosCss.value = { ...res.size }
				show.value = props.defaultShow;
				dom.getComponentRect(proxy?.$refs.content, function (res) {
					if (res?.size) {
						domNvueContentCss.value = { ...res.size }
					}
				})
				
			})

		})
	} catch (e) {
		//TODO handle the exception
	}
	
}
onUpdated(() => {
	// #ifdef APP-PLUS-NVUE
	if (domNvuePosCss.value.width == 0 || !domNvueContentCss.value.height) {
		nvueDomPos()
	}
	// #endif
})
onMounted(() => {
	// #ifdef APP-PLUS-NVUE
	nvueDomPos()
	// #endif
	// #ifndef APP-PLUS-NVUE
	show.value = props.defaultShow;
	// #endif
})
watch(
	() => show.value,
	() => {
		// #ifdef APP-PLUS-NVUE
		clearTimeout(timeid.value)
		if (show.value == true) {
			setTimeout(function () {
				aniDom.value?.play()
			}, 80)
		}
		// #endif
	}
)
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

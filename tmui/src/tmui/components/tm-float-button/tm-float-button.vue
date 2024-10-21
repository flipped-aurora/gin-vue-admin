<template>
	<view class="fixed flex" :style="[_pos?.parent,{zIndex:402}]">
		<!-- 主按钮。 -->
		<view :style="{ width: props.width + 'rpx', height: props.height + 'rpx' }" class="flex flex-row flex-row-center-center">
			<slot>
				<tm-button
					:followTheme="props.followTheme"
					@click="onclick"
					_class="flex flex-col flex-col-center-center"
					:shadow="3"
					:linear="_btn.linear"
					:linear-deep="_btn.linearDeep"
					:color="_btn.color"
					:margin="[0, 0]"
					:round="16"
					:padding="[0, 0]"
					:width="props.width - 24"
					:height="props.height - 24"
					:openType="_btn.openType"
					@getphonenumber="btnCall($event,_btn)"
					@error="btnCall($event,_btn)"
					@opensetting="btnCall($event,_btn)"
					@launchapp="btnCall($event,_btn)"
					@contact="btnCall($event,_btn)"
					@chooseavatar="btnCall($event,_btn)"
				>
					<view class="flex flex-col flex-col-center-center">
						<tm-icon
							:userInteractionEnabled="false"
							:follow-dark="false"
							:color="_btn.fontColor"
							:name="_btn.icon"
							:font-size="_btn.iconSize"
						></tm-icon>
						<tm-text
							:userInteractionEnabled="false"
							:follow-dark="false"
							:color="_btn.fontColor"
							v-if="_btn.label"
							:label="_btn.label"
							:font-size="_btn.fontSize"
						></tm-text>
					</view>
				</tm-button>
			</slot>
		</view>
		<!-- 子按钮。 -->
		<view
			v-if="showActions && _actionsItem.length > 0"
			class="fixed zIndex-12"
			:style="_pos?.children"
			:userInteractionEnabled="showActions"
		>
			<view
				ref="btnChildren"
				:style="{ width: props.width + 'rpx', height: props.height + 'rpx' }"
				:class="[showActions && _actionsItem.length > 0 ? 'scaleNvue' : '']"
				class="flex flex-row flex-row-center-center scale"
				v-for="(item, index) in _actionsItem"
				:key="index"
			>
				<tm-button
					:followTheme="props.followTheme"
					@click="change(index, item)"
					_class="flex flex-col flex-col-center-center"
					:shadow="3"
					:linear="item.linear"
					:linear-deep="item.linearDeep"
					:color="item.color"
					:margin="[0, 0]"
					:round="16"
					:padding="[0, 0]"
					:width="props.width - 24"
					:height="props.height - 24"
					:openType="item.openType"
					@getphonenumber="btnCall($event,item)"
					@error="btnCall($event,item)"
					@opensetting="btnCall($event,item)"
					@launchapp="btnCall($event,item)"
					@contact="btnCall($event,item)"
					@chooseavatar="btnCall($event,item)"
				>
					<view class="flex flex-col flex-col-center-center">
						<tm-icon
							:userInteractionEnabled="false"
							:follow-dark="false"
							:color="item.fontColor"
							:name="item.icon"
							:font-size="item.iconSize"
						></tm-icon>
						<tm-text
							:userInteractionEnabled="false"
							:follow-dark="false"
							:color="item.fontColor"
							v-if="item.label"
							:label="item.label"
							:font-size="item.fontSize"
						></tm-text>
					</view>
				</tm-button>
			</view>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 悬浮按钮
 * @description 总共6个位置，每个位置有四方向的子按钮展开位置，一共24个位置可控制。具体看文档。
 * @example <tm-floatButton :btn="{icon:'tmicon-plus',linear:'top'}"></tm-floatButton>
 * @method click 主按钮被点击，(e:Event)
 * @method change  子按钮被点击， (index:number,item:actionsItem)
 */
import { computed, PropType, ref, inject, getCurrentInstance, onMounted } from 'vue'
import { positionType, popDir, actionsItem } from './interface'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmButton from '../tm-button/tm-button.vue'
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

// #ifdef APP-PLUS-NVUE
const animation = uni.requireNativePlugin('animation')
// #endif
/**
 * 事件说明
 * click：主按钮被点击，
 * change :子按钮被点击，
 */
const emits = defineEmits(['click', 'change'])
const props = defineProps({
	followTheme: {
		type: [Boolean],
		default: true
	},
	//主按钮的位置
	position: {
		type: String as PropType<'bc' | 'bl' | 'br' | 'tc' | 'tl' | 'tr'>,
		default: 'br'
	},
	//子菜单弹出的位置
	actionsPos: {
		type: String as PropType<'left' | 'right' | 'top' | 'bottom'>,
		default: 'top'
	},
	width: {
		type: Number,
		default: 124
	},
	height: {
		type: Number,
		default: 124
	},
	offset: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 32]
	},
	//子按钮组数据
	actions: {
		type: Array as PropType<Array<actionsItem>>,
		default: () => []
	},
	// 主按钮对象数据
	btn: {
		type: Object as PropType<actionsItem>,
		default: () => {},
		required: true
	},
	//是否默认显示子菜单
	showActions: {
		type: Boolean,
		default: false
	},
	//点击子菜单后，是否需要隐藏，如果为false,点击子按钮后不会隐藏按钮。始终保持展开子按钮。
	clickHidnActions: {
		type: Boolean,
		default: true
	},
	disabledScrollTo: {
		type: Boolean,
		default: true
	},
	scrollTo: {
		type: Object as PropType<{ scrollTop: number; selector: string; duration: number }>,
		default: () => {
			return {
				scrollTop: 0,
				selector: '',
				duration: 300
			}
		}
	}
})

const sysinfo = useWindowInfo()
const windowWidth = computed(() => sysinfo.width)
const windowTop = computed(() => sysinfo.top)
const proxy = getCurrentInstance()?.proxy ?? null
const isH5 = ref(false)
// #ifdef H5
isH5.value = true
// #endif
const showActions = ref(props.showActions ?? false)
const BtnPos = computed(() => props.position)
const AcionPos = computed(() => props.actionsPos)
const _offset = computed(() => {
	let ost = props.offset ?? [0, 0]
	ost = [uni.upx2px(props.offset[0]), uni.upx2px(props.offset[1])]
	return ost
})
const centerPosLeft = computed(() => {
	let ps = (windowWidth.value - uni.upx2px(props.width * 2)) / 2 + _offset.value[0] * 2
	return ps
})

const _btn = computed(() => {
	return {
		icon: 'tmicon-plus',
		fontSize: 20,
		color: 'primary',
		linear: '',
		linearDeep: 'accent',
		label: '',
		iconSize: 42,
		fontColor: '',
		...(props.btn ?? {})
	}
})
const _actionsItem = computed(() => {
	let asbtn = props.actions.map((el) => {
		let default_btn: actionsItem = {
			icon: 'tmicon-plus',
			fontSize: 20,
			color: 'primary',
			linear: '',
			linearDeep: 'accent',
			label: '',
			fontColor: '',
			iconSize: 36
		}
		return { ...default_btn, ...el }
	})
	return asbtn
})

const _pos = computed(() => {
	let gutter = uni.upx2px(24)
	let actionwidth_total = _actionsItem.value.length * uni.upx2px(props.width)
	let actionwidth = uni.upx2px(props.width)
	if (AcionPos.value == 'left' && BtnPos.value == 'bc') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				bottom: _offset.value[1] + 'px',
				left: centerPosLeft.value - actionwidth_total + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'right' && BtnPos.value == 'bc') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				bottom: _offset.value[1] + 'px',
				left: centerPosLeft.value + actionwidth + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'bottom' && BtnPos.value == 'bc') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				bottom: _offset.value[1] - actionwidth_total + 'px',
				left: centerPosLeft.value + 'px'
			}
		}
	}
	if (AcionPos.value == 'top' && BtnPos.value == 'bc') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				bottom: _offset.value[1] + actionwidth + 'px',
				left: centerPosLeft.value + 'px'
			}
		}
	}

	if (AcionPos.value == 'right' && BtnPos.value == 'bl') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] + 'px',
				left: _offset.value[0] + actionwidth + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'left' && BtnPos.value == 'bl') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] + 'px',
				left: _offset.value[0] - actionwidth_total + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'top' && BtnPos.value == 'bl') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] + actionwidth + 'px',
				left: _offset.value[0] + 'px'
			}
		}
	}
	if (AcionPos.value == 'bottom' && BtnPos.value == 'bl') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] - actionwidth_total + 'px',
				left: _offset.value[0] + 'px'
			}
		}
	}

	if (AcionPos.value == 'right' && BtnPos.value == 'br') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] + 'px',
				right: _offset.value[0] - actionwidth_total + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'left' && BtnPos.value == 'br') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] + 'px',
				right: _offset.value[0] + actionwidth + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'top' && BtnPos.value == 'br') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] + actionwidth + 'px',
				right: _offset.value[0] + 'px'
			}
		}
	}
	if (AcionPos.value == 'bottom' && BtnPos.value == 'br') {
		return {
			parent: {
				bottom: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				bottom: _offset.value[1] - actionwidth_total + 'px',
				right: _offset.value[0] + 'px'
			}
		}
	}

	if (AcionPos.value == 'left' && BtnPos.value == 'tc') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				top: _offset.value[1] + 'px',
				left: centerPosLeft.value - actionwidth_total + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'right' && BtnPos.value == 'tc') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				top: _offset.value[1] + 'px',
				left: centerPosLeft.value + actionwidth + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'bottom' && BtnPos.value == 'tc') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				top: _offset.value[1] + actionwidth + 'px',
				left: centerPosLeft.value + 'px'
			}
		}
	}
	if (AcionPos.value == 'top' && BtnPos.value == 'tc') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: centerPosLeft.value + 'px'
			},
			children: {
				top: _offset.value[1] - actionwidth_total + 'px',
				left: centerPosLeft.value + 'px'
			}
		}
	}

	if (AcionPos.value == 'right' && BtnPos.value == 'tl') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] + 'px',
				left: _offset.value[0] + actionwidth + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'left' && BtnPos.value == 'tl') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] + 'px',
				left: _offset.value[0] - actionwidth_total + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'top' && BtnPos.value == 'tl') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] - actionwidth_total + 'px',
				left: _offset.value[0] + 'px'
			}
		}
	}
	if (AcionPos.value == 'bottom' && BtnPos.value == 'tl') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				left: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] + actionwidth + 'px',
				left: _offset.value[0] + 'px'
			}
		}
	}

	if (AcionPos.value == 'right' && BtnPos.value == 'tr') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] + 'px',
				right: _offset.value[0] - actionwidth_total + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'left' && BtnPos.value == 'tr') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] + 'px',
				right: _offset.value[0] + actionwidth + 'px',
				display: 'flex',
				'flex-direction': 'row'
			}
		}
	}
	if (AcionPos.value == 'top' && BtnPos.value == 'tr') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] - actionwidth_total + 'px',
				right: _offset.value[0] + 'px'
			}
		}
	}
	if (AcionPos.value == 'bottom' && BtnPos.value == 'tr') {
		return {
			parent: {
				top: _offset.value[1] + 'px',
				right: _offset.value[0] + 'px'
			},
			children: {
				top: _offset.value[1] + actionwidth + 'px',
				right: _offset.value[0] + 'px'
			}
		}
	}
})

function onclick(e: any) {
	if (props.clickHidnActions) {
		showActions.value = !showActions.value
		// #ifdef APP-NVUE
		setTimeout(function () {
			nvueani()
		}, 50)
		// #endif
	} else {
		showActions.value = true
	}
	emits('click', e)
	// #ifndef APP-NVUE
	if (!props.disabledScrollTo) {
		uni.pageScrollTo({
			scrollTop: props.scrollTo.scrollTop,
			duration: props.scrollTo.duration,
			selector: props.scrollTo.selector
		})
	}
	// #endif
}
onMounted(() => {
	if (showActions.value) {
		// #ifdef APP-NVUE
		setTimeout(function () {
			nvueani()
		}, 50)
		// #endif
	}
})
function change(index: number, item: actionsItem) {
	if (props.clickHidnActions) {
		showActions.value = false
	}
	emits('change', index, item)
}

function btnCall(event:any,item:actionsItem){
	if(typeof item?.callback === 'function'){
		item.callback(event,item)
	}
}

function nvueani() {
	let el = proxy?.$refs?.btnChildren ?? null
	if (!el) return
	let pel = []
	if (!Array.isArray(el)) {
		pel = [el]
	} else {
		pel = [...el]
	}
	pel.forEach((element) => {
		animation.transition(
			element,
			{
				styles: {
					opacity: 1,
					transform: 'scale(1,1)',
					transformOrigin: 'center center'
				},
				duration: 340, //ms
				timingFunction: 'ease',
				delay: 0 //ms
			},
			() => {}
		)
	})
}
</script>
<style scoped>
/* #ifdef APP-NVUE */
.scaleNvue {
	opacity: 0;
	transform: scale(0);
}
/* #endif */
/* #ifndef APP-NVUE */
.scale {
	animation: scale 0.34s ease;
}
@keyframes scale {
	0% {
		transform: scale(0);
	}

	40% {
		transform: scale(1.2);
	}
	100% {
		transform: scale(1);
	}
}
/* #endif */
</style>

<template>
	<view
		v-if="isActive && FilterMenuMaxHeight > 0 && !_props.isButton"
		@click.stop="close"
		:style="{ height: FilterMenuMaxHeight + 'px' }"
		class="overflow"
	>
		<view ref="FilterMenuItemBody" class="FilterMenuItemBody">
			<tm-divider :margin="[0, 0]"></tm-divider>
			<tm-sheet :eventPenetrationEnabled="true" :color="props.bgColor" :margin="[0, 0]" :padding="[24, 24]">
				<scroll-view @click.stop="" scroll-y :style="{ height: props.height + 'rpx' }">
					<view @click.stop="stopEvent">
						<slot></slot>
					</view>
				</scroll-view>
				<view v-if="props.footerHeight > 0" :style="{ height: props.footerHeight + 'rpx' }" @click.stop="stopEvent">
					<slot name="footer"></slot>
				</view>
			</tm-sheet>
		</view>
	</view>
</template>
<script lang="ts" setup>
import tmSheet from '@/tmui/components/tm-sheet/tm-sheet.vue'
import tmText from '@/tmui/components/tm-text/tm-text.vue'
import tmIcon from '@/tmui/components/tm-icon/tm-icon.vue'
import tmDivider from '@/tmui/components/tm-divider/tm-divider.vue'
import { computed, getCurrentInstance, onMounted, onUnmounted, inject, ref, watchEffect, ComputedRef, nextTick, watch } from 'vue'
import { FilterMenuType } from '../tm-filterMenu/interface'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null

const props = defineProps({
	title: {
		type: String,
		default: '',
		require: true
	},
	icon: {
		type: String,
		default: 'tmicon-angle-up'
	},
	unIcon: {
		type: String,
		default: 'tmicon-angle-down'
	},
	/**内容高度,非菜单栏高度 */
	height: {
		type: Number,
		default: 500
	},
	/**内容高度,非菜单栏高度 */
	footerHeight: {
		type: Number,
		default: 0
	},
	unit: {
		type: String,
		default: 'rpx'
	},
	fontSize: {
		type: Number,
		default: 28
	},
	fontColor: {
		type: String,
		default: 'primary'
	},
	unFontColor: {
		type: String,
		default: 'black'
	},
	/**是否当作按钮使用,点击时,只触发事件,不展开相关内容的动作. */
	isButton: {
		type: Boolean,
		default: false
	},
	bgColor: {
		type: String,
		default: 'white'
	},
	uuid: {
		type: [String, Number],
		default: ''
	}
})
const _props = computed(() => props)
let tid: any = NaN
const activeIndex = inject(
	'activeIndex',
	computed(() => 0)
)
const FilterMenuMaxHeight = inject(
	'FilterMenuMaxHeight',
	computed(() => 0)
)
const AllList: ComputedRef<FilterMenuType[]> = inject(
	'AllList',
	computed(() => [])
)
const id = 'filterItem_' + uni.$tm.u.getUid(1)
const uuid = computed(() => props.uuid)
const isActive = ref(false)
//父级方法。
let parent: any = proxy?.$parent
while (parent) {
	if (parent?.FilterMenu == 'FilterMenu' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
pushKey()
watch(
	() => _props.value,
	() => {
		clearTimeout(tid)
		tid = setTimeout(function () {
			pushKey()
		}, 150)
	},
	{ deep: true }
)
onMounted(() => {
	setTimeout(() => {
		nextTick(() => {
			isActive.value = getActive()
			watch([() => activeIndex.value, AllList], () => {
				isActive.value = getActive()
			})
		})
	}, 200)
})
// #ifdef APP-NVUE
watch(isActive, () => {
	if (isActive.value) {
		nextTick(() => {
			showNvueAniMation()
		})
	}
})
// #endif

function getActive() {
	let index = AllList.value.findIndex((el) => el.id == id)
	return index === activeIndex.value
}
function pushKey() {
	if (parent) {
		parent?.pushKey({
			id: id,
			height: _props.value.height,
			unit: _props.value.unit,
			icon: _props.value.icon,
			unIcon: _props.value.unIcon,
			fontSize: _props.value.fontSize,
			fontColor: _props.value.fontColor,
			unFontColor: _props.value.unFontColor,
			text: _props.value.title,
			isButton: _props.value.isButton,
			uuid: uuid.value
		})
	}
}

onUnmounted(() => {
	if (parent) {
		parent?.delKey(id)
	}
})

function stopEvent(e: any) {
	e?.preventDefault()
	e?.stopPropagation()
}

function close(e: any) {
	e?.preventDefault()
	e?.stopPropagation()
	if (parent) {
		parent?.close()
	}
}

function showNvueAniMation() {
	var el = proxy.$refs?.FilterMenuItemBody
	animation.transition(
		el,
		{
			styles: {
				transform: 'translateY(-100%)',
				opacity: 0.6
			},
			duration: 1, //ms
			timingFunction: 'ease',
			delay: 0 //ms
		},
		() => {
			animation.transition(
				el,
				{
					styles: {
						transform: 'translateY(0%)',
						opacity: 1
					},
					duration: 300, //ms
					timingFunction: 'ease',
					delay: 0 //ms
				},
				() => {}
			)
		}
	)
}
</script>

<style scoped>
.FilterMenuItemBody {
	/* #ifndef APP-NVUE */
	animation: scaleTop 0.3s;
	/* #endif */
	/* #ifdef APP-NVUE */
	transform: translateY(-100%);
	opacity: 0.4;
	/* #endif */
}

/* #ifndef APP-NVUE */
@keyframes scaleTop {
	0% {
		opacity: 0.4;
		transform: translateY(-100%);
	}

	100% {
		opacity: 1;
		transform: translateY(0%);
	}
}

/* #endif */
</style>

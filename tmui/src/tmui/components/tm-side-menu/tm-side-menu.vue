<template>
	<!-- 侧边栏菜单 导航 。 -->
	<view class="flex flex-row">
		<tm-sheet :color="props.menuBarBgColor" :height="_height" :width="_sliderMenuWidth" :unit="props.unit" :margin="[0, 0]" :padding="[0, 0]">
			<scroll-view :show-scrollbar="false" scroll-y :style="{ height: _height + 'rpx', width: _sliderMenuWidth + 'rpx' }">
				<tm-sheet
					v-for="(item, index) in _list"
					:key="index"
					:margin="[0, 0]"
					:padding="[0, 0]"
					:color="props.menuActiveBgColor"
					:transprent="_active - 1 !== index && _active + 1 !== index"
					no-level
				>
					<tm-sheet
						no-level
						@click="menuClick(item, index)"
						:color="_active === index ? props.menuActiveBgColor : props.menuBarBgColor"
						:isDisabledRoundAndriod="true"
						_class="flex flex-row flex-row-center-center relative"
						:margin="[0, 0]"
						:height="_sliderMenuItemHeight"
						:padding="[0, 0]"
						:width="_sliderMenuWidth"
						:class="(_active - 1 === index ? 'round-br-6' : '') + ' ' + (_active + 1 === index ? 'round-tr-6' : '')"
						:parenClass="(_active - 1 === index ? 'round-br-6' : '') + ' ' + (_active + 1 === index ? 'round-tr-6' : '')"
						borderDirection="bottom"
					>
						<view :userInteractionEnabled="false" class="absolute l-6" :style="{ top: (_sliderMenuItemHeight - 30) / 2 + 'rpx' }">
							<tm-sheet
								:round="2"
								v-if="_active == index"
								:height="30"
								:margin="[0, 0]"
								:padding="[0, 0]"
								:color="props.menuActiveFontColor"
								:width="6"
							></tm-sheet>
						</view>
						<tm-badge v-if="!item.icon" :userInteractionEnabled="false" color="red" :count="item.dotCount">
							<view class="flex flex-col flex-col-center-center py-10" :style="{ width: `${_sliderMenuWidth - 28}rpx` }">
								<tm-text
									:_class="'text-align-center ' + (_active == index ? 'text-weight-b' : '')"
									:color="_active == index ? props.menuActiveFontColor : props.menuFontColor"
									:font-size="props.fontSize"
									:label="item.text"
								></tm-text>
							</view>
						</tm-badge>
						<tm-badge v-if="item.icon" :userInteractionEnabled="false" color="red" :count="item.dotCount">
							<view class="flex flex-row flex-row-center-center py-10" :style="{ width: `${_sliderMenuWidth - 28}rpx` }">
								<tm-icon
									v-if="item.icon"
									:color="_active == index ? props.menuActiveFontColor : props.menuFontColor"
									:font-size="props.fontSize"
									:name="item.icon"
								></tm-icon>
								<view :userInteractionEnabled="false" class="px-10">
									<tm-text
										:_class="'text-align-center ' + (_active == index ? 'text-weight-b' : '')"
										:color="_active == index ? props.menuActiveFontColor : props.menuFontColor"
										:font-size="props.fontSize"
										:label="item.text"
									></tm-text>
								</view>
							</view>
						</tm-badge>
					</tm-sheet>
				</tm-sheet>
			</scroll-view>
		</tm-sheet>
		<tm-sheet no-level class="flex-1" parenClass="flex-1" :show-scrollbar="false" :margin="[0, 0]" :padding="[0, 0]" :color="props.bodyBgColor">
			<scroll-view
				:refresher-triggered="_refresh"
				@scrolltolower="onBottom"
				@refresherrefresh="onRefresh"
				@refresherrestore="onRestore"
				@refresherabort="onAbort"
				:refresher-enabled="!_disabledPull"
				scroll-y
				class="flex-1"
				:style="{ height: _height + 'rpx' }"
				v-if="!props.isScroll"
			>
				<slot name="default"></slot>
				<view v-if="_pullStauts == 'bottom' && _isLoadding" class="my-40 flex flex-row flex-row-center-center">
					<tm-icon spin name="tmicon-shuaxin" :font-size="22" color="primary"></tm-icon>
					<tm-text :font-size="22" color="grey" _class="pl-20" label="加载中..."></tm-text>
				</view>
			</scroll-view>
			<view v-if="props.isScroll">
				<slot name="default"></slot>
			</view>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
import tmText from '@/tmui/components/tm-text/tm-text.vue'
import tmIcon from '@/tmui/components/tm-icon/tm-icon.vue'
import tmSheet from '@/tmui/components/tm-sheet/tm-sheet.vue'
import tmBadge from '@/tmui/components/tm-badge/tm-badge.vue'
import { ref, getCurrentInstance, nextTick, Ref, computed, PropType, watch } from 'vue'
import { propsCustom } from './props'
interface listItem {
	text: string
	dotCount?: number | string
	[prop: string]: any
}
/**事件 */
const emits = defineEmits<{
	/** 点击菜单时触发，返回索引和项目数据 */
	(e: 'click', item: listItem, index: number): void
	/** 索引切换时触发 */
	(e: 'change', index: number): void
	/** 索引切换时触发 */
	(e: 'update:active', index: number): void
	/** 数据加载时触发 type:'pull'|'bottom'|'menu',item:listItem|null=null,index:number=NaN */
	(e: 'load', type: 'loading' | 'success' | 'pull' | 'bottom' | 'menu', index: number): void
	/** 重置下拉时触发 */
	(e: 'restore'): void
}>()

const props = defineProps({ ...propsCustom })
const slidWidthPx = uni.upx2px(props.sideWidth)
const _list = computed(() => {
	return props.list.map((el) => {
		return {
			...el,
			text: el?.text ?? el[props.rangKey] ?? '',
			dotCount: el?.dotCount ?? 0,
			icon: el?.icon ?? ''
		}
	})
})
/**设置当前下拉刷新状态，true 表示下拉刷新已经被触发，false 表示下拉刷新未被触发 */
const _refresh: Ref<boolean | string> = ref(true)
const _isLoadding = ref(false) //是否刷新中.
const _pullStauts: Ref<'pull' | 'bottom' | 'menu' | ''> = ref('')
const _disabledPull = computed(()=>props.disabledPull)
const _disabledBottom = computed(()=>props.disabledBottom)
const _active = ref(props.active)
const _sliderMenuWidth = ref(props.unit == 'px' ? uni.$tm.u.torpx(props.sideWidth) : props.sideWidth)
const _sliderMenuItemHeight = ref(props.unit == 'px' ? uni.$tm.u.torpx(props.itemHeight) : props.itemHeight)
const _width = computed(() => {
	return props.unit == 'px' ? uni.$tm.u.torpx(props.width) : props.width
})
const _height = computed(() => {
	return props.unit == 'px' ? uni.$tm.u.torpx(props.height) : props.height
})
watch(
	() => props.active,
	() => {
		if (props.active == _active.value) return
		_active.value = props.active
	}
)
watch(_active, () => {
	emits('update:active', _active.value)
	emits('change', _active.value)
})
const menuClick = async (item: listItem, index: number) => {
	_active.value = index
	nextTick(async () => {
		emits('click', item, index)
		emits('load', 'loading')
		_isLoadding.value = true
		_pullStauts.value = 'menu'
		await actionLoad('menu', item, index)
		_isLoadding.value = false
		emits('load', 'success')
	})
}
async function onRefresh() {
	if(_disabledPull.value) return;
	if (_isLoadding.value) {
		_refresh.value = false
		return
	}
	emits('restore')
	emits('load', 'loading')
	_isLoadding.value = true
	_pullStauts.value = 'pull'
	await actionLoad('pull')
	// _refresh.value = false;
	_refresh.value = true
	/** 我解释下下面为什么要如此操作：原因是在微信小程序端，部分小米手机会出现无限下拉的情况，我初步判断是手机性能不行，导致原生的下拉复位还没反应过来，接着又执行了其它函数。 */
	setTimeout(function () {
		_isLoadding.value = false
		_refresh.value = false
	}, 300)
}
async function onBottom() {
	if (_isLoadding.value || _disabledPull.value) return
	emits('load', 'loading')
	_isLoadding.value = true
	_pullStauts.value = 'bottom'
	await actionLoad('bottom')
	emits('load', 'success')
	_isLoadding.value = false
}

async function actionLoad(type: 'pull' | 'bottom' | 'menu', item: listItem | null = null, index: number = NaN) {
	let fun = props.load
	let p: any = true
	// uni.showLoading({title:'...',mask:true})
	if (typeof fun === 'function') {
		p = await fun(type, item, index)
		if (typeof p === 'function') {
			await p(type, item, index)
		}
	}

	// uni.hideLoading()
}

function onAbort() {
	_refresh.value = false
}
function onRestore() {
	_refresh.value = false
	// console.error('restore')
}
</script>

<style></style>

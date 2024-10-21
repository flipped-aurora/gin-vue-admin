<template>
	<tm-drawer
		:round="props.round"
		ref="drawer"
		:height="dHeight"
		@update:show="_show = $event"
		:show="_show"
		@close="close"
		:ok-color="props.color"
		@open="open"
		:hide-header="true"
		:closable="true"
		:foot-height="footerBarHeight"
	>
		<template v-slot:default>
			<view class="pa-24">
				<view class="flex flex-row flex-row-center-start">
					<tm-image :preview="true" v-if="nowItemInfo" :width="120" :height="120" :src="nowItemInfo.img"></tm-image>
					<view class="pl-24">
						<view class="flex flex-row flex-row-center-start">
							<tm-text :font-size="24" :color="props.color" label="￥"></tm-text>
							<tm-text :font-size="42" _class="text-weight-b" :color="props.color" :label="nowItemInfo?.salePrice ?? '0'"></tm-text>
							<tm-sheet linear="left" linear-deep="accent" :color="props.color" :margin="[24, 0]" :padding="[18, 4]" :round="24">
								<view class="flex flex-row flex-row-center-start">
									<tm-text :font-size="24" label="优惠价￥"></tm-text>
									<tm-text :font-size="42" _class="text-weight-b" :label="nowItemInfo?.price ?? '0'"></tm-text>
								</view>
							</tm-sheet>
						</view>
						<tm-text
							color="grey"
							:font-size="24"
							:label="nowItemInfo && nowItemInfo?.num > 0 ? '有货' + ' | 库存 ' + nowItemInfo?.num : '无货'"
						></tm-text>
						<tm-text color="grey" :font-size="24" :label="nowItemInfo?.title ?? ''"></tm-text>
					</view>
				</view>
				<tm-divider :margin="[0, 24]"></tm-divider>
				<view class="mb-24">
					<view @click="addNumberClick" class="flex flex-row flex-row-center-between mb-24">
						<tm-text :font-size="28" label="购买数量"></tm-text>
						<tm-stepper
							@change="numberChange"
							:max="maxCount"
							:disabled="nowInputNumber > maxCount"
							v-model="nowInputNumber"
							:default-value="nowInputNumber"
						></tm-stepper>
					</view>
					<tm-divider :margin="[0, 0]"></tm-divider>
				</view>
				<view v-if="_list">
					<view class="" v-for="(item, index) in _list.data" :key="index">
						<tm-text _class="mb-24" :font-size="28" :label="item.title"></tm-text>
						<tm-radio-group v-model="nowSelected[index]" v-if="item?.children" direction="row">
							<tm-radio @change="radioClick(item2,index)" :disabled="!checkPropsStock(item2,index)" :value="item2.id" v-for="(item2, index2) in item.children" :key="index2" :custom="false">
								<template v-slot:default="{ checked }">
									<tm-badge :count="!checkPropsStock(item2,index)?'缺货':0">
										<view :class="[!checkPropsStock(item2,index)?'opacity-6':'','']">
											<tm-tag
												:shadow="0"
												:color="checked.checked&&checkPropsStock(item2,index)?'red':'grey'"
												:round="24"
												:font-size="26"
												size="n"
												outlined
												text
												:label="item2.title"
											></tm-tag>
										</view>
									</tm-badge>
								</template>
							</tm-radio>
						</tm-radio-group>
						<tm-divider :margin="[0, 24]"></tm-divider>
					</view>
				</view>
			</view>
			<slot></slot>
		</template>
		<template v-slot:foot>
			<view class="mb-20 px-24">
				<view style="height: 40rpx" class="flex flex-row flex-row-center-center">
					<tm-text v-if="nowItemInfo?.tip" :color="props.color" :label="nowItemInfo.tip + '，'"></tm-text>
					<tm-text v-if="nowInputNumber > maxCount && nowInputNumber && nowItemInfo" :color="props.color" label="库存不足"></tm-text>
				</view>
				<view class="flex flex-row">
					<view style="width: 363rpx">
						<tm-button
							@click="addGou"
							block
							:is-disabled-round-andriod="true"
							_class="round-l-24 round-r-0"
							linear="left"
							linear-deep="accent"
							:color="props.color"
							:font-size="32"
							label="加购物车"
							:height="80"
						></tm-button>
					</view>
					<view style="width: 363rpx">
						<tm-button
							@click="buyGou"
							block
							:disabled="!nowItemInfo || nowInputNumber == 0 || nowInputNumber > maxCount || !nowItemInfo?.num"
							:is-disabled-round-andriod="true"
							_class="round-r-24 round-l-0"
							linear="left"
							linear-deep="accent"
							:color="props.color"
							:font-size="32"
							:label="
								!nowItemInfo?.num
									? '缺货，提醒我'
									: '购买' + (nowItemInfo?.salePrice ? '￥' + nowItemInfo?.salePrice * nowInputNumber : '')
							"
							:height="80"
						></tm-button>
					</view>
				</view>
			</view>
			<view :style="{ height: sysinfo.bottomSafe + 'px' }"></view>
		</template>
	</tm-drawer>
</template>
<script lang="ts" setup>
import { computed, ref, inject, PropType, toRaw, nextTick, watch, onMounted } from 'vue'
import tmDrawer from '../tm-drawer/tm-drawer.vue'
import tmButton from '../tm-button/tm-button.vue'
import tmText from '../tm-text/tm-text.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import tmTag from '../tm-tag/tm-tag.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmImage from '../tm-image/tm-image.vue'
import tmStepper from '../tm-stepper/tm-stepper.vue'
import tmRadioGroup from '../tm-radio-group/tm-radio-group.vue'
import tmRadio from '../tm-radio/tm-radio.vue'
import tmBadge from '../tm-badge/tm-badge.vue'
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

const emits = defineEmits(['open', 'close', 'update:show', 'add', 'buy'])
const props = defineProps({
	round: {
		type: Number,
		default: 6
	},
	show: {
		type: Boolean,
		default: false
	},
	color: {
		type: String,
		default: 'red'
	},
	height: {
		type: Number,
		default: 650
	},
	//购物车数量
	num: {
		type: Number,
		default: 1
	},
	list: {
		type: Object as PropType<Tmui.sku>,
		default: (): any => null
	},
	/** 
	 * 当前选中的id
	 * id是以-来连接的，与数据结构中的条目id一样.比如a-b-b
	 *  */
	value:{
		type:String,
		default:""
	}
})
const sysinfo = useWindowInfo()
const _show = ref(props.show)
const footerBarHeight = ref(160)
const dHeight = computed(() => {
	return props.height + footerBarHeight.value + uni.$tm.u.torpx(sysinfo.bottomSafe)
})
const nowSelectedItem: any = ref(null)
const nowSelected = ref<string[]>([])

const nowInputNumber = ref(props.num)
let textcm = toRaw(props.list)
const _list = ref(props.list)
const maxCount = computed(() => {
	if (nowSelected.value.length == 0) return 0
	if (!props.list) return 0
	if (!props.list?.data) return 0
	if (props.list.data.length == 0) return 0
	if (!props.list?.product?.length) return 0
	let item = props.list?.product.filter((el) => el.id == nowSelected.value.join('-'))
	if (!item.length) return 0
	return Math.min(item[0].max_buy,item[0].num) || 0
})
const nowItemInfo = computed(() => {
	if (!props.list?.product || !props.list.data) return null
	let item = props.list?.product.filter((el) => el.id == nowSelected.value.join('-'))
	if (!item.length) return null
	return item[0]
})

function init() {
	if (!props.list) return
	let MaxLen = props.list.data.length
	nowSelected.value = new Array(MaxLen).fill('')
}
onMounted(()=>{
	init()
	setValueToSelected();
})
function setValueToSelected(){
	let p = props.value.split("-")
	let len = props.list?.data?.length??0
	for(let i=0;i<len;i++){
		if(i>p.length-1){
			break;
		}
		nowSelected.value[i] = p[i]
	}
}
watch(
	() => props.num,
	() => {
		nowInputNumber.value = props.num
	}
)
watch(
	() => props.value,
	() => {
		setValueToSelected();
	}
)
watch(
	() => props.list,
	() => {
		textcm = uni.$tm.u.deepClone(toRaw(props.list))
		_list.value = textcm
		init()
	},
	{ deep: true }
)
watch(
	() => props.show,
	() => {
		_show.value = props.show
	}
)
//加入购物车触发。
function addGou() {
	if (nowSelected.value.length !== props.list?.data?.length || nowSelectedIsEmpty()) {
		uni.showToast({ title: '未选择完整', icon: 'none' })
		return
	}
	emits('add', {
		buyNumber: nowInputNumber.value,
		data: toRaw(nowItemInfo.value)
	})
	nextTick(() => {
		_show.value = false
	})
}
function nowSelectedIsEmpty(){
	let empty = false;
	for(let key=0; key < nowSelected.value.length;key++){
		if(nowSelected.value[key]===""||nowSelected.value[key]==null||nowSelected.value[key]==undefined){
			empty = true;
			break;
		}
	}
	return empty;
}
//购买时触发。
function buyGou() {
	if (nowSelected.value.length !== props.list?.data?.length ||  nowSelectedIsEmpty()) {
		uni.showToast({ title: '未选择完整', icon: 'none' })
		return
	}
	emits('buy', {
		buyNumber: nowInputNumber.value,
		data: toRaw(nowItemInfo.value)
	})
	nextTick(() => {
		_show.value = false
	})
}
function numberChange(num: number) {}
function addNumberClick() {
	if (nowSelected.value.length !== props.list?.data?.length ||  nowSelectedIsEmpty()) {
		uni.showToast({ title: '未选择完整', icon: 'none' })
		return
	}
}

// 根据选择自动计算属性库存
const checkPropsStock = computed(()=>{
    return (e:any,index:number)=>{
        // 复制选择的数组
        let selectIds = nowSelected.value.concat();
        // 是否选中此属性,没选中则加入传入的属性,检测是否有库存
        selectIds[index]=e.id;

        let stocked:any=false;
        //是否有选中项
        const selectCount=nowSelected.value.filter(value=>value).length;
        if(selectCount>0){
            stocked=props.list?.product.some((item)=>{
                const dataIds=item.id.split("-");
                for (let x = 0; x < selectIds.length; x++) {
                    if(!selectIds[x]) dataIds[x]="";
                }
                return dataIds.join("-")===selectIds.join("-") && item.num>0
            })      
        } else {
            // 属性ID是否在库存表中
            stocked=props.list?.product.some((item)=>{
                return item.id.split("-").includes(e.id,index)
            })
        }
        return e.num || stocked;
    }
})

function radioClick(item2:any,index:number){
	console.log(checkPropsStock.value(item2,index))
	
	
}

function close() {
	emits('close')
	emits('update:show', false)
}
function open() {
	emits('open')
	emits('update:show', true)
}
</script>
<style></style>

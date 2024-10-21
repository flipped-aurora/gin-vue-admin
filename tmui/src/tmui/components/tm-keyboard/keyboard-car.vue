<template>
	<tmSheet
		:follow-theme="false"
		:follow-dark="false"
		:dark="_dark"
		color="white"
		:transprent="true"
		:padding="[4, 4]"
		:margin="[0, 0]"
		_class="flex flex-col"
		paren-class="flex-1"
	>
		<view class="flex-center flex-row" style="height: 62rpx">
			<tm-text v-if="!_value && !props.showInputConten" :font-size="28" _class="text-weight-b" :label="props.title"></tm-text>
			<tm-text v-if="_value && props.showInputContent" :font-size="34" _class="text-weight-b pr-24" :label="_value"></tm-text>
		</view>
		<view class="flex flex-row">
			<view class="flex-9 flex flex-col">
				<view class="flex-row flex flex-1" v-for="(item2, index2) in numberArray" :key="index2">
					<tmSheet
						hover-class="opacity-5 keywordBoradAni"
						no-level
						@click="keydown(item)"
						:follow-theme="false"
						:follow-dark="false"
						:dark="_dark"
						:round="2"
						:height="90"
						_class="flex-center"
						:padding="[0, 0]"
						:margin="[4, 4]"
						v-for="(item, index) in item2"
						:key="index"
						paren-class="flex-1"
						class="flex-1"
					>
						<view style="width: 40rpx" class="flex flex-center flex-row">
							<tmText :userInteractionEnabled="false" :font-size="32" _class="text-weight-b" :label="item"></tmText>
						</view>
					</tmSheet>
				</view>
			</view>
			<view class="flex-1 flex flex-col">
				<tmSheet
					hover-class="opacity-5 keywordBoradAni"
					no-level
					:height="90"
					@click="keydown('学')"
					:follow-theme="false"
					:follow-dark="false"
					:dark="_dark"
					:round="2"
					_class="flex-center"
					:padding="[0, 0]"
					:margin="[4, 4]"
					class="flex flex-col"
					paren-class="flex-row flex-center"
				>
					<tmText :userInteractionEnabled="false" :font-size="32" _class="text-weight-b" label="学"></tmText>
				</tmSheet>
				<tmSheet
					hover-class="opacity-5 keywordBoradAni"
					no-level
					:height="90"
					@click="del"
					:follow-theme="false"
					:follow-dark="false"
					:dark="_dark"
					color="grey-1"
					:round="2"
					_class="flex-center"
					:padding="[0, 0]"
					:margin="[4, 4]"
					class="flex flex-col"
					paren-class="flex-row flex-center"
				>
					<tm-icon :userInteractionEnabled="false" name="tmicon-caret-left"></tm-icon>
				</tmSheet>

				<tmSheet
					hover-class="opacity-5 keywordBoradAni"
					no-level
					:height="90"
					@click="changeEnChart"
					:follow-theme="false"
					:follow-dark="false"
					:dark="_dark"
					color="grey-1"
					:round="2"
					_class="flex-center"
					:padding="[0, 0]"
					:margin="[4, 4]"
					class="flex flex-col"
					paren-class="flex-row flex-center"
				>
					<tmText :userInteractionEnabled="false" :font-size="32" _class="text-weight-b" :label="!changeChart ? 'En' : '简'"></tmText>
				</tmSheet>

				<view class="flex-1 flex flex-row">
					<tmSheet
						hover-class="opacity-5 keywordBoradAni"
						@click="confirm"
						:height="90"
						:follow-theme="props.followTheme"
						:follow-dark="false"
						:dark="_dark"
						:color="props.color"
						:round="2"
						_class="flex-center "
						:padding="[0, 0]"
						:margin="[4, 4]"
						class="flex-1 flex flex-col"
						paren-class="flex-1 flex-row flex-center"
					>
						<tm-icon :userInteractionEnabled="false" name="tmicon-check"></tm-icon>
					</tmSheet>
				</view>
			</view>
		</view>
	</tmSheet>
</template>
<script lang="ts" setup>
import { computed, ref, toRaw, watch, nextTick } from 'vue'
import tmText from '../tm-text/tm-text.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { propsCutom } from './props'
const props = defineProps({ ...propsCutom })
const _maxLength = computed(() => props.maxLength)
const emits = defineEmits(['update:modelValue', 'change', 'confirm', 'success'])

const _dark = computed(() => props.dark)
const numberArray = ref([])
const _value = ref(props.modelValue)

//是否切换为字符
const changeChart = ref(false)
getChart()

function keydown(e: number | string) {
	let k = String(e)
	let estr = _value.value + k
	if (estr.split('').length > _maxLength.value && _maxLength.value > 0) {
		return
	}

	_value.value = estr
	emits('update:modelValue', _value.value)
	emits('change', _value.value)
	if (changeChart.value == false) {
		changeEnChart()
	}
	if (estr.split('').length === _maxLength.value && _maxLength.value > 0) {
		emits('success', estr)
	}
}
function del() {
	if (_value.value == '' || _value.value.length == 0) return
	_value.value = _value.value.substring(0, _value.value.length - 1)
	emits('update:modelValue', _value.value)
	emits('change', _value.value)
	if (_value.value.length == 0) {
		changeChart.value == true
		changeEnChart()
	}
}
function confirm() {
	emits('confirm', _value.value)
}

function changeEnChart() {
	changeChart.value = !changeChart.value
	getChart()
}

function getChart() {
	const numbersfc = ref([
		'京',
		'沪',
		'津',
		'渝',
		'鲁',
		'冀',
		'晋',
		'蒙',
		'辽',
		'吉',
		'黑',
		'苏',
		'浙',
		'皖',
		'闽',
		'赣',
		'豫',
		'湘',
		'鄂',
		'粤',
		'桂',
		'琼',
		'川',
		'贵',
		'云',
		'藏',
		'陕',
		'甘',
		'青',
		'宁',
		'新',
		'港',
		'澳',
		'台',
		'警',
		'使'
	])
	const chartsfc = ref([
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
		'0',
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z'
	])
	//字母
	if (!changeChart.value) {
		if (props.random) {
			numbersfc.value = shuffle(toRaw(numbersfc.value))
		}
		numberArray.value = uni.$tm.u.splitData(toRaw(numbersfc.value), 9)
		//符号
	} else {
		if (props.random) {
			chartsfc.value = shuffle(toRaw(chartsfc.value))
		}
		numberArray.value = uni.$tm.u.splitData(toRaw(chartsfc.value), 9)
	}
}
// 随机数组排序。
function shuffle(arr: Array<any> = []) {
	var i = arr.length,
		t: number,
		j: number
	while (--i) {
		j = Math.floor(Math.random() * i)
		t = arr[i]
		arr[i] = arr[j]
		arr[j] = t
	}
	return arr
}

watch(
	() => props.modelValue,
	() => {
		nextTick(() => (_value.value = props.modelValue))
	}
)
</script>
<style>
@import url(./ani.css);
</style>

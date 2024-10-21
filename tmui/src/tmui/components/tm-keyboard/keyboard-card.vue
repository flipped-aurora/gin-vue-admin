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
			<tm-text
				:color="_isOk ? 'green' : 'red'"
				v-if="_value && props.showInputContent"
				:font-size="34"
				_class="text-weight-b pr-24"
				:label="_value"
			></tm-text>
			<tm-icon v-if="_isOk && props.showInputContent" color="green" :font-size="34" name="tmicon-check-circle-fill"></tm-icon>
		</view>
		<view class="flex flex-row">
			<view class="flex-5 flex flex-col">
				<view class="flex-row flex flex-1" v-for="(item2, index2) in numberArray" :key="index2">
					<tmSheet
						hover-class="opacity-5 keywordBoradAni"
						no-level
						@click="keydown(item)"
						:follow-theme="false"
						:follow-dark="false"
						:dark="_dark"
						:round="2"
						:height="95"
						_class="flex-center"
						:padding="[0, 0]"
						:margin="[4, 4]"
						v-for="(item, index) in item2"
						:key="index"
						:paren-class="index2 == 3 && index == 0 ? 'flex-3' : 'flex-3'"
						:class="index2 == 3 && index == 0 ? 'flex-5' : 'flex-3'"
					>
						<view style="width: 40rpx" class="flex flex-center flex-row">
							<tmText :userInteractionEnabled="false" :font-size="38" _class="text-weight-b" :label="item"></tmText>
						</view>
					</tmSheet>
				</view>
			</view>
			<view class="flex-1 flex flex-col">
				<tmSheet
					hover-class="opacity-5 keywordBoradAni"
					no-level
					:height="95"
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
				<view class="flex-1 flex flex-row">
					<tmSheet
						hover-class="opacity-5 keywordBoradAni"
						no-level
						@click="confirm"
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
import { validateIdCard } from './card'
import { propsCutom } from './props'
const props = defineProps({ ...propsCutom })
const _maxLength = computed(() => props.maxLength)
const emits = defineEmits(['update:modelValue', 'change', 'confirm', 'success'])

const _dark = computed(() => props.dark)
const numbersfc = ref([1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 'X'])
if (props.random) {
	numbersfc.value = shuffle(toRaw(numbersfc.value))
}
const numberArray = ref([])
numberArray.value = uni.$tm.u.splitData(toRaw(numbersfc.value), 3)

const _value = ref(props.modelValue)
const _isOk = computed(() => validateIdCard(_value.value))
function keydown(e: number | string) {
	let k = String(e)
	let estr = _value.value + k
	if (estr.split('').length > _maxLength.value && _maxLength.value > 0) {
		return
	}

	_value.value = estr
	emits('update:modelValue', _value.value)
	emits('change', _value.value)
	if (estr.split('').length === _maxLength.value && _maxLength.value > 0) {
		emits('success', estr)
	}
}
function del() {
	if (_value.value == '' || _value.value.length == 0) return
	_value.value = _value.value.substring(0, _value.value.length - 1)
	emits('update:modelValue', _value.value)
	emits('change', _value.value)
}
function confirm() {
	emits('confirm', _value.value)
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

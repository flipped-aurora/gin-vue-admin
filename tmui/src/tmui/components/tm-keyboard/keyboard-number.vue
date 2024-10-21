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
			<tm-text v-if="!_value && !props.showInputContent" :font-size="28" _class="text-weight-b" :label="props.title"> </tm-text>
			<tm-text v-if="_value && props.showInputContent" :font-size="34" _class="text-weight-b pr-24" :label="_value"> </tm-text>
		</view>
		<view>
			<tm-row :column="7" align="start">
				<tm-col :col="6" :height="0" :transprent="true">
					<tm-row :column="3">
						<tm-col
							@click="keydown(item)"
							hover-class="opacity-5"
							:margin="[8, 0, 4, 8]"
							:round="props.round"
							:height="100"
							:col="1"
							v-for="(item, index) in 9"
							:key="index"
						>
							<tmText :userInteractionEnabled="false" :font-size="38" _class="text-weight-b" :label="item"></tmText>
						</tm-col>
						<tm-col @click="keydown('0')" hover-class="opacity-5 " :margin="[8, 0, 4, 0]" :round="props.round" :height="100" :col="2">
							<tmText :userInteractionEnabled="false" :font-size="38" _class="text-weight-b" label="0"></tmText>
						</tm-col>
						<tm-col @click="keydown('.')" hover-class="opacity-5 " :margin="[8, 0, 4, 0]" :round="props.round" :height="100" :col="1">
							<tmText :userInteractionEnabled="false" :font-size="38" _class="text-weight-b" label="."></tmText>
						</tm-col>
					</tm-row>
				</tm-col>
				<tm-col :col="1" :height="420" align="start" :transprent="true" :margin="[0]">
					<tm-row :column="1">
						<tm-col
							@click="del"
							hover-class="opacity-5 "
							color="grey-1"
							:margin="[8, 0, 4, 8]"
							:round="props.round"
							:height="100"
							:col="1"
						>
							<tm-icon :userInteractionEnabled="false" name="tmicon-caret-left"></tm-icon>
						</tm-col>
						<tm-col
							@click="confirm"
							hover-class="opacity-5 "
							:color="props.color"
							:margin="[8, 0, 0, 0]"
							:round="props.round"
							:height="300"
							:col="1"
						>
							<tm-icon :userInteractionEnabled="false" name="tmicon-check"></tm-icon>
						</tm-col>
					</tm-row>
				</tm-col>
			</tm-row>
		</view>
		<!--  
    <view class="flex flex-row">
      <view class="flex-5 flex flex-col">
        <view
          class="flex-row flex flex-1"
          v-for="(item2, index2) in numberArray"
          :key="index2"
        >
          <tmSheet
            hover-class="opacity-5 keywordBoradAni"
            no-level
            @click="keydown(item)"
            :follow-theme="false"
            :follow-dark="false"
            :dark="_dark"
            :round="2"
            :height="100"
            _class="flex-center"
            :padding="[0, 0]"
            :margin="[4, 4]"
            v-for="(item, index) in item2"
            :key="index"
            :paren-class="index2 == 3 && index == 0 ? 'flex-3' : 'flex-3'"
            :class="index2 == 3 && index == 0 ? 'flex-5' : 'flex-3'"
          >
            <view style="width: 40rpx" class="flex flex-center flex-row">
              <tmText
                :userInteractionEnabled="false"
                :font-size="32"
                _class="text-weight-b"
                :label="item"
              ></tmText>
            </view>
          </tmSheet>
        </view>
      </view>
      <view class="flex-1 flex flex-col">
        <view class="flex flex-row">
          <tmSheet
            hover-class="opacity-5 keywordBoradAni"
            no-level
            :height="100"
            @click="del"
            :follow-theme="false"
            :follow-dark="false"
            :dark="_dark"
            color="grey-1"
            :round="2"
            _class="flex-center"
            :padding="[0, 0]"
            :margin="[4, 4]"
            class="flex-1 flex flex-col"
            paren-class="flex-1 flex-row flex-center"
          >
            <tm-icon :userInteractionEnabled="false" name="tmicon-caret-left"></tm-icon>
          </tmSheet>
        </view>
        <view class="flex-6 flex flex-row">
          <tmSheet
            hover-class="opacity-5 keywordBoradAni"
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
    -->
	</tmSheet>
</template>
<script lang="ts" setup>
import { computed, ref, toRaw, watch, Ref } from 'vue'
import tmText from '../tm-text/tm-text.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmRow from '../tm-row/tm-row.vue'
import tmCol from '../tm-col/tm-col.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { propsCutom } from './props'
const props = defineProps({ ...propsCutom })
const _maxLength = computed(() => props.maxLength)

const emits = defineEmits(['update:modelValue', 'change', 'confirm', 'success'])
const _dark = computed(() => props.dark)
let defaultNum = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0, '.']
if (!props.decimal) {
	defaultNum.pop()
}
const numbersfc = ref(defaultNum)
if (props.random) {
	numbersfc.value = shuffle(toRaw(numbersfc.value))
}
const numberArray: Ref<Array<string | number>> = ref([])
numberArray.value = uni.$tm.u.splitData(toRaw(numbersfc.value), 3)

const _value = ref(props.modelValue)
const _valueArrays = computed(() => _value.value.split(''))
function keydown(e: number | string) {
	let k = String(e)
	if (props.decimal && k == '.' && _valueArrays.value.includes('.')) {
		return
	}

	let estr = _value.value + k
	if (estr.split('').length > _maxLength.value && _maxLength.value > 0) {
		return
	}
	_value.value = estr
	emits('update:modelValue', _value.value)
	emits('change', props.modelValue)
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
		_value.value = props.modelValue
	}
)
</script>

<style>
@import url(./ani.css);
</style>

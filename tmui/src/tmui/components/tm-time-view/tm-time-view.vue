<template>
  <view class="flex relative flex-col" :style="{ height: props.height + 'rpx' }">
	<!-- #ifndef MP-ALIPAY -->
    <picker-view
      v-if="show"
      :value="colIndex"
      @change="colchange"
      :style="{ height: props.height + 'rpx' }"
      :mask-style="maskStyle"
      :immediateChange="props.immediateChange"
      indicator-style="height:50px"
    >
	<!-- #endif -->
	<!-- #ifdef MP-ALIPAY -->
	<picker-view
	  v-if="show"
	  :value="colIndex"
	  @change="colchange"
	  :style="{ height: props.height + 'rpx' }"
	  :mask-style="maskStyle"
	  :immediateChange="props.immediateChange"
	>
	<!-- #endif -->
      <picker-view-column v-if="showCol.year">
        <view
          v-for="(item, index) in _col.year"
          :key="index"
          class="flex itemcel flex-row flex-row-center-center"
          :class="[colIndex[0] == index ? '' : 'UnitemSelected']"
		  
        >
		<!-- #ifdef APP-NVUE -->
          <TmText
            :font-size="30"
            :dark="isDark"
            :label="item + showSuffix['year']"
          ></TmText>
		  <!-- #endif -->
		  <!-- #ifndef APP-NVUE -->
		  <text :style="{color:store.tmStore.dark?'white':'black'}">{{item + showSuffix['year']}}</text>
		  <!-- #endif -->
		  
        </view>
      </picker-view-column>
      <picker-view-column v-if="showCol.month">
        <view
          v-for="(item, index) in _col.month"
          :key="index"
          class="flex itemcel flex-row flex-row-center-center"
          :class="[colIndex[1] == index ? '' : 'UnitemSelected']"
        >
         
		  <!-- #ifdef APP-NVUE -->
		   <TmText
		     :font-size="30"
		     :dark="isDark"
		     :label="item + 1 + showSuffix['month']"
		   ></TmText>
		    <!-- #endif -->
		    <!-- #ifndef APP-NVUE -->
		    <text :style="{color:store.tmStore.dark?'white':'black'}">{{item + 1 + showSuffix['month']}}</text>
		    <!-- #endif -->
		  
        </view>
      </picker-view-column>
      <picker-view-column v-if="showCol.day">
        <view
          v-for="(item, index) in _col.date"
          :key="index"
          class="flex itemcel flex-row flex-row-center-center"
          :class="[colIndex[2] == index ? '' : 'UnitemSelected']"
        >
          
		  <!-- #ifdef APP-NVUE -->
		   <TmText
		     :font-size="30"
		     :dark="isDark"
		     :label="item + showSuffix['date']"
		   ></TmText>
		    <!-- #endif -->
		    <!-- #ifndef APP-NVUE -->
		    <text :style="{color:store.tmStore.dark?'white':'black'}">{{item + showSuffix['date']}}</text>
		    <!-- #endif -->
        </view>
      </picker-view-column>
      <picker-view-column v-if="showCol.hour">
        <view
          v-for="(item, index) in _col.hour"
          :key="index"
          class="flex itemcel flex-row flex-row-center-center"
          :class="[colIndex[3] == index ? '' : 'UnitemSelected']"
        >
          
		  <!-- #ifdef APP-NVUE -->
		   <TmText
		     :font-size="30"
		     :dark="isDark"
		     :label="item + showSuffix['hour']"
		   ></TmText>
		    <!-- #endif -->
		    <!-- #ifndef APP-NVUE -->
		    <text :style="{color:store.tmStore.dark?'white':'black'}">{{item + showSuffix['hour']}}</text>
		    <!-- #endif -->
        </view>
      </picker-view-column>
      <picker-view-column v-if="showCol.minute">
        <view
          v-for="(item, index) in _col.minute"
          :key="index"
          class="flex itemcel flex-row flex-row-center-center"
          :class="[colIndex[4] == index ? '' : 'UnitemSelected']"
        >
          
		  <!-- #ifdef APP-NVUE -->
		   <TmText
		     :font-size="30"
		     :dark="isDark"
		     :label="item + showSuffix['minute']"
		   ></TmText>
		    <!-- #endif -->
		    <!-- #ifndef APP-NVUE -->
		    <text :style="{color:store.tmStore.dark?'white':'black'}">{{item + showSuffix['minute']}}</text>
		    <!-- #endif -->
        </view>
      </picker-view-column>
      <picker-view-column v-if="showCol.second">
        <view
          v-for="(item, index) in _col.second"
          :key="index"
          class="flex itemcel flex-row flex-row-center-center"
          :class="[colIndex[5] == index ? '' : 'UnitemSelected']"
        >
          
		  <!-- #ifdef APP-NVUE -->
		  <TmText
		    :font-size="30"
		    :dark="isDark"
		    :label="item + showSuffix['second']"
		  ></TmText>
		    <!-- #endif -->
		    <!-- #ifndef APP-NVUE -->
		    <text :style="{color:store.tmStore.dark?'white':'black'}">{{item + showSuffix['second']}}</text>
		    <!-- #endif -->
        </view>
      </picker-view-column>
    </picker-view>
    <!-- #ifdef APP-NVUE -->
    <view
      v-if="isDark"
      :userInteractionEnabled="false"
      class="top absolute l-0 t-0"
      :style="{ height: maskHeight + 'px', width: maskWidth + 'px' }"
    ></view>
    <view
      v-if="isDark"
      :userInteractionEnabled="false"
      class="bottom absolute l-0 b-0"
      :style="{ height: maskHeight + 'px', width: maskWidth + 'px' }"
    ></view>
    <TmIcon v-if="!show" spin name="tmicon-shuaxin"></TmIcon>
    <!-- #endif -->
  </view>
</template>

<script lang="ts" setup>
/**
 * 时间选择
 * @description 嵌入在页面的时间选择器。
 */
import { useTmpiniaStore } from "../../tool/lib/tmpinia";
import {
  computed,
  PropType,
  getCurrentInstance,
  ref,
  toRaw,
  onMounted,
  nextTick,
  watch,
  onUpdated,
  Ref,
} from "vue";
import { showDetail, coltimeData, timeDetailType } from "./interface";
import * as dayjs from "../../tool/dayjs/esm/index";
import { propsOpts } from "./props";
import {
  rangeTimeArray,
  getNowbyIndex,
  getIndexNowbydate,
  checkNowDateisBetween,
} from "./time";
import TmText from "../tm-text/tm-text.vue";
import TmIcon from "../tm-icon/tm-icon.vue";
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin("dom");
// #endif
const proxy = getCurrentInstance()?.proxy ?? null;
const store = useTmpiniaStore();
const emits = defineEmits(["update:modelValue", "update:modelStr", "change"]);
const tmTimeViewName = "tmTimeViewName";
const DayJs = dayjs.default;
const props = defineProps({ ...propsOpts });
const _nowtime = ref(
  DayJs(checkNowDateisBetween(props.defaultValue, props.start, props.end))
);
const _nowtimeValue = computed(() => _nowtime.value.format());

const show = ref(true);

const _startTime = computed(() => {
  return DayJs(props.start).isValid()
    ? DayJs(props.start).format()
    : DayJs().subtract(3, "year").format();
});
const _endTime = computed(() => {
  return DayJs(props.end).isValid()
    ? DayJs(props.end).format()
    : DayJs().add(1, "year").format();
});
const showCol = computed<showDetail>(() => {
  return {
    year: props.showDetail?.year ?? true,
    month: props.showDetail?.month ?? true,
    day: props.showDetail?.day ?? true,
    hour: props.showDetail?.hour ?? false,
    minute: props.showDetail?.minute ?? false,
    second: props.showDetail?.second ?? false,
  };
});

const showSuffix = computed(() => {
  return {
    year: props.showSuffix?.year ?? "年",
    month: props.showSuffix?.month ?? "月",
    hour: props.showSuffix?.hour ?? "时",
    minute: props.showSuffix?.minute ?? "分",
    second: props.showSuffix?.second ?? "秒",
    date: props.showSuffix?.day ?? "日",
  };
});
const isDark = computed(() => store.tmStore.dark);
let colIndex: Ref<Array<number>> = ref([0, 0, 0, 0, 0, 0]);
const _col = ref({
  year: [] as Array<number>,
  month: [] as Array<number>,
  date: [] as Array<number>,
  hour: [] as Array<number>,
  minute: [] as Array<number>,
  second: [] as Array<number>,
});

let timid = NaN;
const maskWidth = ref(0);
const maskHeight = computed(() => {
  return (uni.upx2px(props.height) - 50) / 2;
});
const maskStyle = computed(() => {
  let str_white =
    "background-image:linear-gradient(rgba(255,255,255,0.95),rgba(255,255,255,0.6)),linear-gradient(rgba(255,255,255,0.6),rgba(255,255,255,0.95))";
  let str_black =
    "background-image:linear-gradient(rgba(17, 17, 17, 1.0),rgba(106, 106, 106, 0.2)),linear-gradient(rgba(106, 106, 106, 0.2),rgba(17, 17, 17, 1.0))";

  // #ifdef APP-NVUE
  str_black =
    "background-image: linear-gradient(to bottom,rgba(30, 30, 30, 0.9),rgba(104, 104, 104, 0.6))";
  // #endif
  if (!isDark.value) {
    return str_white;
  }
  return str_black;
});
_col.value = rangeTimeArray(
  _nowtimeValue.value,
  _startTime.value,
  _endTime.value,
  showCol.value
);

function colchange(e: any) {
 
  let changedate = getNowbyIndex(_col.value, e.detail.value, showCol.value, _startTime.value,_endTime.value);
  let testDate  = checkNowDateisBetween(changedate, _startTime.value,_endTime.value)
  

   
  let testRang = rangeTimeArray(
      testDate,
      _startTime.value,
      _endTime.value,
      showCol.value
    );
  
  
 
  
  _nowtime.value = DayJs(testDate);
  colIndex.value = getIndexNowbydate(testRang, _nowtime.value, showCol.value);
  emits("update:modelValue", _nowtime.value.format("YYYY/MM/DD HH:mm:ss"));
  emits("update:modelStr", _nowtime.value.format(props.format));
  emits("change", _nowtime.value.format(props.format));
 
  _col.value= testRang;

  
}

watch(
  () => props.modelValue,
  () => {
    if (!DayJs(props.modelValue).isValid()) return;
    let deattime = DayJs(checkNowDateisBetween(props.modelValue, props.start, props.end));

    if (DayJs(deattime).isSame(_nowtime.value)) return;
    _nowtime.value = deattime;
    emits("update:modelStr", _nowtime.value.format(props.format));
    // #ifdef APP-NVUE
    _col.value = rangeTimeArray(
      deattime,
      _startTime.value,
      _endTime.value,
      showCol.value
    );
    show.value = false;
    colIndex.value = getIndexNowbydate(_col.value, _nowtime.value, showCol.value);
    nextTick(() => {
      /**这力着重解释下，uni sdk从3.6.8开始，在nvue下直接对picker view赋值value，页面不会有任何变化，必须刷新下页面才可以显示正确
       * 其它平台没有这问题
       */
      show.value = true;
    });
    // #endif
    // #ifndef APP-NVUE
    _col.value = rangeTimeArray(
      deattime,
      _startTime.value,
      _endTime.value,
      showCol.value
    );
    colIndex.value = getIndexNowbydate(_col.value, _nowtime.value, showCol.value);
    // #endif
  }
);

function nvuegetClientRect() {
  nextTick(function () {
    // #ifdef APP-PLUS-NVUE
    dom.getComponentRect(proxy.$refs.picker, function (res) {
      if (res?.size) {
        maskWidth.value = res.size.width;

        if (res.size.width == 0) {
          nvuegetClientRect();
        }
      }
    });
    // #endif
  });
}

onMounted(() => {
  nvuegetClientRect();
  nextTick(() => {
    emits("update:modelValue", _nowtime.value.format("YYYY/MM/DD HH:mm:ss"));
    emits("update:modelStr", _nowtime.value.format(props.format));
    colIndex.value = getIndexNowbydate(_col.value, _nowtime.value, showCol.value);
  });
});


onUpdated(() => nvuegetClientRect());

// defineExpose({tmTimeViewName,setNowtime})
</script>

<style scoped>
.top {
  background-image: linear-gradient(
    to bottom,
    rgba(17, 17, 17, 1),
    rgba(36, 36, 36, 0.6)
  );
}

.bottom {
  background-image: linear-gradient(to top, rgba(17, 17, 17, 1), rgba(36, 36, 36, 0.6));
}

.itemcel {
  justify-content: center;
  height: 50px;
  align-items: center;
}

.itemSelected {
  background-image: linear-gradient(
    to bottom,
    rgba(0, 0, 0, 0.04),
    rgba(0, 0, 0, 0.01),
    rgba(0, 0, 0, 0.04)
  );
}

.UnitemSelected {
  background-image: rgba(0, 0, 0, 0);
}

.itemSelectedNvue {
  background-image: rgba(0, 0, 0, 0.04);
}
</style>

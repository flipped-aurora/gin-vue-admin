<template>
  <view
    ref="tmIndexes"
    :class="[
      'overflow relative',
      `mx-${_margin[0]} my-${_padding[1]} px-${_margin[0]} py-${_padding[1]}`,
    ]"
    :style="[{ height: `${_height}` + props.unit }]"
  >
    <scroll-view
      :offset-accuracy="5"
      @scroll="scrollChnage"
      :scroll-into-view="'id_' + _cureent_id"
      :scroll-with-animation="false"
      :scroll-top="_cureent_top"
      :scroll-y="true"
      :style="[
        _width ? { width: _width + props.unit } : '',
        _height ? { height: _height + props.unit } : '',
      ]"
    >
      <slot></slot>
    </scroll-view>
    <view
      class="absolute flex flex-col flex-center t-0 r-24"
      :style="[{ height: `${_height}` + props.unit, width: '60rpx' }]"
    >
      <view
        :style="[{ height: `${navright.length * navHeight}rpx`, width: '60rpx' }]"
        @touchstart="touchStart"
        @touchmove.stop.prevent="touchMove"
        @touchend="touchEnd"
        id="navlist"
        ref="navlist"
        class="flex flex-col flex-center"
      >
        <tm-sheet
          :eventPenetrationEnabled="true"
          no-level
          :round="10"
          color="white"
          :shadow="2"
          :margin="[0, 0]"
          :padding="[0, 0]"
          :width="40"
        >
          <view
            @click.stop="navClick(item)"
            hover-class="opacity-5"
            class="flex-center flex"
            v-for="(item, index) in navright"
            :key="index"
            :style="{ width: '40rpx', height: `${navHeight}rpx` }"
          >
            <tm-text
              :followTheme="_cureent_id == item.id ? props.followTheme : false"
              :color="_cureent_id == item.id ? props.color : ''"
              :font-size="20"
              :label="item.subText"
            ></tm-text>
          </view>
        </tm-sheet>
      </view>
    </view>
    <view
      v-if="_showCenterTitle"
      class="absolute l-0 t-0 fulled"
      :style="[
        {
          top: (_height - 70) / 2 + props.unit,
        },
        _isNvue ? { left: parentLeft + 'px' } : { left: 'calc(50% - 70rpx)' },
      ]"
    >
      <tm-sheet
        v-if="_cureent_item != null"
        _class="flex flex-center"
        :shadow="5"
        :margin="[24, 24]"
        :padding="[0, 0]"
        :width="100"
        :height="100"
        :round="20"
      >
        <tm-text
          :followTheme="props.followTheme"
          :color="props.color"
          :font-size="36"
          :label="_cureent_item.subText"
        ></tm-text>
      </tm-sheet>
    </view>
  </view>
</template>

<script lang="ts" setup>
/**
	 * 列表索引
	 * @description 索引列表，内部只能放置tm-indexes-item组件。
	 * @example
	 * <tm-indexes>
			<tm-indexes-item :title="index%10==0?index:''"  v-for="(item,index) in 60" :key="index">
				<tm-text label="想要什么的."></tm-text>
			</tm-indexes-item>
		</tm-indexes>
	 */
import {
  computed,
  PropType,
  ref,
  Ref,
  nextTick,
  onMounted,
  getCurrentInstance,
} from "vue";
import tmSheet from "../tm-sheet/tm-sheet.vue";
import tmText from "../tm-text/tm-text.vue";

// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin("dom");
// #endif
const proxy = getCurrentInstance()?.proxy ?? null;
const emits = defineEmits(["nav-click"]);
const props = defineProps({
  followTheme: {
    type: Boolean,
    default: true,
  },
  width: {
    type: Number,
    default: 0,
  },
  height: {
    type: Number,
    default: 700,
  },
  margin: {
    type: Array as PropType<Array<number>>,
    default: () => [0, 0],
  },
  padding: {
    type: Array as PropType<Array<number>>,
    default: () => [0, 0],
  },
  color: {
    type: String,
    default: "primary",
  },
  unit: {
    type: String,
    default: "rpx",
  },
  menuItemHeight: {
    type: Number,
    default: 40,
  },
});
interface heightItem {
  height: number;
  id: number;
  text: string; //标题
  subText: string; //子标题显示在导航上的。
}
const navHeight = ref(props.menuItemHeight);
const _margin = computed(() => props.margin);
const _padding = computed(() => props.padding);
const _height = computed(() => props.height);
const _width = computed(() => props.width);
const _cureent_id = ref(0);
const _cureent_item: Ref<heightItem | null> = ref(null);
const _isNvue = ref(false);
// #ifdef APP-NVUE
_isNvue.value = true;
// #endif
const _showCenterTitle = ref(false);
let _timeid = Number(uni.$tm.u.getUid(1) || 12);
const parentLeft = ref(0);

const _cacheHeightArrays: Ref<Array<heightItem>> = ref([]);
const _cureent_top = ref(0);
const compentNameId = "tmIndexesId";
const navright = computed(() => {
  return _cacheHeightArrays.value.filter((el) => el.text !== "");
});

//------------
//兼容滑块选中的参数。
const winOffsetY = ref(0);
const itemHeight = Math.ceil(uni.upx2px(navHeight.value));
const winHeight = computed(() => navright.value.length * navHeight.value);
const winHeightPx = computed(() => Math.ceil(uni.upx2px(winHeight.value)));

function throttle(func: Function, delay: number = 40) {
  var prev = Date.now();
  return function () {
    // @ts-ignore
    var context: any = this;
    var args = arguments;
    var now = Date.now();
    if (now - prev >= delay) {
      func.apply(context, args);
      prev = Date.now();
    }
  };
}

//---------

function pushKey(height: number, id: number, text: string, subText: string) {
  let index = _cacheHeightArrays.value.findIndex((el) => el.id == id);
  if (index == -1) {
    _cacheHeightArrays.value.push({
      height: height,
      id: id,
      text: text,
      subText: subText,
    });
  } else {
    _cacheHeightArrays.value.splice(index, 1, {
      height: height,
      id: id,
      text: text,
      subText: subText,
    });
  }
}

function delKey(height: number, id: number) {
  let index = _cacheHeightArrays.value.findIndex((el) => el.id == id);
  if (index > -1) {
    _cacheHeightArrays.value.splice(index, 1);
  }
}
function scrollChnage(e) {
  // uni.$tm.u.debounce(function(){
  // 	let nowitem = getPosItem(e.detail.scrollTop);
  // 	if(nowitem){
  // 		_cureent_id.value = nowitem.id;
  // 		_cureent_item.value = nowitem;
  // 	}
  // },200)
}
onMounted(() => {
  nextTick(() => {
    nvuegetClientRect();
  });
});
function nvuegetClientRect() {
  // #ifdef APP-PLUS-NVUE
  nextTick(function () {
    dom.getComponentRect(proxy.$refs.tmIndexes, function (res) {
      if (res?.size) {
        if (res.size.width > 0) {
          if (props.unit == "rpx") {
            parentLeft.value = Math.ceil(uni.upx2px(res.size.width / 2 - 74));
          } else {
            parentLeft.value = Math.ceil(res.size.width / 2 - uni.upx2px(74));
          }
        }
        if (res.size.height == 0) {
          nvuegetClientRect();
        }
      }
    });
  });
  // #endif

  // #ifndef APP-NVUE
  uni
    .createSelectorQuery()
    .in(proxy)
    .select("#navlist")
    .boundingClientRect()
    .exec((ret) => {
      winOffsetY.value = ret[0].top;
    });
  // #endif
  // #ifdef APP-NVUE
  setTimeout(function () {
    dom.getComponentRect(proxy.$refs["navlist"], (res: any) => {
      winOffsetY.value = res.size.top;
    });
  }, 100);
  // #endif
}
function getPosItem(top: number) {
  let avl: Array<any> = [];
  let nowitem = null;
  navright.value.forEach((el2) => {
    let index = _cacheHeightArrays.value.findIndex((el) => el.id == el2.id);
    if (index > -1) {
      let ar = _cacheHeightArrays.value.slice(0, index);
      let atm = { top: 0, item: el2 };
      ar.forEach((el3) => (atm.top += el3.height));

      atm.top = uni.upx2px(atm.top - 50);
      avl.push(atm);
    }
  });
  let pavl = [...avl];
  let lastitem = pavl[pavl.length - 1];
  if (top >= lastitem.top) {
    return lastitem.item;
  }
  avl.reverse();
  for (let i = 0; i < avl.length; i++) {
    let item = avl[i + 1];
    if (top >= item.top) {
      nowitem = item.item;
      break;
    }
  }
  return nowitem;
}
function navClick(item: heightItem) {
  // #ifdef MP
  let index = _cacheHeightArrays.value.findIndex((el) => el.id == item.id);
  let splitAr = _cacheHeightArrays.value.slice(0, index);
  let scrotop = 0;
  splitAr.forEach((a, b) => (scrotop += a.height));
  _cureent_top.value = 0;
  _cureent_item.value = item;
  _cureent_id.value = item.id;
  nextTick(() => {
    _cureent_top.value = scrotop;
  });
  // #endif

  // #ifndef MP
  if (_cureent_id.value == item.id) return;
  _cureent_id.value = "";
  _cureent_item.value = item;
  _showCenterTitle.value = true;
  clearTimeout(_timeid);
  _timeid = setTimeout(function () {
    _showCenterTitle.value = false;
  }, 800);
  nextTick(() => {
    _cureent_id.value = item.id;
  });
  // #endif
  emits("nav-click", item);
}

function touchStart(e: TouchEvent) {
  if (winOffsetY.value == 0) {
    nvuegetClientRect();
  }
}
function touchMove(e: TouchEvent) {
  let pageY = isPc(e) ? e.pageY : e.touches[0].pageY;
  let index = Math.floor((pageY - winOffsetY.value) / itemHeight);
  let item = navright.value[index];
  if (!item) return;
  let touchmoveIndex = _cacheHeightArrays.value.findIndex(
    (el) => el.id == _cureent_id.value
  );
  if (touchmoveIndex === index) {
    return false;
  }
  if (item) {
    _cureent_item.value = item;
    _cureent_id.value = item.id;
  }
  // #ifdef MP

  if (item) {
    let indexMp = _cacheHeightArrays.value.findIndex((el) => el.id == item.id);
    if (indexMp == touchmoveIndex) return;
    let splitAr = _cacheHeightArrays.value.slice(0, indexMp);
    let scrotop = 0;
    splitAr.forEach((a, b) => (scrotop += a.height));

    _cureent_top.value = scrotop;
  }
  // #endif

  // #ifdef APP-PLUS
  throttle(() => {
    if (item) {
      // #ifdef APP-NVUE
      dom.scrollToElement(this.$refs["id_" + String(_cureent_id.value)][0], {
        animated: false,
      });
      // #endif
    }
  }, 40);
  // #endif
  _showCenterTitle.value = true;
}
function touchEnd(e: TouchEvent) {
  _showCenterTitle.value = false;
}

function isPc(e: TouchEvent) {
  let eventStr = e.type.toLocaleLowerCase();
  if (eventStr.indexOf("mouse") !== -1) return true;
  return false;
}

defineExpose({
  compentNameId,
  pushKey,
  delKey,
});
</script>

<style></style>

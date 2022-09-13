<template>
  <div>
    <span class="gva-icon" style="position: absolute; z-index: 9999; padding: 3px 10px 0; ">
      <el-icon>
        <component :is="metaData.icon" />
      </el-icon>
    </span>
    <el-select v-model="metaData.icon" style="width:100%" clearable filterable class="gva-select" placeholder="请选择">
      <el-option
        v-for="item in options"
        :key="item.key"
        class="select__option_item"
        :label="item.key"
        :value="item.key"
      >
        <span class="gva-icon" style=" padding: 3px 0 0; " :class="item.label">
          <el-icon>
            <component :is="item.label" />
          </el-icon>
        </span>
        <span style="text-align: left">{{ item.key }}</span>
      </el-option>
    </el-select>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const props = defineProps({
  meta: {
    default: function() {
      return {}
    },
    type: Object,
  },
})

const options = reactive([
  {
    'key': 'aim',
    'label': 'aim',
  },
  {
    'key': 'add-location',
    'label': 'add-location',
  },
  {
    'key': 'apple',
    'label': 'apple',
  },
  {
    'key': 'alarm-clock',
    'label': 'alarm-clock',
  },
  {
    'key': 'arrow-down',
    'label': 'arrow-down',
  },
  {
    'key': 'arrow-down-bold',
    'label': 'arrow-down-bold',
  },
  {
    'key': 'arrow-left',
    'label': 'arrow-left',
  },
  {
    'key': 'arrow-left-bold',
    'label': 'arrow-left-bold',
  },
  {
    'key': 'arrow-right-bold',
    'label': 'arrow-right-bold',
  },
  {
    'key': 'arrow-up',
    'label': 'arrow-up',
  },
  {
    'key': 'back',
    'label': 'back',
  },
  {
    'key': 'bell',
    'label': 'bell',
  },
  {
    'key': 'baseball',
    'label': 'baseball',
  },
  {
    'key': 'bicycle',
    'label': 'bicycle',
  },
  {
    'key': 'bell-filled',
    'label': 'bell-filled',
  },
  {
    'key': 'basketball',
    'label': 'basketball',
  },
  {
    'key': 'bottom',
    'label': 'bottom',
  },
  {
    'key': 'box',
    'label': 'box',
  },
  {
    'key': 'briefcase',
    'label': 'briefcase',
  },
  {
    'key': 'brush-filled',
    'label': 'brush-filled',
  },
  {
    'key': 'bowl',
    'label': 'bowl',
  },
  {
    'key': 'avatar',
    'label': 'avatar',
  },
  {
    'key': 'brush',
    'label': 'brush',
  },
  {
    'key': 'burger',
    'label': 'burger',
  },
  {
    'key': 'camera',
    'label': 'camera',
  },
  {
    'key': 'bottom-left',
    'label': 'bottom-left',
  },
  {
    'key': 'calendar',
    'label': 'calendar',
  },
  {
    'key': 'caret-bottom',
    'label': 'caret-bottom',
  },
  {
    'key': 'caret-left',
    'label': 'caret-left',
  },
  {
    'key': 'caret-right',
    'label': 'caret-right',
  },
  {
    'key': 'caret-top',
    'label': 'caret-top',
  },
  {
    'key': 'chat-dot-square',
    'label': 'chat-dot-square',
  },
  {
    'key': 'cellphone',
    'label': 'cellphone',
  },
  {
    'key': 'chat-dot-round',
    'label': 'chat-dot-round',
  },
  {
    'key': 'chat-line-square',
    'label': 'chat-line-square',
  },
  {
    'key': 'chat-line-round',
    'label': 'chat-line-round',
  },
  {
    'key': 'chat-round',
    'label': 'chat-round',
  },
  {
    'key': 'check',
    'label': 'check',
  },
  {
    'key': 'chat-square',
    'label': 'chat-square',
  },
  {
    'key': 'cherry',
    'label': 'cherry',
  },
  {
    'key': 'chicken',
    'label': 'chicken',
  },
  {
    'key': 'circle-check-filled',
    'label': 'circle-check-filled',
  },
  {
    'key': 'circle-check',
    'label': 'circle-check',
  },
  {
    'key': 'checked',
    'label': 'checked',
  },
  {
    'key': 'circle-close-filled',
    'label': 'circle-close-filled',
  },
  {
    'key': 'circle-close',
    'label': 'circle-close',
  },
  {
    'key': 'arrow-right',
    'label': 'arrow-right',
  },
  {
    'key': 'circle-plus',
    'label': 'circle-plus',
  },
  {
    'key': 'clock',
    'label': 'clock',
  },
  {
    'key': 'close-bold',
    'label': 'close-bold',
  },
  {
    'key': 'close',
    'label': 'close',
  },
  {
    'key': 'cloudy',
    'label': 'cloudy',
  },
  {
    'key': 'circle-plus-filled',
    'label': 'circle-plus-filled',
  },
  {
    'key': 'coffee-cup',
    'label': 'coffee-cup',
  },
  {
    'key': 'cold-drink',
    'label': 'cold-drink',
  },
  {
    'key': 'coin',
    'label': 'coin',
  },
  {
    'key': 'arrow-up-bold',
    'label': 'arrow-up-bold',
  },
  {
    'key': 'collection-tag',
    'label': 'collection-tag',
  },
  {
    'key': 'bottom-right',
    'label': 'bottom-right',
  },
  {
    'key': 'coffee',
    'label': 'coffee',
  },
  {
    'key': 'camera-filled',
    'label': 'camera-filled',
  },
  {
    'key': 'collection',
    'label': 'collection',
  },
  {
    'key': 'cpu',
    'label': 'cpu',
  },
  {
    'key': 'crop',
    'label': 'crop',
  },
  {
    'key': 'coordinate',
    'label': 'coordinate',
  },
  {
    'key': 'd-arrow-left',
    'label': 'd-arrow-left',
  },
  {
    'key': 'compass',
    'label': 'compass',
  },
  {
    'key': 'connection',
    'label': 'connection',
  },
  {
    'key': 'credit-card',
    'label': 'credit-card',
  },
  {
    'key': 'data-board',
    'label': 'data-board',
  },
  {
    'key': 'd-arrow-right',
    'label': 'd-arrow-right',
  },
  {
    'key': 'dessert',
    'label': 'dessert',
  },
  {
    'key': 'delete-location',
    'label': 'delete-location',
  },
  {
    'key': 'd-caret',
    'label': 'd-caret',
  },
  {
    'key': 'delete',
    'label': 'delete',
  },
  {
    'key': 'dish',
    'label': 'dish',
  },
  {
    'key': 'dish-dot',
    'label': 'dish-dot',
  },
  {
    'key': 'document-copy',
    'label': 'document-copy',
  },
  {
    'key': 'discount',
    'label': 'discount',
  },
  {
    'key': 'document-checked',
    'label': 'document-checked',
  },
  {
    'key': 'document-add',
    'label': 'document-add',
  },
  {
    'key': 'document-remove',
    'label': 'document-remove',
  },
  {
    'key': 'data-analysis',
    'label': 'data-analysis',
  },
  {
    'key': 'delete-filled',
    'label': 'delete-filled',
  },
  {
    'key': 'download',
    'label': 'download',
  },
  {
    'key': 'drizzling',
    'label': 'drizzling',
  },
  {
    'key': 'eleme',
    'label': 'eleme',
  },
  {
    'key': 'eleme-filled',
    'label': 'eleme-filled',
  },
  {
    'key': 'edit',
    'label': 'edit',
  },
  {
    'key': 'failed',
    'label': 'failed',
  },
  {
    'key': 'expand',
    'label': 'expand',
  },
  {
    'key': 'female',
    'label': 'female',
  },
  {
    'key': 'document',
    'label': 'document',
  },
  {
    'key': 'film',
    'label': 'film',
  },
  {
    'key': 'finished',
    'label': 'finished',
  },
  {
    'key': 'data-line',
    'label': 'data-line',
  },
  {
    'key': 'filter',
    'label': 'filter',
  },
  {
    'key': 'flag',
    'label': 'flag',
  },
  {
    'key': 'folder-checked',
    'label': 'folder-checked',
  },
  {
    'key': 'first-aid-kit',
    'label': 'first-aid-kit',
  },
  {
    'key': 'folder-add',
    'label': 'folder-add',
  },
  {
    'key': 'fold',
    'label': 'fold',
  },
  {
    'key': 'folder-delete',
    'label': 'folder-delete',
  },
  {
    'key': 'document-delete',
    'label': 'document-delete',
  },
  {
    'key': 'folder',
    'label': 'folder',
  },
  {
    'key': 'food',
    'label': 'food',
  },
  {
    'key': 'folder-opened',
    'label': 'folder-opened',
  },
  {
    'key': 'football',
    'label': 'football',
  },
  {
    'key': 'folder-remove',
    'label': 'folder-remove',
  },
  {
    'key': 'fries',
    'label': 'fries',
  },
  {
    'key': 'full-screen',
    'label': 'full-screen',
  },
  {
    'key': 'fork-spoon',
    'label': 'fork-spoon',
  },
  {
    'key': 'goblet',
    'label': 'goblet',
  },
  {
    'key': 'goblet-full',
    'label': 'goblet-full',
  },
  {
    'key': 'goods',
    'label': 'goods',
  },
  {
    'key': 'goblet-square-full',
    'label': 'goblet-square-full',
  },
  {
    'key': 'goods-filled',
    'label': 'goods-filled',
  },
  {
    'key': 'grid',
    'label': 'grid',
  },
  {
    'key': 'grape',
    'label': 'grape',
  },
  {
    'key': 'goblet-square',
    'label': 'goblet-square',
  },
  {
    'key': 'headset',
    'label': 'headset',
  },
  {
    'key': 'comment',
    'label': 'comment',
  },
  {
    'key': 'help-filled',
    'label': 'help-filled',
  },
  {
    'key': 'histogram',
    'label': 'histogram',
  },
  {
    'key': 'home-filled',
    'label': 'home-filled',
  },
  {
    'key': 'help',
    'label': 'help',
  },
  {
    'key': 'house',
    'label': 'house',
  },
  {
    'key': 'ice-cream-round',
    'label': 'ice-cream-round',
  },
  {
    'key': 'hot-water',
    'label': 'hot-water',
  },
  {
    'key': 'ice-cream',
    'label': 'ice-cream',
  },
  {
    'key': 'files',
    'label': 'files',
  },
  {
    'key': 'ice-cream-square',
    'label': 'ice-cream-square',
  },
  {
    'key': 'key',
    'label': 'key',
  },
  {
    'key': 'ice-tea',
    'label': 'ice-tea',
  },
  {
    'key': 'knife-fork',
    'label': 'knife-fork',
  },
  {
    'key': 'iphone',
    'label': 'iphone',
  },
  {
    'key': 'info-filled',
    'label': 'info-filled',
  },
  {
    'key': 'link',
    'label': 'link',
  },
  {
    'key': 'ice-drink',
    'label': 'ice-drink',
  },
  {
    'key': 'lightning',
    'label': 'lightning',
  },
  {
    'key': 'loading',
    'label': 'loading',
  },
  {
    'key': 'lollipop',
    'label': 'lollipop',
  },
  {
    'key': 'location-information',
    'label': 'location-information',
  },
  {
    'key': 'lock',
    'label': 'lock',
  },
  {
    'key': 'location-filled',
    'label': 'location-filled',
  },
  {
    'key': 'magnet',
    'label': 'magnet',
  },
  {
    'key': 'male',
    'label': 'male',
  },
  {
    'key': 'location',
    'label': 'location',
  },
  {
    'key': 'menu',
    'label': 'menu',
  },
  {
    'key': 'magic-stick',
    'label': 'magic-stick',
  },
  {
    'key': 'message-box',
    'label': 'message-box',
  },
  {
    'key': 'map-location',
    'label': 'map-location',
  },
  {
    'key': 'mic',
    'label': 'mic',
  },
  {
    'key': 'message',
    'label': 'message',
  },
  {
    'key': 'medal',
    'label': 'medal',
  },
  {
    'key': 'milk-tea',
    'label': 'milk-tea',
  },
  {
    'key': 'microphone',
    'label': 'microphone',
  },
  {
    'key': 'minus',
    'label': 'minus',
  },
  {
    'key': 'money',
    'label': 'money',
  },
  {
    'key': 'moon-night',
    'label': 'moon-night',
  },
  {
    'key': 'monitor',
    'label': 'monitor',
  },
  {
    'key': 'moon',
    'label': 'moon',
  },
  {
    'key': 'more',
    'label': 'more',
  },
  {
    'key': 'mostly-cloudy',
    'label': 'mostly-cloudy',
  },
  {
    'key': 'more-filled',
    'label': 'more-filled',
  },
  {
    'key': 'mouse',
    'label': 'mouse',
  },
  {
    'key': 'mug',
    'label': 'mug',
  },
  {
    'key': 'mute',
    'label': 'mute',
  },
  {
    'key': 'no-smoking',
    'label': 'no-smoking',
  },
  {
    'key': 'mute-notification',
    'label': 'mute-notification',
  },
  {
    'key': 'notification',
    'label': 'notification',
  },
  {
    'key': 'notebook',
    'label': 'notebook',
  },
  {
    'key': 'odometer',
    'label': 'odometer',
  },
  {
    'key': 'office-building',
    'label': 'office-building',
  },
  {
    'key': 'operation',
    'label': 'operation',
  },
  {
    'key': 'opportunity',
    'label': 'opportunity',
  },
  {
    'key': 'orange',
    'label': 'orange',
  },
  {
    'key': 'open',
    'label': 'open',
  },
  {
    'key': 'paperclip',
    'label': 'paperclip',
  },
  {
    'key': 'pear',
    'label': 'pear',
  },
  {
    'key': 'partly-cloudy',
    'label': 'partly-cloudy',
  },
  {
    'key': 'phone',
    'label': 'phone',
  },
  {
    'key': 'picture-filled',
    'label': 'picture-filled',
  },
  {
    'key': 'phone-filled',
    'label': 'phone-filled',
  },
  {
    'key': 'picture-rounded',
    'label': 'picture-rounded',
  },
  {
    'key': 'guide',
    'label': 'guide',
  },
  {
    'key': 'place',
    'label': 'place',
  },
  {
    'key': 'platform',
    'label': 'platform',
  },
  {
    'key': 'pie-chart',
    'label': 'pie-chart',
  },
  {
    'key': 'pointer',
    'label': 'pointer',
  },
  {
    'key': 'plus',
    'label': 'plus',
  },
  {
    'key': 'position',
    'label': 'position',
  },
  {
    'key': 'postcard',
    'label': 'postcard',
  },
  {
    'key': 'present',
    'label': 'present',
  },
  {
    'key': 'price-tag',
    'label': 'price-tag',
  },
  {
    'key': 'promotion',
    'label': 'promotion',
  },
  {
    'key': 'pouring',
    'label': 'pouring',
  },
  {
    'key': 'reading-lamp',
    'label': 'reading-lamp',
  },
  {
    'key': 'question-filled',
    'label': 'question-filled',
  },
  {
    'key': 'printer',
    'label': 'printer',
  },
  {
    'key': 'picture',
    'label': 'picture',
  },
  {
    'key': 'refresh-right',
    'label': 'refresh-right',
  },
  {
    'key': 'reading',
    'label': 'reading',
  },
  {
    'key': 'refresh-left',
    'label': 'refresh-left',
  },
  {
    'key': 'refresh',
    'label': 'refresh',
  },
  {
    'key': 'refrigerator',
    'label': 'refrigerator',
  },
  {
    'key': 'remove-filled',
    'label': 'remove-filled',
  },
  {
    'key': 'right',
    'label': 'right',
  },
  {
    'key': 'scale-to-original',
    'label': 'scale-to-original',
  },
  {
    'key': 'school',
    'label': 'school',
  },
  {
    'key': 'remove',
    'label': 'remove',
  },
  {
    'key': 'scissor',
    'label': 'scissor',
  },
  {
    'key': 'select',
    'label': 'select',
  },
  {
    'key': 'management',
    'label': 'management',
  },
  {
    'key': 'search',
    'label': 'search',
  },
  {
    'key': 'sell',
    'label': 'sell',
  },
  {
    'key': 'semi-select',
    'label': 'semi-select',
  },
  {
    'key': 'share',
    'label': 'share',
  },
  {
    'key': 'setting',
    'label': 'setting',
  },
  {
    'key': 'service',
    'label': 'service',
  },
  {
    'key': 'ship',
    'label': 'ship',
  },
  {
    'key': 'set-up',
    'label': 'set-up',
  },
  {
    'key': 'shopping-bag',
    'label': 'shopping-bag',
  },
  {
    'key': 'shop',
    'label': 'shop',
  },
  {
    'key': 'shopping-cart',
    'label': 'shopping-cart',
  },
  {
    'key': 'shopping-cart-full',
    'label': 'shopping-cart-full',
  },
  {
    'key': 'soccer',
    'label': 'soccer',
  },
  {
    'key': 'sold-out',
    'label': 'sold-out',
  },
  {
    'key': 'smoking',
    'label': 'smoking',
  },
  {
    'key': 'sort-down',
    'label': 'sort-down',
  },
  {
    'key': 'sort',
    'label': 'sort',
  },
  {
    'key': 'sort-up',
    'label': 'sort-up',
  },
  {
    'key': 'star',
    'label': 'star',
  },
  {
    'key': 'stamp',
    'label': 'stamp',
  },
  {
    'key': 'star-filled',
    'label': 'star-filled',
  },
  {
    'key': 'stopwatch',
    'label': 'stopwatch',
  },
  {
    'key': 'success-filled',
    'label': 'success-filled',
  },
  {
    'key': 'suitcase',
    'label': 'suitcase',
  },
  {
    'key': 'sugar',
    'label': 'sugar',
  },
  {
    'key': 'sunny',
    'label': 'sunny',
  },
  {
    'key': 'sunrise',
    'label': 'sunrise',
  },
  {
    'key': 'switch',
    'label': 'switch',
  },
  {
    'key': 'ticket',
    'label': 'ticket',
  },
  {
    'key': 'sunset',
    'label': 'sunset',
  },
  {
    'key': 'tickets',
    'label': 'tickets',
  },
  {
    'key': 'switch-button',
    'label': 'switch-button',
  },
  {
    'key': 'takeaway-box',
    'label': 'takeaway-box',
  },
  {
    'key': 'toilet-paper',
    'label': 'toilet-paper',
  },
  {
    'key': 'timer',
    'label': 'timer',
  },
  {
    'key': 'tools',
    'label': 'tools',
  },
  {
    'key': 'top-left',
    'label': 'top-left',
  },
  {
    'key': 'top',
    'label': 'top',
  },
  {
    'key': 'top-right',
    'label': 'top-right',
  },
  {
    'key': 'trend-charts',
    'label': 'trend-charts',
  },
  {
    'key': 'turn-off',
    'label': 'turn-off',
  },
  {
    'key': 'unlock',
    'label': 'unlock',
  },
  {
    'key': 'trophy',
    'label': 'trophy',
  },
  {
    'key': 'umbrella',
    'label': 'umbrella',
  },
  {
    'key': 'upload-filled',
    'label': 'upload-filled',
  },
  {
    'key': 'user-filled',
    'label': 'user-filled',
  },
  {
    'key': 'upload',
    'label': 'upload',
  },
  {
    'key': 'user',
    'label': 'user',
  },
  {
    'key': 'van',
    'label': 'van',
  },
  {
    'key': 'copy-document',
    'label': 'copy-document',
  },
  {
    'key': 'video-pause',
    'label': 'video-pause',
  },
  {
    'key': 'video-camera-filled',
    'label': 'video-camera-filled',
  },
  {
    'key': 'view',
    'label': 'view',
  },
  {
    'key': 'wallet',
    'label': 'wallet',
  },
  {
    'key': 'warning-filled',
    'label': 'warning-filled',
  },
  {
    'key': 'watch',
    'label': 'watch',
  },
  {
    'key': 'video-play',
    'label': 'video-play',
  },
  {
    'key': 'watermelon',
    'label': 'watermelon',
  },
  {
    'key': 'video-camera',
    'label': 'video-camera',
  },
  {
    'key': 'wallet-filled',
    'label': 'wallet-filled',
  },
  {
    'key': 'warning',
    'label': 'warning',
  },
  {
    'key': 'list',
    'label': 'list',
  },
  {
    'key': 'zoom-in',
    'label': 'zoom-in',
  },
  {
    'key': 'zoom-out',
    'label': 'zoom-out',
  },
  {
    'key': 'rank',
    'label': 'rank',
  },
  {
    'key': 'wind-power',
    'label': 'wind-power',
  },
])

const metaData = ref(props.meta)
if (!metaData.value.icon) {
  metaData.value.icon = options[0].label
}
</script>

<script>
export default {
  name: 'Icon',
}
</script>

<style lang="scss">
.gva-icon {
  color: rgb(132, 146, 166);
  font-size: 14px;
  margin-right: 10px;
}

.gva-select .el-input__inner {
  padding: 0 30px !important
}

.select__option_item {
  display: flex;
  align-items: center;
  justify-content: flex-start;
}
</style>

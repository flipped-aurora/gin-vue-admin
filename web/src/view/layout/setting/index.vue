<template>
  <el-drawer v-model="drawer" :title="t('system.systemConfig')" direction="rtl" :size="width">
    <div class="flex flex-col">
      <div class="mb-8">
        <div class="text-gray-800 dark:text-gray-100">{{ t('layout.setting.defaultTheme') }}</div>
        <div class="mt-2 text-sm p-2 flex items-center gap-2">
          <el-segmented
            v-model="config.darkMode"
            :options="options"
            size="default"
            @change="appStore.toggleDarkMode"
          />
        </div>
      </div>
      <div class="mb-8">
        <div class="text-gray-800 dark:text-gray-100">{{ t('layout.setting.themeColor') }}</div>
        <div class="mt-2 text-sm p-2 flex items-center gap-2">
          <div
            v-for="item in colors"
            :key="item"
            class="w-5 h-5 rounded cursor-pointer flex items-center justify-center"
            :style="`background:${item}`"
            @click="appStore.togglePrimaryColor(item)"
          >
            <el-icon v-if="config.primaryColor === item">
              <Select />
            </el-icon>
          </div>
          <el-color-picker
            v-model="customColor"
            @change="appStore.togglePrimaryColor"
          />
        </div>
      </div>
      <div class="mb-8">
        <div class="text-gray-800 dark:text-gray-100">{{ t('layout.setting.interfaceDisplay') }}</div>
        <div class="mt-2 text-sm p-2">
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.showWaterMark') }}</div>
            <el-switch
              v-model="config.show_watermark"
              @change="appStore.toggleConfigWatermark"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.grayMode') }}</div>
            <el-switch v-model="config.grey" @change="appStore.toggleGrey" />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.colorFadeMode') }}</div>
            <el-switch
              v-model="config.weakness"
              @change="appStore.toggleWeakness"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.menuMode') }}</div>
            <el-segmented
              v-model="config.side_mode"
              :options="sideModes"
              size="default"
              @change="appStore.toggleSideModel"
            />
            <!-- <el-select
              v-model="config.side_mode"
              @change="handleSideModelChange"
            >
              <el-option value="normal" label="标准模式" />
              <el-option value="head" label="顶部导航栏" />
              <el-option value="multilayer" disabled label="多侧边导航模式" />
            </el-select> -->
          </div>

          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.showTabs') }}</div>
            <el-switch
              v-model="config.showTabs"
              @change="appStore.toggleTabs"
            />
          </div>
        </div>
      </div>

      <div class="mb-8">
        <div class="text-gray-800 dark:text-gray-100">{{ t('layout.setting.layputSizeConfig') }}</div>
        <div class="mt-2 text-sm p-2">
          <div class="flex items-center justify-between mb-2">
            <div>{{ t('layout.setting.sidebarExpandedWidth') }}</div>
            <el-input-number
              v-model="config.layout_side_width"
              :min="150"
              :max="400"
              :step="10"
            ></el-input-number>
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ t('layout.setting.sidebarShrinkWidth') }}</div>
            <el-input-number
              v-model="config.layout_side_collapsed_width"
              :min="60"
              :max="100"
            ></el-input-number>
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ t('layout.setting.sidebarItemHeight') }}</div>
            <el-input-number
              v-model="config.layout_side_item_height"
              :min="30"
              :max="50"
            ></el-input-number>
          </div>
        </div>
      </div>

      <el-alert type="warning" :closable="false">
        {{ t('layout.setting.configSaveNote1') }}
        <el-tag>config.json</el-tag> {{ t('layout.setting.configSaveNote2') }}
      </el-alert>

      <el-button type="primary" class="mt-4" @click="copyConfig"
        >{{ t('layout.setting.copyConfig') }}</el-button
      >
    </div>
  </el-drawer>
</template>

<script setup>
import { useAppStore } from "@/pinia";
import { storeToRefs } from "pinia";
import { ref, computed } from "vue";
import { ElMessage } from "element-plus";
import { useI18n } from 'vue-i18n'; // added by mohamed hassan to support multilanguage

const appStore = useAppStore();
const { config, device } = storeToRefs(appStore);

const { t } = useI18n() // added by mohamed hassan to support multilanguage
defineOptions({
  name: "GvaSetting",
});

const width = computed(() => {
  return device.value === "mobile" ? "100%" : "500px";
});

const colors = [
  "#EB2F96",
  "#3b82f6",
  "#2FEB54",
  "#EBEB2F",
  "#EB2F2F",
  "#2FEBEB",
];

const drawer = defineModel("drawer", {
  default: true,
  type: Boolean,
});

const options = ["dark", "light", "auto"];
const sideModes = [
  {
    label : t('layout.setting.normalMode'),
    value : "normal"
  },
  {
    label : t('layout.setting.topMenuBarMode'),
    value: "head"
  },
  {
    label : t('layout.setting.combinationMode'),
    value: "combination"
  }
];

const copyConfig = () => {
  const input = document.createElement("textarea");
  input.value = JSON.stringify(config.value);
  // 添加回车
  input.value = input.value.replace(/,/g, ",\n");
  document.body.appendChild(input);
  input.select();
  document.execCommand("copy");
  document.body.removeChild(input);
  ElMessage.success(t('layout.setting.copyConfigSuccess'));
};

const customColor = ref("");

const handleSideModelChange = (e) => {
  console.log(e);
};
</script>

<style lang="scss" scoped>
::v-deep(.el-drawer__header) {
  @apply border-gray-400 dark:border-gray-600;
}
</style>

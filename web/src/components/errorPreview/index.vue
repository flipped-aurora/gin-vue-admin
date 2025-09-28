<template>
  <div 
    class="fixed inset-0 bg-black/40 dark:bg-black/60 flex items-center justify-center z-[999]"
    @click.self="closeModal"
  >
    <div class="bg-white dark:bg-gray-800 rounded-xl shadow-dialog dark:shadow-lg w-full max-w-md mx-4 transform transition-all duration-300 ease-in-out border border-transparent dark:border-gray-700">
      <!-- 弹窗头部 -->
      <div class="p-5 border-b border-gray-100 dark:border-gray-700 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-800 dark:text-gray-100">{{ displayData.title }}</h3>
        <div class="text-gray-400 dark:text-gray-300 hover:text-gray-600 dark:hover:text-gray-200 transition-colors cursor-pointer" @click="closeModal">
          <close class="h-6 w-6" />
        </div>
      </div>
      
      <!-- 弹窗内容 -->
      <div class="p-6 pt-0">
        <!-- 错误类型 -->
        <div class="mb-4">
          <div class="text-xs font-medium text-gray-500 dark:text-gray-400 uppercase mb-2">错误类型</div>
          <div class="flex items-center gap-2">
            <lock v-if="displayData.icon === 'lock'" :class="['w-5 h-5', displayData.color]" />
            <warn v-if="displayData.icon === 'warn'" :class="['w-5 h-5', displayData.color]" />
            <server v-if="displayData.icon === 'server'" :class="['w-5 h-5', displayData.color]" />
            <span class="font-medium text-gray-800 dark:text-gray-100">{{ displayData.type }}</span>
          </div>
        </div>
        
        <!-- 具体错误 -->
        <div class="mb-6">
          <div class="text-xs font-medium text-gray-500 dark:text-gray-400 uppercase mb-2">具体错误</div>
          <div class="bg-gray-100 dark:bg-gray-900/40 rounded-lg p-3 text-sm text-gray-700 dark:text-gray-200 leading-relaxed">
            {{ displayData.message }}
          </div>
        </div>
        
        <!-- 提示信息 -->
        <div v-if="displayData.tips">
          <div class="text-xs font-medium text-gray-500 dark:text-gray-400 uppercase mb-2">提示</div>
          <div class="flex items-center gap-2">
            <idea class="text-blue-500 dark:text-blue-400 w-5 h-5" />
            <p class="text-sm text-gray-600 dark:text-gray-300">{{ displayData.tips }}</p>
          </div>
        </div>
      </div>
      
      <!-- 弹窗底部 -->
      <div class="py-2 px-4 border-t border-gray-100 dark:border-gray-700 flex justify-end">
        <div class="px-4 py-2 bg-blue-600 dark:bg-blue-500 text-white dark:text-gray-100 rounded-lg hover:bg-blue-700 dark:hover:bg-blue-600 transition-colors font-medium text-sm shadow-sm cursor-pointer" @click="handleConfirm">
          确定
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, computed } from 'vue';

const props = defineProps({
  errorData: {
    type: Object,
    required: true
  }
});

const emits = defineEmits(['close', 'confirm']);

const presetErrors = {
  500: {
    title: '检测到接口错误',
    type: '服务器发生内部错误',
    icon: 'server',
    color: 'text-red-500 dark:text-red-400',
    tips: '此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存'
  },
  404: {
    title: '资源未找到',
    type: 'Not Found',
    icon: 'warn',
    color: 'text-orange-500 dark:text-orange-400',
    tips: '此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格'
  },
  401: {
    title: '身份认证失败',
    type: '身份令牌无效',
    icon: 'lock',
    color: 'text-purple-500 dark:text-purple-400',
    tips: '您的身份认证已过期或无效，请重新登录。'
  },
  'network': {
    title: '网络错误',
    type: 'Network Error',
    icon: 'fa-wifi-slash',
    color: 'text-gray-500 dark:text-gray-400',
    tips: '无法连接到服务器，请检查您的网络连接。'
  }
};

const displayData = computed(() => {
  const preset = presetErrors[props.errorData.code];
  if (preset) {
    return {
      ...preset,
      message: props.errorData.message || '没有提供额外信息。'
    };
  }

  return {
    title: '未知错误',
    type: '检测到请求错误',
    icon: 'fa-question-circle',
    color: 'text-gray-400 dark:text-gray-300',
    message: props.errorData.message || '发生了一个未知错误。',
    tips: '请检查控制台获取更多信息。'
  };
});

const closeModal = () => {
   emits('close')
};

const handleConfirm = () => {
  emits('confirm', props.errorData.code);
  closeModal();
};
</script>

<template>
  <div>
    <iframe ref="iframe" class="w-full h-screen"></iframe>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from 'vue';

const props = defineProps({
  htmlContent: {
    type: String,
    required: true,
  },
});

const iframe = ref(null);

const updateIframeContent = () => {
  if (iframe.value) {
    const doc = iframe.value.contentDocument || iframe.value.contentWindow.document;
    doc.open();
    doc.write(`
      <html>
        <head>
          <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
        </head>
        <body>${props.htmlContent}</body>
      </html>
    `);
    doc.close();
  }
};

onMounted(() => {
  updateIframeContent();
});

watch(() => props.htmlContent, () => {
  updateIframeContent();
});
</script>

<style scoped>

</style>

<template>
  <iframe
      ref="iframe"
      class="w-full border-0"
      :style="{ height: iframeHeight + 'px' }"
  ></iframe>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'

const props = defineProps({
  html: {
    type: String,
    required: true
  }
})

const iframe = ref(null)
const iframeHeight = ref(400) // Default height

const renderContent = () => {
  if (!iframe.value) return

  const doc = iframe.value.contentDocument || iframe.value.contentWindow.document

  const htmlTemplate = `
    <!DOCTYPE html>
    <html>
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <!-- Production version of Vue -->
      <script src="https://unpkg.com/vue@3/dist/vue.global.prod.js"><\/script>
      <!-- Element Plus -->
      <link rel="stylesheet" href="https://unpkg.com/element-plus/dist/index.css" />
      <script src="https://unpkg.com/element-plus/dist/index.full.min.js"><\/script>
      <!-- Tailwind CSS -->
      <script src="https://cdn.tailwindcss.com"><\/script>
      <style>
        body { margin: 0; padding: 10px; }
      </style>
    </head>
    <body>
      <div id="app"></div>
      <script>
        // Initialize Vue app
        const app = Vue.createApp({
          template: \`${props.html.replace(/`/g, '\\`')}\`,
          data() {
            return {}
          }
        });

        // Use Element Plus
        app.use(ElementPlus);

        // Mount the application
        app.mount('#app');

        // Adjust iframe height
        setTimeout(() => {
          const height = document.body.scrollHeight;
          window.parent.postMessage({
            type: 'resize',
            height: height
          }, '*');
        }, 100);
      <\/script>
    </body>
    </html>
  `

  doc.open()
  doc.write(htmlTemplate)
  doc.close()
}

// Listen for height updates from iframe
onMounted(() => {
  window.addEventListener('message', (event) => {
    if (event.data && event.data.type === 'resize') {
      iframeHeight.value = event.data.height + 20 // Add padding
    }
  })

  // Render with slight delay to ensure iframe is ready
  setTimeout(() => {
    renderContent()
  }, 50)
})

// Re-render when HTML changes
watch(() => props.html, () => {
  setTimeout(() => {
    renderContent()
  }, 50)
})
</script>

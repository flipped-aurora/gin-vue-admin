<template>
  <div class="chat-wrapper">
    <div class="chat">
      <div class="chat-container">
        <div ref="body" class="chat-body">
          <div v-for="(message, index) in messages" :key="index" class="chat-message">
            <el-row>
              <el-col :span="1">
                <el-avatar
                  :icon="message.icon"
                  :size="32"
                  fit="cover"
                />
              </el-col>
              <el-col :span="23">
                <div class="chat-message-text">
                  <pre><div v-html="md.render(message.text)" /></pre>
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
        <div class="chat-footer">
          <el-input
            :model-value="input"
            :size="'large'"
            placeholder="请输入要提问的内容"
            @keydown.enter="sendMessage"
            @update:modelValue="input = $event"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ChatView'
}
</script>

<script setup>
import MarkdownIt from 'markdown-it'
import { chat } from '@/api/chat'
import { ref } from 'vue'
import { UserFilled, Tools } from '@element-plus/icons-vue'

const md = new MarkdownIt()

const messages = ref([
  {
    'username': '用户',
    'text': '提问',
    'icon': UserFilled
  },
  {
    'username': 'Bot',
    'text': '# test\n' +
        '```go\n' +
        '// 中间件函数  \n' +
        'func myMiddleware(next http.Handler) http.Handler {  \n' +
        '\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n' +
        '\t\tlog.Println(r.Method, r.Body)  \n' +
        '\t\t// 在请求被处理之前执行的逻辑代码  \n' +
        '\t\tlog.Println("Middleware processing request...")  \n' +
        '\t\tlog.Println("Config allow Origin")  \n' +
        '\t\tw.Header().Set("Access-Control-Allow-Origin", "*")  \n' +
        '     \n' +
        '\t\tif r.Method == http.MethodOptions {  \n' +
        '\t        w.WriteHeader(http.StatusOK)  \n' +
        '\t    }  \n' +
        '\t    log.Println(r.Method)  \n' +
        '\t    log.Println("Config allow Origin done")  \n' +
        '\t    next.ServeHTTP(w, r)  \n' +
        '\t    // 在请求被处理之后执行的逻辑代码  \n' +
        '\t    // ...\n' +
        '\t})  \n' +
        '}\n' +
        '```\n' +
        '```go\n' +
        '// 中间件函数  \n' +
        'func myMiddleware(next http.Handler) http.Handler {  \n' +
        '\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n' +
        '\t\tlog.Println(r.Method, r.Body)  \n' +
        '\t\t// 在请求被处理之前执行的逻辑代码  \n' +
        '\t\tlog.Println("Middleware processing request...")  \n' +
        '\t\tlog.Println("Config allow Origin")  \n' +
        '\t\tw.Header().Set("Access-Control-Allow-Origin", "*")  \n' +
        '     \n' +
        '\t\tif r.Method == http.MethodOptions {  \n' +
        '\t        w.WriteHeader(http.StatusOK)  \n' +
        '\t    }  \n' +
        '\t    log.Println(r.Method)  \n' +
        '\t    log.Println("Config allow Origin done")  \n' +
        '\t    next.ServeHTTP(w, r)  \n' +
        '\t    // 在请求被处理之后执行的逻辑代码  \n' +
        '\t    // ...\n' +
        '\t})  \n' +
        '}\n' +
        '```',
    'icon': Tools
  },
])

const getAvatarUrl = async(username) => {
  // 返回头像图片的 URL，根据实际情况进行修改
  return (username === 'user') ? UserFilled : Tools
}

const input = ref('')

const sendMessage = async() => {
  if (input.value.trim() === '') {
    return
  }

  try {
    const response = chat({
      username: 'user',
      text: input.value,
    })

    if (response.data.code === 200) {
      messages.value.push({
        username: 'user',
        text: input.value,
        icon: getAvatarUrl('user')
      })
      input.value = ''

      // 检查 response 中是否包含有效数据
      if (response.data && response.data.username && response.data.text) {
        messages.value.push({
          username: response.data.username, // 使用服务器返回的 "BOT" 用户名
          text: response.data.text, // 使用服务器返回的回复内容
          icon: getAvatarUrl(response.data.username)
        })
      }
    } else {
      // 提示用户系统存在问题，要求重新输入
      alert('系统存在问题，请重新输入')
    }
  } catch (error) {
    console.error('Error sending message:', error)
  }
}

</script>

<style scoped lang="scss">
.el-main {
  position: relative;
}

.chat-wrapper {
  position: absolute;
  top: 0;
  left: 50%;
  width: 70%;
  height: 90%;
  transform: translateX(-50%);
  display: flex;
  flex-direction: column;
}

.chat {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-body {
  flex: 1;
  overflow-y: auto;
}

.chat-message-text div {
  display: block;
  padding: 10px;
  border: 1px solid #e5e5e5;
  border-radius: 4px;
  background-color: #f5f5f5;
  color: #333;
  font-size: 14px;
  max-height: calc(1.5em * 20); // 假设行高为 1.5em
  overflow-y: auto;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.chat-footer {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
}

</style>

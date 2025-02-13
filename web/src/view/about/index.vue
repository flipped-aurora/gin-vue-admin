<template>
  <div class="mt-2">
    <div class="flex flex-col md:flex-row gap-4">
      <div class="w-full md:w-1/2">
        <el-card class="min-w-96">
          <template #header>
            <el-divider>gin-vue-admin</el-divider>
          </template>
          <div>
            <div class="w-full flex items-center justify-center">
              <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="org-img dom-center"
                    src="@/assets/logo.png"
                    alt="gin-vue-admin"
                  />
                </a>
            </div>
            <div class="w-full flex items-center justify-around">
              <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/watchers/flipped-aurora/gin-vue-admin.svg?label=Watch"
                    alt=""
                  />
                </a>
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/stars/flipped-aurora/gin-vue-admin.svg?style=social"
                    alt=""
                  />
                </a>
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/forks/flipped-aurora/gin-vue-admin.svg?label=Fork"
                    alt=""
                  />
                </a>
            </div>
          </div>
        </el-card>
        <el-card class="min-w-96 mt-5">
          <template #header>
            <div>flipped-aurora团队</div>
          </template>
          <div>
            <div class="w-full flex items-center justify-center">
                <a href="https://github.com/flipped-aurora">
                  <img
                    class="org-img dom-center"
                    src="@/assets/flipped-aurora.png"
                    alt="flipped-aurora"
                  />
                </a>
              </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4 mt-4">
              <div v-for="(item, index) in members" :key="index" class="min-h-10 flex items-center">
                <a :href="item.html_url" class="flex items-center group">
                  <img class="w-8 h-8 rounded-full" :src="item.avatar_url" />
                  <el-link
                    class="text-blue-700 ml-2 text-lg font-bold font-sans break-all"
                    >{{ item.login }}</el-link
                  >
                </a>
              </div>
            </div>
          </div>
        </el-card>
      </div>
      <div class="w-full md:w-1/2">
        <el-card>
          <template #header>
            <div>提交记录</div>
          </template>
          <div class="h-[calc(100vh-300px)] overflow-y-auto">
            <el-timeline>
              <el-timeline-item
                v-for="(item, index) in dataTimeline"
                :key="index"
                :timestamp="item.from"
                placement="top"
              >
                <el-card>
                  <h4>{{ item.title }}</h4>
                  <p>{{ item.message }}</p>
                </el-card>
              </el-timeline-item>
            </el-timeline>
          </div>
         <div class="w-full flex items-center justify-center">
          <el-button class="load-more" type="primary" link @click="loadMore">
            Load more
          </el-button>
         </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { Commits, Members } from '@/api/github'
  import { formatTimeToStr } from '@/utils/date'
  const page = ref(0)

  defineOptions({
    name: 'About'
  })

  const loadMore = () => {
    page.value++
    loadCommits()
  }

  const dataTimeline = ref([])
  const loadCommits = () => {
    Commits(page.value).then(({ data }) => {
      data.forEach((element) => {
        if (element.commit.message) {
          dataTimeline.value.push({
            from: formatTimeToStr(element.commit.author.date, 'yyyy-MM-dd'),
            title: element.commit.author.name,
            showDayAndMonth: true,
            message: element.commit.message
          })
        }
      })
    })
  }

  const members = ref([])
  const loadMembers = () => {
    Members().then(({ data }) => {
      members.value = data
      members.value.sort()
    })
  }

  loadCommits()
  loadMembers()
</script>

<style scoped>
  .avatar-img {
    float: left;
    height: 40px;
    width: 40px;
    border-radius: 50%;
    -webkit-border-radius: 50%;
    -moz-border-radius: 50%;
    margin-top: 15px;
  }

  .org-img {
    height: 150px;
    width: 150px;
  }

  .dom-center {
    margin-left: 50%;
    transform: translateX(-50%);
  }
</style>

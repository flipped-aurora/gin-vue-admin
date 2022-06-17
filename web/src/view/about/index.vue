<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <template #header>
            <el-divider>gin-vue-admin</el-divider>
          </template>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="org-img dom-center"
                    src="@/assets/logo.png"
                    alt="gin-vue-admin"
                  >
                </a>
              </el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/watchers/flipped-aurora/gin-vue-admin.svg?label=Watch"
                    alt=""
                  >
                </a>
              </el-col>
              <el-col :span="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/stars/flipped-aurora/gin-vue-admin.svg?style=social"
                    alt=""
                  >
                </a>
              </el-col>
              <el-col :span="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/forks/flipped-aurora/gin-vue-admin.svg?label=Fork"
                    alt=""
                  >
                </a>
              </el-col>
            </el-row>
          </div>
        </el-card>
        <el-card style="margin-top: 20px">
          <template #header>
            <div>flipped-aurora团队</div>
          </template>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                <a href="https://github.com/flipped-aurora">
                  <img
                    class="org-img dom-center"
                    src="@/assets/flipped-aurora.png"
                    alt="flipped-aurora"
                  >
                </a>
              </el-col>
            </el-row>
            <el-row style="margin-left: 40px" :gutter="20">
              <el-col v-for="(item, index) in members" :key="index" :span="8">
                <a :href="item.html_url">
                  <img class="avatar-img" :src="item.avatar_url">
                  <a class="author-name" style="">{{ item.login }}</a>
                </a>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <div>提交记录</div>
          </template>
          <div>
            <el-timeline>
              <el-timeline-item
                v-for="(item,index) in dataTimeline"
                :key="index"
                timestamp="2018/4/12"
                placement="top"
              >
                <el-card>
                  <h4>{{ item.title }}</h4>
                  <p>{{ item.message }}</p>
                </el-card>
              </el-timeline-item>
            </el-timeline>
          </div>
          <el-button
            class="load-more"
            type="primary" link
            @click="loadMore"
          >Load more</el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'About',
}
</script>

<script setup>
import { ref } from 'vue'
import { Commits, Members } from '@/api/github'
const page = ref(0)

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
          from: new Date(element.commit.author.date),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
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
.load-more {
  margin-left: 120px;
}

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

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>

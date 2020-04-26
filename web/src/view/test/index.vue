<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <div slot="header">gin-vue-admin</div>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="org-img dom-center"
                    src="@/assets/logo.png"
                    alt="gin-vue-admin"
                  />
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
                  />
                </a>
              </el-col>
              <el-col :span="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/stars/flipped-aurora/gin-vue-admin.svg?style=social"
                    alt=""/></a
              ></el-col>
              <el-col :span="8">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/forks/flipped-aurora/gin-vue-admin.svg?label=Fork"
                    alt=""/></a
              ></el-col>
            </el-row>
          </div>
        </el-card>
        <el-card style="margin-top: 20px">
          <div slot="header">flipped-aurora团队</div>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                <a href="https://github.com/flipped-aurora">
                  <img
                    class="org-img dom-center"
                    src="@/assets/flipped-aurora.png"
                    alt="flipped-aurora"
                  />
                </a>
              </el-col>
            </el-row>
            <el-row style="margin-left: 40px" :gutter="20">
              <template v-for="(item, index) in members">
                <el-col :span="8" :key="index">
                  <a :href="item.html_url">
                    <img class="avatar-img" :src="item.avatar_url" />
                    <a class="author-name" style="">{{ item.login }}</a>
                  </a>
                </el-col>
              </template>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <div slot="header">
            提交记录
          </div>
          <div>
            <Timeline
              :timeline-items="dataTimeline"
              :message-when-no-items="messageWhenNoItems"
              :uniqueTimeline="true"
              :unique-year="true"
              :show-day-and-month="true"
              order="desc"
              dateLocale="zh-CN"
            />
          </div>
          <el-button class="load-more" @click="loadMore" type="text"
            >Load more</el-button
          >
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { Commits, Members } from "@/api/github";
import Timeline from "timeline-vuejs";
export default {
  name: "Test",
  components: {
    Timeline,
  },
  data() {
    return {
      messageWhenNoItems: "There arent commits",
      members: [],
      dataTimeline: [],
      page: 0,
    };
  },
  created() {
    this.loadCommits();
    this.loadMembers();
  },
  methods: {
    loadMore() {
      this.page++;
      this.loadCommits();
    },
    loadCommits() {
      Commits(this.page).then(({ data }) => {
        data.forEach((element) => {
          if (element.commit.message) {
            this.dataTimeline.push({
              from: new Date(element.commit.author.date),
              title: element.commit.author.name,
              showDayAndMonth: true,
              description: `<a style="color: #26191b" href="${element.html_url}">${element.commit.message}</a>`,
            });
          }
        });
      });
    },
    loadMembers() {
      Members().then(({ data }) => {
        this.members = data;
        this.members.sort();
      });
    },
  },
};
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

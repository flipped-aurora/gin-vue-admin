<template>
  <div class="commit-table">
    <div class="commit-table-title">
      更新日志
    </div>
    <div class="log">
      <div v-for="(item,key) in dataTimeline" :key="key" class="log-item">
        <div class="flex-1 flex key-box"><span class="key" :class="key<3&&'top'">{{ key+1 }}</span></div>
        <div class="flex-5 flex message">{{ item.message }}</div>
        <div class="flex-3 flex form">{{ item.from }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { Commits } from '@/api/github'
import { formatTimeToStr } from '@/utils/date.js'
export default {
  data() {
    return {
      loading: true,
      dataTimeline: [],
    }
  },
  created() {
    this.loadCommits()
  },
  methods: {
    loadCommits() {
      Commits(0).then(({ data }) => {
        this.loading = false
        data.forEach((element, index) => {
          if (element.commit.message && index < 10) {
            this.dataTimeline.push({
              from: formatTimeToStr(element.commit.author.date, 'yyyy-MM-dd'),
              title: element.commit.author.name,
              showDayAndMonth: true,
              message: element.commit.message,
            })
          }
        })
      })
    },
  }
}
</script>
<style lang="scss" scoped>
.commit-table{
    background-color: #fff;
    height: 400px;
    &-title{
        font-weight: 600;
        margin-bottom: 12px;
    }
    .log{
      &-item{
        display: flex;
        justify-content: space-between;
        margin-top: 14px;
        .key-box{
          justify-content: center;
        }
        .key{
          &.top{
            background: #314659;
            color: #FFFFFF;;
          }
          display: inline-flex;
          justify-content: center;
          align-items: center;
          width: 20px;
          height: 20px;
          border-radius: 50%;
          background: #F0F2F5;
          text-align: center;
          color:rgba($color: #000000, $alpha: 0.65)
        }
        .message{
          color: rgba(0, 0, 0, 0.65);
        }
        .form{
          color: rgba(0, 0, 0, 0.65);
          margin-left: 12px;
        }
        .flex{
          line-height: 20px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
        .flex-1{
          flex:1;
        }
        .flex-2{
          flex:2;
        }
        .flex-3{
          flex:3;
        }
        .flex-4{
          flex:4;
        }
        .flex-5{
          flex:5;
        }
      }
    }
}
</style>

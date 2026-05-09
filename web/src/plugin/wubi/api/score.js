import service from '@/utils/request'

/**
 * 提交一次五笔练习成绩
 * @param {Object} data
 * @param {string} data.mode 练习模式 char/word/article
 * @param {number} data.totalCount 题目总数
 * @param {number} data.correctCount 正确数
 * @param {number} data.errorCount 错误数
 * @param {number} data.durationSecond 用时(秒)
 * @param {number} data.accuracy 准确率(0-100)
 * @param {number} data.wpm 每分钟字数
 */
export const submitScore = (data) => service({
  url: '/wubi/submitScore',
  method: 'post',
  data
})

/**
 * 删除自己的一条成绩
 * @param {Object} params
 * @param {number|string} params.ID
 */
export const deleteScore = (params) => service({
  url: '/wubi/deleteScore',
  method: 'delete',
  params
})

/**
 * 分页获取我的成绩
 * @param {Object} params
 * @param {number} params.page
 * @param {number} params.pageSize
 * @param {string} [params.mode]
 */
export const getMyScores = (params) => service({
  url: '/wubi/getMyScores',
  method: 'get',
  params
})

/**
 * 排行榜
 * @param {Object} params
 * @param {string} [params.mode]
 * @param {number} [params.limit]
 * @param {number} [params.days]
 * @param {string} [params.order] wpm/accuracy
 */
export const getLeaderboard = (params) => service({
  url: '/wubi/getLeaderboard',
  method: 'get',
  params
})

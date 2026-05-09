// 五笔 86 版常用字字根表 (示例数据,可根据需要扩充)
// 字段含义:
//   ch:   汉字
//   code: 五笔 86 版编码 (可能不完全准确,仅作练习提示)
//   pinyin: 拼音(便于不熟悉五笔者也能看到读音)
// 这里只放了 60 多个高频字, 可以在数据库或后端扩展更全面的题库
export const wubiDict = [
  { ch: '一', code: 'ggll', pinyin: 'yī' },
  { ch: '工', code: 'aaaa', pinyin: 'gōng' },
  { ch: '王', code: 'gggg', pinyin: 'wáng' },
  { ch: '土', code: 'ffff', pinyin: 'tǔ' },
  { ch: '大', code: 'dddd', pinyin: 'dà' },
  { ch: '木', code: 'ssss', pinyin: 'mù' },
  { ch: '目', code: 'hhhh', pinyin: 'mù' },
  { ch: '日', code: 'jjjj', pinyin: 'rì' },
  { ch: '口', code: 'kkkk', pinyin: 'kǒu' },
  { ch: '田', code: 'llll', pinyin: 'tián' },
  { ch: '山', code: 'mmmm', pinyin: 'shān' },
  { ch: '禾', code: 'tttt', pinyin: 'hé' },
  { ch: '白', code: 'rrrr', pinyin: 'bái' },
  { ch: '月', code: 'eeee', pinyin: 'yuè' },
  { ch: '人', code: 'wwww', pinyin: 'rén' },
  { ch: '金', code: 'qqqq', pinyin: 'jīn' },
  { ch: '言', code: 'yyyy', pinyin: 'yán' },
  { ch: '立', code: 'uuuu', pinyin: 'lì' },
  { ch: '水', code: 'iiii', pinyin: 'shuǐ' },
  { ch: '火', code: 'oooo', pinyin: 'huǒ' },
  { ch: '已', code: 'nnnn', pinyin: 'yǐ' },
  { ch: '子', code: 'bbbb', pinyin: 'zǐ' },
  { ch: '女', code: 'vvvv', pinyin: 'nǚ' },
  { ch: '又', code: 'cccc', pinyin: 'yòu' },

  { ch: '中', code: 'khk', pinyin: 'zhōng' },
  { ch: '国', code: 'lgyi', pinyin: 'guó' },
  { ch: '我', code: 'trnt', pinyin: 'wǒ' },
  { ch: '你', code: 'wqiy', pinyin: 'nǐ' },
  { ch: '他', code: 'wbn', pinyin: 'tā' },
  { ch: '好', code: 'vbg', pinyin: 'hǎo' },
  { ch: '是', code: 'jghu', pinyin: 'shì' },
  { ch: '不', code: 'gii', pinyin: 'bù' },
  { ch: '在', code: 'dhfd', pinyin: 'zài' },
  { ch: '上', code: 'hhgg', pinyin: 'shàng' },
  { ch: '下', code: 'ghi', pinyin: 'xià' },
  { ch: '来', code: 'goi', pinyin: 'lái' },
  { ch: '去', code: 'fcu', pinyin: 'qù' },
  { ch: '了', code: 'bnh', pinyin: 'le' },
  { ch: '的', code: 'rqyy', pinyin: 'de' },
  { ch: '有', code: 'def', pinyin: 'yǒu' },
  { ch: '会', code: 'wfcu', pinyin: 'huì' },
  { ch: '说', code: 'yuk', pinyin: 'shuō' },
  { ch: '想', code: 'shnu', pinyin: 'xiǎng' },
  { ch: '看', code: 'rhf', pinyin: 'kàn' },
  { ch: '听', code: 'krh', pinyin: 'tīng' },
  { ch: '吃', code: 'ktnn', pinyin: 'chī' },
  { ch: '喝', code: 'kjqn', pinyin: 'hē' },
  { ch: '走', code: 'fhu', pinyin: 'zǒu' },
  { ch: '跑', code: 'khqn', pinyin: 'pǎo' },
  { ch: '学', code: 'ipbf', pinyin: 'xué' },
  { ch: '生', code: 'tgd', pinyin: 'shēng' },
  { ch: '老', code: 'ftxb', pinyin: 'lǎo' },
  { ch: '师', code: 'jgmh', pinyin: 'shī' },
  { ch: '家', code: 'peu', pinyin: 'jiā' },
  { ch: '电', code: 'jnv', pinyin: 'diàn' },
  { ch: '脑', code: 'eybh', pinyin: 'nǎo' },
  { ch: '车', code: 'lgnh', pinyin: 'chē' },
  { ch: '门', code: 'uyhn', pinyin: 'mén' },
  { ch: '心', code: 'nyny', pinyin: 'xīn' },
  { ch: '手', code: 'rtgh', pinyin: 'shǒu' },
  { ch: '春', code: 'dwjf', pinyin: 'chūn' },
  { ch: '夏', code: 'dhtu', pinyin: 'xià' },
  { ch: '秋', code: 'toy', pinyin: 'qiū' },
  { ch: '冬', code: 'tuu', pinyin: 'dōng' }
]

// shuffle 简单的洗牌函数
export const shuffle = (arr) => {
  const list = arr.slice()
  for (let i = list.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1))
    ;[list[i], list[j]] = [list[j], list[i]]
  }
  return list
}

// pickRandom 抽取 n 个题
export const pickRandom = (n = 20) => {
  return shuffle(wubiDict).slice(0, Math.min(n, wubiDict.length))
}

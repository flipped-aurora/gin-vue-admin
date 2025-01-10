/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  important: true,
  theme: {
    extend: {
      backgroundColor: {
        main: '#F5F5F5'
      },
      textColor: {
        active: 'var(--el-color-primary)'
      },
      boxShadowColor: {
        active: 'var(--el-color-primary)'
      },
      borderColor: {
        'table-border': 'var(--el-border-color-lighter)'
      }
    }
  },
  safelist: [
    /*
      1) 常见的自定义颜色写法，如 bg-[#xxxxxx]、text-[#xxxxxx]、border-[#xxxxxx]
         如果 LLM/接口频繁返回各种 [#RRGGBB] 形式，这个 pattern 可以保留它们的CSS。
    */
    { pattern: /^bg-\[.*\]$/ },
    { pattern: /^text-\[.*\]$/ },
    { pattern: /^border-\[.*\]$/ },

    /*
      2) Tailwind 默认调色板里的常见前缀
         下面以 (red|green|blue|gray|indigo|...) 为例，你可根据自己项目加减。
         同时包括不同深度 (50,100,200,...,900)，也可再详细拆分。
    */
    {
      pattern: /^(bg|text|border)-(red|green|blue|gray|indigo|yellow|purple|pink|cyan|teal|orange|amber|lime|emerald|fuchsia|rose|sky|violet|stone|neutral)-(50|100|200|300|400|500|600|700|800|900)$/
    },

    /*
      3) 大小相关（padding、margin等）：
         m-*, p-* 以及更精细的 mt-*, mb-*...
         匹配数字、1/2、px、等常见写法（也包括 m-auto）。
         你可以根据自己需求加减。
    */
    { pattern: /^m-[0-9]+$/ },
    { pattern: /^mx-[0-9]+$/ },
    { pattern: /^my-[0-9]+$/ },
    { pattern: /^mt-[0-9]+$/ },
    { pattern: /^mr-[0-9]+$/ },
    { pattern: /^mb-[0-9]+$/ },
    { pattern: /^ml-[0-9]+$/ },
    { pattern: /^m-(auto|px)$/ },

    { pattern: /^p-[0-9]+$/ },
    { pattern: /^px-[0-9]+$/ },
    { pattern: /^py-[0-9]+$/ },
    { pattern: /^pt-[0-9]+$/ },
    { pattern: /^pr-[0-9]+$/ },
    { pattern: /^pb-[0-9]+$/ },
    { pattern: /^pl-[0-9]+$/ },
    { pattern: /^p-(auto|px)$/ },

    /*
      4) 宽高相关: w-*, h-*，以及自定义的 w-[300px]、h-[50vh] 等。
    */
    { pattern: /^w-.*$/ },
    { pattern: /^h-.*$/ },

    /*
      5) 文本尺寸/排版
         常见如 text-sm, text-lg, text-xl, text-2xl... 也可加正则覆盖 text-[数字]xl
    */
    { pattern: /^text-(xs|sm|base|lg|xl|2xl|3xl|4xl|5xl|6xl)$/ },
    { pattern: /^text-\d+xl$/ },

    /*
      6) Flex 相关
    */
    "flex",
    { pattern: /^flex-(row|col|wrap|nowrap|row-reverse|col-reverse)$/ },
    { pattern: /^items-(start|end|center|baseline|stretch)$/ },
    { pattern: /^justify-(start|end|center|between|around|evenly)$/ },
    { pattern: /^content-(start|end|center|between|around|evenly)$/ },

    /*
      7) 边框 & 圆角 & 阴影
    */
    { pattern: /^rounded(-(none|sm|md|lg|xl|2xl|3xl|full))?$/ },
    { pattern: /^rounded-[trbl]{1,2}(-(none|sm|md|lg|xl|2xl|3xl|full))?$/ }, // 形如 rounded-t-lg、rounded-bl-md
    { pattern: /^shadow(-(sm|md|lg|xl|2xl|inner|none))?$/ },
    { pattern: /^border(-(0|2|4|8))?$/ },
    { pattern: /^border-[trbl]{1,2}(-(0|2|4|8))?$/ },

    /*
      8) 文本对齐 & 显示模式
    */
    { pattern: /^text-(left|center|right|justify)$/ },
    { pattern: /^(block|inline|inline-block|inline-flex|hidden)$/ },

    /*
      9) 状态变体（如 hover:, focus:, active: 等）
         允许 hover:bg-red-500、focus:border-blue-500 等
         这里用 .+ 去捕获前缀后所有东西，不过要小心可能会保留过多无用CSS
    */
    { pattern: /^hover:.+$/ },
    { pattern: /^focus:.+$/ },
    { pattern: /^active:.+$/ },

    /*
      10) 你自己项目中最常出现的其他 patterns (可自行添加)：
          - z-[数字]
          - absolute / relative / fixed / sticky
          - top-[数字] / left-[数字]
          - grid / gap-[数字]
          - ...
    */
    // { pattern: /^z-\d+$/ },
    // "absolute", "relative", "fixed", "sticky",
    // { pattern: /^top-\[.*\]$/ },
    // { pattern: /^left-\[.*\]$/ },
    // "grid", { pattern: /^gap-\d+$/ },
  ],
  darkMode: 'class',
  plugins: []
}

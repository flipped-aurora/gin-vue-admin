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
  darkMode: 'class',
  plugins: []
}

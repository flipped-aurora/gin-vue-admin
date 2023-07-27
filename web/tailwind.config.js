/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      backgroundColor: {
        'base': '#f0f2f5'
      },
      width: {
        'aside': '220px',
        'aside-title': '60px',
      },
      margin: {
        'aside': '220px',
      },

      minHeight: {
        'base': 'calc(100vh - 160px)'
      }
    },
  },
  plugins: [],
  corePlugins: {
    preflight: false
  }
}


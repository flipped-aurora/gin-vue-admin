const {
  uniPostcssPlugin
} = require('@dcloudio/uni-cli-shared')
module.exports = {
  plugins: [
    uniPostcssPlugin(),
    require('autoprefixer')()
  ]
}

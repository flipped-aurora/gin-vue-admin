var pkg = require('./package.json');

export default {
  entry: 'src/index.js',
  targets: [
    {
      dest: pkg['main'],
      format: 'umd',
      moduleName: 'QrCode',
      sourceMap: true
    }
  ]
};

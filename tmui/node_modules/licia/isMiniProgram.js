var isFn = require('./isFn');

exports = typeof wx !== 'undefined' && isFn(wx.openLocation);

module.exports = exports;

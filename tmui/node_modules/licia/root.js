var isBrowser = require('./isBrowser');
exports = isBrowser ? window : global;

module.exports = exports;

var toStr = require('./toStr');
exports = function(str) {
    return toStr(str).toLocaleUpperCase();
};

module.exports = exports;

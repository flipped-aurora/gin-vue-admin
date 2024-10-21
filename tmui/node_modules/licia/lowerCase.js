var toStr = require('./toStr');
exports = function(str) {
    return toStr(str).toLocaleLowerCase();
};

module.exports = exports;

var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object WeakMap]';
};

module.exports = exports;

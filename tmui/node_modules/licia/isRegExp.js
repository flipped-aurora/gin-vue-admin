var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object RegExp]';
};

module.exports = exports;

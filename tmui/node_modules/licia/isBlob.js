var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object Blob]';
};

module.exports = exports;

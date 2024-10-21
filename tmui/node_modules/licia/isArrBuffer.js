var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object ArrayBuffer]';
};

module.exports = exports;

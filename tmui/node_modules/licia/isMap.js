var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object Map]';
};

module.exports = exports;

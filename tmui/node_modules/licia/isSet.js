var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object Set]';
};

module.exports = exports;

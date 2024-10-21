var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object Date]';
};

module.exports = exports;

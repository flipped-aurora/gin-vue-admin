var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object Number]';
};

module.exports = exports;

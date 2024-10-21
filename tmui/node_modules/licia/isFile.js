var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object File]';
};

module.exports = exports;

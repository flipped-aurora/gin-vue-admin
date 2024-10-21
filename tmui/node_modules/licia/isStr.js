var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object String]';
};

module.exports = exports;

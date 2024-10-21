var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object Arguments]';
};

module.exports = exports;

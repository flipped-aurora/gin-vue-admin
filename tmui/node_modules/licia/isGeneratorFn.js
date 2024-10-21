var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object GeneratorFunction]';
};

module.exports = exports;

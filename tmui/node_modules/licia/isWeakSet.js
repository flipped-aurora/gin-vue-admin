var objToStr = require('./objToStr');
exports = function(val) {
    return objToStr(val) === '[object WeakSet]';
};

module.exports = exports;

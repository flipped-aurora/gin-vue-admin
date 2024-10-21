var objToStr = require('./objToStr');
exports = function(val) {
    switch (objToStr(val)) {
        case '[object Error]':
        case '[object DOMException]':
            return true;
        default:
            return val instanceof Error;
    }
};

module.exports = exports;

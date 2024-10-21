var objToStr = require('./objToStr');
if (Array.isArray && !false) {
    exports = Array.isArray;
} else {
    exports = function(val) {
        return objToStr(val) === '[object Array]';
    };
}

module.exports = exports;

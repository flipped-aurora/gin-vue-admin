var isDate = require('./isDate');
var isStr = require('./isStr');
exports = function(val) {
    if (!val) return new Date();
    if (isDate(val)) return val;
    if (isStr(val)) {
        var match = val.match(regDate);
        if (match) return new Date(match[1], match[2] - 1, match[3]);
    }
    return new Date(val);
};
var regDate = /^(\d{4})-?(\d{2})-?(\d{1,2})$/;

module.exports = exports;

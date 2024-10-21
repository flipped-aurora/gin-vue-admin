var isNum = require('./isNum');
exports = function(val) {
    return isNum(val) && val !== +val;
};

module.exports = exports;

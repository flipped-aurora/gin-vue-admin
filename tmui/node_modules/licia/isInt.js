var isNum = require('./isNum');
exports = function(val) {
    return isNum(val) && val % 1 === 0;
};

module.exports = exports;

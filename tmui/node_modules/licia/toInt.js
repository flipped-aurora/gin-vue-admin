var toNum = require('./toNum');
exports = function(val) {
    if (!val) return val === 0 ? val : 0;
    val = toNum(val);
    return val - (val % 1);
};

module.exports = exports;

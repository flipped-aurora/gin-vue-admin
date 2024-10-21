var isUndef = require('./isUndef');
exports = function(n, lower, upper) {
    if (isUndef(upper)) {
        upper = lower;
        lower = undefined;
    }
    if (!isUndef(lower) && n < lower) return lower;
    if (n > upper) return upper;
    return n;
};

module.exports = exports;

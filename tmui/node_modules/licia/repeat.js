exports = function(str, n) {
    var ret = '';
    if (n < 1) return '';
    while (n > 0) {
        if (n & 1) ret += str;
        n >>= 1;
        str += str;
    }
    return ret;
};

module.exports = exports;

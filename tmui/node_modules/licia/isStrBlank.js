exports = function(str) {
    for (var i = 0, len = str.length; i < len; i++) {
        var c = str[i];
        if (
            c !== ' ' &&
            c !== '\n' &&
            c !== '\r' &&
            c !== '\t' &&
            c !== '\f' &&
            c !== '\v'
        ) {
            return false;
        }
    }
    return true;
};

module.exports = exports;

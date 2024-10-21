var vector = [];
var bChars = [];

exports = function(a, b) {
    if (a === b) return 0;

    if (a.length > b.length) {
        var tmp = a;
        a = b;
        b = tmp;
    }
    var aLen = a.length;
    var bLen = b.length;
    if (!aLen) return bLen;
    if (!bLen) return aLen;

    while (aLen > 0 && a.charCodeAt(aLen - 1) === b.charCodeAt(bLen - 1)) {
        aLen--;
        bLen--;
    }
    if (!aLen) return bLen;

    var start = 0;
    while (start < aLen && a.charCodeAt(start) === b.charCodeAt(start)) {
        start++;
    }
    aLen -= start;
    bLen -= start;
    if (!aLen) return bLen;
    var current = 0;
    var left;
    var above;
    var charA;
    var i = 0;
    while (i < bLen) {
        bChars[i] = b.charCodeAt(start + i);
        vector[i] = ++i;
    }

    for (var _i = 0; _i < aLen; _i++) {
        left = _i;
        current = _i + 1;
        charA = a.charCodeAt(start + _i);
        for (var j = 0; j < bLen; j++) {
            above = current;
            current = left;
            left = vector[j];
            if (charA !== bChars[j]) {
                if (left < current) current = left;
                if (above < current) current = above;
                current++;
            }
            vector[j] = current;
        }
    }
    return current;
};

module.exports = exports;

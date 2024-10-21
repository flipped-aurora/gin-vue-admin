var chunk = require('./chunk');
var map = require('./map');

exports = {
    encode: function(arr) {
        if (arr.length < 32768) {
            return String.fromCodePoint.apply(String, arr);
        }
        return map(chunk(arr, 32767), function(nums) {
            return String.fromCodePoint.apply(String, nums);
        }).join('');
    },
    decode: function(str) {
        var ret = [];
        var i = 0;
        var len = str.length;
        while (i < len) {
            var c = str.charCodeAt(i++);

            if (c >= 0xd800 && c <= 0xdbff && i < len) {
                var tail = str.charCodeAt(i++);

                if ((tail & 0xfc00) === 0xdc00) {
                    ret.push(((c & 0x3ff) << 10) + (tail & 0x3ff) + 0x10000);
                } else {
                    ret.push(c);
                    i--;
                }
            } else {
                ret.push(c);
            }
        }
        return ret;
    }
};

module.exports = exports;

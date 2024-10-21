var toArr = require('./toArr');

exports = {
    encode: function(arr) {
        arr = toArr(arr);
        var ret = '';
        for (var i = 0, len = arr.length; i < len; i++) {
            ret += encode(arr[i]);
        }
        return ret;
    },
    decode: function(str) {
        var ret = [];
        var i = 0;
        var len = str.length;
        while (i < len) {
            var value = 0;
            var continuation = false;
            var shift = 0;
            do {
                var digit = charToInt[str[i++]];
                continuation = (digit & VLQ_CONTINUATION_BIT) !== 0;
                digit &= VLQ_BASE_MASK;
                value = value + (digit << shift);
                shift = shift + VLQ_BASE_SHIFT;
            } while (continuation);
            ret.push(fromVLQSigned(value));
        }
        return ret;
    }
};
var charToInt = {};
var intToChar = {};
var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
for (var i = 0, len = chars.length; i < len; i++) {
    charToInt[chars[i]] = i;
    intToChar[i] = chars[i];
}
var VLQ_BASE_SHIFT = 5;
var VLQ_BASE = 1 << VLQ_BASE_SHIFT;
var VLQ_BASE_MASK = VLQ_BASE - 1;
var VLQ_CONTINUATION_BIT = VLQ_BASE;
function encode(value) {
    var ret = '';
    value = toVLQSigned(value);
    do {
        var digit = value & VLQ_BASE_MASK;
        value >>>= VLQ_BASE_SHIFT;
        if (value > 0) {
            digit |= VLQ_CONTINUATION_BIT;
        }
        ret += intToChar[digit];
    } while (value > 0);
    return ret;
}
function toVLQSigned(value) {
    if (value < 0) {
        return (-value << 1) + 1;
    } else {
        return (value << 1) + 0;
    }
}
function fromVLQSigned(value) {
    var negate = (value & 1) === 1;
    value = value >> 1;
    return negate ? -value : value;
}

module.exports = exports;

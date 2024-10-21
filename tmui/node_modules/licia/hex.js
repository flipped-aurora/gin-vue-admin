var isOdd = require('./isOdd');
exports = {
    encode: function(bytes) {
        var hex = [];
        for (var i = 0, len = bytes.length; i < len; i++) {
            var byte = bytes[i];
            hex.push((byte >>> 4).toString(16));
            hex.push((byte & 0xf).toString(16));
        }
        return hex.join('');
    },
    decode: function(str) {
        var bytes = [];
        var len = str.length;
        if (isOdd(len)) len--;
        for (var i = 0; i < len; i += 2) {
            bytes.push(parseInt(str.substr(i, 2), 16));
        }
        return bytes;
    }
};

module.exports = exports;

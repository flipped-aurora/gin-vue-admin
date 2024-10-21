var utf8 = require('./utf8');
var hex = require('./hex');
var base64 = require('./base64');
exports = function(bytes) {
    var encoding =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : 'utf8';
    if (encoding === 'hex') return hex.encode(bytes);
    if (encoding === 'base64') return base64.encode(bytes);
    var str = [];
    for (var i = 0, len = bytes.length; i < len; i++) {
        str.push(String.fromCharCode(bytes[i]));
    }
    str = str.join('');
    if (encoding === 'utf8') {
        str = utf8.decode(str);
    }
    return str;
};

module.exports = exports;

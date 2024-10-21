var utf8 = require('./utf8');
var hex = require('./hex');
var base64 = require('./base64');
exports = function(str) {
    var encoding =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : 'utf8';
    if (encoding === 'hex') return hex.decode(str);
    if (encoding === 'base64') return base64.decode(str);
    var bytes = [];
    if (encoding === 'utf8') {
        str = utf8.encode(str);
    }
    for (var i = 0, len = str.length; i < len; i++) {
        bytes.push(str.charCodeAt(i) & 0xff);
    }
    return bytes;
};

module.exports = exports;

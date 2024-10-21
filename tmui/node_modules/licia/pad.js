var repeat = require('./repeat');
var toStr = require('./toStr');
exports = function(str, len, chars) {
    str = toStr(str);
    var strLen = str.length;
    chars = chars || ' ';
    if (strLen < len) {
        var padStr = repeat(chars, Math.ceil((len - strLen) / 2));
        str = padStr + str + padStr;
        str = str.substr(Math.ceil((str.length - len) / 2), len);
    }
    return str;
};

module.exports = exports;

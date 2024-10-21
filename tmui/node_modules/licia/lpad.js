var repeat = require('./repeat');
var toStr = require('./toStr');
exports = function(str, len, chars) {
    str = toStr(str);
    var strLen = str.length;
    chars = chars || ' ';
    if (strLen < len) str = (repeat(chars, len - strLen) + str).slice(-len);
    return str;
};

module.exports = exports;

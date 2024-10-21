var each = require('./each');
var ucs2 = require('./ucs2');
var map = require('./map');
var utf8 = require('./utf8');
exports = function(str) {
    try {
        return decodeURIComponent(str);
    } catch (e) {
        var matches = str.match(regMatcher);
        if (!matches) {
            return str;
        }
        each(matches, function(match) {
            str = str.replace(match, decode(match));
        });
        return str;
    }
};
function decode(str) {
    str = str.split('%').slice(1);
    var bytes = map(str, hexToInt);
    str = ucs2.encode(bytes);
    str = utf8.decode(str, true);
    return str;
}
function hexToInt(numStr) {
    return +('0x' + numStr);
}
var regMatcher = /(%[a-f0-9]{2})+/gi;

module.exports = exports;

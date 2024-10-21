var stripAnsi = require('./stripAnsi');
var isFullWidth = require('./isFullWidth');
exports = function(str) {
    str = stripAnsi(str);
    var width = 0;
    for (var i = 0, len = str.length; i < len; i++) {
        var c = str.codePointAt(i);

        if (c <= 31 || c === 127) {
            continue;
        }
        width += isFullWidth(c) ? 2 : 1;
    }
    return width;
};

module.exports = exports;

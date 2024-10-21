var defaults = require('./defaults');
var isUndef = require('./isUndef');
exports = function(txt, width) {
    var options =
        arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : {};
    defaults(options, defOptions);
    var ellipsis = options.ellipsis,
        separator = options.separator;
    var len = txt.length;
    if (width > len) return txt;
    var end = width - ellipsis.length;
    if (end < 1) return ellipsis;
    var ret = txt.slice(0, end);
    if (isUndef(separator)) return ret + ellipsis;
    if (txt.indexOf(separator, end) !== end) {
        var idx = ret.lastIndexOf(separator);
        if (idx > -1) {
            ret = ret.slice(0, idx);
        }
    }
    return ret + ellipsis;
};
var defOptions = {
    ellipsis: '...'
};

module.exports = exports;

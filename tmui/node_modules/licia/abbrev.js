var restArgs = require('./restArgs');
var isSorted = require('./isSorted');
exports = restArgs(function(names) {
    names = names.sort(isSorted.defComparator);
    var ret = {};
    var idleMap = {};
    for (var i = 0, len = names.length; i < len; i++) {
        var str = names[i];
        var nextStr = names[i + 1] || '';
        if (str === nextStr) continue;
        var start = false;
        var abbrev = '';
        for (var j = 0, strLen = str.length; j < strLen; j++) {
            abbrev += str[j];
            if (!start && (str[j] !== nextStr[j] || j === strLen - 1)) {
                start = true;
            }
            if (!start) {
                idleMap[abbrev] = str;
            } else if (!ret[abbrev] && !idleMap[abbrev]) {
                ret[abbrev] = str;
            }
        }
    }
    return ret;
});

module.exports = exports;

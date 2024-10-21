var restArgs = require('./restArgs');
exports = restArgs(function(first, arrays) {
    var end = first.length;
    for (var i = 0, len = arrays.length; i < len; i++) {
        var arr = arrays[i];
        for (var j = 0, _len = arr.length; j < _len; j++) {
            first[end++] = arr[j];
        }
    }
    first.length = end;
    return first;
});

module.exports = exports;

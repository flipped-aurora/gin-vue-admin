var map = require('./map');
var pluck = require('./pluck');
var max = require('./max');
exports = function(arr) {
    var len = max.apply(
        null,
        map(arr, function(arr) {
            return arr.length;
        })
    );
    var ret = Array(len);
    for (var i = 0; i < len; i++) {
        ret[i] = pluck(arr, i);
    }
    return ret;
};

module.exports = exports;

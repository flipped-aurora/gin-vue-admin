var isArrLike = require('./isArrLike');
var clone = require('./clone');
var values = require('./values');
var random = require('./random');
var swap = require('./swap');
exports = function(obj, n) {
    var sample = isArrLike(obj) ? clone(obj) : values(obj);
    var len = sample.length;
    n = Math.max(Math.min(n, len), 0);
    var last = len - 1;
    for (var i = 0; i < n; i++) {
        var rand = random(i, last);
        swap(sample, i, rand);
    }
    return sample.slice(0, n);
};

module.exports = exports;

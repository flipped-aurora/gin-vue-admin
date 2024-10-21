var safeCb = require('./safeCb');
var pluck = require('./pluck');
var map = require('./map');
var isUndef = require('./isUndef');
exports = function(obj, iteratee, ctx) {
    iteratee = safeCb(iteratee, ctx);
    var idx = 0;
    return pluck(
        map(obj, function(val, key) {
            return {
                val: val,
                idx: idx++,
                criteria: iteratee(val, key, obj)
            };
        }).sort(function(left, right) {
            var a = left.criteria;
            var b = right.criteria;
            if (a !== b) {
                if (a > b || isUndef(a)) return 1;
                if (a < b || isUndef(b)) return -1;
            }
            return left.idx - right.idx;
        }),
        'val'
    );
};

module.exports = exports;

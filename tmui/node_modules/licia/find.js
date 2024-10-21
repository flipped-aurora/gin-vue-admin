var findKey = require('./findKey');
var findIdx = require('./findIdx');
var isArrLike = require('./isArrLike');
var isUndef = require('./isUndef');
exports = function(obj, predicate, ctx) {
    var keyFinder = isArrLike(obj) ? findIdx : findKey;
    var key = keyFinder(obj, predicate, ctx);
    if (!isUndef(key) && key !== -1) return obj[key];
};

module.exports = exports;

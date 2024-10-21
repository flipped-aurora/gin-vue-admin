var findIdx = require('./findIdx');
exports = function(arr, predicate, ctx) {
    return findIdx(arr, predicate, ctx, -1);
};

module.exports = exports;

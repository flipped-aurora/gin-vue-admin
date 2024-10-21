var isArrLike = require('./isArrLike');
var keys = require('./keys');
exports = function(obj) {
    return isArrLike(obj) ? obj.length : keys(obj).length;
};

module.exports = exports;

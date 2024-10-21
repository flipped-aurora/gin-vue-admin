var extendOwn = require('./extendOwn');
var isMatch = require('./isMatch');
exports = function(attrs) {
    attrs = extendOwn({}, attrs);
    return function(obj) {
        return isMatch(obj, attrs);
    };
};

module.exports = exports;

var isFn = require('./isFn');
exports = function(val) {
    if (val == null) return false;
    if (val._isBuffer) return true;
    return (
        val.constructor &&
        isFn(val.constructor.isBuffer) &&
        val.constructor.isBuffer(val)
    );
};

module.exports = exports;

var root = require('./root');
var nativeIsFinite = root.isFinite;
var nativeIsNaN = root.isNaN;
exports = function(val) {
    return nativeIsFinite(val) && !nativeIsNaN(parseFloat(val));
};

module.exports = exports;

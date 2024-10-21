var isUndef = require('./isUndef');
exports = function() {
    for (var i = 0, len = arguments.length; i < len; i++) {
        if (!isUndef(arguments[i])) return arguments[i];
    }
};

module.exports = exports;

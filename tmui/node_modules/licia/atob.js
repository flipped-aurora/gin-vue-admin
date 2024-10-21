var root = require('./root');
var isNode = require('./isNode');
var base64 = require('./base64');
var map = require('./map');
if (isNode) {
    exports = function(str) {
        return new Buffer(str, 'base64').toString('binary');
    };
} else {
    if (root.atob && !false) {
        exports = root.atob;
    } else {
        exports = function(str) {
            return map(base64.decode(str), function(c) {
                return String.fromCharCode(c);
            }).join('');
        };
    }
}

module.exports = exports;

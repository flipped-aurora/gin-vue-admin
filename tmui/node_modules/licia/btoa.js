var root = require('./root');
var isNode = require('./isNode');
var base64 = require('./base64');
var map = require('./map');
if (isNode) {
    exports = function(str) {
        return new Buffer(str, 'binary').toString('base64');
    };
} else {
    if (root.btoa && !false) {
        exports = root.btoa;
    } else {
        exports = function(str) {
            return base64.encode(
                map(str, function(c) {
                    return c.charCodeAt(0);
                })
            );
        };
    }
}

module.exports = exports;

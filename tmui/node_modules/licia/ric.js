var root = require('./root');
var now = require('./now');
exports =
    root.requestIdleCallback ||
    function(cb) {
        var start = now();
        return setTimeout(function() {
            cb({
                didTimeout: false,
                timeRemaining: function() {
                    return Math.max(0, 50 - (now() - start));
                }
            });
        }, 1);
    };
exports.cancel =
    root.cancelIdleCallback ||
    function(id) {
        clearTimeout(id);
    };

module.exports = exports;

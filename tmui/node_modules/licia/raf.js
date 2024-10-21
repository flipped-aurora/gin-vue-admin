var now = require('./now');
var isBrowser = require('./isBrowser');
var raf, cancel;
var lastTime = 0;
if (isBrowser) {
    raf = window.requestAnimationFrame;
    cancel = window.cancelAnimationFrame;
    var vendors = ['ms', 'moz', 'webkit', 'o'];
    for (var i = 0, len = vendors.length; i < len && !raf; i++) {
        raf = window[vendors[i] + 'RequestAnimationFrame'];
        cancel =
            window[vendors[i] + 'CancelAnimationFrame'] ||
            window[vendors[i] + 'CancelRequestAnimationFrame'];
    }
    if (raf) {
        raf = raf.bind(window);
        cancel = cancel.bind(window);
    }
}
raf =
    raf ||
    function(cb) {
        var curTime = now();
        var timeToCall = Math.max(0, 16 - (curTime - lastTime));
        var id = setTimeout(function() {
            cb(curTime + timeToCall);
        }, timeToCall);
        lastTime = curTime + timeToCall;
        return id;
    };
cancel =
    cancel ||
    function(id) {
        clearTimeout(id);
    };
raf.cancel = cancel;
exports = raf;

module.exports = exports;

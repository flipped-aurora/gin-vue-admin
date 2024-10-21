var Emitter = require('./Emitter');
var safeGet = require('./safeGet');
var screen = window.screen;
exports = {
    get: function() {
        if (screen) {
            var orientation = safeGet(screen, 'orientation.type');
            if (orientation) return orientation.split('-').shift();
        }
        return window.innerWidth > window.innerHeight
            ? 'landscape'
            : 'portrait';
    }
};
Emitter.mixin(exports);
window.addEventListener(
    'orientationchange',
    function() {
        setTimeout(function() {
            exports.emit('change', exports.get());
        }, 200);
    },
    false
);

module.exports = exports;

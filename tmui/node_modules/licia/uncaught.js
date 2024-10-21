var isBrowser = require('./isBrowser');
var SingleEmitter = require('./SingleEmitter');
var isOn = false;
exports = {
    start: function() {
        isOn = true;
    },
    stop: function() {
        isOn = false;
    }
};
SingleEmitter.mixin(exports);
if (isBrowser) {
    window.addEventListener('error', function(event) {
        if (event.error) {
            callListeners(event.error);
        } else if (event.message) {
            var e = new Error(event.message);
            e.stack = 'Error: '
                .concat(event.message, ' \n at ')
                .concat(event.filename, ':')
                .concat(event.lineno, ':')
                .concat(event.colno);
            callListeners(e);
        }
    });
    window.addEventListener('unhandledrejection', function(e) {
        callListeners(e.reason);
    });
} else {
    process.on('uncaughtException', callListeners);
    process.on('unhandledRejection', callListeners);
}
function callListeners(err) {
    if (!isOn) return;
    exports.emit(err);
}

module.exports = exports;

var Emitter = require('./Emitter');
var keyCode = require('./keyCode');
var each = require('./each');
var unique = require('./unique');
var trim = require('./trim');
var map = require('./map');
var isFn = require('./isFn');
exports = {
    on: function(keys, options, listener) {
        if (isFn(options)) {
            listener = options;
            options = {};
        }
        keys = keys.split(regComma);
        each(keys, function(key) {
            key = normalizeKey(key);
            if (options.element) {
                var _options = options,
                    element = _options.element;
                var hotkeyListeners = element._hotkeyListeners || {};
                element._hotkeyListeners = hotkeyListeners;
                hotkeyListeners[key] = hotkeyListeners[key] || [];
                var hotkeyListener = function(e) {
                    if (key === getKeysFromEvent(e)) {
                        listener(e);
                    }
                };
                hotkeyListeners[key].push({
                    listener: hotkeyListener,
                    origin: listener
                });
                element.addEventListener('keydown', hotkeyListener);
            } else {
                emitter.on(key, listener);
            }
        });
    },
    off: function(keys, options, listener) {
        if (isFn(options)) {
            listener = options;
            options = {};
        }
        keys = keys.split(regComma);
        each(keys, function(key) {
            key = normalizeKey(key);
            if (options.element) {
                var _options2 = options,
                    element = _options2.element;
                var hotkeyListeners = element._hotkeyListeners;
                if (hotkeyListeners && hotkeyListeners[key]) {
                    var listeners = hotkeyListeners[key];
                    var hotkeyListener;
                    for (var i = 0, len = listeners.length; i < len; i++) {
                        if (listeners[i].origin === listener) {
                            hotkeyListener = listeners[i].listener;
                            listeners.splice(i, 1);
                        }
                    }
                    if (hotkeyListener) {
                        element.removeEventListener('keydown', hotkeyListener);
                    }
                }
            } else {
                emitter.off(key, listener);
            }
        });
    }
};
var emitter = new Emitter();
document.addEventListener('keydown', function(e) {
    emitter.emit(getKeysFromEvent(e), e);
});
function getKeysFromEvent(e) {
    var keys = [];
    if (e.ctrlKey) keys.push('ctrl');
    if (e.shiftKey) keys.push('shift');
    keys.push(keyCode(e.keyCode));
    return normalizeKey(keys.join('+'));
}
function normalizeKey(keyStr) {
    var keys = keyStr.split(regPlus);
    keys = map(keys, function(key) {
        return trim(key);
    });
    keys = unique(keys);
    keys.sort();
    return keys.join('+');
}
var regComma = /,/g;
var regPlus = /\+/g;

module.exports = exports;

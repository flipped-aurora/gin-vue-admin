var isNode = require('./isNode');
var root = require('./root');
var memStorage = require('./memStorage');
if (isNode) {
    exports = eval('require')('util').deprecate;
} else {
    var localStorage = root.localStorage || memStorage;
    exports = function(fn, msg) {
        if (localStorage.getItem('noDeprecation')) {
            return fn;
        }
        var warned = false;
        function deprecated() {
            if (!warned) {
                warned = true;

                console.warn(msg);
            }
            for (
                var _len = arguments.length, args = new Array(_len), _key = 0;
                _key < _len;
                _key++
            ) {
                args[_key] = arguments[_key];
            }
            return fn.apply(this, args);
        }
        Object.setPrototypeOf(deprecated, fn);
        if (fn.prototype) {
            deprecated.prototype = fn.prototype;
        }
        return deprecated;
    };
}

module.exports = exports;

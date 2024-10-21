var promisify = require('./promisify');
var callbackify = require('./callbackify');
var last = require('./last');
var isFn = require('./isFn');
exports = function(fn, type) {
    var callbackFn;
    var promiseFn;
    if (type === 'callback') {
        callbackFn = fn;
        promiseFn = promisify(fn);
    } else {
        promiseFn = fn;
        callbackFn = callbackify(fn);
    }
    return function() {
        for (
            var _len = arguments.length, args = new Array(_len), _key = 0;
            _key < _len;
            _key++
        ) {
            args[_key] = arguments[_key];
        }
        if (isFn(last(args))) {
            callbackFn.apply(this, args);
        } else {
            return promiseFn.apply(this, args);
        }
    };
};

module.exports = exports;

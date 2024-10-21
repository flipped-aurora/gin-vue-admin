var toArr = require('./toArr');
var isGeneratorFn = require('./isGeneratorFn');
var isPromise = require('./isPromise');
var toStr = require('./toStr');

exports = function(fn) {
    if (!isGeneratorFn(fn)) {
        throw new TypeError('Expected a generator function');
    }
    return function() {
        var _this = this;
        var args = toArr(arguments);
        return new Promise(function(resolve, reject) {
            var generator = fn.apply(_this, args);
            function onFulfilled(res) {
                var ret;
                try {
                    ret = generator.next(res);
                } catch (e) {
                    return reject(e);
                }
                next(ret);
            }
            function onRejected(err) {
                var ret;
                try {
                    ret = generator.throw(err);
                } catch (e) {
                    return reject(e);
                }
                next(ret);
            }
            function next(ret) {
                if (ret.done) return resolve(ret.value);
                if (isPromise(ret.value)) {
                    return ret.value.then(onFulfilled, onRejected);
                }
                return onRejected(
                    new TypeError(
                        'You may only yield a promise, '.concat(
                            toStr(ret.value),
                            ' is passed'
                        )
                    )
                );
            }
            onFulfilled();
        });
    };
};

module.exports = exports;

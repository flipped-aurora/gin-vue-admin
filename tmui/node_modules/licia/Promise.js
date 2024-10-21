var Class = require('./Class');
var isObj = require('./isObj');
var isFn = require('./isFn');
var State = require('./State');
var bind = require('./bind');
var nextTick = require('./nextTick');
var noop = require('./noop');
var toArr = require('./toArr');
var Promise = (exports = Class(
    {
        initialize: function Promise(fn) {
            if (!isObj(this))
                throw new TypeError('Promises must be constructed via new');
            if (!isFn(fn)) throw new TypeError(fn + ' is not a function');
            var self = this;
            this._state = new State('pending', {
                fulfill: {
                    from: 'pending',
                    to: 'fulfilled'
                },
                reject: {
                    from: 'pending',
                    to: 'rejected'
                },
                adopt: {
                    from: 'pending',
                    to: 'adopted'
                }
            })
                .on('fulfill', assignVal)
                .on('reject', assignVal)
                .on('adopt', assignVal);
            function assignVal(val) {
                self._value = val;
            }
            this._handled = false;
            this._value = undefined;
            this._deferreds = [];
            doResolve(fn, this);
        },
        catch: function(onRejected) {
            return this.then(null, onRejected);
        },
        then: function(onFulfilled, onRejected) {
            var promise = new Promise(noop);
            handle(this, new Handler(onFulfilled, onRejected, promise));
            return promise;
        }
    },
    {
        all: function(arr) {
            var args = toArr(arr);
            return new Promise(function(resolve, reject) {
                if (args.length === 0) return resolve([]);
                var remaining = args.length;
                function res(i, val) {
                    try {
                        if (val && (isObj(val) || isFn(val))) {
                            var then = val.then;
                            if (isFn(then)) {
                                then.call(
                                    val,
                                    function(val) {
                                        res(i, val);
                                    },
                                    reject
                                );
                                return;
                            }
                        }
                        args[i] = val;
                        if (--remaining === 0) resolve(args);
                    } catch (e) {
                        reject(e);
                    }
                }
                for (var i = 0; i < args.length; i++) res(i, args[i]);
            });
        },
        resolve: function(val) {
            if (val && isObj(val) && val.constructor === Promise) return val;
            return new Promise(function(resolve) {
                resolve(val);
            });
        },
        reject: function(val) {
            return new Promise(function(resolve, reject) {
                reject(val);
            });
        },
        race: function(values) {
            return new Promise(function(resolve, reject) {
                for (var i = 0, len = values.length; i < len; i++) {
                    values[i].then(resolve, reject);
                }
            });
        }
    }
));
var Handler = Class({
    initialize: function Handler(onFulfilled, onRejected, promise) {
        this.onFulfilled = isFn(onFulfilled) ? onFulfilled : null;
        this.onRejected = isFn(onRejected) ? onRejected : null;
        this.promise = promise;
    }
});
function reject(self, err) {
    self._state.reject(err);
    finale(self);
}
function resolve(self, val) {
    try {
        if (val === self)
            throw new TypeError('A promise cannot be resolved with itself');
        if (val && (isObj(val) || isFn(val))) {
            var then = val.then;
            if (val instanceof Promise) {
                self._state.adopt(val);
                return finale(self);
            }
            if (isFn(then)) return doResolve(bind(then, val), self);
        }
        self._state.fulfill(val);
        finale(self);
    } catch (e) {
        reject(self, e);
    }
}
function finale(self) {
    for (var i = 0, len = self._deferreds.length; i < len; i++) {
        handle(self, self._deferreds[i]);
    }
    self._deferreds = null;
}
function handle(self, deferred) {
    while (self._state.is('adopted')) self = self._value;
    if (self._state.is('pending')) return self._deferreds.push(deferred);
    self._handled = true;
    nextTick(function() {
        var isFulfilled = self._state.is('fulfilled');
        var cb = isFulfilled ? deferred.onFulfilled : deferred.onRejected;
        if (cb === null)
            return (isFulfilled ? resolve : reject)(
                deferred.promise,
                self._value
            );
        var ret;
        try {
            ret = cb(self._value);
        } catch (e) {
            return reject(deferred.promise, e);
        }
        resolve(deferred.promise, ret);
    });
}
function doResolve(fn, self) {
    var done = false;
    try {
        fn(
            function(val) {
                if (done) return;
                done = true;
                resolve(self, val);
            },
            function(reason) {
                if (done) return;
                done = true;
                reject(self, reason);
            }
        );
    } catch (e) {
        if (done) return;
        done = true;
        reject(self, e);
    }
}

module.exports = exports;

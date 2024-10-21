var restArgs = require('./restArgs');
var uniqId = require('./uniqId');
var toSrc = require('./toSrc');
var createUrl = require('./createUrl');
var isStr = require('./isStr');
var root = require('./root');
exports = function(fn) {
    var promises = {};
    var src = [
        toSrc(isPromise),
        'onmessage=(',
        toSrc(function(fn) {
            return function(e) {
                var data = e.data;
                var id = data[0];
                var args = data[1];
                var value;
                try {
                    value = fn.apply(fn, args);
                    if (isPromise(value)) {
                        value.then(
                            function(value) {
                                postMessage([id, null, value]);
                            },
                            function(err) {
                                postMessage([id, err.message]);
                            }
                        );
                    } else {
                        postMessage([id, null, value]);
                    }
                } catch (e) {
                    postMessage([id, e.message]);
                }
            };
        }),
        ')(' + toSrc(fn) + ')'
    ].join('\n');
    var worker = new Worker(createUrl(src));
    worker.onmessage = function(e) {
        var data = e.data;
        var id = data[0];
        var err = data[1];
        var value = data[2];
        if (isStr(err)) err = new Error(err);
        promises[id](err, value);
        delete promises[id];
    };
    return restArgs(function(args) {
        var id = uniqId('workerize');
        return new root.Promise(function(resolve, reject) {
            promises[id] = function(err, value) {
                if (err) return reject(err);
                resolve(value);
            };
            worker.postMessage([id, args]);
        });
    });
};

function isPromise(val) {
    return (
        !!val &&
        (typeof val === 'object' || typeof val === 'function') &&
        typeof val.then === 'function' &&
        typeof val.catch === 'function'
    );
}

module.exports = exports;

var now = require('./now');
exports = function(condition) {
    var timeout =
        arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 0;
    var interval =
        arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : 250;
    function evalCondition() {
        return new Promise(function(resolve, reject) {
            try {
                resolve(condition());
            } catch (e) {
                reject(e);
            }
        });
    }
    return new Promise(function(resolve, reject) {
        var startTime = now();
        var pollCondition = function() {
            evalCondition().then(function(val) {
                var elapsed = now() - startTime;
                if (val) {
                    resolve(val);
                } else if (timeout && elapsed >= timeout) {
                    reject(
                        Error('Wait timed out after '.concat(timeout, ' ms'))
                    );
                } else {
                    setTimeout(pollCondition, interval);
                }
            }, reject);
        };
        pollCondition();
    });
};

module.exports = exports;

var isFn = require('./isFn');
var restArgs = require('./restArgs');
exports = function(fn) {
    if (isFn(fn)) {
        return restArgs(function(args) {
            return fn
                .apply(this, args)
                .then(function(v) {
                    return [v, null];
                })
                .catch(function(err) {
                    return [void 0, err];
                });
        });
    } else {
        return fn
            .then(function(v) {
                return [v, null];
            })
            .catch(function(err) {
                return [void 0, err];
            });
    }
};

module.exports = exports;

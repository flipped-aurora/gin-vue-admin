var restArgs = require('./restArgs');
exports = restArgs(function(fn, ctx, args) {
    return restArgs(function(callArgs) {
        return fn.apply(ctx, args.concat(callArgs));
    });
});

module.exports = exports;

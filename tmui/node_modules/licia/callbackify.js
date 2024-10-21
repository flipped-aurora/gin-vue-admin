var restArgs = require('./restArgs');
exports = function(fn) {
    return restArgs(function(args) {
        var cb = args.pop();
        fn.apply(this, args).then(
            function(value) {
                cb(null, value);
            },
            function(err) {
                if (err === null) err = new Error();
                cb(err);
            }
        );
    });
};

module.exports = exports;

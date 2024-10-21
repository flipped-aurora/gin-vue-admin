var restArgs = require('./restArgs');
var root = require('./root');
exports = function(fn, multiArgs) {
    return restArgs(function(args) {
        return new root.Promise(function(resolve, reject) {
            args.push(
                restArgs(function callback(err, values) {
                    if (err) return reject(err);
                    if (!multiArgs) return resolve(values[0]);
                    resolve(values);
                })
            );
            fn.apply(this, args);
        });
    });
};

module.exports = exports;

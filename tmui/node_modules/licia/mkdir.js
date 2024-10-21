const isFn = require('./isFn');
const noop = require('./noop');

const fs = require('fs');
const path = require('path');

const _0777 = parseInt('0777', 8);

exports = function(p, mode, cb) {
    if (isFn(mode)) {
        cb = mode;
        mode = _0777;
    }
    cb = cb || noop;
    p = path.resolve(p);

    fs.mkdir(p, mode, function(err) {
        if (!err) return cb();

        if (err.code === 'ENOENT') {
            exports(path.dirname(p), mode, function(err) {
                if (err) return cb(err);

                exports(p, mode, cb);
            });
        } else {
            fs.stat(p, function(errStat, stat) {
                if (errStat || !stat.isDirectory()) return cb(errStat);

                cb();
            });
        }
    });
};

exports.sync = function(p, mode = _0777) {
    p = path.resolve(p);

    try {
        fs.mkdirSync(p, mode);
    } catch (err) {
        if (err.code === 'ENOENT') {
            exports.sync(path.dirname(p), mode);
            exports.sync(p, mode);
        } else {
            try {
                if (!fs.statSync(p).isDirectory()) {
                    throw err;
                }
            } catch (_) {
                throw err;
            }
        }
    }
};

module.exports = exports;

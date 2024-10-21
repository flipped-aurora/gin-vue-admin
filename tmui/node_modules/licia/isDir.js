const root = require('./root');

const fs = require('fs');

exports = function(path) {
    return new root.Promise(function(resolve, reject) {
        fs.stat(path, function(err, stats) {
            if (err) {
                if (err.code === 'ENOENT') {
                    resolve(false);
                } else {
                    reject(err);
                }
            } else {
                resolve(stats.isDirectory());
            }
        });
    });
};

module.exports = exports;

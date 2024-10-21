const promisify = require('./promisify');
const root = require('./root');
const each = require('./each');
const toArr = require('./toArr');

const fs = require('fs');

each(
    [
        'access',
        'appendFile',
        'chmod',
        'chown',
        'close',
        'fchmod',
        'fchown',
        'fdatasync',
        'fstat',
        'fsync',
        'ftruncate',
        'futimes',
        'link',
        'lstat',
        'mkdir',
        'mkdtemp',
        'open',
        'read',
        'readFile',
        'readdir',
        'readlink',
        'realpath',
        'rename',
        'rmdir',
        'stat',
        'symlink',
        'truncate',
        'unlink',
        'utimes',
        'write',
        'writeFile'
    ],
    function(method) {
        exports[method] = promisify(fs[method]);
    }
);

exports.exists = function() {
    const args = toArr(arguments);

    return new root.Promise(function(resolve) {
        args.push(resolve);
        fs.exists.apply(null, args);
    });
};

module.exports = exports;

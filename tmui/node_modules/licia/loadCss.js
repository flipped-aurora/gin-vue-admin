var noop = require('./noop');
exports = function(src, cb) {
    cb = cb || noop;
    var link = document.createElement('link');
    link.rel = 'stylesheet';
    link.type = 'text/css';
    link.onerror = function() {
        cb(false);
    };
    link.onload = function() {
        cb(true);
    };
    link.href = src;
    document.head.appendChild(link);
};

module.exports = exports;

var noop = require('./noop');
exports = function(src, cb) {
    cb = cb || noop;
    var img = new Image();
    img.onload = function() {
        cb(null, img);
    };
    img.onerror = function(err) {
        cb(err);
    };
    img.src = src;
};

module.exports = exports;

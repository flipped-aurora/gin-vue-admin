var normalizePath = require('./normalizePath');
exports = function(path) {
    path = normalizePath(path);
    if (path[0] !== '/') {
        path = '/'.concat(path);
    }
    return encodeURI('file://'.concat(path)).replace(
        /[?#]/g,
        encodeURIComponent
    );
};

module.exports = exports;

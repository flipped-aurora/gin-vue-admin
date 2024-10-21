exports = function(url) {
    return regAbsolute.test(url);
};
var regAbsolute = /^[a-z][a-z0-9+.-]*:/;

module.exports = exports;

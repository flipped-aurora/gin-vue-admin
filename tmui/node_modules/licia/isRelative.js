exports = function(path) {
    return !regAbsolute.test(path);
};
var regAbsolute = /^([a-z]+:)?[\\/]/i;

module.exports = exports;

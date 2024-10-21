exports = function(path) {
    return path.replace(regSlashes, '/');
};
var regSlashes = /[\\/]+/g;

module.exports = exports;

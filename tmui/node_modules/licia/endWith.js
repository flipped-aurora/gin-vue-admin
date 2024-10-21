exports = function(str, suffix) {
    var idx = str.length - suffix.length;
    return idx >= 0 && str.indexOf(suffix, idx) === idx;
};

module.exports = exports;

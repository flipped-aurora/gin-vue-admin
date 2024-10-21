exports = function(val) {
    var type = typeof val;
    return !!val && (type === 'function' || type === 'object');
};

module.exports = exports;

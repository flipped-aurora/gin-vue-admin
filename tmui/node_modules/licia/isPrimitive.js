exports = function(val) {
    var type = typeof val;
    return val == null || (type !== 'function' && type !== 'object');
};

module.exports = exports;

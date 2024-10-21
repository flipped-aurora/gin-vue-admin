var idCounter = 0;
exports = function(prefix) {
    var id = ++idCounter + '';
    return prefix ? prefix + id : id;
};

module.exports = exports;

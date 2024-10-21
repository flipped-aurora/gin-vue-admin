exports = function(num) {
    var precision =
        arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 12;
    return parseFloat(num.toPrecision(precision));
};

module.exports = exports;

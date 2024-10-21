exports = function(seed) {
    var min =
        arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 0;
    var max =
        arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : 1;
    var floating =
        arguments.length > 3 && arguments[3] !== undefined
            ? arguments[3]
            : true;
    return function() {
        seed = (seed * 9301 + 49297) % 233280;
        var rnd = seed / 233280.0;
        rnd = min + rnd * (max - min);
        return floating ? rnd : Math.floor(rnd);
    };
};

module.exports = exports;

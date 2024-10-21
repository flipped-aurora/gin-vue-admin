var each = require('./each');
var upperFirst = require('./upperFirst');
exports.linear = function(t) {
    return t;
};
var pow = Math.pow;
var sqrt = Math.sqrt;
var sin = Math.sin;
var min = Math.min;
var asin = Math.asin;
var PI = Math.PI;
var fns = {
    sine: function(t) {
        return 1 + sin((PI / 2) * t - PI / 2);
    },
    circ: function(t) {
        return 1 - sqrt(1 - t * t);
    },
    elastic: function(t, m) {
        m = m || DEFAULT_ELASTICITY;
        if (t === 0 || t === 1) return t;
        var p = 1 - min(m, 998) / 1000;
        var st = t / 1;
        var st1 = st - 1;
        var s = (p / (2 * PI)) * asin(1);
        return -(pow(2, 10 * st1) * sin(((st1 - s) * (2 * PI)) / p));
    },
    back: function(t) {
        return t * t * (3 * t - 2);
    },
    bounce: function(t) {
        var pow2,
            bounce = 4;

        while (t < ((pow2 = pow(2, --bounce)) - 1) / 11) {}
        return (
            1 / pow(4, 3 - bounce) - 7.5625 * pow((pow2 * 3 - 2) / 22 - t, 2)
        );
    }
};
each(['quad', 'cubic', 'quart', 'quint', 'expo'], function(name, i) {
    fns[name] = function(t) {
        return pow(t, i + 2);
    };
});
var DEFAULT_ELASTICITY = 400;
each(fns, function(fn, name) {
    name = upperFirst(name);
    exports['in' + name] = fn;
    exports['out' + name] = function(t, m) {
        return 1 - fn(1 - t, m);
    };
    exports['inOut' + name] = function(t, m) {
        return t < 0.5 ? fn(t * 2, m) / 2 : 1 - fn(t * -2 + 2, m) / 2;
    };
    exports['outIn' + name] = function(t, m) {
        return t < 0.5
            ? (1 - fn(1 - 2 * t, m)) / 2
            : (fn(t * 2 - 1, m) + 1) / 2;
    };
});

module.exports = exports;

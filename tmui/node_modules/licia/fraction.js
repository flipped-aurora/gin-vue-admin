var gcd = require('./gcd');
var precision = require('./precision');
exports = function(num) {
    if (num === 0) return '0';
    var _precision = precision(num);
    _precision = pow(10, _precision);
    var numerator = num * _precision,
        denominator = _precision;
    var _gcd = abs(gcd(numerator, denominator));
    numerator /= _gcd;
    denominator /= _gcd;
    return numerator + '/' + denominator;
};
var abs = Math.abs;
var pow = Math.pow;

module.exports = exports;

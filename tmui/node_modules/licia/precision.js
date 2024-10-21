exports = function(num) {
    num = num.toExponential().match(regExponential);
    var coefficient = num[1];
    var exponent = parseInt(num[2], 10);
    var places = (coefficient.split('.')[1] || '').length;
    var ret = places - exponent;
    return ret < 0 ? 0 : ret;
};
var regExponential = /^(-?\d?\.?\d+)e([+-]\d)+/;

module.exports = exports;

var isNum = require('./isNum');
exports = function(a, b, relTol, absTol) {
    if (!isNum(relTol)) relTol = 1e-9;
    if (!isNum(absTol)) absTol = 0;
    return abs(a - b) <= max(relTol * max(abs(a), abs(b)), absTol);
};
var abs = Math.abs;
var max = Math.max;

module.exports = exports;

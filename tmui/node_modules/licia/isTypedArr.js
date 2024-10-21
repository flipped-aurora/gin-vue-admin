var objToStr = require('./objToStr');
var each = require('./each');
exports = function(val) {
    return !!map[objToStr(val)];
};
var map = {};
each(
    [
        'Int8Array',
        'Int16Array',
        'Int32Array',
        'Uint8Array',
        'Uint8ClampedArray',
        'Uint16Array',
        'Uint32Array',
        'Float32Array',
        'Float64Array'
    ],
    function(val) {
        map['[object ' + val + ']'] = true;
    }
);

module.exports = exports;

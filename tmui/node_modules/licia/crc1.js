var isStr = require('./isStr');
var strToBytes = require('./strToBytes');

exports = function(input, previous) {
    return exports.signed(input, previous) >>> 0;
};
exports.signed = function(input, previous) {
    if (isStr(input)) input = strToBytes(input);
    var crc = ~~previous;
    var accum = 0;
    for (var i = 0, len = input.length; i < len; i++) {
        var byte = input[i];
        accum += byte;
    }
    crc += accum % 256;
    return crc % 256;
};

module.exports = exports;

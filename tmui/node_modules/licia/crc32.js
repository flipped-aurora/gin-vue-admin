var isStr = require('./isStr');
var strToBytes = require('./strToBytes');
var TABLE = [];
for (var n = 0; n < 256; n++) {
    var c = n;
    for (var k = 0; k < 8; k++) {
        if (c & 1) {
            c = 0xedb88320 ^ (c >>> 1);
        } else {
            c = c >>> 1;
        }
    }
    TABLE[n] = c >>> 0;
}
if (typeof Int32Array !== 'undefined') TABLE = new Int32Array(TABLE);
exports = function(input, previous) {
    return exports.signed(input, previous) >>> 0;
};
exports.signed = function(input, previous) {
    if (isStr(input)) input = strToBytes(input);
    var crc = previous === 0 ? 0 : ~~previous ^ -1;
    for (var i = 0, len = input.length; i < len; i++) {
        var byte = input[i];
        crc = TABLE[(crc ^ byte) & 0xff] ^ (crc >>> 8);
    }
    return crc ^ -1;
};

module.exports = exports;

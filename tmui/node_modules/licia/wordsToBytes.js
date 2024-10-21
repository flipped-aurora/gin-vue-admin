exports = function(words) {
    var bytes = [];
    for (var b = 0, len = words.length * 32; b < len; b += 8) {
        bytes.push((words[b >>> 5] >>> (24 - (b % 32))) & 0xff);
    }
    return bytes;
};

module.exports = exports;

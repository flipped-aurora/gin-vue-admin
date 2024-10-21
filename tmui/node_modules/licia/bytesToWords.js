exports = function(bytes) {
    var words = [];
    for (var i = 0, len = bytes.length; i < len; i++) {
        words[i >>> 2] |= bytes[i] << (24 - (i % 4) * 8);
    }
    return words;
};

module.exports = exports;

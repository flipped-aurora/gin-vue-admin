var BASE = 0x811c9dc5;
exports = function(str) {
    var ret = BASE;
    for (var i = 0, len = str.length; i < len; i++) {
        ret ^= str.charCodeAt(i);
        ret += (ret << 1) + (ret << 4) + (ret << 7) + (ret << 8) + (ret << 24);
    }
    return ret >>> 0;
};

module.exports = exports;

exports = function(str) {
    var hash = 5381;
    var i = str.length;
    while (i) {
        hash = (hash << 5) + hash + str.charCodeAt(--i);
    }
    return hash >>> 0;
};

module.exports = exports;

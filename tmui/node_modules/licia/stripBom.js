exports = function(str) {
    if (str.charCodeAt(0) === 0xfeff) {
        return str.slice(1);
    }
    return str;
};

module.exports = exports;

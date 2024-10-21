exports = function(val) {
    if (window.ShadowRoot) {
        return val instanceof ShadowRoot;
    }
    return false;
};

module.exports = exports;

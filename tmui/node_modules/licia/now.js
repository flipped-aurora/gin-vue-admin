if (Date.now && !false) {
    exports = Date.now;
} else {
    exports = function() {
        return new Date().getTime();
    };
}

module.exports = exports;

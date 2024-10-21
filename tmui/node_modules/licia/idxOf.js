exports = function(arr, val, fromIdx) {
    return Array.prototype.indexOf.call(arr, val, fromIdx);
};

module.exports = exports;

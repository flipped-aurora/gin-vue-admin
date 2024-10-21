var hasOwnProp = Object.prototype.hasOwnProperty;
exports = function(obj, key) {
    return hasOwnProp.call(obj, key);
};

module.exports = exports;

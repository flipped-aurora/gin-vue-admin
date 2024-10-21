var ObjToStr = Object.prototype.toString;
exports = function(val) {
    return ObjToStr.call(val);
};

module.exports = exports;

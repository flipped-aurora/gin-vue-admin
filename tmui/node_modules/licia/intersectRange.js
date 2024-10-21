exports = function(a, b) {
    var min = a.start < b.start ? a : b;
    var max = min === a ? b : a;
    if (min.end < max.start) return;
    return {
        start: max.start,
        end: min.end < max.end ? min.end : max.end
    };
};

module.exports = exports;

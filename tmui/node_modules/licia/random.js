exports = function(min, max, floating) {
    if (max == null) {
        max = min;
        min = 0;
    }
    var rand = Math.random();
    if (floating || min % 1 || max % 1) {
        return Math.min(
            min +
                rand *
                    (max - min + parseFloat('1e-' + ((rand + '').length - 1))),
            max
        );
    }
    return min + Math.floor(rand * (max - min + 1));
};

module.exports = exports;

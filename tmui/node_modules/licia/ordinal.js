exports = function(num) {
    var j = num % 10;
    var k = num % 100;
    var indicator = 'th';
    if (j == 1 && k != 11) {
        indicator = 'st';
    }
    if (j == 2 && k != 12) {
        indicator = 'nd';
    }
    if (j == 3 && k != 13) {
        indicator = 'rd';
    }
    return num + indicator;
};

module.exports = exports;

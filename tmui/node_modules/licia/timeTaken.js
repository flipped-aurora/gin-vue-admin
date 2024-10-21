var perfNow = require('./perfNow');
exports = function(fn) {
    var start = perfNow();
    fn();
    return perfNow() - start;
};

module.exports = exports;

var isStr = require('./isStr');
exports = function(val) {
    if (isStr(val)) {
        val = val.toLowerCase();
        return val !== '0' && val !== '' && val !== 'false';
    }
    return !!val;
};

module.exports = exports;

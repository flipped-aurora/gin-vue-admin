const isObj = require('./isObj');
const isFn = require('./isFn');

exports = function(val) {
    return val !== null && isObj(val) && isFn(val.pipe);
};

module.exports = exports;

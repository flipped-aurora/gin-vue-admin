var pick = require('./pick');
exports = function(obj, filter) {
    return pick(obj, filter, true);
};

module.exports = exports;

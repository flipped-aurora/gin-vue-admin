var createAssigner = require('./createAssigner');
var allKeys = require('./allKeys');
exports = createAssigner(allKeys, true);

module.exports = exports;

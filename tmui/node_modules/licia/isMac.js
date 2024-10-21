var detectOs = require('./detectOs');
exports = detectOs() === 'os x';

module.exports = exports;

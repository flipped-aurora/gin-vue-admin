var detectOs = require('./detectOs');
exports = detectOs() === 'windows';

module.exports = exports;

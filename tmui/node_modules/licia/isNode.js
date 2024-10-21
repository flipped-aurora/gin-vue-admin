var objToStr = require('./objToStr');
exports =
    typeof process !== 'undefined' && objToStr(process) === '[object process]';

module.exports = exports;

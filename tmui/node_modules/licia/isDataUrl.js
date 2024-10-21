var trim = require('./trim');
exports = function(str) {
    return regDataUrl.test(trim(str));
};

var regDataUrl = /^data:([a-z]+\/[a-z0-9-+.]+(;[a-z0-9-.!#$%*+.{}|~`]+=[a-z0-9-.!#$%*+.{}|~`]+)*)?(;base64)?,([a-z0-9!$&',()*+;=\-._~:@/?%\s]*?)$/i;

module.exports = exports;

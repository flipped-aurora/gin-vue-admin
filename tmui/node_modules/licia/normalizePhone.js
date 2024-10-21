var trim = require('./trim');
exports = function(phone, options) {
    phone = trim(phone);
    var countryCode = options.countryCode,
        _options$trunkPrefix = options.trunkPrefix,
        trunkPrefix =
            _options$trunkPrefix === void 0 ? false : _options$trunkPrefix;
    var plusSign = regPlusSign.test(phone);
    phone = phone.replace(regNotDigit, '');
    if (plusSign) {
        phone = phone.replace(new RegExp('^'.concat(countryCode)), '');
    }
    if (trunkPrefix) {
        phone = phone.replace(regTrunkPrefix, '');
    }
    return '+'.concat(countryCode + phone);
};
var regPlusSign = /^\+/;
var regNotDigit = /\D/g;
var regTrunkPrefix = /^\d/;

module.exports = exports;

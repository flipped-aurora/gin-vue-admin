exports = function(val) {
    return regEmail.test(val);
};
var regEmail = /.+@.+\..+/;

module.exports = exports;

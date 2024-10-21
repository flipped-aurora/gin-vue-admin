exports = function(val) {
    return regUrl.test(val);
};
var regUrl = /^(?:\w+:)?\/\/([^\s.]+\.\S{2}|localhost[:?\d]*)\S*$/;

module.exports = exports;

var regColor = /\x1B\[([0-9]{1,2}(;[0-9]{1,2})?)?[m|K]/g;
exports = function(str) {
    return str.replace(regColor, '');
};

module.exports = exports;

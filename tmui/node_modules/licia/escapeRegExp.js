exports = function(str) {
    return str.replace(/\W/g, '\\$&');
};

module.exports = exports;

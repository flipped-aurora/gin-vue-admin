var regHtmlTag = /<[^>]*>/g;
exports = function(str) {
    return str.replace(regHtmlTag, '');
};

module.exports = exports;

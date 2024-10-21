var extractUrls = require('./extractUrls');
var each = require('./each');
var escapeRegExp = require('./escapeRegExp');
exports = function(str, hyperlink) {
    hyperlink = hyperlink || defHyperlink;
    var urlList = extractUrls(str);
    each(urlList, function(url) {
        str = str.replace(new RegExp(escapeRegExp(url), 'g'), hyperlink);
    });
    return str;
};
function defHyperlink(url) {
    return '<a href="' + url + '">' + url + '</a>';
}

module.exports = exports;

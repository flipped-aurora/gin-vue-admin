var Url = require('./Url');
exports = function(url1, url2) {
    url1 = new Url(url1);
    url2 = new Url(url2);
    url1.port = url1.port | 0 || (url1.protocol === 'https' ? 443 : 80);
    url2.port = url2.port | 0 || (url2.protocol === 'https' ? 443 : 80);
    return (
        url1.protocol === url2.protocol &&
        url1.hostname === url2.hostname &&
        url1.port === url2.port
    );
};

module.exports = exports;

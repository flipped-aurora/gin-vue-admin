var map = require('./map');
var capitalize = require('./capitalize');
exports = function(header) {
    var ret = specialHeaders[header.toLowerCase()];
    if (!ret) {
        ret = map(header.split('-'), capitalize).join('-');
    }
    return ret;
};
var specialHeaders = {
    'content-md5': 'Content-MD5',
    dnt: 'DNT',
    etag: 'ETag',
    'last-event-id': 'Last-Event-ID',
    tcn: 'TCN',
    te: 'TE',
    'www-authenticate': 'WWW-Authenticate',
    'x-dnsprefetch-control': 'X-DNSPrefetch-Control'
};

module.exports = exports;

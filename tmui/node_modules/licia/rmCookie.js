var cookie = require('./cookie');
exports = function(key) {
    var location = window.location;
    var hostname = location.hostname;
    var pathname = location.pathname;
    var hostNames = hostname.split('.');
    var pathNames = pathname.split('/');
    var domain = '';
    var pathLen = pathNames.length;
    var path;
    if (del()) return;
    for (var i = hostNames.length - 1; i >= 0; i--) {
        var hostName = hostNames[i];
        if (hostName === '') continue;
        domain = domain === '' ? hostName : hostName + '.' + domain;
        path = '/';
        if (
            del({
                domain: domain,
                path: path
            }) ||
            del({
                domain: domain
            })
        )
            return;
        for (var j = 0; j < pathLen; j++) {
            var pathName = pathNames[j];
            if (pathName === '') continue;
            path += pathName;
            if (
                del({
                    domain: domain,
                    path: path
                }) ||
                del({
                    path: path
                })
            )
                return;
            path += '/';
            if (
                del({
                    domain: domain,
                    path: path
                }) ||
                del({
                    path: path
                })
            )
                return;
        }
    }
    function del(options) {
        options = options || {};
        cookie.remove(key, options);
        return !cookie.get(key);
    }
};

module.exports = exports;

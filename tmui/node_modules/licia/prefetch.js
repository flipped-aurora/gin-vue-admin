function prefetchByLink(url) {
    return new Promise(function(resolve, reject) {
        var link = document.createElement('link');
        link.rel = 'prefetch';
        link.href = url;
        link.onload = resolve;
        link.onerror = reject;
        document.head.appendChild(link);
    });
}
function prefetchByXhr(url) {
    return new Promise(function(resolve, reject) {
        var req = new XMLHttpRequest();
        req.open('GET', url, (req.withCredentials = true));
        req.onload = function() {
            req.status === 200 ? resolve() : reject();
        };
        req.send();
    });
}
var link = document.createElement('link');
var supportLink =
    (link.relList || {}).supports && link.relList.supports('prefetch');
exports = supportLink ? prefetchByLink : prefetchByXhr;

module.exports = exports;

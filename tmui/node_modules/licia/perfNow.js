var now = require('./now');
var root = require('./root');
var performance = root.performance;
var process = root.process;
var loadTime;
if (performance && performance.now) {
    exports = function() {
        return performance.now();
    };
} else if (process && process.hrtime) {
    var getNanoSeconds = function() {
        var hr = process.hrtime();
        return hr[0] * 1e9 + hr[1];
    };
    loadTime = getNanoSeconds() - process.uptime() * 1e9;
    exports = function() {
        return (getNanoSeconds() - loadTime) / 1e6;
    };
} else {
    loadTime = now();
    exports = function() {
        return now() - loadTime;
    };
}

module.exports = exports;

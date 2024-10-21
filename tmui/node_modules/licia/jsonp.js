var loadJs = require('./loadJs');
var defaults = require('./defaults');
var noop = require('./noop');
var uniqId = require('./uniqId');
var query = require('./query');
exports = function(options) {
    defaults(options, exports.settings);
    var name = options.name || uniqId('jsonp');
    var param = options.param;
    var timeout = options.timeout;
    var error = options.error;
    var success = options.success;
    var complete = options.complete;
    var data = options.data;
    var url = options.url;
    var timer;
    var isTimeout = false;
    if (timeout > 0) {
        timer = setTimeout(function() {
            isTimeout = true;
            error(new Error('Timeout'));
            complete();
        }, timeout);
    }
    window[name] = function(data) {
        success(data);
        complete();
        window[name] = noop;
    };

    if (!false) {
        data[param] = name;
        data = query.stringify(data);
        url += url.indexOf('?') > -1 ? '&' + data : '?' + data;
    }
    loadJs(url, function(isLoaded) {
        if (isTimeout) return;
        if (timer) clearTimeout(timer);
        if (!isLoaded) {
            error(new Error());
            complete();
        }
    });
};
exports.settings = {
    data: {},
    param: 'callback',
    success: noop,
    error: noop,
    complete: noop,
    timeout: 0
};

module.exports = exports;

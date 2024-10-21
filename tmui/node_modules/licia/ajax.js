var isFn = require('./isFn');
var noop = require('./noop');
var defaults = require('./defaults');
var isObj = require('./isObj');
var query = require('./query');
exports = function(options) {
    defaults(options, exports.setting);
    var type = options.type;
    var url = options.url;
    var data = options.data;
    var dataType = options.dataType;
    var success = options.success;
    var error = options.error;
    var timeout = options.timeout;
    var complete = options.complete;
    var xhr = options.xhr();
    var abortTimeout;
    xhr.onreadystatechange = function() {
        if (xhr.readyState !== 4) return;
        clearTimeout(abortTimeout);
        var result;
        var status = xhr.status;
        if ((status >= 200 && status < 300) || status === 304) {
            result = xhr.responseText;
            if (dataType === 'xml') result = xhr.responseXML;
            try {
                if (dataType === 'json') result = JSON.parse(result);
            } catch (e) {}
            success(result, xhr);
        } else {
            error(xhr);
        }
        complete(xhr);
    };
    if (type === 'GET') {
        data = query.stringify(data);
        if (data) url += url.indexOf('?') > -1 ? '&' + data : '?' + data;
    } else if (options.contentType === 'application/x-www-form-urlencoded') {
        if (isObj(data)) data = query.stringify(data);
    } else if (options.contentType === 'application/json') {
        if (isObj(data)) data = JSON.stringify(data);
    }
    xhr.open(type, url, true);
    xhr.setRequestHeader('Content-Type', options.contentType);
    if (timeout > 0) {
        abortTimeout = setTimeout(function() {
            xhr.onreadystatechange = noop;
            xhr.abort();
            error(xhr, 'timeout');
            complete(xhr);
        }, timeout);
    }
    xhr.send(type === 'GET' ? null : data);
    return xhr;
};
exports.setting = {
    type: 'GET',
    success: noop,
    error: noop,
    complete: noop,
    dataType: 'json',
    contentType: 'application/x-www-form-urlencoded',
    data: {},
    xhr: function() {
        return new XMLHttpRequest();
    },
    timeout: 0
};
exports.get = function() {
    return exports(parseArgs.apply(null, arguments));
};
exports.post = function() {
    var options = parseArgs.apply(null, arguments);
    options.type = 'POST';
    return exports(options);
};
function parseArgs(url, data, success, dataType) {
    if (isFn(data)) {
        dataType = success;
        success = data;
        data = {};
    }
    return {
        url: url,
        data: data,
        success: success,
        dataType: dataType
    };
}

module.exports = exports;

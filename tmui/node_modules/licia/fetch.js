var each = require('./each');
var defaults = require('./defaults');
var noop = require('./noop');
var has = require('./has');
var root = require('./root');
var Promise = root.Promise;
exports = function(url, options) {
    options = options || {};
    defaults(options, exports.setting);
    return new Promise(function(resolve, reject) {
        var xhr = options.xhr();
        var headers = options.headers;
        var body = options.body;
        var timeout = options.timeout;
        var abortTimer;
        xhr.withCredentials = options.credentials == 'include';
        xhr.onload = function() {
            clearTimeout(abortTimer);
            resolve(getRes(xhr));
        };
        xhr.onerror = reject;
        xhr.open(options.method, url, true);
        each(headers, function(val, key) {
            xhr.setRequestHeader(key, val);
        });
        if (timeout > 0) {
            setTimeout(function() {
                xhr.onload = noop;
                xhr.abort();
                reject(Error('timeout'));
            }, timeout);
        }
        xhr.send(body);
    });
};
var regHeaders = /^(.*?):\s*([\s\S]*?)$/gm;
function getRes(xhr) {
    var keys = [];
    var all = [];
    var headers = {};
    var header;
    xhr.getAllResponseHeaders().replace(regHeaders, function(m, key, val) {
        key = key.toLowerCase();
        keys.push(key);

        all.push([key, val]);
        header = headers[key];
        headers[key] = header ? header + ',' + val : val;
    });
    return {
        ok: xhr.status >= 200 && xhr.status < 400,
        status: xhr.status,
        statusText: xhr.statusText,
        url: xhr.responseURL,
        clone: function() {
            return getRes(xhr);
        },
        text: function() {
            return Promise.resolve(xhr.responseText);
        },
        json: function() {
            return Promise.resolve(xhr.responseText).then(JSON.parse);
        },
        xml: function() {
            return Promise.resolve(xhr.responseXML);
        },
        blob: function() {
            return Promise.resolve(new Blob([xhr.response]));
        },
        headers: {
            keys: function() {
                return keys;
            },
            entries: function() {
                return all;
            },
            get: function(name) {
                return headers[name.toLowerCase()];
            },
            has: function(name) {
                return has(headers, name);
            }
        }
    };
}
exports.setting = {
    method: 'GET',
    headers: {},
    timeout: 0,
    xhr: function() {
        return new XMLHttpRequest();
    }
};

module.exports = exports;

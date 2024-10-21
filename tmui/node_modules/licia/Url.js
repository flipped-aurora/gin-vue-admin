var Class = require('./Class');
var extend = require('./extend');
var trim = require('./trim');
var query = require('./query');
var isEmpty = require('./isEmpty');
var each = require('./each');
var isArr = require('./isArr');
var toArr = require('./toArr');
var isBrowser = require('./isBrowser');
var isObj = require('./isObj');
var toStr = require('./toStr');
exports = Class(
    {
        className: 'Url',
        initialize: function(url) {
            if (!url && isBrowser) url = window.location.href;
            extend(this, exports.parse(url || ''));
        },
        setQuery: function(name, val) {
            var query = this.query;
            if (isObj(name)) {
                each(name, function(val, key) {
                    query[key] = toStr(val);
                });
            } else {
                query[name] = toStr(val);
            }
            return this;
        },
        rmQuery: function(name) {
            var query = this.query;
            if (!isArr(name)) name = toArr(name);
            each(name, function(key) {
                delete query[key];
            });
            return this;
        },
        toString: function() {
            return exports.stringify(this);
        }
    },
    {
        parse: function(url) {
            var ret = {
                protocol: '',
                auth: '',
                hostname: '',
                hash: '',
                query: {},
                port: '',
                pathname: '',
                slashes: false
            };
            var rest = trim(url);
            var slashes = false;
            var proto = rest.match(regProto);
            if (proto) {
                proto = proto[0];
                ret.protocol = proto.toLowerCase();
                rest = rest.substr(proto.length);
            }
            if (proto) {
                slashes = rest.substr(0, 2) === '//';
                if (slashes) {
                    rest = rest.slice(2);
                    ret.slashes = true;
                }
            }
            if (slashes) {
                var host = rest;
                var hostEnd = -1;
                for (var i = 0, len = hostEndingChars.length; i < len; i++) {
                    var pos = rest.indexOf(hostEndingChars[i]);
                    if (pos !== -1 && (hostEnd === -1 || pos < hostEnd))
                        hostEnd = pos;
                }
                if (hostEnd > -1) {
                    host = rest.slice(0, hostEnd);
                    rest = rest.slice(hostEnd);
                }
                var atSign = host.lastIndexOf('@');
                if (atSign !== -1) {
                    ret.auth = decodeURIComponent(host.slice(0, atSign));
                    host = host.slice(atSign + 1);
                }
                ret.hostname = host;
                var port = host.match(regPort);
                if (port) {
                    port = port[0];
                    if (port !== ':') ret.port = port.substr(1);
                    ret.hostname = host.substr(0, host.length - port.length);
                }
            }
            var hash = rest.indexOf('#');
            if (hash !== -1) {
                ret.hash = rest.substr(hash);
                rest = rest.slice(0, hash);
            }
            var queryMark = rest.indexOf('?');
            if (queryMark !== -1) {
                ret.query = query.parse(rest.substr(queryMark + 1));
                rest = rest.slice(0, queryMark);
            }
            ret.pathname = rest || '/';
            return ret;
        },
        stringify: function(obj) {
            var ret =
                obj.protocol +
                (obj.slashes ? '//' : '') +
                (obj.auth ? encodeURIComponent(obj.auth) + '@' : '') +
                obj.hostname +
                (obj.port ? ':' + obj.port : '') +
                obj.pathname;
            if (!isEmpty(obj.query)) ret += '?' + query.stringify(obj.query);
            if (obj.hash) ret += obj.hash;
            return ret;
        }
    }
);
var regProto = /^([a-z0-9.+-]+:)/i;
var regPort = /:[0-9]*$/;
var hostEndingChars = ['/', '?', '#'];

module.exports = exports;

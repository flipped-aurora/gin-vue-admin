var defaults = require('./defaults');
var Emitter = require('./Emitter');
exports = Emitter.extend({
    initialize: function Socket(url) {
        var options =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {};
        this.callSuper(Emitter, 'initialize');
        defaults(options, defOpts);
        this._options = options;
        this._url = url;
        this.connect();
    },
    send: function(message) {
        this._ws.send(message);
    },
    close: function(code, reason) {
        this._ws.close(code || 1e3, reason);
    },
    connect: function() {
        var _this = this;
        var options = this._options;
        var ws = new WebSocket(this._url, options.protocols);
        ws.onmessage = function(e) {
            return _this.emit('message', e);
        };
        ws.onopen = function(e) {
            return _this.emit('open', e);
        };
        ws.onclose = function(e) {
            var code = e.code;
            if (
                code !== 1e3 &&
                code !== 1001 &&
                code !== 1005 &&
                options.reconnect
            ) {
                _this.connect();
            }
            _this.emit('close', e);
        };
        ws.onerror = function(e) {
            if (e && e.code === 'ECONNREFUSED' && options.reconnect) {
                _this.connect();
            } else {
                _this.emit('error', e);
            }
        };
        this._ws = ws;
    }
});
var defOpts = {
    protocols: [],
    reconnect: true
};

module.exports = exports;

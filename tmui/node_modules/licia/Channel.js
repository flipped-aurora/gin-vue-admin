var Emitter = require('./Emitter');
var each = require('./each');
var remove = require('./remove');
var some = require('./some');
exports = Emitter.extend({
    initialize: function Channel() {
        this._connections = [];
        this.callSuper(Emitter, 'initialize');
    },
    send: function(msg) {
        var _this = this;
        each(this._connections, function(connection) {
            connection.emit('message', msg, _this);
        });
    },
    connect: function(connection) {
        if (this.isConnected(connection)) {
            return;
        }
        this._connections.push(connection);
        connection.connect(this);
    },
    disconnect: function(connection) {
        if (!this.isConnected(connection)) {
            return;
        }
        remove(this._connections, function(item) {
            return item === connection;
        });
        connection.disconnect(this);
    },
    isConnected: function(connection) {
        if (connection === this) {
            throw new Error('Channel cannot be connected to itself.');
        }
        return some(this._connections, function(item) {
            return item === connection;
        });
    },
    destroy: function() {
        var _this2 = this;
        each(this._connections, function(connection) {
            _this2.disconnect(connection);
        });
    }
});

module.exports = exports;

var Emitter = require('./Emitter');
var root = require('./root');
var each = require('./each');
var Notification = root.Notification;
exports = function(title, options) {
    var notification = new exports.Notification(title, options);
    notification.show();
};
exports.Notification = Emitter.extend({
    initialize: function Notification(title) {
        var options =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {};
        this._options = options;
        this._title = title;
        this.callSuper(Emitter, 'initialize', arguments);
    },
    handleEvent: function(e) {
        this.emit(e.type, e);
    },
    show: function() {
        var _this = this;
        if (!Notification) {
            return this.emit('error', Error('Notification is not supported'));
        }
        if (Notification.permission === 'granted') {
            this._show();
        } else {
            Notification.requestPermission(function(permission) {
                switch (permission) {
                    case 'granted':
                        _this._show();
                        break;
                    case 'denied':
                        _this.emit(
                            'error',
                            Error('Notification permission is denied')
                        );
                        break;
                }
            });
        }
    },
    _show: function() {
        var _this2 = this;
        var notification = new Notification(this._title, this._options);
        each(['show', 'close', 'click'], function(type) {
            notification.addEventListener(type, _this2, false);
        });
    }
});

module.exports = exports;

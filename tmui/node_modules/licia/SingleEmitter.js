var Class = require('./Class');
var clone = require('./clone');
var each = require('./each');
var toArr = require('./toArr');
exports = Class(
    {
        initialize: function SingleEmitter() {
            this._listeners = [];
        },
        addListener: function(listener) {
            this._listeners.push(listener);
        },
        rmListener: function(listener) {
            var idx = this._listeners.indexOf(listener);
            if (idx > -1) {
                this._listeners.splice(idx, 1);
            }
        },
        rmAllListeners: function() {
            this._listeners = [];
        },
        emit: function() {
            var _this = this;
            var args = toArr(arguments);
            var listeners = clone(this._listeners);
            each(
                listeners,
                function(listener) {
                    return listener.apply(_this, args);
                },
                this
            );
        }
    },
    {
        mixin: function(obj) {
            each(
                ['addListener', 'rmListener', 'emit', 'rmAllListeners'],
                function(val) {
                    obj[val] = exports.prototype[val];
                }
            );
            obj._listeners = obj._listeners || [];
        }
    }
);

module.exports = exports;

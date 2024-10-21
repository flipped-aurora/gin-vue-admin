var Class = require('./Class');
var has = require('./has');
var each = require('./each');
var slice = require('./slice');
var once = require('./once');
var clone = require('./clone');
exports = Class(
    {
        initialize: function Emitter() {
            this._events = this._events || {};
        },
        on: function(event, listener) {
            this._events[event] = this._events[event] || [];
            this._events[event].push(listener);
            return this;
        },
        off: function(event, listener) {
            var events = this._events;
            if (!has(events, event)) return;
            var idx = events[event].indexOf(listener);
            if (idx > -1) {
                events[event].splice(idx, 1);
            }
            return this;
        },
        once: function(event, listener) {
            this.on(event, once(listener));
            return this;
        },
        emit: function(event) {
            var _this = this;
            if (!has(this._events, event)) return;
            var args = slice(arguments, 1);
            var events = clone(this._events[event]);
            each(
                events,
                function(val) {
                    return val.apply(_this, args);
                },
                this
            );
            return this;
        },
        removeAllListeners: function(event) {
            if (!event) {
                this._events = {};
            } else {
                delete this._events[event];
            }
            return this;
        }
    },
    {
        mixin: function(obj) {
            each(['on', 'off', 'once', 'emit', 'removeAllListeners'], function(
                val
            ) {
                obj[val] = exports.prototype[val];
            });
            obj._events = obj._events || {};
        }
    }
);

module.exports = exports;

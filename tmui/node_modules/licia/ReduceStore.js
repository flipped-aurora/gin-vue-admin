var Class = require('./Class');
var clone = require('./clone');
var remove = require('./remove');
exports = Class({
    initialize: function ReduceStore(reducer, initialState) {
        this._reducer = reducer;
        this._state = initialState;
        this._curListeners = [];
        this._nextListeners = this._curListeners;
    },
    subscribe: function(listener) {
        var isSubscribed = true;
        this._ensureCanMutateNextListeners();
        this._nextListeners.push(listener);
        var self = this;
        return function() {
            if (!isSubscribed) return;
            isSubscribed = false;
            self._ensureCanMutateNextListeners();
            remove(self._nextListeners, function(val) {
                return val === listener;
            });
        };
    },
    dispatch: function(action) {
        this._state = this._reducer(this._state, action);
        var listeners = (this._curListeners = this._nextListeners);
        for (var i = 0, len = listeners.length; i < len; i++) listeners[i]();
        return action;
    },
    getState: function() {
        return this._state;
    },
    _ensureCanMutateNextListeners: function() {
        if (this._nextListeners === this._curListeners) {
            this._nextListeners = clone(this._curListeners);
        }
    }
});

module.exports = exports;

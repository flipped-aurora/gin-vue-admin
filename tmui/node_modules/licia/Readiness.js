var Class = require('./Class');
var toArr = require('./toArr');
var each = require('./each');
var map = require('./map');
var noop = require('./noop');
var some = require('./some');
exports = Class({
    initialize: function Readiness() {
        this._promises = {};
        this._resolves = {};
        this._states = {};
    },
    signal: function(tasks) {
        var states = this._states;
        each(this._getPromises(toArr(tasks)), function(val) {
            if (!val.state) {
                states[val.task] = true;
                val.resolve();
            }
        });
    },
    isReady: function(tasks) {
        return !some(this._getPromises(toArr(tasks)), function(val) {
            return !val.state;
        });
    },
    ready: function(tasks) {
        var fn =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : noop;
        return Promise.all(
            map(this._getPromises(toArr(tasks)), function(val) {
                return val.promise;
            })
        ).then(fn);
    },
    _getPromises: function(tasks) {
        var promises = this._promises;
        var resolves = this._resolves;
        var states = this._states;
        return map(tasks, function(task) {
            if (!promises[task]) {
                promises[task] = new Promise(function(resolve) {
                    resolves[task] = resolve;
                    states[task] = false;
                });
            }
            return {
                task: task,
                promise: promises[task],
                resolve: resolves[task],
                state: states[task]
            };
        });
    }
});

module.exports = exports;

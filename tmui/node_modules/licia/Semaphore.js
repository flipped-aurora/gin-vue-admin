var Class = require('./Class');
var Queue = require('./Queue');
exports = Class({
    initialize: function Semaphore() {
        var counter =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : 1;
        this._counter = counter;
        this._tasks = new Queue();
    },
    wait: function(fn) {
        if (this._counter > 0) {
            this._counter--;
            return fn();
        }
        this._tasks.enqueue(fn);
    },
    signal: function() {
        var task = this._tasks.dequeue();
        if (task) {
            return task();
        }
        this._counter++;
    }
});

module.exports = exports;

var Class = require('./Class');
var Heap = require('./Heap');
var isSorted = require('./isSorted');
var wrap = require('./wrap');
exports = Class({
    initialize: function PriorityQueue() {
        var cmp =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : isSorted.defComparator;
        this._heap = new Heap(
            wrap(cmp, function(fn, a, b) {
                return fn(a, b) * -1;
            })
        );
        this.size = 0;
    },
    clear: function() {
        this._heap.clear();
        this.size = 0;
    },
    enqueue: function(item) {
        this._heap.add(item);
        this.size++;
        return this.size;
    },
    dequeue: function() {
        var item = this._heap.poll();
        if (item) {
            this.size--;
            return item;
        }
    },
    peek: function() {
        return this._heap.peek();
    }
});

module.exports = exports;

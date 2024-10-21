var Class = require('./Class');
exports = Class({
    initialize: function Queue() {
        this.clear();
    },
    clear: function() {
        this._items = [];
        this.size = 0;
    },
    enqueue: function(item) {
        this._items.push(item);
        return ++this.size;
    },
    dequeue: function() {
        if (!this.size) return;
        this.size--;
        return this._items.shift();
    },
    peek: function() {
        if (!this.size) return;
        return this._items[0];
    },
    forEach: function(iterator, ctx) {
        ctx = arguments.length > 1 ? ctx : this;
        var items = this._items;
        for (var i = 0, size = this.size; i < size; i++) {
            iterator.call(ctx, items[i], i, this);
        }
    },
    toArr: function() {
        return this._items.slice(0);
    }
});

module.exports = exports;

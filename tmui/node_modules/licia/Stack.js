var Class = require('./Class');
var reverse = require('./reverse');
exports = Class({
    initialize: function Stack() {
        this.clear();
    },
    clear: function() {
        this._items = [];
        this.size = 0;
    },
    push: function(item) {
        this._items.push(item);
        return ++this.size;
    },
    pop: function() {
        if (!this.size) return;
        this.size--;
        return this._items.pop();
    },
    peek: function() {
        return this._items[this.size - 1];
    },
    forEach: function(iterator, ctx) {
        ctx = arguments.length > 1 ? ctx : this;
        var items = this._items;
        for (var i = this.size - 1, j = 0; i >= 0; i--, j++) {
            iterator.call(ctx, items[i], j, this);
        }
    },
    toArr: function() {
        return reverse(this._items);
    }
});

module.exports = exports;

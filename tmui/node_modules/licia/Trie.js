var Class = require('./Class');
var each = require('./each');

exports = Class({
    initialize: function Trie() {
        this.clear();
    },
    add: function(word) {
        var edges = this._edges;
        var node = this._root;
        this._wordsInSubtree[node]++;
        for (var i = 0, len = word.length; i < len; i++) {
            var edge = word[i];
            var next = edges[node][edge];
            if (!next) {
                if (this._freeNodes.length) {
                    next = this._freeNodes.pop();
                } else {
                    next = this._idx++;
                    this._isWord.push(false);
                    this._wordsInSubtree.push(0);
                    edges.push({});
                }
                edges[node][edge] = next;
            }
            this._wordsInSubtree[next]++;
            node = next;
        }
        this._isWord[node] = true;
    },
    remove: function(word) {
        if (!this.has(word)) {
            return;
        }
        var node = this._root;
        this._wordsInSubtree[node]--;
        for (var i = 0, len = word.length; i < len; i++) {
            var edge = word[i];
            var next = this._edges[node][edge];
            if (!--this._wordsInSubtree[next]) {
                delete this._edges[node][edge];
                this._freeNodes.push(next);
            }
            node = next;
        }
        this._isWord[node] = false;
    },
    has: function(word) {
        var node = this._root;
        for (var i = 0, len = word.length; i < len; i++) {
            node = this._edges[node][word[i]];
            if (!node) {
                return false;
            }
        }
        return this._isWord[node];
    },
    words: function(prefix) {
        var node = this._root;
        for (var i = 0, len = prefix.length; i < len; i++) {
            node = this._edges[node][prefix[i]];
            if (!node) {
                return [];
            }
        }
        var result = [];
        this._dfs(node, prefix, result);
        return result;
    },
    clear: function() {
        this._idx = 1;
        this._root = 0;
        this._edges = [{}];
        this._isWord = [false];
        this._wordsInSubtree = [0];
        this._freeNodes = [];
    },
    _dfs: function(node, prefix, result) {
        var _this = this;
        if (this._isWord[node]) {
            result.push(prefix);
        }
        var edges = this._edges[node];
        each(edges, function(node, edge) {
            return _this._dfs(node, prefix + edge, result);
        });
    }
});

module.exports = exports;

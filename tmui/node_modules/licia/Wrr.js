var Class = require('./Class');
var max = require('./max');
var map = require('./map');
var reduce = require('./reduce');
var gcd = require('./gcd');
var filter = require('./filter');
exports = Class({
    initialize: function Wrr() {
        this._peers = [];
    },
    set: function(val, weight) {
        var peers = this._peers;
        var size = this.size;
        for (var i = 0; i < size; i++) {
            var peer = peers[i];
            if (peer.val === val) {
                peer.weight = weight;
                this._reset();
                return;
            }
        }
        peers.push({
            val: val,
            weight: weight
        });
        this._reset();
    },
    get: function(val) {
        var peers = this._peers;
        var size = this.size;
        for (var i = 0; i < size; i++) {
            var peer = peers[i];
            if (peer.val === val) {
                return peer.weight;
            }
        }
    },
    remove: function(val) {
        this._peers = filter(this._peers, function(peer) {
            return peer.val !== val;
        });
        this._reset();
    },
    next: function() {
        var peers = this._peers;
        var size = this.size;
        if (size === 0) return;

        while (true) {
            this._i = (this._i + 1) % size;
            if (this._i === 0) {
                this._cw = this._cw - this._gcdS;
                if (this._cw <= 0) {
                    this._cw = this._maxS;
                }
            }
            if (this._cw === 0) return;
            if (peers[this._i].weight >= this._cw) {
                return peers[this._i].val;
            }
        }
    },
    clear: function() {
        this._peers = [];
        this._reset();
    },
    _reset: function() {
        var peers = this._peers;
        this.size = peers.length;
        var weights = map(peers, function(peer) {
            return peer.weight;
        });
        this._i = -1;
        this._cw = 0;
        this._maxS = max.apply(null, weights);
        this._gcdS = reduce(
            weights,
            function(prev, weight) {
                return gcd(prev, weight);
            },
            0
        );
    }
});

module.exports = exports;

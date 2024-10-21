var Emitter = require('./Emitter');
exports = Emitter.extend({
    className: 'MediaQuery',
    initialize: function(query) {
        var _this = this;
        this.callSuper(Emitter, 'initialize');
        this._listener = function() {
            _this.emit(_this.isMatch() ? 'match' : 'unmatch');
        };
        this.setQuery(query);
    },
    setQuery: function(query) {
        if (this._mql) {
            this._mql.removeListener(this._listener);
        }
        this._mql = window.matchMedia(query);
        this._mql.addListener(this._listener);
    },
    isMatch: function() {
        return this._mql.matches;
    }
});

module.exports = exports;

var Store = require('./Store');
var safeStorage = require('./safeStorage');
var isEmpty = require('./isEmpty');
var stringify = require('./stringify');
var defaults = require('./defaults');
var isObj = require('./isObj');
var localStorage = safeStorage('local');
exports = Store.extend({
    initialize: function LocalStore(name, data) {
        this._name = name;
        data = data || {};
        var localData = localStorage.getItem(name);
        try {
            localData = JSON.parse(localData);
        } catch (e) {
            localData = {};
        }
        if (!isObj(localData)) localData = {};
        data = defaults(localData, data);
        this.callSuper(Store, 'initialize', [data]);
    },
    save: function(data) {
        if (isEmpty(data)) return localStorage.removeItem(this._name);
        localStorage.setItem(this._name, stringify(data));
    }
});

module.exports = exports;

var Store = require('./Store');
var safeStorage = require('./safeStorage');
var isEmpty = require('./isEmpty');
var stringify = require('./stringify');
var defaults = require('./defaults');
var isObj = require('./isObj');
var sessionStorage = safeStorage('session');
exports = Store.extend({
    initialize: function SessionStore(name, data) {
        this._name = name;
        data = data || {};
        var sessionData = sessionStorage.getItem(name);
        try {
            sessionData = JSON.parse(sessionData);
        } catch (e) {
            sessionData = {};
        }
        if (!isObj(sessionData)) sessionData = {};
        data = defaults(sessionData, data);
        this.callSuper(Store, 'initialize', [data]);
    },
    save: function(data) {
        if (isEmpty(data)) return sessionStorage.removeItem(this._name);
        sessionStorage.setItem(this._name, stringify(data));
    }
});

module.exports = exports;

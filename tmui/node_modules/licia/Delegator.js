var Class = require('./Class');
var safeGet = require('./safeGet');
var defineProp = require('./defineProp');
var isStr = require('./isStr');
exports = Class({
    initialize: function Delegator(host, target) {
        this._host = host;
        if (isStr(target)) {
            target = safeGet(host, target);
        }
        this._target = target;
    },
    method: function(name, targetName) {
        var target = this._target;
        var fn = target[targetName || name];
        this._host[name] = function() {
            return fn.apply(target, arguments);
        };
        return this;
    },
    getter: function(name, targetName) {
        var target = this._target;
        targetName = targetName || name;
        defineProp(this._host, name, {
            get: function() {
                return target[targetName];
            },
            configurable: true
        });
        return this;
    },
    setter: function(name, targetName) {
        var target = this._target;
        targetName = targetName || name;
        defineProp(this._host, name, {
            set: function(val) {
                return (target[targetName] = val);
            },
            configurable: true
        });
        return this;
    },
    access: function(name, targetName) {
        return this.getter(name, targetName).setter(name, targetName);
    }
});

module.exports = exports;

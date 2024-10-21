var extend = require('./extend');
var toArr = require('./toArr');
var inherits = require('./inherits');
var safeGet = require('./safeGet');
var isMiniProgram = require('./isMiniProgram');
exports = function(methods, statics) {
    return Base.extend(methods, statics);
};
function makeClass(parent, methods, statics) {
    statics = statics || {};
    var className =
        methods.className || safeGet(methods, 'initialize.name') || '';
    delete methods.className;
    var ctor = function() {
        var args = toArr(arguments);
        return this.initialize
            ? this.initialize.apply(this, args) || this
            : this;
    };
    if (!isMiniProgram) {
        try {
            ctor = new Function(
                'toArr',
                'return function ' +
                    className +
                    '()' +
                    '{' +
                    'var args = toArr(arguments);' +
                    'return this.initialize ? this.initialize.apply(this, args) || this : this;' +
                    '};'
            )(toArr);
        } catch (e) {}
    }
    inherits(ctor, parent);
    ctor.prototype.constructor = ctor;
    ctor.extend = function(methods, statics) {
        return makeClass(ctor, methods, statics);
    };
    ctor.inherits = function(Class) {
        inherits(ctor, Class);
    };
    ctor.methods = function(methods) {
        extend(ctor.prototype, methods);
        return ctor;
    };
    ctor.statics = function(statics) {
        extend(ctor, statics);
        return ctor;
    };
    ctor.methods(methods).statics(statics);
    return ctor;
}
var Base = (exports.Base = makeClass(Object, {
    className: 'Base',
    callSuper: function(parent, name, args) {
        var superMethod = parent.prototype[name];
        return superMethod.apply(this, args);
    },
    toString: function() {
        return this.constructor.name;
    }
}));

module.exports = exports;

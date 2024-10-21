var isObj = require('./isObj');
exports = function(proto) {
    if (!isObj(proto)) return {};
    if (objCreate && !false) return objCreate(proto);
    function noop() {}
    noop.prototype = proto;
    return new noop();
};
var objCreate = Object.create;

module.exports = exports;

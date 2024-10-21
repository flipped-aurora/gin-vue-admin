var toArr = require('./toArr');
exports = function(name, requires, method) {
    if (arguments.length === 2) {
        method = requires;
        requires = [];
    }
    define(name, requires, method);
};
var modules = (exports._modules = {});
function define(name, requires, method) {
    modules[name] = {
        requires: toArr(requires),
        body: method
    };
}

module.exports = exports;

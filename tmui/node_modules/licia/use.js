var map = require('./map');
var define = require('./define');
var has = require('./has');
var toArr = require('./toArr');
exports = function(requires, method) {
    if (method == null) {
        method = requires;
        requires = [];
    }
    requires = map(toArr(requires), function(val) {
        return req(val);
    });
    method.apply(null, requires);
};
var modules = define._modules;
var requireMarks = {};
function req(name) {
    if (has(requireMarks, name)) return modules[name];
    var requires = modules[name].requires;
    var body = modules[name].body;
    var len = requires.length;
    for (var i = 0; i < len; i++) requires[i] = req(requires[i]);
    var exports = body.apply(null, requires);
    if (exports) modules[name] = exports;
    requireMarks[name] = true;
    return modules[name];
}

module.exports = exports;

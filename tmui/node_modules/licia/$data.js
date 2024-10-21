var $attr = require('./$attr');
var isStr = require('./isStr');
var isObj = require('./isObj');
var each = require('./each');
var $safeEls = require('./$safeEls');
exports = function(nodes, name, val) {
    var dataName = name;
    if (isStr(name)) dataName = 'data-' + name;
    if (isObj(name)) {
        dataName = {};
        each(name, function(val, key) {
            dataName['data-' + key] = val;
        });
    }
    return $attr(nodes, dataName, val);
};

module.exports = exports;

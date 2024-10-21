var each = require('./each');
var trim = require('./trim');
var safeSet = require('./safeSet');
var safeGet = require('./safeGet');
var endWith = require('./endWith');
var isArr = require('./isArr');
var isObj = require('./isObj');

var regSection = /^\[([^\]]*)\]$/i;
var regKeyVal = /^([^=]+)(=(.*))?$/i;
var regComment = /^\s*[;#]/;
function parse(ini) {
    var ret = {};
    var section = ret;
    each(ini.split('\n'), function(line) {
        line = trim(line);
        if (!line || line.match(regComment)) return;
        var match = line.match(regSection);
        if (match && match[1]) {
            var _key = match[1];
            section = safeGet(ret, _key) || {};
            return safeSet(ret, _key, section);
        }
        match = line.match(regKeyVal);
        if (!match) return;
        var key = trim(match[1]);
        var val = match[2] ? trim(match[3]) : true;
        if (val === 'true') val = true;
        if (val === 'false') val = false;
        if (val === 'null') val = null;
        if (endWith(key, '[]')) {
            key = key.substring(0, key.length - 2);
            if (!section[key]) section[key] = [];
        }
        isArr(section[key]) ? section[key].push(val) : (section[key] = val);
    });
    return ret;
}
function stringify(obj) {
    var options =
        arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
    var ret = '';
    var section = options.section;
    var whitespace = options.whitespace;
    var separator = whitespace ? ' = ' : '=';
    var children = [];
    each(obj, function(val, key) {
        if (isArr(val)) {
            each(val, function(item) {
                ret += ''
                    .concat(key, '[]')
                    .concat(separator)
                    .concat(item, '\n');
            });
        } else if (isObj(val)) {
            children.push({
                key: key,
                val: val
            });
        } else {
            ret += ''
                .concat(key)
                .concat(separator)
                .concat(val, '\n');
        }
    });
    if (section && ret) {
        ret = '['.concat(section, ']\n') + ret;
    }
    section = section ? section + '.' : '';
    each(children, function(child) {
        child = stringify(child.val, {
            section: section + child.key,
            whitespace: options.whitespace
        });
        if (child) {
            if (ret) {
                ret += '\n';
            }
            ret += child;
        }
    });
    return ret;
}
exports = {
    parse: parse,
    stringify: stringify
};

module.exports = exports;

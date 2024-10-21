var defaults = require('./defaults');
var toNum = require('./toNum');
var invert = require('./invert');
var toBool = require('./toBool');
exports = function(args, opts) {
    opts = opts || {};
    defaults(opts, defOpts);
    var names = opts.names;
    var shorthands = invert(opts.shorthands);
    var remain = [];
    var ret = {
        remain: remain
    };
    var name;
    var type;
    for (var i = 0, len = args.length; i < len; i++) {
        var arg = args[i];
        var nextArg = args[i + 1];
        var match = arg.match(regDoubleDash);
        if (match) {
            name = match[1];
            type = names[name];
            if (!type) {
                remain.push(arg);
            } else if (nextArg && !regDashStart.test(nextArg)) {
                setArg(name, nextArg);
                i++;
            } else if (type === 'boolean') {
                setArg(name, true);
                i++;
            }
            continue;
        }
        match = arg.match(regSingleDash);
        if (match) {
            var letters = match[1];
            for (var j = 0; j < letters.length; j++) {
                var letter = letters[j];
                name = shorthands[letter];
                if (!name) continue;
                type = names[name];
                if (type === 'boolean') setArg(shorthands[letter], true);
            }
            continue;
        }
        remain.push(arg);
    }
    function setArg(name, val) {
        var type = names[name];
        switch (type) {
            case 'number':
                val = toNum(val);
                break;
            case 'boolean':
                val = toBool(val);
                break;
            default:
                break;
        }
        ret[name] = val;
    }
    return ret;
};
var defOpts = {
    names: {},
    shorthands: {}
};
var regDoubleDash = /^--(.+)/;
var regSingleDash = /^-([^-]+)/;
var regDashStart = /^-/;

module.exports = exports;

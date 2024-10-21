var escape = require('./escape');
var defaults = require('./defaults');
var escapeJsStr = require('./escapeJsStr');

var regMatcher = /<%-([\s\S]+?)%>|<%=([\s\S]+?)%>|<%([\s\S]+?)%>|$/g;
exports = function(str, util) {
    if (!util) {
        util =
            typeof _ === 'object'
                ? _
                : {
                      escape: escape
                  };
    } else {
        defaults(util, {
            escape: escape
        });
    }
    var index = 0;
    var src = "__p+='";
    str.replace(regMatcher, function(
        match,
        escape,
        interpolate,
        evaluate,
        offset
    ) {
        src += escapeJsStr(str.slice(index, offset));
        index = offset + match.length;
        if (escape) {
            src += "'+\n((__t=(".concat(
                escape,
                "))==null?'':util.escape(__t))+\n'"
            );
        } else if (interpolate) {
            src += "'+\n((__t=(".concat(interpolate, "))==null?'':__t)+\n'");
        } else if (evaluate) {
            src += "';\n".concat(evaluate, "\n__p+='");
        }
        return match;
    });
    src += "';\n";
    src = 'with(obj||{}){\n'.concat(src, '}\n');
    src = "var __t,__p='',__j=Array.prototype.join,print=function(){__p+=__j.call(arguments,'');};\n".concat(
        src,
        'return __p;\n'
    );
    var render = new Function('obj', 'util', src);
    return function(data) {
        return render.call(null, data, util);
    };
};

module.exports = exports;

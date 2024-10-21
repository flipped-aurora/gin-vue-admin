var template = require('./template');
var each = require('./each');
var map = require('./map');
var rpad = require('./rpad');
var ansiColor = require('./ansiColor');
var toArr = require('./toArr');
var cloneDeep = require('./cloneDeep');
var strWidth = require('./strWidth');
var max = require('./max');
exports = function(data) {
    data = cloneDeep(data);
    data.usage = toArr(data.usage);
    if (data.commands) {
        var cmdNameWidths = map(data.commands, function(command) {
            return strWidth(command.name);
        });
        data.maxNameWidth = max.apply(null, cmdNameWidths);
        return helpTpl(data);
    }
    each(data.options, function(option) {
        option.name =
            (option.shorthand ? '-' + option.shorthand + ', ' : '    ') +
            '--' +
            option.name;
    });
    var optNameWidths = map(data.options, function(option) {
        return strWidth(option.name);
    });
    data.maxNameWidth = max.apply(null, optNameWidths);
    return cmdTpl(data);
};
var tplUtil = {
    each: each
};
tplUtil.rpad = function(text, len) {
    return rpad(text, len, ' ');
};
each(['yellow', 'green', 'cyan', 'red', 'white', 'magenta'], function(color) {
    tplUtil[color] = function(text) {
        return ansiColor[color](text);
    };
});
var cmdTpl = template(
    [
        'Usage:',
        '',
        "<% util.each(usage, function (value) { %>  <%=util.cyan(name)%> <%=value%><%='\\n'%><% }); %>",
        '<% if (options) { %>Options:',
        '',
        "<%     util.each(options, function (option) { %>  <%=util.yellow(util.rpad(option.name, maxNameWidth))%> <%=option.desc%><%='\\n'%><% }); %>",
        '<% } %>Description:',
        '',
        '  <%=desc%>'
    ].join('\n'),
    tplUtil
);
var helpTpl = template(
    [
        'Usage:',
        '',
        "<% util.each(usage, function (value) { %>  <%=util.cyan(name)%> <%=value%><%='\\n'%><% }); %>",
        'Commands:',
        '',
        "<% util.each(commands, function (command) { %>  <%=util.yellow(util.rpad(command.name, maxNameWidth))%> <%=command.desc%><%='\\n'%><% }); %>",
        "Run '<%=util.cyan(name + ' help <command>')%>' for more information on a specific command"
    ].join('\n'),
    tplUtil
);

module.exports = exports;

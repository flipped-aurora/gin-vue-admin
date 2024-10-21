var each = require('./each');
var strWidth = require('./strWidth');
var map = require('./map');
var repeat = require('./repeat');
var cloneDeep = require('./cloneDeep');
exports = function(rows) {
    rows = cloneDeep(rows);
    var options = {
        border: defBorder
    };
    options.columns = getColumns(rows);
    padData(rows, options);
    return render(rows, options);
};
function padData(rows, options) {
    var columnCount = options.columns.length;
    for (var i = 0, len = rows.length; i < len; i++) {
        while (rows[i].length < columnCount) {
            rows[i].push('');
        }
    }
    return loopData(rows, function(data, row, column) {
        var _options$columns$colu = options.columns[column],
            paddingLeft = _options$columns$colu.paddingLeft,
            width = _options$columns$colu.width,
            paddingRight = _options$columns$colu.paddingRight;
        return (
            repeat(' ', paddingLeft) +
            data +
            repeat(' ', width - strWidth(data) - paddingRight)
        );
    });
}
function loopData(rows, handler) {
    for (var i = 0, len = rows.length; i < len; i++) {
        var row = rows[i];
        for (var j = 0, _len = row.length; j < _len; j++) {
            var data = handler(row[j], i, j);
            if (data) {
                row[j] = data;
            }
        }
    }
}
function getColumns(rows) {
    var columns = [];
    var paddingLeft = 1;
    var paddingRight = 1;
    loopData(rows, function(data, row, column) {
        columns[column] = columns[column] || {
            width: paddingLeft + paddingRight,
            paddingLeft: paddingLeft,
            paddingRight: paddingRight
        };
        var width = strWidth(data) + paddingLeft + paddingRight;
        if (width > columns[column].width) {
            columns[column].width = width;
        }
    });
    return columns;
}
function render(rows, options) {
    var ret = '';
    ret += renderBorder('top', options);
    each(rows, function(row, idx) {
        ret += renderRow(row, options);
        if (idx === rows.length - 1) {
            ret += renderBorder('bottom', options);
        } else {
            ret += renderBorder('join', options);
        }
    });
    return ret;
}
function renderRow(columns, options) {
    var border = options.border;
    return (
        border.bodyLeft +
        columns.join(border.bodyJoin) +
        border.bodyRight +
        '\n'
    );
}
function renderBorder(type, options) {
    var border = options.border,
        columns = options.columns;
    var left = border[type + 'Left'];
    var right = border[type + 'Right'];
    var body = border[type + 'Body'];
    var join = border[type + 'Join'];
    var ret = map(columns, function(column) {
        return repeat(body, column.width);
    }).join(join);
    ret = left + ret + right;
    if (type !== 'bottom') {
        ret += '\n';
    }
    return ret;
}
var defBorder = {
    topBody: '─',
    topJoin: '┬',
    topLeft: '┌',
    topRight: '┐',
    bottomBody: '─',
    bottomJoin: '┴',
    bottomLeft: '└',
    bottomRight: '┘',
    bodyLeft: '│',
    bodyRight: '│',
    bodyJoin: '│',
    joinBody: '─',
    joinLeft: '├',
    joinRight: '┤',
    joinJoin: '┼'
};

module.exports = exports;

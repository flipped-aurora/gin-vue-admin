var selector = require('./selector');
var each = require('./each');
var startWith = require('./startWith');
var contain = require('./contain');
var cmpVersion = require('./cmpVersion');

exports = function(sel) {
    var _ref =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {},
        _ref$important = _ref.important,
        important = _ref$important === void 0 ? false : _ref$important,
        _ref$inlineStyle = _ref.inlineStyle,
        inlineStyle = _ref$inlineStyle === void 0 ? false : _ref$inlineStyle,
        _ref$position = _ref.position,
        position = _ref$position === void 0 ? 0 : _ref$position;
    var ret = [0, 0, 0, 0, 0, position];
    if (important) ret[0] = 1;
    if (inlineStyle) ret[1] = 1;
    var group = selector.parse(sel)[0];
    each(group, function(_ref2) {
        var type = _ref2.type,
            value = _ref2.value;
        switch (type) {
            case 'id':
                ret[2]++;
                break;
            case 'class':
            case 'attribute':
                ret[3]++;
                break;
            case 'pseudo':
                if (contain(PSEUDO_ELEMS, value.replace(/:/g, ''))) {
                    ret[4]++;
                } else if (!startWith(value, '::')) {
                    ret[3]++;
                }
                break;
            case 'tag':
                if (value !== '*') {
                    ret[4]++;
                }
                break;
        }
    });
    return ret;
};
var PSEUDO_ELEMS = [
    'first-letter',
    'last-letter',
    'first-line',
    'last-line',
    'first-child',
    'last-child',
    'before',
    'after'
];
exports.compare = function(p1, p2) {
    return cmpVersion(p1.join('.'), p2.join('.'));
};

module.exports = exports;

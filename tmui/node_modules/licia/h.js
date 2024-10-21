var isEl = require('./isEl');
var isStr = require('./isStr');
var startWith = require('./startWith');
var $class = require('./$class');
var $css = require('./$css');
var each = require('./each');
var isFn = require('./isFn');
exports = function(tag, attrs) {
    for (
        var _len = arguments.length,
            children = new Array(_len > 2 ? _len - 2 : 0),
            _key = 2;
        _key < _len;
        _key++
    ) {
        children[_key - 2] = arguments[_key];
    }
    if (isEl(attrs) || isStr(attrs)) {
        children.unshift(attrs);
        attrs = null;
    }
    if (!attrs) attrs = {};
    var _parseTag = parseTag(tag),
        tagName = _parseTag.tagName,
        id = _parseTag.id,
        classes = _parseTag.classes;
    var el = document.createElement(tagName);
    if (id) el.setAttribute('id', id);
    $class.add(el, classes);
    each(children, function(child) {
        if (isStr(child)) {
            el.appendChild(document.createTextNode(child));
        } else if (isEl(child)) {
            el.appendChild(child);
        }
    });
    each(attrs, function(val, key) {
        if (isStr(val)) {
            el.setAttribute(key, val);
        } else if (isFn(val) && startWith(key, 'on')) {
            el.addEventListener(key.slice(2), val, false);
        } else if (key === 'style') {
            $css(el, val);
        }
    });
    return el;
};
function parseTag(tag) {
    var tagName = 'div';
    var id = '';
    var classes = [];
    var words = [];
    var word = '';
    for (var i = 0, len = tag.length; i < len; i++) {
        var c = tag[i];
        if (c === '#' || c === '.') {
            words.push(word);
            word = c;
        } else {
            word += c;
        }
    }
    words.push(word);
    for (var _i = 0, _len2 = words.length; _i < _len2; _i++) {
        word = words[_i];
        if (!word) continue;
        if (startWith(word, '#')) {
            id = word.slice(1);
        } else if (startWith(word, '.')) {
            classes.push(word.slice(1));
        } else {
            tagName = word;
        }
    }
    return {
        tagName: tagName,
        id: id,
        classes: classes
    };
}

module.exports = exports;

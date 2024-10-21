var Select = require('./Select');
var $offset = require('./$offset');
var $show = require('./$show');
var $css = require('./$css');
var $attr = require('./$attr');
var $property = require('./$property');
var last = require('./last');
var $remove = require('./$remove');
var $data = require('./$data');
var $event = require('./$event');
var $class = require('./$class');
var $insert = require('./$insert');
var isUndef = require('./isUndef');
var isStr = require('./isStr');
exports = function(selector) {
    return new Select(selector);
};
Select.methods({
    offset: function() {
        return $offset(this);
    },
    hide: function() {
        return this.css('display', 'none');
    },
    show: function() {
        $show(this);
        return this;
    },
    first: function() {
        return exports(this[0]);
    },
    last: function() {
        return exports(last(this));
    },
    get: function(idx) {
        return this[idx];
    },
    eq: function(idx) {
        return exports(this[idx]);
    },
    on: function(event, selector, handler) {
        $event.on(this, event, selector, handler);
        return this;
    },
    off: function(event, selector, handler) {
        $event.off(this, event, selector, handler);
        return this;
    },
    html: function(val) {
        var result = $property.html(this, val);
        if (isUndef(val)) return result;
        return this;
    },
    text: function(val) {
        var result = $property.text(this, val);
        if (isUndef(val)) return result;
        return this;
    },
    val: function(val) {
        var result = $property.val(this, val);
        if (isUndef(val)) return result;
        return this;
    },
    css: function(name, val) {
        var result = $css(this, name, val);
        if (isGetter(name, val)) return result;
        return this;
    },
    attr: function(name, val) {
        var result = $attr(this, name, val);
        if (isGetter(name, val)) return result;
        return this;
    },
    data: function(name, val) {
        var result = $data(this, name, val);
        if (isGetter(name, val)) return result;
        return this;
    },
    rmAttr: function(name) {
        $attr.remove(this, name);
        return this;
    },
    remove: function() {
        $remove(this);
        return this;
    },
    addClass: function(name) {
        $class.add(this, name);
        return this;
    },
    rmClass: function(name) {
        $class.remove(this, name);
        return this;
    },
    toggleClass: function(name) {
        $class.toggle(this, name);
        return this;
    },
    hasClass: function(name) {
        return $class.has(this, name);
    },
    parent: function() {
        return exports(this[0].parentNode);
    },
    append: function(val) {
        $insert.append(this, val);
        return this;
    },
    prepend: function(val) {
        $insert.prepend(this, val);
        return this;
    },
    before: function(val) {
        $insert.before(this, val);
        return this;
    },
    after: function(val) {
        $insert.after(this, val);
        return this;
    }
});
var isGetter = function(name, val) {
    return isUndef(val) && isStr(name);
};

module.exports = exports;

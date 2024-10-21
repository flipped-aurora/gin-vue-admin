var each = require('./each');
var isStr = require('./isStr');
var isUndef = require('./isUndef');
var contain = require('./contain');
var isArr = require('./isArr');
var isObj = require('./isObj');
var toArr = require('./toArr');
exports = function(name, content) {
    if (isUndef(name)) return getAllMeta();
    var isGetter = (isStr(name) && isUndef(content)) || isArr(name);
    if (isGetter) return getMeta(name);
    var metas = name;
    if (!isObj(metas)) {
        metas = {};
        metas[name] = content;
    }
    setMeta(metas);
};
exports.remove = function(nameList) {
    nameList = toArr(nameList);
    each(nameList, function(name) {
        var meta = selectMeta(name);
        if (meta) doc.head.removeChild(meta);
    });
};
var doc = document;
function getAllMeta() {
    var ret = {};
    metaEach(function(name, content) {
        ret[name] = content;
    });
    return ret;
}
function getMeta(name) {
    if (isStr(name)) {
        var meta = selectMeta(name);
        if (meta) return meta.getAttribute('content');
    } else {
        var ret = {};
        metaEach(function(key, val) {
            if (contain(name, key)) ret[key] = val;
        });
        return ret;
    }
}
function setMeta(metas) {
    each(metas, function(content, name) {
        var meta = selectMeta(name);
        if (meta) return meta.setAttribute('content', content);
        meta = doc.createElement('meta');
        meta.setAttribute('name', name);
        meta.setAttribute('content', content);
        doc.head.appendChild(meta);
    });
}
function metaEach(fn) {
    var metaList = doc.querySelectorAll('meta');
    each(metaList, function(meta) {
        var name = meta.getAttribute('name');
        var content = meta.getAttribute('content');
        if (!name || !content) return;
        fn(name, content);
    });
}
function selectMeta(name) {
    return doc.querySelector('meta[name="' + name + '"]');
}

module.exports = exports;

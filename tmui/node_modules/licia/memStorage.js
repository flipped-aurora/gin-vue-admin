var keys = require('./keys');
exports = {
    getItem: function(key) {
        return (API_KEYS[key] ? cloak[key] : this[key]) || null;
    },
    setItem: function(key, val) {
        API_KEYS[key] ? (cloak[key] = val) : (this[key] = val);
    },
    removeItem: function(key) {
        API_KEYS[key] ? delete cloak[key] : delete this[key];
    },
    key: function(i) {
        var keys = enumerableKeys();
        return i >= 0 && i < keys.length ? keys[i] : null;
    },
    clear: function() {
        var keys = uncloakedKeys();

        for (var i = 0, key; (key = keys[i]); i++) delete this[key];
        keys = cloakedKeys();

        for (var _i = 0, _key; (_key = keys[_i]); _i++) delete cloak[_key];
    }
};
Object.defineProperty(exports, 'length', {
    enumerable: false,
    configurable: true,
    get: function() {
        return enumerableKeys().length;
    }
});
var cloak = {};
var API_KEYS = {
    getItem: 1,
    setItem: 1,
    removeItem: 1,
    key: 1,
    clear: 1,
    length: 1
};
function enumerableKeys() {
    return uncloakedKeys().concat(cloakedKeys());
}
function uncloakedKeys() {
    return keys(exports).filter(function(key) {
        return !API_KEYS[key];
    });
}
function cloakedKeys() {
    return keys(cloak);
}

module.exports = exports;

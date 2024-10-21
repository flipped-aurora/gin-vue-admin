var Class = require('./Class');
var safeGet = require('./safeGet');
var extend = require('./extend');
var strTpl = require('./strTpl');
var isStr = require('./isStr');
var isFn = require('./isFn');
exports = Class({
    initialize: function I18n(locale, langs) {
        this._locale = locale;
        this._langs = langs;
    },
    set: function(locale, lang) {
        if (this._langs[locale]) {
            extend(this._langs[locale], lang);
        } else {
            this._langs[locale] = lang;
        }
    },
    t: function(path, data) {
        var val = '';
        var lang = this._langs[this._locale];
        if (!lang) return '';
        val = safeGet(lang, path);
        if (data) {
            if (isStr(val)) {
                val = strTpl(val, data);
            } else if (isFn(val)) {
                val = val(data);
            }
        }
        return val || '';
    },
    locale: function(locale) {
        this._locale = locale;
    }
});

module.exports = exports;

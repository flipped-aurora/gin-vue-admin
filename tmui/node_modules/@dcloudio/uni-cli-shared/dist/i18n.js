"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.resolveI18nLocale = exports.initLocales = exports.getLocaleFiles = exports.isUniAppLocaleFile = exports.initI18nOptionsOnce = exports.initI18nOptions = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const fast_glob_1 = require("fast-glob");
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const json_1 = require("./json");
const messages_1 = require("./messages");
function initI18nOptions(platform, inputDir, warning = false, withMessages = true) {
    const locales = initLocales(path_1.default.resolve(inputDir, 'locale'), withMessages);
    if (!Object.keys(locales).length) {
        return;
    }
    const manifestJson = (0, json_1.parseManifestJsonOnce)(inputDir);
    let fallbackLocale = manifestJson.fallbackLocale || 'en';
    const locale = resolveI18nLocale(platform, Object.keys(locales), fallbackLocale);
    if (warning) {
        if (!fallbackLocale) {
            console.warn(messages_1.M['i18n.fallbackLocale.default'].replace('{locale}', locale));
        }
        else if (locale !== fallbackLocale) {
            console.warn(messages_1.M['i18n.fallbackLocale.missing'].replace('{locale}', fallbackLocale));
        }
    }
    return {
        locale,
        locales,
        delimiters: uni_shared_1.I18N_JSON_DELIMITERS,
    };
}
exports.initI18nOptions = initI18nOptions;
exports.initI18nOptionsOnce = (0, uni_shared_1.once)(initI18nOptions);
const localeJsonRE = /uni-app.*.json/;
function isUniAppLocaleFile(filepath) {
    if (!filepath) {
        return false;
    }
    return localeJsonRE.test(path_1.default.basename(filepath));
}
exports.isUniAppLocaleFile = isUniAppLocaleFile;
function parseLocaleJson(filepath) {
    let jsonObj = (0, json_1.parseJson)(fs_1.default.readFileSync(filepath, 'utf8'));
    if (isUniAppLocaleFile(filepath)) {
        jsonObj = jsonObj.common || {};
    }
    return jsonObj;
}
function getLocaleFiles(cwd) {
    return (0, fast_glob_1.sync)('*.json', { cwd, absolute: true });
}
exports.getLocaleFiles = getLocaleFiles;
function initLocales(dir, withMessages = true) {
    if (!fs_1.default.existsSync(dir)) {
        return {};
    }
    return fs_1.default.readdirSync(dir).reduce((res, filename) => {
        if (path_1.default.extname(filename) === '.json') {
            try {
                const locale = path_1.default
                    .basename(filename)
                    .replace(/(uni-app.)?(.*).json/, '$2');
                if (withMessages) {
                    (0, shared_1.extend)(res[locale] || (res[locale] = {}), parseLocaleJson(path_1.default.join(dir, filename)));
                }
                else {
                    res[locale] = {};
                }
            }
            catch (e) { }
        }
        return res;
    }, {});
}
exports.initLocales = initLocales;
function resolveI18nLocale(platform, locales, locale) {
    if (locale && locales.includes(locale)) {
        return locale;
    }
    const defaultLocales = ['zh-Hans', 'zh-Hant'];
    if (platform === 'app' || platform === 'h5') {
        defaultLocales.unshift('en');
    }
    else {
        // 小程序
        defaultLocales.push('en');
    }
    return defaultLocales.find((locale) => locales.includes(locale)) || locales[0];
}
exports.resolveI18nLocale = resolveI18nLocale;

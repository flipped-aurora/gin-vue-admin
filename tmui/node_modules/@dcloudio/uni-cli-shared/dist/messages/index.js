"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.M = void 0;
const os_locale_s_fix_1 = require("os-locale-s-fix");
const en_1 = __importDefault(require("./en"));
const zh_CN_1 = __importDefault(require("./zh_CN"));
function format(lang) {
    const array = lang.split(/[.,]/)[0].split(/[_-]/);
    array[0] = array[0].toLowerCase();
    if (array[0] === 'zh') {
        array[1] = (array[1] || 'CN').toUpperCase();
    }
    array.length = Math.min(array.length, 2);
    return array.join('_');
}
const locale = format(process.env.UNI_HBUILDERX_LANGID ||
    os_locale_s_fix_1.osLocale.sync({ spawn: true, cache: false }) ||
    'en');
exports.M = locale === 'zh_CN' ? zh_CN_1.default : en_1.default;

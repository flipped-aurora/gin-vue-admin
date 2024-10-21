"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initI18n = void 0;
const uni_i18n_1 = require("@dcloudio/uni-i18n");
const i18n_1 = require("../../../i18n");
function initI18n(manifestJson) {
    const i18nOptions = (0, i18n_1.initI18nOptions)(process.env.UNI_PLATFORM, process.env.UNI_INPUT_DIR, true);
    if (i18nOptions) {
        manifestJson = JSON.parse((0, uni_i18n_1.compileI18nJsonStr)(JSON.stringify(manifestJson), i18nOptions));
        manifestJson.fallbackLocale = i18nOptions.locale;
    }
    return manifestJson;
}
exports.initI18n = initI18n;

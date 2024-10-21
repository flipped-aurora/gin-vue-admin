"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformScopedCss = void 0;
const tags_1 = require("./tags");
const logs_1 = require("../logs");
function transformScopedCss(cssCode) {
    checkHtmlTagSelector(cssCode);
    return cssCode.replace(/\[(data-v-[a-f0-9]{8})\]/gi, (_, scopedId) => {
        return '.' + scopedId;
    });
}
exports.transformScopedCss = transformScopedCss;
function checkHtmlTagSelector(cssCode) {
    for (const tag in tags_1.HTML_TO_MINI_PROGRAM_TAGS) {
        if (new RegExp(`( |\n|\t|,|})${tag}( *)(,|{)`, 'g').test(cssCode)) {
            (0, logs_1.output)('warn', `小程序端 style 暂不支持 ${tag} 标签选择器，推荐使用 class 选择器，详情参考：https://uniapp.dcloud.net.cn/tutorial/migration-to-vue3.html#style`);
            break;
        }
    }
}

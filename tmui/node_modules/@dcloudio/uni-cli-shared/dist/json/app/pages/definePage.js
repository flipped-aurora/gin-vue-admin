"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineNVuePageCode = exports.definePageCode = void 0;
const utils_1 = require("../../../utils");
function definePageCode(pagesJson, platform = 'app') {
    const importPagesCode = [];
    const definePagesCode = [];
    pagesJson.pages.forEach((page) => {
        if (platform === 'app' && page.style.isNVue) {
            return;
        }
        const pagePath = page.path;
        const pageIdentifier = (0, utils_1.normalizeIdentifier)(pagePath);
        const pagePathWithExtname = (0, utils_1.normalizePagePath)(pagePath, platform);
        if (pagePathWithExtname) {
            if (process.env.UNI_APP_CODE_SPLITING) {
                // 拆分页面
                importPagesCode.push(`const ${pageIdentifier} = ()=>import('./${pagePathWithExtname}')`);
            }
            else {
                importPagesCode.push(`import ${pageIdentifier} from './${pagePathWithExtname}'`);
            }
            definePagesCode.push(`__definePage('${pagePath}',${pageIdentifier})`);
        }
    });
    return importPagesCode.join('\n') + '\n' + definePagesCode.join('\n');
}
exports.definePageCode = definePageCode;
function defineNVuePageCode(pagesJson) {
    const importNVuePagesCode = [];
    pagesJson.pages.forEach((page) => {
        if (!page.style.isNVue) {
            return;
        }
        const pagePathWithExtname = (0, utils_1.normalizePagePath)(page.path, 'app');
        if (pagePathWithExtname) {
            importNVuePagesCode.push(`import('./${pagePathWithExtname}').then((res)=>{res.length})`);
        }
    });
    return importNVuePagesCode.join('\n');
}
exports.defineNVuePageCode = defineNVuePageCode;

"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppPagesPlugin = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../utils");
const utils_2 = require("../android/utils");
function uniAppPagesPlugin() {
    const pagesJsonPath = path_1.default.resolve(process.env.UNI_INPUT_DIR, 'pages.json');
    const pagesJsonUTSPath = path_1.default.resolve(process.env.UNI_INPUT_DIR, uni_cli_shared_1.PAGES_JSON_UTS);
    let allPagePaths = [];
    let isFirst = true;
    return {
        name: 'uni:app-pages',
        apply: 'build',
        resolveId(id) {
            if ((0, utils_1.isPages)(id)) {
                return pagesJsonUTSPath;
            }
        },
        load(id) {
            if ((0, utils_1.isPages)(id)) {
                return fs_extra_1.default.readFileSync(pagesJsonPath, 'utf8');
            }
        },
        transform(code, id) {
            if (isFirst && allPagePaths.length) {
                const { filename } = (0, uni_cli_shared_1.parseVueRequest)(id);
                if ((0, utils_2.isVue)(filename)) {
                    const vueFilename = (0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizePath)(path_1.default.relative(process.env.UNI_INPUT_DIR, filename)));
                    // 项目内的
                    if (!vueFilename.startsWith('.')) {
                        // const index = allPagePaths.indexOf(pagePath)
                        // if (index > -1) {
                        if ((0, uni_cli_shared_1.runByHBuilderX)()) {
                            console.log(`当前工程${allPagePaths.length}个页面，正在编译${vueFilename}...${'\u200b'}`);
                        }
                        // }
                    }
                }
            }
            if ((0, utils_1.isPages)(id)) {
                this.addWatchFile(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'pages.json'));
                // dark mode
                this.addWatchFile(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'theme.json'));
                // pages.json
                const pagesJson = (0, uni_cli_shared_1.normalizeUniAppXAppPagesJson)(code);
                // add themeConfig - can move to uni-x/index.ts
                pagesJson.themeConfig = readThemeJSONFile();
                (0, utils_1.setGlobalPageOrientation)(pagesJson.globalStyle?.pageOrientation || '');
                allPagePaths = pagesJson.pages.map((p) => p.path);
                this.emitFile({
                    fileName: uni_cli_shared_1.APP_CONFIG,
                    type: 'asset',
                    // 生成 app-config.js
                    source: (0, uni_cli_shared_1.normalizeUniAppXAppConfig)(pagesJson, (0, uni_cli_shared_1.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR)),
                });
                return {
                    code: (0, uni_cli_shared_1.normalizeAppPagesJson)(pagesJson),
                    map: { mappings: '' },
                };
            }
        },
        buildEnd() {
            isFirst = false;
        },
    };
}
exports.uniAppPagesPlugin = uniAppPagesPlugin;
function readThemeJSONFile() {
    try {
        // 后续读取 theme location
        const themeJsonPath = path_1.default.resolve(process.env.UNI_INPUT_DIR, 'theme.json');
        let content = '{}';
        if (fs_extra_1.default.existsSync(themeJsonPath)) {
            content = fs_extra_1.default.readFileSync(themeJsonPath, 'utf8');
        }
        return JSON.parse(content);
    }
    catch (error) {
        console.error('read theme.json error:', error);
    }
}

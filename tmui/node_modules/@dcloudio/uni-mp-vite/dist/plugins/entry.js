"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniEntryPlugin = exports.isUniComponentUrl = exports.isUniPageUrl = exports.parseVirtualComponentPath = exports.parseVirtualPagePath = exports.virtualComponentPath = exports.virtualPagePath = void 0;
const path_1 = __importDefault(require("path"));
const fs_1 = __importDefault(require("fs"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uniPagePrefix = 'uniPage://';
const uniComponentPrefix = 'uniComponent://';
function virtualPagePath(filepath) {
    return uniPagePrefix + (0, uni_cli_shared_1.encodeBase64Url)(filepath);
}
exports.virtualPagePath = virtualPagePath;
function virtualComponentPath(filepath) {
    return uniComponentPrefix + (0, uni_cli_shared_1.encodeBase64Url)(filepath);
}
exports.virtualComponentPath = virtualComponentPath;
function parseVirtualPagePath(uniPageUrl) {
    return (0, uni_cli_shared_1.decodeBase64Url)(uniPageUrl.replace(uniPagePrefix, ''));
}
exports.parseVirtualPagePath = parseVirtualPagePath;
function parseVirtualComponentPath(uniComponentUrl) {
    return (0, uni_cli_shared_1.decodeBase64Url)(uniComponentUrl.replace(uniComponentPrefix, ''));
}
exports.parseVirtualComponentPath = parseVirtualComponentPath;
function isUniPageUrl(id) {
    return id.startsWith(uniPagePrefix);
}
exports.isUniPageUrl = isUniPageUrl;
function isUniComponentUrl(id) {
    return id.startsWith(uniComponentPrefix);
}
exports.isUniComponentUrl = isUniComponentUrl;
const styleIsolationRE = /export\s+default\s+[\s\S]*?styleIsolation\s*:\s*['|"](isolated|apply-shared|shared)['|"]/;
function parseComponentStyleIsolation(file) {
    const content = fs_1.default.readFileSync(file, 'utf-8');
    const matches = content.match(styleIsolationRE);
    if (matches) {
        return matches[1];
    }
}
function uniEntryPlugin({ global, }) {
    const inputDir = process.env.UNI_INPUT_DIR;
    const manifestJson = (0, uni_cli_shared_1.parseManifestJsonOnce)(inputDir);
    const platformOptions = manifestJson[process.env.UNI_PLATFORM] || {};
    return {
        name: 'uni:virtual',
        enforce: 'pre',
        resolveId(id) {
            if (isUniPageUrl(id) || isUniComponentUrl(id)) {
                return id;
            }
        },
        load(id) {
            if (isUniPageUrl(id)) {
                const filepath = (0, uni_cli_shared_1.normalizePath)(path_1.default.resolve(inputDir, parseVirtualPagePath(id)));
                this.addWatchFile(filepath);
                return {
                    code: `import MiniProgramPage from '${filepath}'
${global}.createPage(MiniProgramPage)`,
                };
            }
            else if (isUniComponentUrl(id)) {
                const filepath = (0, uni_cli_shared_1.normalizePath)(path_1.default.resolve(inputDir, parseVirtualComponentPath(id)));
                this.addWatchFile(filepath);
                const json = {
                    component: true,
                    styleIsolation: undefined,
                };
                if (process.env.UNI_PLATFORM === 'mp-alipay') {
                    json.styleIsolation =
                        parseComponentStyleIsolation(filepath) ||
                            platformOptions.styleIsolation ||
                            'apply-shared';
                }
                // 微信小程序json文件中的styleIsolation优先级比options中的高，为了兼容旧版本，不能设置默认值，并且只有在manifest.json中配置styleIsolation才会静态分析组件的styleIsolation
                if (process.env.UNI_PLATFORM === 'mp-weixin') {
                    if (platformOptions.styleIsolation) {
                        json.styleIsolation =
                            parseComponentStyleIsolation(filepath) ||
                                platformOptions.styleIsolation;
                    }
                }
                (0, uni_cli_shared_1.addMiniProgramComponentJson)((0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizeMiniProgramFilename)(filepath, inputDir)), json);
                return {
                    code: `import Component from '${filepath}'
${global}.createComponent(Component)`,
                };
            }
        },
    };
}
exports.uniEntryPlugin = uniEntryPlugin;

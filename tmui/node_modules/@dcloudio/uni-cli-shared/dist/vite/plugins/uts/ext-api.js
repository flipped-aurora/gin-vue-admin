"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.injectsToAutoImports = exports.uniUTSExtApiReplace = exports.parseUniExtApisOnce = void 0;
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const vite_1 = __importDefault(require("unplugin-auto-import/vite"));
const uni_modules_1 = require("../../../uni_modules");
const url_1 = require("../../utils/url");
const escape = (str) => str.replace(/[-[\]/{}()*+?.\\^$|]/g, '\\$&');
exports.parseUniExtApisOnce = (0, uni_shared_1.once)(uni_modules_1.parseUniExtApis);
function uniUTSExtApiReplace() {
    const injects = (0, exports.parseUniExtApisOnce)(true, process.env.UNI_UTS_PLATFORM, process.env.UNI_UTS_TARGET_LANGUAGE);
    const injectApis = Object.keys(injects);
    const firstPass = new RegExp(`(?:${injectApis.map(escape).join('|')})`, 'g');
    return {
        name: 'uni:uts-ext-api-replace',
        configResolved(config) {
            const index = config.plugins.findIndex((p) => p.name === 'uts');
            if (index > -1) {
                if (Object.keys(injects).length) {
                    // @ts-expect-error
                    config.plugins.splice(index, 0, (0, vite_1.default)({
                        include: [/\.[u]?ts$/, /\.[u]?vue/],
                        exclude: [/[\\/]\.git[\\/]/],
                        imports: injectsToAutoImports(injects),
                        dts: false,
                    }));
                }
            }
        },
        transform(code, id) {
            if (!injectApis.length) {
                return;
            }
            if (!(0, url_1.isJsFile)(id)) {
                return;
            }
            if (code.search(firstPass) === -1) {
                return;
            }
            injectApis.forEach((api) => {
                code = code.replaceAll(api, api.replace('.', '_'));
            });
            return {
                code,
                map: { mappings: '' },
            };
        },
    };
}
exports.uniUTSExtApiReplace = uniUTSExtApiReplace;
/**
 * { 'uni.getBatteryInfo': ['@/uni_modules/uni-getbatteryinfo/utssdk/web/index.uts','getBatteryInfo'] }
 * { '@/uni_modules/uni-getbatteryinfo/utssdk/web/index.ts': [['getBatteryInfo', 'uni_getBatteryInfo']] }
 * @param injects
 */
function injectsToAutoImports(injects) {
    const autoImports = {};
    Object.keys(injects).forEach((api) => {
        const options = injects[api];
        if ((0, shared_1.isArray)(options) && options.length >= 2) {
            const source = options[0];
            const name = options[1];
            if (!autoImports[source]) {
                autoImports[source] = [];
            }
            autoImports[source].push([name, api.replace('.', '_')]);
        }
    });
    return Object.keys(autoImports).map((source) => {
        return {
            from: source,
            imports: autoImports[source],
        };
    });
}
exports.injectsToAutoImports = injectsToAutoImports;

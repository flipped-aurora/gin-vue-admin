"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniRuntimeHooksPlugin = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniRuntimeHooksPlugin() {
    return {
        name: 'uni:mp-runtime-hooks',
        enforce: 'post',
        async transform(source, id) {
            const isSetupJs = (0, uni_cli_shared_1.isUniPageSfcFile)(id);
            const isSetupTs = !isSetupJs && (0, uni_cli_shared_1.isUniPageSetupAndTs)(id);
            if (!isSetupJs && !isSetupTs) {
                return null;
            }
            if (isSetupJs && !source.includes('_sfc_main')) {
                return null;
            }
            if (isSetupTs && !source.includes('defineComponent')) {
                return null;
            }
            const matches = source.match(new RegExp(`(${Object.keys(uni_shared_1.MINI_PROGRAM_PAGE_RUNTIME_HOOKS).join('|')})`, 'g'));
            if (!matches) {
                return null;
            }
            if (matches.includes('onShareTimeline')) {
                matches.push('onShareAppMessage');
            }
            const hooks = new Set(matches);
            let flag = 0;
            for (const hook of hooks) {
                flag |= uni_shared_1.MINI_PROGRAM_PAGE_RUNTIME_HOOKS[hook];
            }
            if (isSetupJs) {
                source = source + `;_sfc_main.__runtimeHooks = ${flag};`;
            }
            else if (isSetupTs) {
                source =
                    require('@vue/compiler-sfc').rewriteDefault(source, '_sfc_defineComponent') +
                        `\n_sfc_defineComponent.__runtimeHooks = ${flag};\nexport default _sfc_defineComponent`;
            }
            return {
                code: source,
                map: { mappings: '' },
            };
        },
    };
}
exports.uniRuntimeHooksPlugin = uniRuntimeHooksPlugin;

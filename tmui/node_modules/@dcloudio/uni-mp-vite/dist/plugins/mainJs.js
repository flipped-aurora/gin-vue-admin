"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniMainJsPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const usingComponents_1 = require("./usingComponents");
function uniMainJsPlugin(options = {}) {
    const normalizeComponentName = options.normalizeComponentName || ((name) => name);
    return (0, uni_cli_shared_1.defineUniMainJsPlugin)((opts) => {
        return {
            name: 'uni:mp-main-js',
            enforce: 'pre',
            async transform(source, id) {
                if (opts.filter(id)) {
                    source = source.includes('createSSRApp')
                        ? createApp(source)
                        : createLegacyApp(source);
                    const inputDir = process.env.UNI_INPUT_DIR;
                    const { imports } = await (0, uni_cli_shared_1.updateMiniProgramGlobalComponents)(id, (0, uni_cli_shared_1.parseProgram)(source, id, {
                        babelParserPlugins: options.babelParserPlugins,
                    }), {
                        inputDir,
                        resolve: this.resolve,
                        normalizeComponentName,
                    });
                    const { code, map } = await (0, uni_cli_shared_1.transformDynamicImports)(source, imports, {
                        id,
                        sourceMap: (0, uni_cli_shared_1.withSourcemap)(opts.resolvedConfig),
                        dynamicImport: usingComponents_1.dynamicImport,
                    });
                    return {
                        code: `import '\0plugin-vue:export-helper';import 'uni-mp-runtime';import './${uni_cli_shared_1.PAGES_JSON_JS}';` +
                            code,
                        map,
                    };
                }
            },
        };
    });
}
exports.uniMainJsPlugin = uniMainJsPlugin;
function createApp(code) {
    return `${code};createApp().app.mount("#app");`;
}
function createLegacyApp(code) {
    return code;
}

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppIOSMainPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniAppIOSMainPlugin() {
    const mainUTS = (0, uni_cli_shared_1.resolveMainPathOnce)(process.env.UNI_INPUT_DIR);
    return {
        name: 'uni:app-main',
        apply: 'build',
        async transform(code, id) {
            if ((0, uni_cli_shared_1.normalizePath)(id) === mainUTS) {
                return {
                    code: `
          import './${uni_cli_shared_1.MANIFEST_JSON_UTS}'
          import './${uni_cli_shared_1.PAGES_JSON_UTS}'
          const __global__ = typeof globalThis === 'undefined' ? Function('return this')() : globalThis
          __global__.__uniX = true
          ${code}
          createApp().app.mount("#app");
          `,
                    map: null,
                };
            }
        },
    };
}
exports.uniAppIOSMainPlugin = uniAppIOSMainPlugin;

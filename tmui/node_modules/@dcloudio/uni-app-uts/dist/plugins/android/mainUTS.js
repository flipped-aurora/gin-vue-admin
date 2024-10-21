"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppMainPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("./utils");
function uniAppMainPlugin() {
    const mainUTS = (0, uni_cli_shared_1.resolveMainPathOnce)(process.env.UNI_INPUT_DIR);
    return {
        name: 'uni:app-main',
        apply: 'build',
        async transform(code, id) {
            if ((0, uni_cli_shared_1.normalizePath)(id) === mainUTS) {
                code = await (0, utils_1.parseImports)(code, (0, utils_1.createTryResolve)(id, this.resolve.bind(this)));
                return {
                    code: `import './${uni_cli_shared_1.MANIFEST_JSON_UTS}'
          import './${uni_cli_shared_1.PAGES_JSON_UTS}'
          ${code}
          `,
                    map: null,
                };
            }
        },
    };
}
exports.uniAppMainPlugin = uniAppMainPlugin;

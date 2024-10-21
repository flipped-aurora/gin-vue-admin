"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniRenderjsPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniRenderjsPlugin() {
    return {
        name: 'uni:app-nvue-renderjs',
        async transform(code, id) {
            const { type } = (0, uni_cli_shared_1.parseRenderjs)(id);
            if (!type) {
                return;
            }
            return {
                code: `export default {}`,
                map: { mappings: '' },
            };
        },
    };
}
exports.uniRenderjsPlugin = uniRenderjsPlugin;

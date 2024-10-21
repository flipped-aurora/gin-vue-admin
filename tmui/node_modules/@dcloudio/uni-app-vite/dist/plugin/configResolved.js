"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createConfigResolved = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function createConfigResolved({ createCssPostPlugin, }) {
    return (config) => {
        (0, uni_cli_shared_1.injectCssPlugin)(config, process.env.UNI_COMPILE_TARGET === 'uni_modules'
            ? {
                createUrlReplacer: uni_cli_shared_1.createEncryptCssUrlReplacer,
            }
            : {});
        (0, uni_cli_shared_1.injectCssPostPlugin)(config, createCssPostPlugin(config));
        // 强制不inline
        config.build.assetsInlineLimit = 0;
        (0, uni_cli_shared_1.injectAssetPlugin)(config);
    };
}
exports.createConfigResolved = createConfigResolved;

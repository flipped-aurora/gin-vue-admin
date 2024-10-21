"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createUni = exports.compilerOptions = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
function realIsH5CustomElement(tag) {
    return (0, uni_shared_1.isH5CustomElement)(tag, process.env.UNI_APP_X === 'true');
}
exports.compilerOptions = {
    isNativeTag: uni_shared_1.isH5NativeTag,
    isCustomElement: realIsH5CustomElement,
    nodeTransforms: [
        uni_cli_shared_1.transformRefresherSlot,
        uni_cli_shared_1.transformH5BuiltInComponents,
        uni_cli_shared_1.transformTapToClick,
        uni_cli_shared_1.transformMatchMedia,
        uni_cli_shared_1.transformPageHead,
    ],
};
function createUni() {
    return {
        copyOptions: {
            assets: ['hybrid/html/**/*', 'uni_modules/*/hybrid/html/**/*'],
        },
        compilerOptions: exports.compilerOptions,
        jsxOptions: {
            babelPlugins: [uni_cli_shared_1.transformUniH5Jsx],
        },
    };
}
exports.createUni = createUni;

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createDefine = void 0;
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function createDefine(_) {
    return (0, shared_1.extend)({
        __VUE_I18N_FULL_INSTALL__: true,
        __VUE_I18N_LEGACY_API__: true,
        __VUE_I18N_PROD_DEVTOOLS__: false,
        __INTLIFY_PROD_DEVTOOLS__: false,
    }, (0, uni_cli_shared_1.initDefine)());
}
exports.createDefine = createDefine;

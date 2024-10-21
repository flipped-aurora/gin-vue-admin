"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformVue = exports.genClassName = void 0;
const plugins_1 = require("./plugins");
exports.default = () => {
    return process.env.UNI_UTS_PLATFORM === 'app-android'
        ? (0, plugins_1.initAndroid)()
        : (0, plugins_1.initIOS)();
};
var uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
Object.defineProperty(exports, "genClassName", { enumerable: true, get: function () { return uni_cli_shared_1.genUTSClassName; } });
var main_1 = require("./plugins/android/uvue/sfc/main");
Object.defineProperty(exports, "transformVue", { enumerable: true, get: function () { return main_1.transformMain; } });

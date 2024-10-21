"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initEnv = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function initEnv(config) {
    if (!process.env.UNI_PLATFORM) {
        process.env.UNI_PLATFORM = 'h5';
    }
    if (!process.env.UNI_CLI_CONTEXT) {
        process.env.UNI_CLI_CONTEXT = process.cwd();
    }
    if (!process.env.UNI_INPUT_DIR) {
        process.env.UNI_INPUT_DIR = path_1.default.resolve(config.root, 'src');
    }
    if (!process.env.UNI_OUTPUT_DIR) {
        process.env.UNI_OUTPUT_DIR = path_1.default.resolve(config.root, config.build.outDir);
    }
    process.env.VUE_APP_DARK_MODE =
        (0, uni_cli_shared_1.getPlatformManifestJsonOnce)().darkmode || false;
    process.env.BROWSERSLIST_CONFIG = [
        path_1.default.resolve(process.env.UNI_INPUT_DIR, '.browserslistrc'),
        path_1.default.resolve(process.env.UNI_CLI_CONTEXT, 'package.json'),
        path_1.default.resolve(process.cwd(), 'package.json'),
    ].find((file) => fs_1.default.existsSync(file));
}
exports.initEnv = initEnv;

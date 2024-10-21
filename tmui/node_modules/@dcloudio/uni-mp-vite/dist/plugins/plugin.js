"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniMiniProgramPluginPlugin = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const subpackage_1 = require("./subpackage");
const shared_1 = require("@vue/shared");
const build_1 = require("../plugin/build");
function uniMiniProgramPluginPlugin({ style: { extname }, }) {
    const entry = initPluginEntry();
    const rollupOptions = {};
    if (entry) {
        rollupOptions.input = (0, shared_1.extend)({
            app: (0, uni_cli_shared_1.resolveMainPathOnce)(process.env.UNI_INPUT_DIR),
        }, entry);
    }
    return {
        name: 'uni:mp-plugin',
        enforce: 'post',
        config() {
            return {
                build: {
                    rollupOptions,
                },
            };
        },
        generateBundle: (0, subpackage_1.createNonAppGenerateBundle)(extname),
    };
}
exports.uniMiniProgramPluginPlugin = uniMiniProgramPluginPlugin;
function initPluginEntry() {
    const pluginJsonFilename = path_1.default.resolve(process.env.UNI_INPUT_DIR, uni_cli_shared_1.MP_PLUGIN_JSON_NAME);
    if (!fs_1.default.existsSync(pluginJsonFilename)) {
        (0, build_1.notFound)(pluginJsonFilename);
    }
    const pluginJson = (0, uni_cli_shared_1.parseJson)(fs_1.default.readFileSync(pluginJsonFilename, 'utf8'), true);
    if (!pluginJson.main) {
        return;
    }
    const mainFilename = path_1.default.resolve(process.env.UNI_INPUT_DIR, pluginJson.main);
    if (!fs_1.default.existsSync(mainFilename)) {
        (0, build_1.notFound)(mainFilename);
    }
    return {
        [(0, uni_cli_shared_1.removeExt)(pluginJson.main)]: mainFilename,
    };
}

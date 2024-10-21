"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineUniManifestJsonPlugin = exports.defineUniPagesJsonPlugin = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const constants_1 = require("../../constants");
const utils_1 = require("../../utils");
exports.defineUniPagesJsonPlugin = createDefineJsonJsPlugin('pages.json');
exports.defineUniManifestJsonPlugin = createDefineJsonJsPlugin('manifest.json');
function createDefineJsonJsPlugin(name) {
    const JSON_JS = constants_1.JSON_JS_MAP[name];
    return function (createVitePlugin) {
        const opts = {
            resolvedConfig: {},
            filter(id) {
                return id.endsWith(JSON_JS);
            },
        };
        const plugin = createVitePlugin(opts);
        const origLoad = plugin.load;
        const origResolveId = plugin.resolveId;
        const origConfigResolved = plugin.configResolved;
        let jsonPath = '';
        let jsonJsPath = '';
        plugin.resolveId = function (id, importer, options) {
            const res = origResolveId && origResolveId.call(this, id, importer, options);
            if (res) {
                return res;
            }
            if (id.endsWith(JSON_JS)) {
                return jsonJsPath;
            }
        };
        plugin.configResolved = function (config) {
            opts.resolvedConfig = config;
            jsonPath = (0, utils_1.normalizePath)(path_1.default.join(process.env.UNI_INPUT_DIR, name));
            jsonJsPath = (0, utils_1.normalizePath)(path_1.default.join(process.env.UNI_INPUT_DIR, JSON_JS));
            return origConfigResolved && origConfigResolved(config);
        };
        plugin.load = function (id, ssr) {
            const res = origLoad && origLoad.call(this, id, ssr);
            if (res) {
                return res;
            }
            if (!opts.filter(id)) {
                return;
            }
            return fs_1.default.readFileSync(jsonPath, 'utf8');
        };
        return plugin;
    };
}

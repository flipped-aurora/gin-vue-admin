"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.copyMiniProgramThemeJson = exports.copyMiniProgramPluginJson = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const json_1 = require("../json/json");
const manifest_1 = require("../json/manifest");
exports.copyMiniProgramPluginJson = {
    src: ['plugin.json'],
    get dest() {
        return process.env.UNI_OUTPUT_DIR;
    },
    transform(source) {
        return JSON.stringify((0, json_1.parseJson)(source.toString(), true), null, 2);
    },
};
const copyMiniProgramThemeJson = () => {
    if (!process.env.UNI_INPUT_DIR)
        return [];
    const manifestJson = (0, manifest_1.getPlatformManifestJsonOnce)();
    const themeLocation = manifestJson.themeLocation || 'theme.json';
    const hasThemeJson = fs_1.default.existsSync(path_1.default.resolve(process.env.UNI_INPUT_DIR, themeLocation));
    if (hasThemeJson) {
        return [
            {
                src: [(manifestJson.themeLocation = themeLocation)],
                get dest() {
                    return process.env.UNI_OUTPUT_DIR;
                },
                transform(source) {
                    return JSON.stringify((0, json_1.parseJson)(source.toString(), true), null, 2);
                },
            },
        ];
    }
    return [];
};
exports.copyMiniProgramThemeJson = copyMiniProgramThemeJson;

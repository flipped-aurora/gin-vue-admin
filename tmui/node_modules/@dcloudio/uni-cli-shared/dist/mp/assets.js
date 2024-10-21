"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.isMiniProgramAssetFile = void 0;
const path_1 = __importDefault(require("path"));
const EXTNAMES = [
    '.png',
    '.jpg',
    '.jpeg',
    '.gif',
    '.svg',
    '.json',
    '.cer',
    '.mp3',
    '.aac',
    '.m4a',
    '.mp4',
    '.wav',
    '.ogg',
    '.silk',
    '.wasm',
    '.br',
    '.cert',
];
function isMiniProgramAssetFile(filename) {
    if (!path_1.default.isAbsolute(filename)) {
        return false;
    }
    return EXTNAMES.includes(path_1.default.extname(filename));
}
exports.isMiniProgramAssetFile = isMiniProgramAssetFile;

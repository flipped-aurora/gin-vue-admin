"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.emptyDir = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
function emptyDir(dir, skip = []) {
    try {
        for (const file of fs_1.default.readdirSync(dir)) {
            if (skip.includes(file)) {
                continue;
            }
            // node >= 14.14.0
            fs_1.default.rmSync(path_1.default.resolve(dir, file), { recursive: true, force: true });
        }
    }
    catch (e) { }
}
exports.emptyDir = emptyDir;

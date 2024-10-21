"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteExistsSyncHasRootFile = exports.cssTarget = exports.getIsStaticFile = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
var static_1 = require("./plugins/vitejs/plugins/static");
Object.defineProperty(exports, "getIsStaticFile", { enumerable: true, get: function () { return static_1.getIsStaticFile; } });
exports.cssTarget = 'chrome53';
__exportStar(require("./utils"), exports);
__exportStar(require("./plugins"), exports);
__exportStar(require("./features"), exports);
__exportStar(require("./autoImport"), exports);
__exportStar(require("./cloud"), exports);
// https://github.com/vitejs/vite/blob/aac2ef77521f66ddd908f9d97020b8df532148cf/packages/vite/src/node/server/searchRoot.ts#L38
// vite 在初始化阶段会执行 initTSConfck，此时会 searchForWorkspaceRoot，如果找到了 pnpm-workspace.yaml 文件，会将其作为 root
// HBuilderX 项目，root 一定是 UNI_INPUT_DIR，所以需要重写 fs.existsSync，不重写的话，可能会找错，
// 一旦找错目录，而该目录下有 N 多文件目录，会导致遍历及其缓慢
function rewriteExistsSyncHasRootFile() {
    const existsSync = fs_1.default.existsSync;
    const pnpmWorkspaceYaml = path_1.default.join(process.env.UNI_INPUT_DIR, 'pnpm-workspace.yaml');
    fs_1.default.existsSync = (path) => {
        if (path === pnpmWorkspaceYaml) {
            return true;
        }
        return existsSync(path);
    };
}
exports.rewriteExistsSyncHasRootFile = rewriteExistsSyncHasRootFile;

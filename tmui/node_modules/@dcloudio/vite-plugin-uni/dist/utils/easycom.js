"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initEasycom = void 0;
const debug_1 = __importDefault(require("debug"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const debugEasycom = (0, debug_1.default)('uni:easycom');
const initEasycom = (watcher) => {
    const { filter, refresh, options } = (0, uni_cli_shared_1.initEasycomsOnce)(process.env.UNI_INPUT_DIR, {
        dirs: [(0, uni_cli_shared_1.resolveComponentsLibPath)()],
        platform: process.env.UNI_PLATFORM,
        isX: process.env.UNI_APP_X === 'true',
    });
    if (!watcher) {
        // build 模式，手动初始化 watcher
        debugEasycom('initWatch', options.dirs);
        watcher = uni_cli_shared_1.chokidar.watch(options.dirs, {
            ignored: ['**/node_modules/**', '**/.git/**'],
            ignoreInitial: true,
            ignorePermissionErrors: true,
            disableGlobbing: true,
        });
    }
    const refreshEasycom = (0, uni_shared_1.debounce)(refresh, 100, { setTimeout, clearTimeout });
    watcher.on('all', (eventName, path) => {
        if (!['add', 'unlink'].includes(eventName)) {
            return;
        }
        if (filter(path)) {
            debugEasycom('watch', eventName, path);
            refreshEasycom();
        }
    });
};
exports.initEasycom = initEasycom;

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPostSourceMapPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const path_1 = require("path");
function uniPostSourceMapPlugin() {
    return {
        name: 'uni:post-sourcemap',
        apply: 'serve',
        enforce: 'post',
        configureServer(server) {
            // cli 工程呢？
            // 重要：hack 了 _pendingRequests，来修改 map
            const pendingRequests = new PendingRequests();
            pendingRequests._server = server;
            pendingRequests._inputDir = (0, uni_cli_shared_1.normalizePath)(process.env.UNI_INPUT_DIR);
            // @ts-expect-error
            server._pendingRequests = pendingRequests;
        },
    };
}
exports.uniPostSourceMapPlugin = uniPostSourceMapPlugin;
function isSourceMap(map) {
    return map && map.sources;
}
class PendingRequests extends Map {
    set(key, value) {
        const then = value.request.then;
        // @ts-expect-error
        value.request.then = (onFulfilled, onRejected) => {
            // @ts-expect-error
            return then.call(value.request, (request) => {
                const map = request?.map;
                if (map) {
                    // @ts-expect-error
                    const mod = this._server.moduleGraph._getUnresolvedUrlToModule(key);
                    if (mod && mod.file && (0, path_1.isAbsolute)(mod.file)) {
                        const dir = (0, uni_cli_shared_1.normalizePath)((0, path_1.dirname)(mod.file));
                        if (dir.startsWith(this._inputDir) && isSourceMap(map)) {
                            for (let sourcesIndex = 0; sourcesIndex < map.sources.length; ++sourcesIndex) {
                                const sourcePath = map.sources[sourcesIndex];
                                if (sourcePath) {
                                    // 将相对路径转换为绝对路径
                                    if (!(0, path_1.isAbsolute)(sourcePath)) {
                                        map.sources[sourcesIndex] = (0, uni_cli_shared_1.normalizePath)((0, path_1.join)(dir, sourcePath));
                                    }
                                }
                            }
                        }
                    }
                }
                return onFulfilled?.(request);
            }, onRejected);
        };
        return super.set(key, value);
    }
}

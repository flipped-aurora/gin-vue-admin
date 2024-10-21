"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniStatsPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const emittedHashMap = new WeakMap();
function uniStatsPlugin() {
    let resolvedConfig;
    return {
        name: 'uni:app-stats',
        enforce: 'post',
        configResolved(config) {
            resolvedConfig = config;
            emittedHashMap.set(resolvedConfig, new Map());
        },
        writeBundle(_, bundle) {
            if (resolvedConfig.isProduction) {
                // 仅dev生效
                return;
            }
            const emittedHash = emittedHashMap.get(resolvedConfig);
            const changedFiles = [];
            Object.keys(bundle).forEach((filename) => {
                const outputFile = bundle[filename];
                let outputFileHash = '';
                if (outputFile.type === 'asset') {
                    outputFileHash = (0, uni_cli_shared_1.hash)(outputFile.source);
                }
                else {
                    outputFileHash = (0, uni_cli_shared_1.hash)(outputFile.code);
                }
                if (emittedHash.get(filename) !== outputFileHash) {
                    emittedHash.set(filename, outputFileHash);
                    changedFiles.push(filename);
                }
            });
            process.env.UNI_APP_CHANGED_FILES = changedFiles.length
                ? JSON.stringify(changedFiles)
                : '';
        },
    };
}
exports.uniStatsPlugin = uniStatsPlugin;

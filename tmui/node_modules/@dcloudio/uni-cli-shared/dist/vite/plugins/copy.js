"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniViteCopyPlugin = void 0;
const watcher_1 = require("../../watcher");
const messages_1 = require("../../messages");
const logs_1 = require("../../logs");
const uni_shared_1 = require("@dcloudio/uni-shared");
function uniViteCopyPlugin({ targets, }) {
    let resolvedConfig;
    let initialized = false;
    return {
        name: 'uni:copy',
        apply: 'build',
        configResolved(config) {
            resolvedConfig = config;
        },
        writeBundle() {
            if (initialized) {
                return;
            }
            if (resolvedConfig.build.ssr) {
                return;
            }
            initialized = true;
            const is_prod = process.env.NODE_ENV !== 'development' ||
                process.env.UNI_AUTOMATOR_CONFIG;
            const onChange = is_prod
                ? undefined
                : (0, uni_shared_1.debounce)(() => {
                    (0, logs_1.resetOutput)('log');
                    (0, logs_1.output)('log', messages_1.M['dev.watching.end']);
                }, 100, { setTimeout, clearTimeout });
            return new Promise((resolve) => {
                Promise.all(targets.map(({ watchOptions, ...target }) => {
                    return new Promise((resolve) => {
                        new watcher_1.FileWatcher({
                            ...target,
                        }).watch({
                            cwd: process.env.UNI_INPUT_DIR,
                            persistent: is_prod ? false : true,
                            ...watchOptions,
                        }, () => {
                            resolve(void 0);
                        }, onChange);
                    });
                })).then(() => resolve());
            });
        },
    };
}
exports.uniViteCopyPlugin = uniViteCopyPlugin;

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseArguments = exports.initArguments = void 0;
function initArguments(manifestJson, pagesJson) {
    const args = parseArguments(pagesJson);
    if (args) {
        manifestJson.plus.arguments = args;
    }
}
exports.initArguments = initArguments;
function parseArguments(pagesJson) {
    if (process.env.NODE_ENV !== 'development') {
        return;
    }
    // 指定了入口
    if (process.env.UNI_CLI_LAUNCH_PAGE_PATH) {
        return JSON.stringify({
            path: process.env.UNI_CLI_LAUNCH_PAGE_PATH,
            query: process.env.UNI_CLI_LAUNCH_PAGE_QUERY,
        });
    }
    const condition = pagesJson.condition;
    if (condition && condition.list?.length) {
        const list = condition.list;
        let current = condition.current || 0;
        if (current < 0) {
            current = 0;
        }
        if (current >= list.length) {
            current = 0;
        }
        return JSON.stringify(list[current]);
    }
}
exports.parseArguments = parseArguments;

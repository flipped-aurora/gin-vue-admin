"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initLaunchwebview = void 0;
const shared_1 = require("@vue/shared");
function initLaunchwebview(manifestJson, pagesJson) {
    let entryPagePath = pagesJson.pages[0].path;
    // 依赖前置执行initArguments
    if (manifestJson.plus.arguments) {
        try {
            const args = JSON.parse(manifestJson.plus.arguments);
            if (args.path) {
                entryPagePath = args.path;
            }
        }
        catch (e) { }
    }
    manifestJson.plus.useragent.value = 'uni-app';
    (0, shared_1.extend)(manifestJson.plus.launchwebview, {
        id: '1',
        kernel: 'WKWebview',
    });
    // 首页为nvue
    const entryPage = pagesJson.pages.find((p) => p.path === entryPagePath);
    if (entryPage?.style.isNVue) {
        manifestJson.plus.launchwebview.uniNView = { path: entryPagePath + '.js' };
    }
    else {
        manifestJson.launch_path = '__uniappview.html';
    }
    return entryPagePath;
}
exports.initLaunchwebview = initLaunchwebview;

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initTabBar = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
function initTabBar(entryPagePath, manifestJson, pagesJson) {
    if (!pagesJson.tabBar?.list?.length) {
        return;
    }
    const tabBar = JSON.parse(JSON.stringify(pagesJson.tabBar));
    tabBar.borderStyle = (0, uni_shared_1.normalizeTabBarStyles)(tabBar.borderStyle);
    if (!tabBar.selectedColor) {
        tabBar.selectedColor = uni_shared_1.SELECTED_COLOR;
    }
    tabBar.height = `${parseFloat(tabBar.height) || uni_shared_1.TABBAR_HEIGHT}px`;
    // 首页是 tabBar 页面
    const item = tabBar.list.find((page) => page.pagePath === entryPagePath);
    if (item) {
        ;
        tabBar.child = ['lauchwebview'];
        tabBar.selected = tabBar.list.indexOf(item);
    }
    manifestJson.plus.tabBar = tabBar;
}
exports.initTabBar = initTabBar;

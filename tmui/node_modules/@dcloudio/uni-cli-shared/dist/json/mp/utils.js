"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseTabBar = exports.parseWindowOptions = void 0;
const shared_1 = require("@vue/shared");
const pages_1 = require("../pages");
function trimJson(json) {
    delete json.maxWidth;
    delete json.topWindow;
    delete json.leftWindow;
    delete json.rightWindow;
    if (json.tabBar) {
        delete json.tabBar.matchMedia;
    }
    return json;
}
function convert(to, from, map) {
    Object.keys(map).forEach((key) => {
        if ((0, shared_1.hasOwn)(from, map[key])) {
            to[key] = from[map[key]];
        }
    });
    return to;
}
function parseWindowOptions(style, platform, windowOptionsMap) {
    if (!style) {
        return {};
    }
    const platformStyle = style[platform] || {};
    (0, pages_1.removePlatformStyle)(trimJson(style));
    const res = {};
    if (windowOptionsMap) {
        return (0, shared_1.extend)(convert(res, style, windowOptionsMap), platformStyle);
    }
    return (0, shared_1.extend)(res, style, platformStyle);
}
exports.parseWindowOptions = parseWindowOptions;
function trimTabBarJson(tabBar) {
    ;
    [
        'fontSize',
        'height',
        'iconWidth',
        'midButton',
        'selectedIndex',
        'spacing',
    ].forEach((name) => {
        delete tabBar[name];
    });
    return tabBar;
}
function parseTabBar(tabBar, platform, tabBarOptionsMap, tabBarItemOptionsMap) {
    const platformStyle = tabBar[platform] || {};
    (0, pages_1.removePlatformStyle)(trimTabBarJson(tabBar));
    const res = {};
    if (tabBarOptionsMap) {
        if (tabBarItemOptionsMap && tabBar.list) {
            tabBar.list = tabBar.list.map((item) => {
                return convert({}, item, tabBarItemOptionsMap);
            });
        }
        convert(res, tabBar, tabBarOptionsMap);
        return (0, shared_1.extend)(res, platformStyle);
    }
    return (0, shared_1.extend)(res, tabBar, platformStyle);
}
exports.parseTabBar = parseTabBar;

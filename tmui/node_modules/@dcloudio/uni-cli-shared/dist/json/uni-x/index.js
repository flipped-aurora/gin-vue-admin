"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.normalizeUniAppXAppConfig = exports.normalizeUniAppXAppPagesJson = exports.checkPagesJson = exports.parseUniXSplashScreen = exports.parseUniXFlexDirection = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const shared_1 = require("@vue/shared");
const jsonc_parser_1 = require("jsonc-parser");
const json_1 = require("../json");
const pages_1 = require("../pages");
const utils_1 = require("../../utils");
const uniRoutes_1 = require("../app/pages/uniRoutes");
const uniConfig_1 = require("./uniConfig");
const utils_2 = require("../../vite/plugins/vitejs/utils");
const preprocess_1 = require("../../preprocess");
var manifest_1 = require("./manifest");
Object.defineProperty(exports, "parseUniXFlexDirection", { enumerable: true, get: function () { return manifest_1.parseUniXFlexDirection; } });
Object.defineProperty(exports, "parseUniXSplashScreen", { enumerable: true, get: function () { return manifest_1.parseUniXSplashScreen; } });
function checkPagesJson(jsonStr, inputDir) {
    if (!inputDir) {
        return false;
    }
    const errors = [];
    const root = (0, jsonc_parser_1.parseTree)(jsonStr, errors);
    if (!root) {
        if (errors.length) {
            for (const error of errors) {
                const { line, column } = (0, utils_2.offsetToLineColumn)(jsonStr, error.offset);
                throw {
                    name: 'SyntaxError',
                    code: error.error,
                    message: (0, jsonc_parser_1.printParseErrorCode)(error.error),
                    loc: {
                        start: { line, column },
                    },
                    offsetStart: error.offset,
                    offsetEnd: error.offset + error.length,
                };
            }
        }
        return false;
    }
    const pagePathNodes = walkNodes(findRootNode(root, ['pages']));
    findRootNode(root, ['subPackages', 'subpackages']).forEach((node) => {
        const subPackageRoot = findSubPackageRoot(node);
        if (subPackageRoot) {
            findRootNode(node, ['pages']).forEach((subNode) => {
                walkNodes(subNode.children ?? []).forEach((node) => {
                    pagePathNodes.push({
                        ...node,
                        value: (0, utils_1.normalizePath)(path_1.default.join(subPackageRoot, node.value)),
                    });
                });
            });
        }
    });
    if (pagePathNodes.length) {
        for (const node of pagePathNodes) {
            const pagePath = node.value ?? '';
            if (pageExistsWithCaseSync(path_1.default.join(inputDir, pagePath))) {
                continue;
            }
            const { line, column } = (0, utils_2.offsetToLineColumn)(jsonStr, node.offset);
            throw {
                name: 'CompilerError',
                code: 'CompilerError',
                message: `The page path "${pagePath}" does not exist`,
                loc: {
                    start: { line, column },
                },
                offsetStart: node.offset,
                offsetEnd: node.offset + node.length,
            };
        }
    }
    return true;
}
exports.checkPagesJson = checkPagesJson;
function pageExistsWithCaseSync(pagePath) {
    try {
        const files = fs_1.default.readdirSync(path_1.default.dirname(pagePath));
        const basename = path_1.default.basename(pagePath);
        const uvuePage = basename + '.uvue';
        const vuePage = basename + '.vue';
        return files.some((file) => file === uvuePage || file === vuePage);
    }
    catch (e) {
        return false;
    }
}
function findSubPackageRoot(node) {
    const child = node.children?.find((child) => child.type === 'property' &&
        child.children &&
        child.children.find((child) => child.type === 'string' && child.value === 'root'));
    if (child && child.children?.length === 2) {
        return child.children[1].value;
    }
    return '';
}
function findRootNode(node, property) {
    const { type, children } = node;
    if (type === 'object' && children) {
        const child = children.find((child) => child.type === 'property' &&
            child.children &&
            child.children.find((child) => child.type === 'string' && property.includes(child.value)));
        if (child) {
            const node = child.children.find((child) => child.type === 'array');
            return node?.children ?? [];
        }
    }
    return [];
}
function walkNodes(node) {
    const pagePathNodes = [];
    node.forEach((node) => walkNode(node, pagePathNodes));
    return pagePathNodes;
}
function walkNode(node, pagePathNodes) {
    const { type, children } = node;
    if (type === 'property' && children && children.length === 2) {
        const maybePagePathNode = children[0];
        const maybePagePathValueNode = children[1];
        if (maybePagePathNode.type === 'string' &&
            maybePagePathNode.value === 'path' &&
            maybePagePathValueNode.type === 'string' &&
            (0, shared_1.isString)(maybePagePathValueNode.value)) {
            pagePathNodes.push(maybePagePathValueNode);
        }
    }
    if (children) {
        children.forEach((node) => walkNode(node, pagePathNodes));
    }
}
function normalizeUniAppXAppPagesJson(jsonStr) {
    // 先条件编译
    jsonStr = (0, preprocess_1.preUVueJson)(jsonStr);
    checkPagesJson(jsonStr, process.env.UNI_INPUT_DIR);
    const pagesJson = {
        pages: [],
        globalStyle: {},
    };
    let userPagesJson = {
        pages: [],
        globalStyle: {},
    };
    try {
        // 此处不需要条件编译了
        userPagesJson = (0, json_1.parseJson)(jsonStr, false);
    }
    catch (e) {
        console.error(`[vite] Error: pages.json parse failed.\n`, jsonStr, e);
    }
    // pages
    (0, pages_1.validatePages)(userPagesJson, jsonStr);
    userPagesJson.subPackages =
        userPagesJson.subPackages || userPagesJson.subpackages;
    // subPackages
    if (userPagesJson.subPackages) {
        userPagesJson.pages.push(...normalizeSubPackages(userPagesJson.subPackages));
    }
    pagesJson.pages = userPagesJson.pages;
    // pageStyle
    normalizePages(pagesJson.pages);
    // globalStyle
    pagesJson.globalStyle = normalizePageStyle(userPagesJson.globalStyle);
    // tabBar
    if (userPagesJson.tabBar) {
        pagesJson.tabBar = userPagesJson.tabBar;
    }
    // condition
    if (userPagesJson.condition) {
        pagesJson.condition = userPagesJson.condition;
    }
    // uniIdRouter
    if (userPagesJson.uniIdRouter) {
        pagesJson.uniIdRouter = userPagesJson.uniIdRouter;
    }
    // 是否应该用 process.env.UNI_UTS_PLATFORM
    (0, pages_1.filterPlatformPages)(process.env.UNI_PLATFORM, pagesJson);
    return pagesJson;
}
exports.normalizeUniAppXAppPagesJson = normalizeUniAppXAppPagesJson;
function normalizeSubPackages(subPackages) {
    const pages = [];
    if ((0, shared_1.isArray)(subPackages)) {
        subPackages.forEach(({ root, pages: subPages }) => {
            if (root && subPages.length) {
                subPages.forEach((subPage) => {
                    subPage.path = (0, utils_1.normalizePath)(path_1.default.join(root, subPage.path));
                    subPage.style = subPage.style;
                    pages.push(subPage);
                });
            }
        });
    }
    return pages;
}
function normalizePages(pages) {
    pages.forEach((page) => {
        page.style = normalizePageStyle(page.style);
    });
}
function normalizePageStyle(pageStyle) {
    if (pageStyle) {
        (0, shared_1.extend)(pageStyle, pageStyle['app']);
        (0, pages_1.removePlatformStyle)(pageStyle);
        return pageStyle;
    }
    return {};
}
/**
 * TODO 应该闭包，通过globalThis赋值？
 * @param pagesJson
 * @param manifestJson
 * @returns
 */
function normalizeUniAppXAppConfig(pagesJson, manifestJson) {
    return `const __uniConfig = ${(0, uniConfig_1.normalizeAppXUniConfig)(pagesJson, manifestJson)};
const __uniRoutes = ${(0, uniRoutes_1.normalizeAppUniRoutes)(pagesJson)}.map(uniRoute=>(uniRoute.meta.route=uniRoute.path,__uniConfig.pages.push(uniRoute.path),uniRoute.path='/'+uniRoute.path,uniRoute));

`;
}
exports.normalizeUniAppXAppConfig = normalizeUniAppXAppConfig;

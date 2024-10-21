"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.ThemeSassParser = exports.initTheme = exports.normalizeThemeConfigOnce = exports.parseThemeJson = exports.hasThemeJson = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const json_1 = require("./json");
const uni_shared_1 = require("@dcloudio/uni-shared");
const manifest_1 = require("./manifest");
function hasThemeJson(themeLocation) {
    if (!fs_1.default.existsSync(themeLocation)) {
        return false;
    }
    return true;
}
exports.hasThemeJson = hasThemeJson;
const parseThemeJson = (themeLocation = 'theme.json') => {
    if (!themeLocation || !process.env.UNI_INPUT_DIR) {
        return {};
    }
    themeLocation = path_1.default.join(process.env.UNI_INPUT_DIR, themeLocation);
    if (!hasThemeJson(themeLocation)) {
        return {};
    }
    const jsonStr = fs_1.default.readFileSync(themeLocation, 'utf8');
    return (0, json_1.parseJson)(jsonStr, true);
};
exports.parseThemeJson = parseThemeJson;
const SCHEME_RE = /^([a-z-]+:)?\/\//i;
const DATA_RE = /^data:.*,.*/;
function normalizeFilepath(filepath) {
    if (!(SCHEME_RE.test(filepath) || DATA_RE.test(filepath)) &&
        filepath.indexOf('/') !== 0) {
        return (0, uni_shared_1.addLeadingSlash)(filepath);
    }
    return filepath;
}
const parseThemeJsonIconPath = (themeConfig, curPathThemeJsonKey) => {
    if (themeConfig.light?.[curPathThemeJsonKey]) {
        themeConfig.light[curPathThemeJsonKey] = normalizeFilepath(themeConfig.light[curPathThemeJsonKey]);
    }
    if (themeConfig.dark?.[curPathThemeJsonKey]) {
        themeConfig.dark[curPathThemeJsonKey] = normalizeFilepath(themeConfig.dark[curPathThemeJsonKey]);
    }
};
function normalizeUniConfigThemeJsonIconPath(themeConfig, pagesJson) {
    // 处理 tabbar 下 list -> item -> iconPath、selectedIconPath; midButton -> backgroundImage 路径 / 不开头的情况
    const tabBar = pagesJson.tabBar;
    if (tabBar && tabBar.list && tabBar.list.length) {
        tabBar.list.forEach((item) => {
            if (item.iconPath && item.iconPath.indexOf('@') === 0) {
                parseThemeJsonIconPath(themeConfig, item.iconPath.replace(/^@/, ''));
            }
            if (item.selectedIconPath && item.selectedIconPath.indexOf('@') === 0) {
                parseThemeJsonIconPath(themeConfig, item.selectedIconPath.replace(/^@/, ''));
            }
        });
        const midButtonBackgroundImage = tabBar.midButton?.backgroundImage;
        if (midButtonBackgroundImage &&
            midButtonBackgroundImage.indexOf('@') === 0) {
            parseThemeJsonIconPath(themeConfig, midButtonBackgroundImage.replace(/^@/, ''));
        }
    }
    return themeConfig;
}
const getPagesJson = (inputDir) => {
    const pagesFilename = path_1.default.join(inputDir, 'pages.json');
    if (!fs_1.default.existsSync(pagesFilename)) {
        return {
            pages: [],
            globalStyle: { navigationBar: {} },
        };
    }
    const jsonStr = fs_1.default.readFileSync(pagesFilename, 'utf8');
    return (0, json_1.parseJson)(jsonStr, true);
};
exports.normalizeThemeConfigOnce = (0, uni_shared_1.once)((manifestJsonPlatform = {}) => {
    const themeConfig = (0, exports.parseThemeJson)(manifestJsonPlatform.themeLocation);
    if (process.env.UNI_INPUT_DIR) {
        normalizeUniConfigThemeJsonIconPath(themeConfig, getPagesJson(process.env.UNI_INPUT_DIR));
    }
    return themeConfig;
});
function initTheme(manifestJson, pagesJson) {
    const platform = process.env.UNI_PLATFORM === 'app' ? 'app-plus' : process.env.UNI_PLATFORM;
    const manifestPlatform = (0, manifest_1.getPlatformManifestJson)(manifestJson, platform) || {};
    const themeConfig = (0, exports.normalizeThemeConfigOnce)(manifestPlatform);
    return (0, uni_shared_1.normalizeStyles)(pagesJson, themeConfig);
}
exports.initTheme = initTheme;
// TODO
class ThemeSassParser {
    constructor() {
        this._index = 0;
        this._input = '';
        this._theme = {};
    }
    parse(input) {
        this._index = 0;
        this._input = input;
        this._theme = {};
        this._theme['light'] = {};
        this._theme['dark'] = {};
        this.parseVariable();
        return this._theme;
    }
    parseVariable() {
        this.skipWhiteSpaceAndComments();
        this.consume('$');
        this.skipWhiteSpaceAndComments();
        let key = this.parseString();
        this.skipWhiteSpaceAndComments();
        this.consume(':');
        this.skipWhiteSpace();
        const value = this.parseVariableValue();
        if (Array.isArray(value)) {
            this.pushThemeValue(key, value);
        }
        this.consume(';');
        this.skipWhiteSpaceAndComments();
        if (this._index < this._input.length) {
            this.parseVariable();
        }
    }
    parseVariableValue() {
        switch (this.currentChar) {
            case 'l':
                return this.parseFunction();
            default:
                return this.skipOtherValue();
        }
    }
    parseFunction() {
        let functionName = '';
        while (this.currentChar != '(') {
            functionName += this.currentChar;
            if (this._index + 1 < this._input.length) {
                ++this._index;
            }
            else {
                break;
            }
        }
        if (functionName != 'light-dark') {
            return this.skipOtherValue();
        }
        let valuePair = new Array(2);
        valuePair[0] = '';
        valuePair[1] = '';
        this.consume('(');
        let index = 0;
        // TODO rgb?
        while (this.currentChar != ')') {
            valuePair[index] += this.currentChar;
            if (this.currentChar === ',') {
                index++;
            }
            ++this._index;
        }
        this.consume(')');
        valuePair[0] = valuePair[0].trim();
        valuePair[1] = valuePair[1].trim();
        return valuePair;
    }
    skipOtherValue() {
        while (this.currentChar != ';') {
            if (this._index + 1 < this._input.length) {
                ++this._index;
            }
            else {
                break;
            }
        }
    }
    parseString() {
        let str = '';
        while (this.currentChar != ':') {
            if (this.currentChar == '\\') {
                str += this.currentChar;
                if (this._index + 1 < this._input.length) {
                    str += this._input[++this._index];
                }
            }
            str += this.currentChar;
            if (this._index + 1 < this._input.length) {
                ++this._index;
            }
            else {
                break;
            }
        }
        return str;
    }
    pushThemeValue(key, valuePair) {
        this._theme['light'][key] = valuePair[0];
        this._theme['dark'][key] = valuePair[1];
    }
    consume(expected) {
        if (this.currentChar != expected) {
            throw new Error('Unexpected character ' +
                expected +
                ' index=' +
                this._index +
                ' ' +
                this.currentChar);
        }
        ++this._index;
    }
    get currentChar() {
        if (this._index >= this._input.length) {
            throw new Error('Unexpected end of input');
        }
        return this._input[this._index];
    }
    skipWhiteSpaceAndComments() {
        while (this._index < this._input.length) {
            const c = this._input[this._index];
            if (this.isspace(c)) {
                ++this._index;
            }
            else if (c == '/') {
                this.skipComment();
            }
            else {
                break;
            }
        }
    }
    skipComment() {
        if (this.currentChar != '/') {
            return;
        }
        this.consume('/');
        let nextChar = this.currentChar;
        if (nextChar == '/') {
            // Single line comment
            while (
            // @ts-expect-error
            this.currentChar !== '\n' &&
                this._index < this._input.length - 1) {
                ++this._index;
            }
            this.skipWhiteSpace();
        }
        else if (nextChar === '*') {
            // Multi-line comment
            while (true) {
                if (this._index + 1 >= this._input.length) {
                    throw new Error('Unterminated multi-line comment');
                }
                ++this._index;
                // @ts-expect-error
                if (this.currentChar === '*' && this._input[this._index + 1] === '/') {
                    this._index += 2;
                    break;
                }
            }
            this.skipWhiteSpace();
        }
        else {
            throw new Error('Invalid comment');
        }
    }
    skipWhiteSpace() {
        while (this._index < this._input.length &&
            this.isspace(this._input[this._index])) {
            ++this._index;
        }
    }
    isspace(str) {
        return str == ' ' || str == '\n' || str == '\r' || str == "'";
    }
}
exports.ThemeSassParser = ThemeSassParser;

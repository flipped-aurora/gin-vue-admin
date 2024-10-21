"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.validateThemeValue = exports.getPlatformManifestJsonOnce = exports.getPlatformManifestJson = exports.getDevServerOptions = exports.isEnableTreeShaking = exports.getRouterOptions = exports.isUniPushOffline = exports.hasPushModule = exports.isEnableSecureNetwork = exports.isEnableUniPushV2 = exports.isEnableUniPushV1 = exports.getUniStatistics = exports.normalizeNetworkTimeout = exports.parseCompatConfigOnce = exports.parseRpx2UnitOnce = exports.parseManifestJsonOnce = exports.parseManifestJson = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const json_1 = require("./json");
const parseManifestJson = (inputDir) => {
    const manifestFilename = path_1.default.join(inputDir, 'manifest.json');
    if (!fs_1.default.existsSync(manifestFilename)) {
        if (process.env.UNI_COMPILE_TARGET === 'uni_modules') {
            return {};
        }
    }
    return (0, json_1.parseJson)(fs_1.default.readFileSync(path_1.default.join(inputDir, 'manifest.json'), 'utf8'));
};
exports.parseManifestJson = parseManifestJson;
exports.parseManifestJsonOnce = (0, uni_shared_1.once)(exports.parseManifestJson);
exports.parseRpx2UnitOnce = (0, uni_shared_1.once)((inputDir, platform = 'h5') => {
    const rpx2unit = platform === 'h5' || platform === 'app' || platform === 'app-harmony'
        ? uni_shared_1.defaultRpx2Unit
        : uni_shared_1.defaultMiniProgramRpx2Unit;
    const manifestJson = (0, exports.parseManifestJsonOnce)(inputDir);
    let platformOptions = getPlatformManifestJson(manifestJson, platform);
    if (platformOptions && platformOptions.rpx) {
        return (0, shared_1.extend)({}, rpx2unit, platformOptions);
    }
    return (0, shared_1.extend)({}, rpx2unit);
});
function parseCompatConfig(_inputDir) {
    // 不支持 mode:2
    return { MODE: 3 }; //parseManifestJsonOnce(inputDir).compatConfig || {}
}
exports.parseCompatConfigOnce = (0, uni_shared_1.once)(parseCompatConfig);
const defaultNetworkTimeout = {
    request: 60000,
    connectSocket: 60000,
    uploadFile: 60000,
    downloadFile: 60000,
};
function normalizeNetworkTimeout(networkTimeout) {
    return {
        ...defaultNetworkTimeout,
        ...networkTimeout,
    };
}
exports.normalizeNetworkTimeout = normalizeNetworkTimeout;
function getUniStatistics(inputDir, platform) {
    const manifest = (0, exports.parseManifestJsonOnce)(inputDir);
    let platformManifest = getPlatformManifestJson(manifest, platform);
    return (0, shared_1.extend)({}, manifest.uniStatistics, platformManifest && platformManifest.uniStatistics);
}
exports.getUniStatistics = getUniStatistics;
function isEnableUniPushV1(inputDir, platform) {
    if (isEnableUniPushV2(inputDir, platform)) {
        return false;
    }
    const manifest = (0, exports.parseManifestJsonOnce)(inputDir);
    if (platform === 'app') {
        const push = manifest['app-plus']?.distribute?.sdkConfigs?.push;
        if (push && (0, shared_1.hasOwn)(push, 'unipush')) {
            return true;
        }
    }
    return false;
}
exports.isEnableUniPushV1 = isEnableUniPushV1;
function isEnableUniPushV2(inputDir, platform) {
    const manifest = (0, exports.parseManifestJsonOnce)(inputDir);
    const platformManifest = getPlatformManifestJson(manifest, platform);
    if (platform === 'app') {
        return (platformManifest?.distribute?.sdkConfigs?.push?.unipush?.version == '2');
    }
    return platformManifest?.unipush?.enable === true;
}
exports.isEnableUniPushV2 = isEnableUniPushV2;
function isEnableSecureNetwork(inputDir, platform) {
    const manifest = (0, exports.parseManifestJsonOnce)(inputDir);
    const platformManifest = getPlatformManifestJson(manifest, platform);
    if (platform === 'app') {
        return !!platformManifest?.modules?.SecureNetwork;
    }
    return platformManifest?.secureNetwork?.enable === true;
}
exports.isEnableSecureNetwork = isEnableSecureNetwork;
function hasPushModule(inputDir) {
    const manifest = (0, exports.parseManifestJsonOnce)(inputDir);
    return !!manifest['app-plus']?.modules?.Push;
}
exports.hasPushModule = hasPushModule;
function isUniPushOffline(inputDir) {
    const manifest = (0, exports.parseManifestJsonOnce)(inputDir);
    return (manifest['app-plus']?.distribute?.sdkConfigs?.push?.unipush?.offline ===
        true);
}
exports.isUniPushOffline = isUniPushOffline;
function getRouterOptions(manifestJson) {
    return (0, shared_1.extend)({}, getPlatformManifestJson(manifestJson, 'h5')?.router);
}
exports.getRouterOptions = getRouterOptions;
function isEnableTreeShaking(manifestJson) {
    // 自动化测试时，一定不摇树
    if (process.env.UNI_AUTOMATOR_WS_ENDPOINT) {
        return false;
    }
    const platformManifest = getPlatformManifestJson(manifestJson, 'h5');
    return platformManifest?.optimization?.treeShaking?.enable !== false;
}
exports.isEnableTreeShaking = isEnableTreeShaking;
function getDevServerOptions(manifestJson) {
    const platformManifest = getPlatformManifestJson(manifestJson, 'h5');
    return (0, shared_1.extend)({}, platformManifest?.devServer);
}
exports.getDevServerOptions = getDevServerOptions;
function getPlatformManifestJson(manifestJson, platform) {
    if (!platform) {
        platform = process.env.UNI_PLATFORM;
    }
    if (platform === 'app') {
        return manifestJson['app-plus'] || manifestJson['plus'] || {};
    }
    if (platform === 'h5') {
        return manifestJson.web || manifestJson.h5 || {};
    }
    return manifestJson[platform] || {};
}
exports.getPlatformManifestJson = getPlatformManifestJson;
function getPlatformManifestJsonOnce() {
    const manifestJson = !process.env.UNI_INPUT_DIR
        ? {}
        : (0, exports.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR);
    return getPlatformManifestJson(manifestJson);
}
exports.getPlatformManifestJsonOnce = getPlatformManifestJsonOnce;
const themeValues = ['dark', 'light', 'auto'];
function validateThemeValue(value) {
    return themeValues.indexOf(value) !== -1;
}
exports.validateThemeValue = validateThemeValue;

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initRecursiveMerge = void 0;
const merge_1 = require("merge");
function initRecursiveMerge(manifestJson, userManifestJson) {
    return (0, merge_1.recursive)(true, manifestJson, {
        id: userManifestJson.appid || '',
        name: userManifestJson.name || '',
        description: userManifestJson.description || '',
        version: {
            name: userManifestJson.versionName,
            code: userManifestJson.versionCode,
        },
        locale: userManifestJson.locale,
        uniStatistics: userManifestJson.uniStatistics,
    }, { plus: userManifestJson['app-plus'] });
}
exports.initRecursiveMerge = initRecursiveMerge;

"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformComponentLink = exports.transformTapToClick = exports.transformMatchMedia = exports.transformH5BuiltInComponents = exports.matchTransformModel = exports.createTransformModel = exports.matchTransformOn = exports.createTransformOn = exports.ATTR_DATASET_EVENT_OPTS = exports.STRINGIFY_JSON = exports.createSrcsetTransformWithOptions = exports.createAssetUrlTransformWithOptions = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
const transformTag_1 = require("./transformTag");
const transformEvent_1 = require("./transformEvent");
const transformComponent_1 = require("./transformComponent");
const constants_1 = require("../../mp/constants");
__exportStar(require("./transformRef"), exports);
__exportStar(require("./transformPageHead"), exports);
__exportStar(require("./transformComponent"), exports);
__exportStar(require("./transformEvent"), exports);
__exportStar(require("./transformTag"), exports);
__exportStar(require("./transformUTSComponent"), exports);
__exportStar(require("./transformRefresherSlot"), exports);
var templateTransformAssetUrl_1 = require("./templateTransformAssetUrl");
Object.defineProperty(exports, "createAssetUrlTransformWithOptions", { enumerable: true, get: function () { return templateTransformAssetUrl_1.createAssetUrlTransformWithOptions; } });
var templateTransformSrcset_1 = require("./templateTransformSrcset");
Object.defineProperty(exports, "createSrcsetTransformWithOptions", { enumerable: true, get: function () { return templateTransformSrcset_1.createSrcsetTransformWithOptions; } });
var vOn_1 = require("./vOn");
Object.defineProperty(exports, "STRINGIFY_JSON", { enumerable: true, get: function () { return vOn_1.STRINGIFY_JSON; } });
Object.defineProperty(exports, "ATTR_DATASET_EVENT_OPTS", { enumerable: true, get: function () { return vOn_1.ATTR_DATASET_EVENT_OPTS; } });
Object.defineProperty(exports, "createTransformOn", { enumerable: true, get: function () { return vOn_1.createTransformOn; } });
Object.defineProperty(exports, "matchTransformOn", { enumerable: true, get: function () { return vOn_1.defaultMatch; } });
var vModel_1 = require("./vModel");
Object.defineProperty(exports, "createTransformModel", { enumerable: true, get: function () { return vModel_1.createTransformModel; } });
Object.defineProperty(exports, "matchTransformModel", { enumerable: true, get: function () { return vModel_1.defaultMatch; } });
exports.transformH5BuiltInComponents = (0, transformTag_1.createTransformTag)(uni_shared_1.BUILT_IN_TAG_NAMES.reduce((tags, tag) => ((tags[tag] = uni_shared_1.COMPONENT_PREFIX + tag), tags), {}));
exports.transformMatchMedia = (0, transformTag_1.createTransformTag)({
    'match-media': 'uni-match-media',
});
exports.transformTapToClick = (0, transformEvent_1.createTransformEvent)({
    tap: (node) => {
        // 地图组件有自己特定的 tap 事件
        if (node.tag === 'map' || node.tag === 'v-uni-map') {
            return 'tap';
        }
        return 'click';
    },
});
exports.transformComponentLink = (0, transformComponent_1.createTransformComponentLink)(constants_1.COMPONENT_BIND_LINK);

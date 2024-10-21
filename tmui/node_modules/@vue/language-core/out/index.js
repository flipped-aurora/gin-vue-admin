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
exports.tsCodegen = exports.sharedTypes = exports.scriptRanges = void 0;
__exportStar(require("./generators/template"), exports);
__exportStar(require("./languageModule"), exports);
__exportStar(require("./parsers/scriptSetupRanges"), exports);
__exportStar(require("./plugins"), exports);
__exportStar(require("./virtualFile/vueFile"), exports);
__exportStar(require("./types"), exports);
__exportStar(require("./utils/ts"), exports);
__exportStar(require("./utils/parseSfc"), exports);
exports.scriptRanges = require("./parsers/scriptRanges");
exports.sharedTypes = require("./utils/globalTypes");
__exportStar(require("./utils/shared"), exports);
var vue_tsx_1 = require("./plugins/vue-tsx");
Object.defineProperty(exports, "tsCodegen", { enumerable: true, get: function () { return vue_tsx_1.tsCodegen; } });
__exportStar(require("@volar/language-core"), exports);
__exportStar(require("@volar/source-map"), exports);
//# sourceMappingURL=index.js.map
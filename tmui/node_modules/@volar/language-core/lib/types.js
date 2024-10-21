"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FileKind = exports.MirrorBehaviorCapabilities = exports.FileRangeCapabilities = exports.FileCapabilities = void 0;
var FileCapabilities;
(function (FileCapabilities) {
    FileCapabilities.full = {
        diagnostic: true,
        foldingRange: true,
        documentFormatting: true,
        documentSymbol: true,
        codeAction: true,
        inlayHint: true,
    };
})(FileCapabilities || (exports.FileCapabilities = FileCapabilities = {}));
var FileRangeCapabilities;
(function (FileRangeCapabilities) {
    FileRangeCapabilities.full = {
        hover: true,
        references: true,
        definition: true,
        rename: true,
        completion: true,
        diagnostic: true,
        semanticTokens: true,
    };
})(FileRangeCapabilities || (exports.FileRangeCapabilities = FileRangeCapabilities = {}));
var MirrorBehaviorCapabilities;
(function (MirrorBehaviorCapabilities) {
    MirrorBehaviorCapabilities.full = {
        references: true,
        definition: true,
        rename: true,
    };
})(MirrorBehaviorCapabilities || (exports.MirrorBehaviorCapabilities = MirrorBehaviorCapabilities = {}));
var FileKind;
(function (FileKind) {
    FileKind[FileKind["TextFile"] = 0] = "TextFile";
    FileKind[FileKind["TypeScriptHostFile"] = 1] = "TypeScriptHostFile";
})(FileKind || (exports.FileKind = FileKind = {}));
//# sourceMappingURL=types.js.map
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.VueEmbeddedFile = void 0;
const language_core_1 = require("@volar/language-core");
class VueEmbeddedFile {
    constructor(fileName, content, contentStacks) {
        this.fileName = fileName;
        this.content = content;
        this.contentStacks = contentStacks;
        this.kind = language_core_1.FileKind.TextFile;
        this.capabilities = {};
        this.mirrorBehaviorMappings = [];
    }
}
exports.VueEmbeddedFile = VueEmbeddedFile;
//# sourceMappingURL=embeddedFile.js.map
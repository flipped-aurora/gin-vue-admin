"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getDocumentRegistry = void 0;
const documentRegistries = [];
function getDocumentRegistry(ts, useCaseSensitiveFileNames, currentDirectory) {
    let documentRegistry = documentRegistries.find(item => item[0] === useCaseSensitiveFileNames && item[1] === currentDirectory)?.[2];
    if (!documentRegistry) {
        documentRegistry = ts.createDocumentRegistry(useCaseSensitiveFileNames, currentDirectory);
        documentRegistries.push([useCaseSensitiveFileNames, currentDirectory, documentRegistry]);
    }
    return documentRegistry;
}
exports.getDocumentRegistry = getDocumentRegistry;
//# sourceMappingURL=documentRegistry.js.map
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.computedFiles = void 0;
const source_map_1 = require("@volar/source-map");
const muggle = require("muggle-string");
const embeddedFile_1 = require("./embeddedFile");
const computeds_1 = require("computeds");
function computedFiles(plugins, fileName, sfc, codegenStack) {
    const nameToBlock = (0, computeds_1.computed)(() => {
        const blocks = {};
        if (sfc.template) {
            blocks[sfc.template.name] = sfc.template;
        }
        if (sfc.script) {
            blocks[sfc.script.name] = sfc.script;
        }
        if (sfc.scriptSetup) {
            blocks[sfc.scriptSetup.name] = sfc.scriptSetup;
        }
        for (const block of sfc.styles) {
            blocks[block.name] = block;
        }
        for (const block of sfc.customBlocks) {
            blocks[block.name] = block;
        }
        return blocks;
    });
    const pluginsResult = plugins.map(plugin => compiledPluginFiles(plugins, plugin, fileName, sfc, nameToBlock, codegenStack));
    const flatResult = (0, computeds_1.computed)(() => pluginsResult.map(r => r()).flat());
    const structuredResult = (0, computeds_1.computed)(() => {
        const embeddedFiles = [];
        let remain = [...flatResult()];
        while (remain.length) {
            const beforeLength = remain.length;
            consumeRemain();
            if (beforeLength === remain.length) {
                break;
            }
        }
        for (const { file, snapshot, mappings, codegenStacks } of remain) {
            embeddedFiles.push({
                ...file,
                snapshot,
                mappings,
                codegenStacks,
                embeddedFiles: [],
            });
            console.error('Unable to resolve embedded: ' + file.parentFileName + ' -> ' + file.fileName);
        }
        return embeddedFiles;
        function consumeRemain() {
            for (let i = remain.length - 1; i >= 0; i--) {
                const { file, snapshot, mappings, codegenStacks } = remain[i];
                if (!file.parentFileName) {
                    embeddedFiles.push({
                        ...file,
                        snapshot,
                        mappings,
                        codegenStacks,
                        embeddedFiles: [],
                    });
                    remain.splice(i, 1);
                }
                else {
                    const parent = findParentStructure(file.parentFileName, embeddedFiles);
                    if (parent) {
                        parent.embeddedFiles.push({
                            ...file,
                            snapshot,
                            mappings,
                            codegenStacks,
                            embeddedFiles: [],
                        });
                        remain.splice(i, 1);
                    }
                }
            }
        }
        function findParentStructure(fileName, current) {
            for (const child of current) {
                if (child.fileName === fileName) {
                    return child;
                }
                let parent = findParentStructure(fileName, child.embeddedFiles);
                if (parent) {
                    return parent;
                }
            }
        }
    });
    return structuredResult;
}
exports.computedFiles = computedFiles;
function compiledPluginFiles(plugins, plugin, fileName, sfc, nameToBlock, codegenStack) {
    const embeddedFiles = {};
    const files = (0, computeds_1.computed)(() => {
        try {
            if (!plugin.getEmbeddedFileNames) {
                return Object.values(embeddedFiles);
            }
            const embeddedFileNames = plugin.getEmbeddedFileNames(fileName, sfc);
            for (const oldFileName of Object.keys(embeddedFiles)) {
                if (!embeddedFileNames.includes(oldFileName)) {
                    delete embeddedFiles[oldFileName];
                }
            }
            for (const embeddedFileName of embeddedFileNames) {
                if (!embeddedFiles[embeddedFileName]) {
                    embeddedFiles[embeddedFileName] = (0, computeds_1.computed)(() => {
                        const [content, stacks] = codegenStack ? muggle.track([]) : [[], []];
                        const file = new embeddedFile_1.VueEmbeddedFile(embeddedFileName, content, stacks);
                        for (const plugin of plugins) {
                            if (!plugin.resolveEmbeddedFile) {
                                continue;
                            }
                            try {
                                plugin.resolveEmbeddedFile(fileName, sfc, file);
                            }
                            catch (e) {
                                console.error(e);
                            }
                        }
                        const newText = (0, source_map_1.toString)(file.content);
                        const changeRanges = new Map();
                        const snapshot = {
                            getText: (start, end) => newText.slice(start, end),
                            getLength: () => newText.length,
                            getChangeRange(oldSnapshot) {
                                if (!changeRanges.has(oldSnapshot)) {
                                    changeRanges.set(oldSnapshot, undefined);
                                    const oldText = oldSnapshot.getText(0, oldSnapshot.getLength());
                                    const changeRange = fullDiffTextChangeRange(oldText, newText);
                                    if (changeRange) {
                                        changeRanges.set(oldSnapshot, changeRange);
                                    }
                                }
                                return changeRanges.get(oldSnapshot);
                            },
                        };
                        return {
                            file,
                            snapshot,
                        };
                    });
                }
            }
        }
        catch (e) {
            console.error(e);
        }
        return Object.values(embeddedFiles);
    });
    return (0, computeds_1.computed)(() => {
        return files().map(_file => {
            const { file, snapshot } = _file();
            const mappings = (0, source_map_1.buildMappings)(file.content);
            for (const mapping of mappings) {
                if (mapping.source !== undefined) {
                    const block = nameToBlock()[mapping.source];
                    if (block) {
                        mapping.sourceRange = [
                            mapping.sourceRange[0] + block.startTagEnd,
                            mapping.sourceRange[1] + block.startTagEnd,
                        ];
                    }
                    else {
                        // ignore
                    }
                    mapping.source = undefined;
                }
            }
            return {
                file,
                snapshot,
                mappings,
                codegenStacks: (0, source_map_1.buildStacks)(file.content, file.contentStacks),
            };
        });
    });
}
function fullDiffTextChangeRange(oldText, newText) {
    for (let start = 0; start < oldText.length && start < newText.length; start++) {
        if (oldText[start] !== newText[start]) {
            let end = oldText.length;
            for (let i = 0; i < oldText.length - start && i < newText.length - start; i++) {
                if (oldText[oldText.length - i - 1] !== newText[newText.length - i - 1]) {
                    break;
                }
                end--;
            }
            let length = end - start;
            let newLength = length + (newText.length - oldText.length);
            if (newLength < 0) {
                length -= newLength;
                newLength = 0;
            }
            return {
                span: { start, length },
                newLength,
            };
        }
    }
}
//# sourceMappingURL=computedFiles.js.map
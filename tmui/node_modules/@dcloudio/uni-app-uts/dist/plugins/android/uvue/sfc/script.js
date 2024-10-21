"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.resolveScript = exports.scriptIdentifier = exports.isUseInlineTemplate = exports.setResolvedScript = exports.getResolvedScript = exports.invalidateScript = void 0;
const template_1 = require("./template");
const descriptorCache_1 = require("../descriptorCache");
const compileScript_1 = require("./compiler/compileScript");
const scriptCache = new WeakMap();
function invalidateScript(filename) {
    const desc = descriptorCache_1.cache.get(filename);
    if (desc) {
        scriptCache.delete(desc);
    }
}
exports.invalidateScript = invalidateScript;
function getResolvedScript(descriptor) {
    return scriptCache.get(descriptor);
}
exports.getResolvedScript = getResolvedScript;
function setResolvedScript(descriptor, script) {
    scriptCache.set(descriptor, script);
}
exports.setResolvedScript = setResolvedScript;
// Check if we can use compile template as inlined render function
// inside <script setup>. This can only be done for build because
// inlined template cannot be individually hot updated.
function isUseInlineTemplate(descriptor, isProd) {
    return isProd && !!descriptor.scriptSetup && !descriptor.template?.src;
}
exports.isUseInlineTemplate = isUseInlineTemplate;
exports.scriptIdentifier = `__sfc__`;
function resolveScript(descriptor, options) {
    if (!descriptor.script && !descriptor.scriptSetup) {
        return null;
    }
    const cached = getResolvedScript(descriptor);
    if (cached) {
        return cached;
    }
    let resolved = null;
    resolved = (0, compileScript_1.compileScript)(descriptor, {
        ...options.script,
        className: options.className || '',
        root: options.root,
        id: descriptor.id,
        isProd: options.isProduction,
        inlineTemplate: true,
        hoistStatic: false,
        templateOptions: (0, template_1.resolveTemplateCompilerOptions)(descriptor, options),
        sourceMap: options.sourceMap,
        defineModel: true,
        componentType: options.componentType,
        genDefaultAs: exports.scriptIdentifier,
    });
    setResolvedScript(descriptor, resolved);
    return resolved;
}
exports.resolveScript = resolveScript;

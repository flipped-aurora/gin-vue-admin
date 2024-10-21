"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.resolveParserPlugins = exports.ScriptCompileContext = void 0;
const shared_1 = require("@vue/shared");
const parser_1 = require("@babel/parser");
const magic_string_1 = __importDefault(require("magic-string"));
class ScriptCompileContext {
    helper(key) {
        this.helperImports.add(key);
        return `${key}`;
    }
    constructor(descriptor, options) {
        this.descriptor = descriptor;
        this.options = options;
        this.source = this.descriptor.source;
        this.filename = this.descriptor.filename;
        this.s = new magic_string_1.default(this.source);
        this.startOffset = this.descriptor.scriptSetup?.loc.start.offset;
        this.endOffset = this.descriptor.scriptSetup?.loc.end.offset;
        this.userImports = Object.create(null);
        // macros presence check
        this.hasDefinePropsCall = false;
        this.hasDefineEmitCall = false;
        this.hasDefineExposeCall = false;
        this.hasDefaultExportName = false;
        this.hasDefaultExportRender = false;
        this.hasDefineOptionsCall = false;
        this.hasDefineSlotsCall = false;
        this.hasDefineModelCall = false;
        this.propsDestructuredBindings = Object.create(null);
        // defineModel
        this.modelDecls = Object.create(null);
        // codegen
        this.bindingMetadata = {};
        this.helperImports = new Set();
        // resolve parser plugins
        const plugins = resolveParserPlugins('ts', options.babelParserPlugins);
        function parse(input, offset, startLine) {
            try {
                return (0, parser_1.parse)(input, {
                    plugins,
                    sourceType: 'module',
                    // 阻止语法解析报错，不影响后续的语法解析，比如
                    // This member cannot have an 'override' modifier because its containing class does not extend another class.
                    errorRecovery: true,
                }).program;
            }
            catch (e) {
                if (e.loc && startLine) {
                    e.loc.line = e.loc.line + (startLine - 1);
                }
                e.message = `[vue/compiler-sfc] ${e.message}\n\n${descriptor.filename}\n${(0, shared_1.generateCodeFrame)(descriptor.source, e.pos + offset, e.pos + offset + 1)}`;
                throw e;
            }
        }
        this.scriptAst =
            descriptor.script &&
                parse(descriptor.script.content, descriptor.script.loc.start.offset, descriptor.script.loc.start.line);
        this.scriptSetupAst =
            descriptor.scriptSetup &&
                parse(descriptor.scriptSetup.content, this.startOffset, descriptor.scriptSetup.loc.start.line);
    }
    getString(node, scriptSetup = true) {
        const block = scriptSetup
            ? this.descriptor.scriptSetup
            : this.descriptor.script;
        return block.content.slice(node.start, node.end);
    }
    error(msg, node, scope) {
        const offset = scope ? scope.offset : this.startOffset;
        throw new Error(`[@vue/compiler-sfc] ${msg}\n\n${(scope || this.descriptor).filename}\n${(0, shared_1.generateCodeFrame)((scope || this.descriptor).source, node.start + offset, node.end + offset)}`);
    }
}
exports.ScriptCompileContext = ScriptCompileContext;
function resolveParserPlugins(lang, userPlugins, dts = false) {
    const plugins = [];
    if (!userPlugins ||
        !userPlugins.some((p) => p === 'importAssertions' ||
            p === 'importAttributes' ||
            ((0, shared_1.isArray)(p) && p[0] === 'importAttributes'))) {
        plugins.push('importAttributes');
    }
    if (lang === 'jsx' || lang === 'tsx') {
        plugins.push('jsx');
    }
    else if (userPlugins) {
        // If don't match the case of adding jsx
        // should remove the jsx from user options
        userPlugins = userPlugins.filter((p) => p !== 'jsx');
    }
    if (lang === 'ts' || lang === 'tsx') {
        plugins.push(['typescript', { dts }], 'explicitResourceManagement');
        if (!userPlugins || !userPlugins.includes('decorators')) {
            plugins.push('decorators-legacy');
        }
    }
    if (userPlugins) {
        plugins.push(...userPlugins);
    }
    return plugins;
}
exports.resolveParserPlugins = resolveParserPlugins;

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genDefaultScriptCode = exports.genScript = void 0;
function genScript({ script }, { genDefaultAs }) {
    if (!script) {
        return genDefaultScriptCode(genDefaultAs);
    }
    return ('\n'.repeat(script.loc.start.line - 1) +
        `${script.content}
`);
}
exports.genScript = genScript;
function genDefaultScriptCode(genDefaultAs) {
    if (genDefaultAs) {
        return `const ${genDefaultAs} = defineComponent({})
export default ${genDefaultAs}`;
    }
    return `
export default {}
`;
}
exports.genDefaultScriptCode = genDefaultScriptCode;

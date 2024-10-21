"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseScriptRanges = void 0;
const scriptSetupRanges_1 = require("./scriptSetupRanges");
function parseScriptRanges(ts, ast, hasScriptSetup, withNode) {
    let exportDefault;
    const bindings = hasScriptSetup ? (0, scriptSetupRanges_1.parseBindingRanges)(ts, ast) : [];
    ast.forEachChild(raw => {
        if (ts.isExportAssignment(raw)) {
            let node = raw;
            while (ts.isAsExpression(node.expression) || ts.isParenthesizedExpression(node.expression)) { // fix https://github.com/vuejs/language-tools/issues/1882
                node = node.expression;
            }
            let obj;
            if (ts.isObjectLiteralExpression(node.expression)) {
                obj = node.expression;
            }
            else if (ts.isCallExpression(node.expression) && node.expression.arguments.length) {
                const arg0 = node.expression.arguments[0];
                if (ts.isObjectLiteralExpression(arg0)) {
                    obj = arg0;
                }
            }
            if (obj) {
                let componentsOptionNode;
                let nameOptionNode;
                obj.forEachChild(node => {
                    if (ts.isPropertyAssignment(node) && ts.isIdentifier(node.name)) {
                        if (node.name.escapedText === 'components' && ts.isObjectLiteralExpression(node.initializer)) {
                            componentsOptionNode = node.initializer;
                        }
                        if (node.name.escapedText === 'name') {
                            nameOptionNode = node.initializer;
                        }
                    }
                });
                exportDefault = {
                    ..._getStartEnd(raw),
                    expression: _getStartEnd(node.expression),
                    args: _getStartEnd(obj),
                    argsNode: withNode ? obj : undefined,
                    componentsOption: componentsOptionNode ? _getStartEnd(componentsOptionNode) : undefined,
                    componentsOptionNode: withNode ? componentsOptionNode : undefined,
                    nameOption: nameOptionNode ? _getStartEnd(nameOptionNode) : undefined,
                };
            }
        }
    });
    return {
        exportDefault,
        bindings,
    };
    function _getStartEnd(node) {
        return (0, scriptSetupRanges_1.getStartEnd)(node, ast);
    }
}
exports.parseScriptRanges = parseScriptRanges;
//# sourceMappingURL=scriptRanges.js.map
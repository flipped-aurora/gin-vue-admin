"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genWxsCallMethodsCode = exports.parseWxsCallMethods = void 0;
const types_1 = require("@babel/types");
const estree_walker_1 = require("estree-walker");
const ast_1 = require("./ast");
function parseWxsCallMethods(code) {
    if (!code.includes('callMethod')) {
        return [];
    }
    const ast = (0, ast_1.parseProgram)(code, '', {});
    const wxsCallMethods = new Set();
    estree_walker_1.walk(ast, {
        enter(child) {
            if (!(0, types_1.isCallExpression)(child)) {
                return;
            }
            const { callee } = child;
            // .callMethod
            if (!(0, types_1.isMemberExpression)(callee) ||
                !(0, types_1.isIdentifier)(callee.property) ||
                callee.property.name !== 'callMethod') {
                return;
            }
            // .callMethod('test',...)
            const args = child.arguments;
            if (!args.length) {
                return;
            }
            const [name] = args;
            if (!(0, types_1.isStringLiteral)(name)) {
                return;
            }
            wxsCallMethods.add(name.value);
        },
    });
    return [...wxsCallMethods];
}
exports.parseWxsCallMethods = parseWxsCallMethods;
function genWxsCallMethodsCode(code) {
    const wxsCallMethods = parseWxsCallMethods(code);
    if (!wxsCallMethods.length) {
        return `export default {}`;
    }
    return `export default (Component) => {
  if(!Component.wxsCallMethods){
    Component.wxsCallMethods = []
  }
  Component.wxsCallMethods.push(${wxsCallMethods
        .map((m) => `'${m}'`)
        .join(', ')})
}
`;
}
exports.genWxsCallMethodsCode = genWxsCallMethodsCode;

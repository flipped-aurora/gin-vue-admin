"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseExternalClasses = exports.updateMiniProgramComponentExternalClasses = exports.findMiniProgramComponentExternalClasses = exports.hasExternalClasses = void 0;
const types_1 = require("@babel/types");
const estree_walker_1 = require("estree-walker");
const externalClassesCache = new Map();
function hasExternalClasses(code) {
    return code.includes('externalClasses');
}
exports.hasExternalClasses = hasExternalClasses;
function findMiniProgramComponentExternalClasses(filename) {
    return externalClassesCache.get(filename);
}
exports.findMiniProgramComponentExternalClasses = findMiniProgramComponentExternalClasses;
function updateMiniProgramComponentExternalClasses(filename, classes) {
    externalClassesCache.set(filename, classes);
}
exports.updateMiniProgramComponentExternalClasses = updateMiniProgramComponentExternalClasses;
function parseExternalClasses(ast) {
    const classes = [];
    estree_walker_1.walk(ast, {
        enter(child, parent) {
            if (!(0, types_1.isIdentifier)(child) || child.name !== 'externalClasses') {
                return;
            }
            // export default { externalClasses: ['my-class'] }
            if (!(0, types_1.isObjectProperty)(parent)) {
                return;
            }
            if (!(0, types_1.isArrayExpression)(parent.value)) {
                return;
            }
            parent.value.elements.forEach((element) => {
                if ((0, types_1.isStringLiteral)(element)) {
                    classes.push(element.value);
                }
            });
        },
    });
    return classes;
}
exports.parseExternalClasses = parseExternalClasses;

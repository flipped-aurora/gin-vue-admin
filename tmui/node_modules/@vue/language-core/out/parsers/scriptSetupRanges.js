"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getStartEnd = exports.findBindingVars = exports.parseBindingRanges = exports.parseScriptSetupRanges = void 0;
function parseScriptSetupRanges(ts, ast, vueCompilerOptions) {
    let foundNonImportExportNode = false;
    let importSectionEndOffset = 0;
    const props = {};
    const slots = {};
    const emits = {};
    const expose = {};
    const definePropProposalA = vueCompilerOptions.experimentalDefinePropProposal === 'kevinEdition' || ast.getFullText().trimStart().startsWith('// @experimentalDefinePropProposal=kevinEdition');
    const definePropProposalB = vueCompilerOptions.experimentalDefinePropProposal === 'johnsonEdition' || ast.getFullText().trimStart().startsWith('// @experimentalDefinePropProposal=johnsonEdition');
    const defineProp = [];
    const bindings = parseBindingRanges(ts, ast);
    const text = ast.getFullText();
    const leadingCommentEndOffset = ts.getLeadingCommentRanges(text, 0)?.reverse()[0].end ?? 0;
    ast.forEachChild(node => {
        const isTypeExport = (ts.isTypeAliasDeclaration(node) || ts.isInterfaceDeclaration(node)) && node.modifiers?.some(mod => mod.kind === ts.SyntaxKind.ExportKeyword);
        if (!foundNonImportExportNode
            && !ts.isImportDeclaration(node)
            && !isTypeExport
            && !ts.isEmptyStatement(node)
            // fix https://github.com/vuejs/language-tools/issues/1223
            && !ts.isImportEqualsDeclaration(node)) {
            const commentRanges = ts.getLeadingCommentRanges(text, node.getFullStart());
            if (commentRanges?.length) {
                const commentRange = commentRanges.sort((a, b) => a.pos - b.pos)[0];
                importSectionEndOffset = commentRange.pos;
            }
            else {
                importSectionEndOffset = node.getStart(ast);
            }
            foundNonImportExportNode = true;
        }
    });
    ast.forEachChild(child => visitNode(child, [ast]));
    return {
        leadingCommentEndOffset,
        importSectionEndOffset,
        bindings,
        props,
        slots,
        emits,
        expose,
        defineProp,
    };
    function _getStartEnd(node) {
        return getStartEnd(node, ast);
    }
    function parseDefineFunction(node) {
        return {
            ..._getStartEnd(node),
            arg: node.arguments.length ? _getStartEnd(node.arguments[0]) : undefined,
            typeArg: node.typeArguments?.length ? _getStartEnd(node.typeArguments[0]) : undefined,
        };
    }
    function visitNode(node, parents) {
        const parent = parents[parents.length - 1];
        if (ts.isCallExpression(node)
            && ts.isIdentifier(node.expression)) {
            const callText = node.expression.getText(ast);
            if (vueCompilerOptions.macros.defineModel.includes(callText)) {
                let name;
                let options;
                if (node.arguments.length >= 2) {
                    name = _getStartEnd(node.arguments[0]);
                    options = node.arguments[1];
                }
                else if (node.arguments.length >= 1) {
                    if (ts.isStringLiteral(node.arguments[0])) {
                        name = _getStartEnd(node.arguments[0]);
                    }
                    else {
                        options = node.arguments[0];
                    }
                }
                let required = false;
                if (options && ts.isObjectLiteralExpression(options)) {
                    for (const property of options.properties) {
                        if (ts.isPropertyAssignment(property) && ts.isIdentifier(property.name) && property.name.getText(ast) === 'required' && property.initializer.kind === ts.SyntaxKind.TrueKeyword) {
                            required = true;
                            break;
                        }
                    }
                }
                defineProp.push({
                    name,
                    nameIsString: true,
                    type: node.typeArguments?.length ? _getStartEnd(node.typeArguments[0]) : undefined,
                    defaultValue: undefined,
                    required,
                });
            }
            else if (callText === 'defineProp') {
                if (definePropProposalA) {
                    let required = false;
                    if (node.arguments.length >= 2) {
                        const secondArg = node.arguments[1];
                        if (ts.isObjectLiteralExpression(secondArg)) {
                            for (const property of secondArg.properties) {
                                if (ts.isPropertyAssignment(property) && ts.isIdentifier(property.name) && property.name.getText(ast) === 'required' && property.initializer.kind === ts.SyntaxKind.TrueKeyword) {
                                    required = true;
                                    break;
                                }
                            }
                        }
                    }
                    if (node.arguments.length >= 1) {
                        defineProp.push({
                            name: _getStartEnd(node.arguments[0]),
                            nameIsString: true,
                            type: node.typeArguments?.length ? _getStartEnd(node.typeArguments[0]) : undefined,
                            defaultValue: undefined,
                            required,
                        });
                    }
                    else if (ts.isVariableDeclaration(parent)) {
                        defineProp.push({
                            name: _getStartEnd(parent.name),
                            nameIsString: false,
                            type: node.typeArguments?.length ? _getStartEnd(node.typeArguments[0]) : undefined,
                            defaultValue: undefined,
                            required,
                        });
                    }
                }
                else if (definePropProposalB && ts.isVariableDeclaration(parent)) {
                    defineProp.push({
                        name: _getStartEnd(parent.name),
                        nameIsString: false,
                        defaultValue: node.arguments.length >= 1 ? _getStartEnd(node.arguments[0]) : undefined,
                        type: node.typeArguments?.length ? _getStartEnd(node.typeArguments[0]) : undefined,
                        required: node.arguments.length >= 2 && node.arguments[1].kind === ts.SyntaxKind.TrueKeyword,
                    });
                }
            }
            else if (vueCompilerOptions.macros.defineSlots.includes(callText)) {
                slots.define = parseDefineFunction(node);
                if (ts.isVariableDeclaration(parent)) {
                    slots.name = parent.name.getText(ast);
                }
            }
            else if (vueCompilerOptions.macros.defineEmits.includes(callText)) {
                emits.define = parseDefineFunction(node);
                if (ts.isVariableDeclaration(parent)) {
                    emits.name = parent.name.getText(ast);
                }
            }
            else if (vueCompilerOptions.macros.defineExpose.includes(callText)) {
                expose.define = parseDefineFunction(node);
            }
            else if (vueCompilerOptions.macros.defineProps.includes(callText)) {
                let statementRange;
                for (let i = parents.length - 1; i >= 0; i--) {
                    if (ts.isStatement(parents[i])) {
                        const statement = parents[i];
                        statement.forEachChild(child => {
                            const range = _getStartEnd(child);
                            statementRange ??= range;
                            statementRange.end = range.end;
                        });
                        break;
                    }
                }
                if (!statementRange) {
                    statementRange = _getStartEnd(node);
                }
                props.define = {
                    ...parseDefineFunction(node),
                    statement: statementRange,
                };
                if (ts.isVariableDeclaration(parent)) {
                    props.name = parent.name.getText(ast);
                }
                if (node.arguments.length) {
                    props.define.arg = _getStartEnd(node.arguments[0]);
                }
                if (node.typeArguments?.length) {
                    props.define.typeArg = _getStartEnd(node.typeArguments[0]);
                }
            }
            else if (vueCompilerOptions.macros.withDefaults.includes(callText)) {
                props.withDefaults = _getStartEnd(node);
                if (node.arguments.length >= 2) {
                    const arg = node.arguments[1];
                    props.withDefaults.arg = _getStartEnd(arg);
                }
                if (ts.isVariableDeclaration(parent)) {
                    props.name = parent.name.getText(ast);
                }
            }
        }
        node.forEachChild(child => {
            parents.push(node);
            visitNode(child, parents);
            parents.pop();
        });
    }
}
exports.parseScriptSetupRanges = parseScriptSetupRanges;
function parseBindingRanges(ts, sourceFile) {
    const bindings = [];
    sourceFile.forEachChild(node => {
        if (ts.isVariableStatement(node)) {
            for (const node_2 of node.declarationList.declarations) {
                const vars = _findBindingVars(node_2.name);
                for (const _var of vars) {
                    bindings.push(_var);
                }
            }
        }
        else if (ts.isFunctionDeclaration(node)) {
            if (node.name && ts.isIdentifier(node.name)) {
                bindings.push(_getStartEnd(node.name));
            }
        }
        else if (ts.isClassDeclaration(node)) {
            if (node.name) {
                bindings.push(_getStartEnd(node.name));
            }
        }
        else if (ts.isEnumDeclaration(node)) {
            bindings.push(_getStartEnd(node.name));
        }
        if (ts.isImportDeclaration(node)) {
            if (node.importClause && !node.importClause.isTypeOnly) {
                if (node.importClause.name) {
                    bindings.push(_getStartEnd(node.importClause.name));
                }
                if (node.importClause.namedBindings) {
                    if (ts.isNamedImports(node.importClause.namedBindings)) {
                        for (const element of node.importClause.namedBindings.elements) {
                            bindings.push(_getStartEnd(element.name));
                        }
                    }
                    else if (ts.isNamespaceImport(node.importClause.namedBindings)) {
                        bindings.push(_getStartEnd(node.importClause.namedBindings.name));
                    }
                }
            }
        }
    });
    return bindings;
    function _getStartEnd(node) {
        return getStartEnd(node, sourceFile);
    }
    function _findBindingVars(left) {
        return findBindingVars(ts, left, sourceFile);
    }
}
exports.parseBindingRanges = parseBindingRanges;
function findBindingVars(ts, left, sourceFile) {
    const vars = [];
    worker(left);
    return vars;
    function worker(_node) {
        if (ts.isIdentifier(_node)) {
            vars.push(getStartEnd(_node, sourceFile));
        }
        // { ? } = ...
        // [ ? ] = ...
        else if (ts.isObjectBindingPattern(_node) || ts.isArrayBindingPattern(_node)) {
            for (const property of _node.elements) {
                if (ts.isBindingElement(property)) {
                    worker(property.name);
                }
            }
        }
        // { foo: ? } = ...
        else if (ts.isPropertyAssignment(_node)) {
            worker(_node.initializer);
        }
        // { foo } = ...
        else if (ts.isShorthandPropertyAssignment(_node)) {
            vars.push(getStartEnd(_node.name, sourceFile));
        }
        // { ...? } = ...
        // [ ...? ] = ...
        else if (ts.isSpreadAssignment(_node) || ts.isSpreadElement(_node)) {
            worker(_node.expression);
        }
    }
}
exports.findBindingVars = findBindingVars;
function getStartEnd(node, sourceFile) {
    return {
        start: node.getStart(sourceFile),
        end: node.getEnd(),
    };
}
exports.getStartEnd = getStartEnd;
//# sourceMappingURL=scriptSetupRanges.js.map
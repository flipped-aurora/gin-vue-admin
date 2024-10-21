import { parse } from 'acorn';
import { walk } from 'estree-walker';
import { l as getMagicString } from '../shared/unimport.85ddadbb.mjs';
import 'node:path';
import 'node:process';
import 'pathe';
import 'scule';
import 'magic-string';
import 'mlly';
import 'strip-literal';

async function detectImportsAcorn(code, ctx, options) {
  const s = getMagicString(code);
  const map = await ctx.getImportMap();
  let matchedImports = [];
  const enableAutoImport = options?.autoImport !== false;
  const enableTransformVirtualImports = options?.transformVirtualImports !== false && ctx.options.virtualImports?.length;
  if (enableAutoImport || enableTransformVirtualImports) {
    const ast = parse(s.original, {
      sourceType: "module",
      ecmaVersion: "latest",
      locations: true
    });
    const occurrenceMap = /* @__PURE__ */ new Map();
    const virtualImports = createVirtualImportsAcronWalker(map, ctx.options.virtualImports);
    const scopes = traveseScopes(
      ast,
      enableTransformVirtualImports ? virtualImports.walk : {}
    );
    if (enableAutoImport) {
      const identifiers = new Set(occurrenceMap.keys());
      matchedImports.push(
        ...Array.from(scopes.unmatched).map((name) => {
          const item = map.get(name);
          if (item && !item.disabled)
            return item;
          occurrenceMap.delete(name);
          return null;
        }).filter(Boolean)
      );
      for (const addon of ctx.addons)
        matchedImports = await addon.matchImports?.call(ctx, identifiers, matchedImports) || matchedImports;
    }
    virtualImports.ranges.forEach(([start, end]) => {
      s.remove(start, end);
    });
    matchedImports.push(...virtualImports.imports);
  }
  return {
    s,
    strippedCode: code.toString(),
    matchedImports,
    isCJSContext: false,
    firstOccurrence: 0
    // TODO:
  };
}
function traveseScopes(ast, additionalWalk) {
  const scopes = [];
  let scopeCurrent = void 0;
  const scopesStack = [];
  function pushScope(node) {
    scopeCurrent = {
      node,
      parent: scopeCurrent,
      declarations: /* @__PURE__ */ new Set(),
      references: /* @__PURE__ */ new Set()
    };
    scopes.push(scopeCurrent);
    scopesStack.push(scopeCurrent);
  }
  function popScope(node) {
    const scope = scopesStack.pop();
    if (scope?.node !== node)
      throw new Error("Scope mismatch");
    scopeCurrent = scopesStack[scopesStack.length - 1];
  }
  pushScope(void 0);
  walk(ast, {
    enter(node, parent, prop, index) {
      additionalWalk?.enter?.call(this, node, parent, prop, index);
      switch (node.type) {
        case "ImportSpecifier":
        case "ImportDefaultSpecifier":
        case "ImportNamespaceSpecifier":
          scopeCurrent.declarations.add(node.local.name);
          return;
        case "FunctionDeclaration":
        case "ClassDeclaration":
          if (node.id)
            scopeCurrent.declarations.add(node.id.name);
          return;
        case "VariableDeclarator":
          if (node.id.type === "Identifier") {
            scopeCurrent.declarations.add(node.id.name);
          } else {
            walk(node.id, {
              enter(node2) {
                if (node2.type === "ObjectPattern") {
                  node2.properties.forEach((i) => {
                    if (i.type === "Property" && i.value.type === "Identifier")
                      scopeCurrent.declarations.add(i.value.name);
                    else if (i.type === "RestElement" && i.argument.type === "Identifier")
                      scopeCurrent.declarations.add(i.argument.name);
                  });
                } else if (node2.type === "ArrayPattern") {
                  node2.elements.forEach((i) => {
                    if (i?.type === "Identifier")
                      scopeCurrent.declarations.add(i.name);
                    if (i?.type === "RestElement" && i.argument.type === "Identifier")
                      scopeCurrent.declarations.add(i.argument.name);
                  });
                }
              }
            });
          }
          return;
        case "BlockStatement":
          pushScope(node);
          return;
        case "Identifier":
          switch (parent?.type) {
            case "CallExpression":
              if (parent.callee === node || parent.arguments.includes(node))
                scopeCurrent.references.add(node.name);
              return;
            case "MemberExpression":
              if (parent.object === node)
                scopeCurrent.references.add(node.name);
              return;
            case "VariableDeclarator":
              if (parent.init === node)
                scopeCurrent.references.add(node.name);
              return;
            case "SpreadElement":
              if (parent.argument === node)
                scopeCurrent.references.add(node.name);
              return;
            case "ClassDeclaration":
              if (parent.superClass === node)
                scopeCurrent.references.add(node.name);
              return;
            case "Property":
              if (parent.value === node)
                scopeCurrent.references.add(node.name);
              return;
            case "TemplateLiteral":
              if (parent.expressions.includes(node))
                scopeCurrent.references.add(node.name);
              return;
            case "AssignmentExpression":
              if (parent.right === node)
                scopeCurrent.references.add(node.name);
              return;
            case "IfStatement":
            case "WhileStatement":
            case "DoWhileStatement":
              if (parent.test === node)
                scopeCurrent.references.add(node.name);
              return;
            case "SwitchStatement":
              if (parent.discriminant === node)
                scopeCurrent.references.add(node.name);
              return;
          }
          if (parent?.type.includes("Expression"))
            scopeCurrent.references.add(node.name);
      }
    },
    leave(node, parent, prop, index) {
      additionalWalk?.leave?.call(this, node, parent, prop, index);
      switch (node.type) {
        case "BlockStatement":
          popScope(node);
      }
    }
  });
  const unmatched = /* @__PURE__ */ new Set();
  for (const scope of scopes) {
    for (const name of scope.references) {
      let defined = false;
      let parent = scope;
      while (parent) {
        if (parent.declarations.has(name)) {
          defined = true;
          break;
        }
        parent = parent?.parent;
      }
      if (!defined)
        unmatched.add(name);
    }
  }
  return {
    unmatched,
    scopes
  };
}
function createVirtualImportsAcronWalker(importMap, virtualImports = []) {
  const imports = [];
  const ranges = [];
  return {
    imports,
    ranges,
    walk: {
      enter(node) {
        if (node.type === "ImportDeclaration") {
          if (virtualImports.includes(node.source.value)) {
            ranges.push([node.start, node.end]);
            node.specifiers.forEach((i) => {
              if (i.type === "ImportSpecifier" && i.imported.type === "Identifier") {
                const original = importMap.get(i.imported.name);
                if (!original)
                  throw new Error(`[unimport] failed to find "${i.imported.name}" imported from "${node.source.value}"`);
                imports.push({
                  from: original.from,
                  name: original.name,
                  as: i.local.name
                });
              }
            });
          }
        }
      }
    }
  };
}

export { createVirtualImportsAcronWalker, detectImportsAcorn, traveseScopes };

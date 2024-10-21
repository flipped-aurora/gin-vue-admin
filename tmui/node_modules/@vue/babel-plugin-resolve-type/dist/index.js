"use strict";
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

// src/index.ts
var src_exports = {};
__export(src_exports, {
  default: () => src_default
});
module.exports = __toCommonJS(src_exports);
var import_parser = require("@babel/parser");
var import_compiler_sfc = require("@vue/compiler-sfc");
var import_code_frame = require("@babel/code-frame");
var import_helper_module_imports = require("@babel/helper-module-imports");
var import_helper_plugin_utils = require("@babel/helper-plugin-utils");
var src_default = (0, import_helper_plugin_utils.declare)(({ types: t }, options) => {
  let ctx;
  let helpers;
  return {
    name: "babel-plugin-resolve-type",
    pre(file) {
      const filename = file.opts.filename || "unknown.js";
      helpers = /* @__PURE__ */ new Set();
      ctx = {
        filename,
        source: file.code,
        options,
        ast: file.ast.program.body,
        isCE: false,
        error(msg, node) {
          throw new Error(
            `[@vue/babel-plugin-resolve-type] ${msg}

${filename}
${(0, import_code_frame.codeFrameColumns)(
              file.code,
              {
                start: {
                  line: node.loc.start.line,
                  column: node.loc.start.column + 1
                },
                end: {
                  line: node.loc.end.line,
                  column: node.loc.end.column + 1
                }
              }
            )}`
          );
        },
        helper(key) {
          helpers.add(key);
          return `_${key}`;
        },
        getString(node) {
          return file.code.slice(node.start, node.end);
        },
        propsTypeDecl: void 0,
        propsRuntimeDefaults: void 0,
        propsDestructuredBindings: {},
        emitsTypeDecl: void 0
      };
    },
    visitor: {
      CallExpression(path) {
        if (!ctx) {
          throw new Error(
            "[@vue/babel-plugin-resolve-type] context is not loaded."
          );
        }
        const { node } = path;
        if (!t.isIdentifier(node.callee, { name: "defineComponent" })) return;
        if (!checkDefineComponent(path)) return;
        const comp = node.arguments[0];
        if (!comp || !t.isFunction(comp)) return;
        let options2 = node.arguments[1];
        if (!options2) {
          options2 = t.objectExpression([]);
          node.arguments.push(options2);
        }
        node.arguments[1] = processProps(comp, options2) || options2;
        node.arguments[1] = processEmits(comp, node.arguments[1]) || options2;
      },
      VariableDeclarator(path) {
        inferComponentName(path);
      }
    },
    post(file) {
      for (const helper of helpers) {
        (0, import_helper_module_imports.addNamed)(file.path, `_${helper}`, "vue");
      }
    }
  };
  function inferComponentName(path) {
    var _a;
    const id = path.get("id");
    const init = path.get("init");
    if (!id || !id.isIdentifier() || !init || !init.isCallExpression()) return;
    if (!((_a = init.get("callee")) == null ? void 0 : _a.isIdentifier({ name: "defineComponent" }))) return;
    if (!checkDefineComponent(init)) return;
    const nameProperty = t.objectProperty(
      t.identifier("name"),
      t.stringLiteral(id.node.name)
    );
    const { arguments: args } = init.node;
    if (args.length === 0) return;
    if (args.length === 1) {
      init.node.arguments.push(t.objectExpression([]));
    }
    args[1] = addProperty(t, args[1], nameProperty);
  }
  function processProps(comp, options2) {
    const props = comp.params[0];
    if (!props) return;
    if (props.type === "AssignmentPattern") {
      ctx.propsTypeDecl = getTypeAnnotation(props.left);
      ctx.propsRuntimeDefaults = props.right;
    } else {
      ctx.propsTypeDecl = getTypeAnnotation(props);
    }
    if (!ctx.propsTypeDecl) return;
    const runtimeProps = (0, import_compiler_sfc.extractRuntimeProps)(ctx);
    if (!runtimeProps) {
      return;
    }
    const ast = (0, import_parser.parseExpression)(runtimeProps);
    return addProperty(
      t,
      options2,
      t.objectProperty(t.identifier("props"), ast)
    );
  }
  function processEmits(comp, options2) {
    var _a;
    const setupCtx = comp.params[1] && getTypeAnnotation(comp.params[1]);
    if (!setupCtx || !t.isTSTypeReference(setupCtx) || !t.isIdentifier(setupCtx.typeName, { name: "SetupContext" }))
      return;
    const emitType = (_a = setupCtx.typeParameters) == null ? void 0 : _a.params[0];
    if (!emitType) return;
    ctx.emitsTypeDecl = emitType;
    const runtimeEmits = (0, import_compiler_sfc.extractRuntimeEmits)(ctx);
    const ast = t.arrayExpression(
      Array.from(runtimeEmits).map((e) => t.stringLiteral(e))
    );
    return addProperty(
      t,
      options2,
      t.objectProperty(t.identifier("emits"), ast)
    );
  }
});
function getTypeAnnotation(node) {
  if ("typeAnnotation" in node && node.typeAnnotation && node.typeAnnotation.type === "TSTypeAnnotation") {
    return node.typeAnnotation.typeAnnotation;
  }
}
function checkDefineComponent(path) {
  var _a;
  const defineCompImport = (_a = path.scope.getBinding("defineComponent")) == null ? void 0 : _a.path.parent;
  if (!defineCompImport) return true;
  return defineCompImport.type === "ImportDeclaration" && /^@?vue(\/|$)/.test(defineCompImport.source.value);
}
function addProperty(t, object, property) {
  if (t.isObjectExpression(object)) {
    object.properties.unshift(property);
  } else if (t.isExpression(object)) {
    return t.objectExpression([property, t.spreadElement(object)]);
  }
  return object;
}

// src/index.ts
import { parseExpression } from "@babel/parser";
import {
  extractRuntimeEmits,
  extractRuntimeProps
} from "@vue/compiler-sfc";
import { codeFrameColumns } from "@babel/code-frame";
import { addNamed } from "@babel/helper-module-imports";
import { declare } from "@babel/helper-plugin-utils";
var src_default = declare(({ types: t }, options) => {
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
${codeFrameColumns(
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
        addNamed(file.path, `_${helper}`, "vue");
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
    const runtimeProps = extractRuntimeProps(ctx);
    if (!runtimeProps) {
      return;
    }
    const ast = parseExpression(runtimeProps);
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
    const runtimeEmits = extractRuntimeEmits(ctx);
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
export {
  src_default as default
};

import { createHash } from 'node:crypto';
import path from 'node:path';
import * as babel from '@babel/core';
import jsx from '@vue/babel-plugin-jsx';
import { createFilter, normalizePath } from 'vite';

const ssrRegisterHelperId = "/__vue-jsx-ssr-register-helper";
const ssrRegisterHelperCode = `import { useSSRContext } from "vue"
export ${ssrRegisterHelper.toString()}`;
function ssrRegisterHelper(comp, filename) {
  const setup = comp.setup;
  comp.setup = (props, ctx) => {
    const ssrContext = useSSRContext();
    (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add(filename);
    if (setup) {
      return setup(props, ctx);
    }
  };
}
function vueJsxPlugin(options = {}) {
  let root = "";
  let needHmr = false;
  let needSourceMap = true;
  const { include, exclude, babelPlugins = [], ...babelPluginOptions } = options;
  const filter = createFilter(include || /\.[jt]sx$/, exclude);
  return {
    name: "vite:vue-jsx",
    config(config) {
      return {
        // only apply esbuild to ts files
        // since we are handling jsx and tsx now
        esbuild: {
          include: /\.ts$/
        },
        define: {
          __VUE_OPTIONS_API__: config.define?.__VUE_OPTIONS_API__ ?? true,
          __VUE_PROD_DEVTOOLS__: config.define?.__VUE_PROD_DEVTOOLS__ ?? false
        }
      };
    },
    configResolved(config) {
      needHmr = config.command === "serve" && !config.isProduction;
      needSourceMap = config.command === "serve" || !!config.build.sourcemap;
      root = config.root;
    },
    resolveId(id) {
      if (id === ssrRegisterHelperId) {
        return id;
      }
    },
    load(id) {
      if (id === ssrRegisterHelperId) {
        return ssrRegisterHelperCode;
      }
    },
    async transform(code, id, opt) {
      const ssr = opt?.ssr === true;
      const [filepath] = id.split("?");
      if (filter(id) || filter(filepath)) {
        const plugins = [[jsx, babelPluginOptions], ...babelPlugins];
        if (id.endsWith(".tsx") || filepath.endsWith(".tsx")) {
          plugins.push([
            // @ts-ignore missing type
            await import('@babel/plugin-transform-typescript').then(
              (r) => r.default
            ),
            // @ts-ignore
            { isTSX: true, allowExtensions: true }
          ]);
        }
        if (!ssr && !needHmr) {
          plugins.push(() => {
            return {
              visitor: {
                CallExpression: {
                  enter(_path) {
                    if (isDefineComponentCall(_path.node)) {
                      const callee = _path.node.callee;
                      callee.name = `/* @__PURE__ */ ${callee.name}`;
                    }
                  }
                }
              }
            };
          });
        }
        const result = babel.transformSync(code, {
          babelrc: false,
          ast: true,
          plugins,
          sourceMaps: needSourceMap,
          sourceFileName: id,
          configFile: false
        });
        if (!ssr && !needHmr) {
          if (!result.code)
            return;
          return {
            code: result.code,
            map: result.map
          };
        }
        const declaredComponents = [];
        const hotComponents = [];
        let hasDefault = false;
        for (const node of result.ast.program.body) {
          if (node.type === "VariableDeclaration") {
            const names = parseComponentDecls(node);
            if (names.length) {
              declaredComponents.push(...names);
            }
          }
          if (node.type === "ExportNamedDeclaration") {
            if (node.declaration && node.declaration.type === "VariableDeclaration") {
              hotComponents.push(
                ...parseComponentDecls(node.declaration).map((name) => ({
                  local: name,
                  exported: name,
                  id: getHash(id + name)
                }))
              );
            } else if (node.specifiers.length) {
              for (const spec of node.specifiers) {
                if (spec.type === "ExportSpecifier" && spec.exported.type === "Identifier") {
                  const matched = declaredComponents.find(
                    (name) => name === spec.local.name
                  );
                  if (matched) {
                    hotComponents.push({
                      local: spec.local.name,
                      exported: spec.exported.name,
                      id: getHash(id + spec.exported.name)
                    });
                  }
                }
              }
            }
          }
          if (node.type === "ExportDefaultDeclaration") {
            if (node.declaration.type === "Identifier") {
              const _name = node.declaration.name;
              const matched = declaredComponents.find((name) => name === _name);
              if (matched) {
                hotComponents.push({
                  local: _name,
                  exported: "default",
                  id: getHash(id + "default")
                });
              }
            } else if (isDefineComponentCall(node.declaration)) {
              hasDefault = true;
              hotComponents.push({
                local: "__default__",
                exported: "default",
                id: getHash(id + "default")
              });
            }
          }
        }
        if (hotComponents.length) {
          if (hasDefault && (needHmr || ssr)) {
            result.code = result.code.replace(
              /export default defineComponent/g,
              `const __default__ = defineComponent`
            ) + `
export default __default__`;
          }
          if (needHmr && !ssr && !/\?vue&type=script/.test(id)) {
            let code2 = result.code;
            let callbackCode = ``;
            for (const { local, exported, id: id2 } of hotComponents) {
              code2 += `
${local}.__hmrId = "${id2}"
__VUE_HMR_RUNTIME__.createRecord("${id2}", ${local})`;
              callbackCode += `
__VUE_HMR_RUNTIME__.reload("${id2}", __${exported})`;
            }
            const newCompNames = hotComponents.map((c) => `${c.exported}: __${c.exported}`).join(",");
            code2 += `
import.meta.hot.accept(({${newCompNames}}) => {${callbackCode}
})`;
            result.code = code2;
          }
          if (ssr) {
            const normalizedId = normalizePath(path.relative(root, id));
            let ssrInjectCode = `
import { ssrRegisterHelper } from "${ssrRegisterHelperId}"
const __moduleId = ${JSON.stringify(normalizedId)}`;
            for (const { local } of hotComponents) {
              ssrInjectCode += `
ssrRegisterHelper(${local}, __moduleId)`;
            }
            result.code += ssrInjectCode;
          }
        }
        if (!result.code)
          return;
        return {
          code: result.code,
          map: result.map
        };
      }
    }
  };
}
function parseComponentDecls(node) {
  const names = [];
  for (const decl of node.declarations) {
    if (decl.id.type === "Identifier" && isDefineComponentCall(decl.init)) {
      names.push(decl.id.name);
    }
  }
  return names;
}
function isDefineComponentCall(node) {
  return node && node.type === "CallExpression" && node.callee.type === "Identifier" && node.callee.name === "defineComponent";
}
function getHash(text) {
  return createHash("sha256").update(text).digest("hex").substring(0, 8);
}

export { vueJsxPlugin as default };

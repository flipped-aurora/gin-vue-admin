"use strict";Object.defineProperty(exports, "__esModule", {value: true}); function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }


var _chunk6UOGCAOYcjs = require('./chunk-6UOGCAOY.cjs');

// src/core/unplugin.ts
var _minimatch = require('minimatch');
var _utils = require('@antfu/utils');
var _unplugin = require('unplugin');

// src/core/ctx.ts
var _path = require('path');
var _fs = require('fs');
var _process = require('process'); var _process2 = _interopRequireDefault(_process);

var _pluginutils = require('@rollup/pluginutils');
var _localpkg = require('local-pkg');
var _unimport = require('unimport');
var _fastglob = require('fast-glob'); var _fastglob2 = _interopRequireDefault(_fastglob);
var _addons = require('unimport/addons');
var _magicstring = require('magic-string'); var _magicstring2 = _interopRequireDefault(_magicstring);

// src/core/eslintrc.ts
function generateESLintConfigs(imports, eslintrc, globals = {}) {
  const eslintConfigs = { globals };
  imports.map((i) => {
    var _a;
    return (_a = i.as) != null ? _a : i.name;
  }).filter(Boolean).sort().forEach((name) => {
    eslintConfigs.globals[name] = eslintrc.globalsPropValue;
  });
  const jsonBody = JSON.stringify(eslintConfigs, null, 2);
  return jsonBody;
}

// src/core/resolvers.ts

function normalizeImport(info, name) {
  if (typeof info === "string") {
    return {
      name: "default",
      as: name,
      from: info
    };
  }
  if ("path" in info) {
    return {
      from: info.path,
      as: info.name,
      name: info.importName,
      sideEffects: info.sideEffects
    };
  }
  return _chunk6UOGCAOYcjs.__spreadValues.call(void 0, {
    name,
    as: name
  }, info);
}
async function firstMatchedResolver(resolvers, fullname) {
  let name = fullname;
  for (const resolver of resolvers) {
    if (typeof resolver === "object" && resolver.type === "directive") {
      if (name.startsWith("v"))
        name = name.slice(1);
      else
        continue;
    }
    const resolved = await (typeof resolver === "function" ? resolver(name) : resolver.resolve(name));
    if (resolved)
      return normalizeImport(resolved, fullname);
  }
}
function resolversAddon(resolvers) {
  return {
    async matchImports(names, matched) {
      if (!resolvers.length)
        return;
      const dynamic = [];
      const sideEffects = [];
      await Promise.all([...names].map(async (name) => {
        const matchedImport = matched.find((i) => i.as === name);
        if (matchedImport) {
          if ("sideEffects" in matchedImport)
            sideEffects.push(..._utils.toArray.call(void 0, matchedImport.sideEffects).map((i) => normalizeImport(i, "")));
          return;
        }
        const resolved = await firstMatchedResolver(resolvers, name);
        if (resolved)
          dynamic.push(resolved);
        if (resolved == null ? void 0 : resolved.sideEffects)
          sideEffects.push(..._utils.toArray.call(void 0, resolved == null ? void 0 : resolved.sideEffects).map((i) => normalizeImport(i, "")));
      }));
      if (dynamic.length) {
        this.dynamicImports.push(...dynamic);
        this.invalidate();
      }
      if (dynamic.length || sideEffects.length)
        return [...matched, ...dynamic, ...sideEffects];
    }
  };
}

// src/core/ctx.ts
function resolveGlobsExclude(root, glob) {
  const excludeReg = /^!/;
  return `${excludeReg.test(glob) ? "!" : ""}${_path.resolve.call(void 0, root, glob.replace(excludeReg, ""))}`;
}
async function scanDirExports(dirs, root) {
  const result = await _fastglob2.default.call(void 0, dirs, {
    absolute: true,
    cwd: root,
    onlyFiles: true,
    followSymbolicLinks: true
  });
  const files = Array.from(new Set(result.flat())).map(_utils.slash);
  return (await Promise.all(files.map((i) => _unimport.scanExports.call(void 0, i, false)))).flat();
}
function createContext(options = {}, root = _process2.default.cwd()) {
  var _a;
  const {
    dts: preferDTS = _localpkg.isPackageExists.call(void 0, "typescript")
  } = options;
  const dirs = (_a = options.dirs) == null ? void 0 : _a.concat(options.dirs.map((dir) => _path.join.call(void 0, dir, "*.{tsx,jsx,ts,js,mjs,cjs,mts,cts}"))).map((dir) => _utils.slash.call(void 0, resolveGlobsExclude(root, dir)));
  const eslintrc = options.eslintrc || {};
  eslintrc.enabled = eslintrc.enabled === void 0 ? false : eslintrc.enabled;
  eslintrc.filepath = eslintrc.filepath || "./.eslintrc-auto-import.json";
  eslintrc.globalsPropValue = eslintrc.globalsPropValue === void 0 ? true : eslintrc.globalsPropValue;
  const resolvers = options.resolvers ? [options.resolvers].flat(2) : [];
  const injectAtEnd = options.injectAtEnd !== false;
  const unimport = _unimport.createUnimport.call(void 0, {
    imports: [],
    presets: [],
    injectAtEnd,
    addons: [
      ...options.vueTemplate ? [_addons.vueTemplateAddon.call(void 0, )] : [],
      resolversAddon(resolvers),
      {
        declaration(dts2) {
          return `${`
/* eslint-disable */
/* prettier-ignore */
// @ts-nocheck
// noinspection JSUnusedGlobalSymbols
// Generated by unplugin-auto-import
${dts2}`.trim()}
`;
        }
      }
    ]
  });
  const importsPromise = flattenImports(options.imports).then((imports) => {
    var _a2;
    if (!imports.length && !resolvers.length && !(dirs == null ? void 0 : dirs.length))
      console.warn("[auto-import] plugin installed but no imports has defined, see https://github.com/antfu/unplugin-auto-import#configurations for configurations");
    (_a2 = options.ignore) == null ? void 0 : _a2.forEach((name) => {
      const i = imports.find((i2) => i2.as === name);
      if (i)
        i.disabled = true;
    });
    return unimport.getInternalContext().replaceImports(imports);
  });
  const filter = _pluginutils.createFilter.call(void 0, 
    options.include || [/\.[jt]sx?$/, /\.vue$/, /\.vue\?vue/, /\.svelte$/],
    options.exclude || [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/]
  );
  const dts = preferDTS === false ? false : preferDTS === true ? _path.resolve.call(void 0, root, "auto-imports.d.ts") : _path.resolve.call(void 0, root, preferDTS);
  const multilineCommentsRE = new RegExp("\\/\\*.*?\\*\\/", "gms");
  const singlelineCommentsRE = /\/\/.*$/gm;
  const dtsReg = new RegExp("declare\\s+global\\s*{(.*?)[\\n\\r]}", "s");
  function parseDTS(dts2) {
    var _a2;
    dts2 = dts2.replace(multilineCommentsRE, "").replace(singlelineCommentsRE, "");
    const code = (_a2 = dts2.match(dtsReg)) == null ? void 0 : _a2[0];
    if (!code)
      return;
    return Object.fromEntries(Array.from(code.matchAll(/['"]?(const\s*[^\s'"]+)['"]?\s*:\s*(.+?)[,;\r\n]/g)).map((i) => [i[1], i[2]]));
  }
  async function generateDTS(file) {
    await importsPromise;
    const dir = _path.dirname.call(void 0, file);
    const originalContent = _fs.existsSync.call(void 0, file) ? await _fs.promises.readFile(file, "utf-8") : "";
    const originalDTS = parseDTS(originalContent);
    const currentContent = await unimport.generateTypeDeclarations({
      resolvePath: (i) => {
        if (i.from.startsWith(".") || _path.isAbsolute.call(void 0, i.from)) {
          const related = _utils.slash.call(void 0, _path.relative.call(void 0, dir, i.from).replace(/\.ts(x)?$/, ""));
          return !related.startsWith(".") ? `./${related}` : related;
        }
        return i.from;
      }
    });
    const currentDTS = parseDTS(currentContent);
    if (originalDTS) {
      Object.keys(currentDTS).forEach((key) => {
        originalDTS[key] = currentDTS[key];
      });
      const dtsList = Object.keys(originalDTS).sort().map((k) => `  ${k}: ${originalDTS[k]}`);
      return currentContent.replace(dtsReg, () => `declare global {
${dtsList.join("\n")}
}`);
    }
    return currentContent;
  }
  async function parseESLint() {
    const configStr = _fs.existsSync.call(void 0, eslintrc.filepath) ? await _fs.promises.readFile(eslintrc.filepath, "utf-8") : "";
    const config = JSON.parse(configStr || '{ "globals": {} }');
    return config.globals;
  }
  async function generateESLint() {
    return generateESLintConfigs(await unimport.getImports(), eslintrc, await parseESLint());
  }
  const writeConfigFilesThrottled = _utils.throttle.call(void 0, 500, writeConfigFiles, { noLeading: false });
  async function writeFile(filePath, content = "") {
    await _fs.promises.mkdir(_path.dirname.call(void 0, filePath), { recursive: true });
    return await _fs.promises.writeFile(filePath, content, "utf-8");
  }
  let lastDTS;
  let lastESLint;
  async function writeConfigFiles() {
    const promises = [];
    if (dts) {
      promises.push(
        generateDTS(dts).then((content) => {
          if (content !== lastDTS) {
            lastDTS = content;
            return writeFile(dts, content);
          }
        })
      );
    }
    if (eslintrc.enabled && eslintrc.filepath) {
      promises.push(
        generateESLint().then((content) => {
          content = `${content}
`;
          if (content.trim() !== (lastESLint == null ? void 0 : lastESLint.trim())) {
            lastESLint = content;
            return writeFile(eslintrc.filepath, content);
          }
        })
      );
    }
    return Promise.all(promises);
  }
  async function scanDirs() {
    if (dirs == null ? void 0 : dirs.length) {
      await unimport.modifyDynamicImports(async (imports) => {
        const exports_ = await scanDirExports(dirs, root);
        exports_.forEach((i) => i.__source = "dir");
        return modifyDefaultExportsAlias([
          ...imports.filter((i) => i.__source !== "dir"),
          ...exports_
        ], options);
      });
    }
    writeConfigFilesThrottled();
  }
  async function transform(code, id) {
    await importsPromise;
    const s = new (0, _magicstring2.default)(code);
    await unimport.injectImports(s, id);
    if (!s.hasChanged())
      return;
    writeConfigFilesThrottled();
    return {
      code: s.toString(),
      map: s.generateMap({ source: id, includeContent: true, hires: true })
    };
  }
  return {
    root,
    dirs,
    filter,
    scanDirs,
    writeConfigFiles,
    writeConfigFilesThrottled,
    transform,
    generateDTS,
    generateESLint
  };
}
async function flattenImports(map) {
  const promises = await Promise.all(_utils.toArray.call(void 0, map).map(async (definition) => {
    if (typeof definition === "string") {
      if (!_chunk6UOGCAOYcjs.presets[definition])
        throw new Error(`[auto-import] preset ${definition} not found`);
      const preset = _chunk6UOGCAOYcjs.presets[definition];
      definition = typeof preset === "function" ? preset() : preset;
    }
    if ("from" in definition && "imports" in definition) {
      return await _unimport.resolvePreset.call(void 0, definition);
    } else {
      const resolved = [];
      for (const mod of Object.keys(definition)) {
        for (const id of definition[mod]) {
          const meta = {
            from: mod
          };
          if (Array.isArray(id)) {
            meta.name = id[0];
            meta.as = id[1];
          } else {
            meta.name = id;
            meta.as = id;
          }
          resolved.push(meta);
        }
      }
      return resolved;
    }
  }));
  return promises.flat();
}
function modifyDefaultExportsAlias(imports, options) {
  if (options.defaultExportByFilename) {
    imports.forEach((i) => {
      var _a, _b, _c;
      if (i.name === "default")
        i.as = (_c = (_b = (_a = i.from.split("/").pop()) == null ? void 0 : _a.split(".")) == null ? void 0 : _b.shift()) != null ? _c : i.as;
    });
  }
  return imports;
}

// src/core/unplugin.ts
var unplugin_default = _unplugin.createUnplugin.call(void 0, (options) => {
  let ctx = createContext(options);
  return {
    name: "unplugin-auto-import",
    enforce: "post",
    transformInclude(id) {
      return ctx.filter(id);
    },
    async transform(code, id) {
      return ctx.transform(code, id);
    },
    async buildStart() {
      await ctx.scanDirs();
    },
    async buildEnd() {
      await ctx.writeConfigFiles();
    },
    vite: {
      async handleHotUpdate({ file }) {
        var _a;
        if ((_a = ctx.dirs) == null ? void 0 : _a.some((glob) => _minimatch.minimatch.call(void 0, _utils.slash.call(void 0, file), _utils.slash.call(void 0, glob))))
          await ctx.scanDirs();
      },
      async configResolved(config) {
        if (ctx.root !== config.root) {
          ctx = createContext(options, config.root);
          await ctx.scanDirs();
        }
      }
    }
  };
});



exports.unplugin_default = unplugin_default;

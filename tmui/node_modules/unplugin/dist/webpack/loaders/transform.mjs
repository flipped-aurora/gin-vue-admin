// src/webpack/context.ts
import { resolve } from "path";
import { Buffer } from "buffer";
import process from "process";
import { createRequire } from "module";
import { Parser } from "acorn";
function createBuildContext(options, compiler, compilation, loaderContext) {
  const require2 = createRequire(import.meta.url);
  const sources = require2("webpack-sources");
  return {
    parse(code, opts = {}) {
      return Parser.parse(code, {
        sourceType: "module",
        ecmaVersion: "latest",
        locations: true,
        ...opts
      });
    },
    addWatchFile(id) {
      options.addWatchFile(resolve(process.cwd(), id));
    },
    emitFile(emittedFile) {
      const outFileName = emittedFile.fileName || emittedFile.name;
      if (emittedFile.source && outFileName) {
        if (!compilation)
          throw new Error("unplugin/webpack: emitFile outside supported hooks  (buildStart, buildEnd, load, transform, watchChange)");
        compilation.emitAsset(
          outFileName,
          sources ? new sources.RawSource(
            // @ts-expect-error types mismatch
            typeof emittedFile.source === "string" ? emittedFile.source : Buffer.from(emittedFile.source)
          ) : {
            source: () => emittedFile.source,
            size: () => emittedFile.source.length
          }
        );
      }
    },
    getWatchFiles() {
      return options.getWatchFiles();
    },
    getNativeBuildContext() {
      return { framework: "webpack", compiler, compilation, loaderContext };
    }
  };
}
function createContext(loader) {
  return {
    error: (error) => loader.emitError(normalizeMessage(error)),
    warn: (message) => loader.emitWarning(normalizeMessage(message))
  };
}
function normalizeMessage(error) {
  const err = new Error(typeof error === "string" ? error : error.message);
  if (typeof error === "object") {
    err.stack = error.stack;
    err.cause = error.meta;
  }
  return err;
}

// src/utils.ts
function resolveQuery(query) {
  if (typeof query === "string") {
    return new URLSearchParams(query).get("unpluginName");
  } else {
    return query.unpluginName;
  }
}

// src/webpack/loaders/transform.ts
async function transform(source, map) {
  var _a;
  const callback = this.async();
  const unpluginName = resolveQuery(this.query);
  const plugin = (_a = this._compiler) == null ? void 0 : _a.$unpluginContext[unpluginName];
  if (!(plugin == null ? void 0 : plugin.transform))
    return callback(null, source, map);
  const context = createContext(this);
  const res = await plugin.transform.call(
    Object.assign({}, createBuildContext({
      addWatchFile: (file) => {
        this.addDependency(file);
      },
      getWatchFiles: () => {
        return this.getDependencies();
      }
    }, this._compiler, this._compilation, this), context),
    source,
    this.resource
  );
  if (res == null)
    callback(null, source, map);
  else if (typeof res !== "string")
    callback(null, res.code, map == null ? map : res.map || map);
  else
    callback(null, res, map);
}
export {
  transform as default
};

"use strict";Object.defineProperty(exports, "__esModule", {value: true});

var _chunkERF3N54Scjs = require('./chunk-ERF3N54S.cjs');
require('./chunk-6UOGCAOY.cjs');

// src/astro.ts
function astro_default(options) {
  return {
    name: "unplugin-auto-import",
    hooks: {
      "astro:config:setup": async (astro) => {
        var _a;
        (_a = astro.config.vite).plugins || (_a.plugins = []);
        astro.config.vite.plugins.push(_chunkERF3N54Scjs.unplugin_default.vite(options));
      }
    }
  };
}


module.exports = astro_default;
exports.default = module.exports;
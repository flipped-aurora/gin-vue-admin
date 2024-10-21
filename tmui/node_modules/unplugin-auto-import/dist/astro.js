import {
  unplugin_default
} from "./chunk-6AAI2DNE.js";
import "./chunk-EZINZJYF.js";

// src/astro.ts
function astro_default(options) {
  return {
    name: "unplugin-auto-import",
    hooks: {
      "astro:config:setup": async (astro) => {
        var _a;
        (_a = astro.config.vite).plugins || (_a.plugins = []);
        astro.config.vite.plugins.push(unplugin_default.vite(options));
      }
    }
  };
}
export {
  astro_default as default
};

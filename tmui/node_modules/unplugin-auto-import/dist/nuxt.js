import {
  unplugin_default
} from "./chunk-6AAI2DNE.js";
import "./chunk-EZINZJYF.js";

// src/nuxt.ts
import { addVitePlugin, addWebpackPlugin, defineNuxtModule } from "@nuxt/kit";
var nuxt_default = defineNuxtModule({
  setup(options) {
    options.exclude = options.exclude || [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/, /[\\/]\.nuxt[\\/]/];
    addWebpackPlugin(unplugin_default.webpack(options));
    addVitePlugin(unplugin_default.vite(options));
  }
});
export {
  nuxt_default as default
};

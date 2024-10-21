import type { Plugin } from 'vite';
declare module '@vue/compiler-sfc' {
    interface SFCDescriptor {
        id: string;
    }
}
export declare const APP_CSS_JS = "./app.css.js";
export declare function uniAppCssPlugin(): Plugin;

import type { Plugin } from 'vite';
export declare function initNVueNodeTransforms(): import("@vue/compiler-core").NodeTransform[];
export declare function initNVueDirectiveTransforms(): {
    model: import("@vue/compiler-core").DirectiveTransform;
    show: import("@vue/compiler-core").DirectiveTransform;
};
export declare function uniAppNVuePlugin({ appService, }: {
    appService: boolean;
}): Plugin;

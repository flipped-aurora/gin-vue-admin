export declare function initNVue(manifestJson: Record<string, any>, pagesJson: UniApp.PagesJson): void;
export declare function getNVueCompiler(manifestJson: Record<string, any>): "uni-app" | "vue" | "weex" | "vite";
export declare function getNVueStyleCompiler(manifestJson: Record<string, any>): "uni-app" | "weex";
export declare function getNVueFlexDirection(manifestJson: Record<string, any>): "column" | "row" | "row-reverse" | "column-reverse";

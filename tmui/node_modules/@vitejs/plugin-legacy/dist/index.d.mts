import { Plugin } from 'vite';

interface Options {
    /**
     * default: 'defaults'
     */
    targets?: string | string[] | {
        [key: string]: string;
    };
    /**
     * default: 'edge>=80, firefox>=72, chrome>=80, safari>=13.1, chromeAndroid>=80, iOS>=13.1'
     */
    modernTargets?: string | string[];
    /**
     * default: true
     */
    polyfills?: boolean | string[];
    additionalLegacyPolyfills?: string[];
    /**
     * default: false
     */
    modernPolyfills?: boolean | string[];
    /**
     * default: true
     */
    renderLegacyChunks?: boolean;
    /**
     * default: false
     */
    externalSystemJS?: boolean;
    /**
     * default: true
     */
    renderModernChunks?: boolean;
}

declare function viteLegacyPlugin(options?: Options): Plugin[];
declare function detectPolyfills(code: string, targets: any, list: Set<string>): Promise<void>;
declare const cspHashes: string[];

export { type Options, cspHashes, viteLegacyPlugin as default, detectPolyfills };

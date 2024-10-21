import { type NodeTransform } from '@vue/compiler-core';
export interface AssetURLTagConfig {
    [name: string]: string[];
}
export interface AssetURLOptions {
    /**
     * If base is provided, instead of transforming relative asset urls into
     * imports, they will be directly rewritten to absolute urls.
     */
    base?: string | null;
    /**
     * If true, also processes absolute urls.
     */
    includeAbsolute?: boolean;
    tags?: AssetURLTagConfig;
}
export declare const defaultAssetUrlOptions: Required<AssetURLOptions>;
export declare const normalizeOptions: (options: AssetURLOptions | AssetURLTagConfig) => Required<AssetURLOptions>;
export declare const createAssetUrlTransformWithOptions: (options: Required<AssetURLOptions>) => NodeTransform;
/**
 * A `@vue/compiler-core` plugin that transforms relative asset urls into
 * either imports or absolute urls.
 *
 * ``` js
 * // Before
 * createVNode('img', { src: './logo.png' })
 *
 * // After
 * import _imports_0 from './logo.png'
 * createVNode('img', { src: _imports_0 })
 * ```
 */
export declare const transformAssetUrl: NodeTransform;

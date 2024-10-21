import { type NodeTransform } from '@vue/compiler-core';
import { type AssetURLOptions } from './templateTransformAssetUrl';
export declare const createSrcsetTransformWithOptions: (options: Required<AssetURLOptions>) => NodeTransform;
export declare const transformSrcset: NodeTransform;

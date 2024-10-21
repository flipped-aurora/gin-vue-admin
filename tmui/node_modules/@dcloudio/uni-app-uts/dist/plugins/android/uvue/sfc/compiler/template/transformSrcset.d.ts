import { type NodeTransform } from '@vue/compiler-core';
import { type AssetURLOptions } from './transformAssetUrl';
export declare const createSrcsetTransformWithOptions: (options: Required<AssetURLOptions>) => NodeTransform;
export declare const transformSrcset: NodeTransform;

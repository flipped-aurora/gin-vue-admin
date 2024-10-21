import type { SFCDescriptor } from '@vue/compiler-sfc';
export declare function genScript({ script }: SFCDescriptor, { genDefaultAs }: {
    genDefaultAs?: string;
}): string;
export declare function genDefaultScriptCode(genDefaultAs?: string): string;

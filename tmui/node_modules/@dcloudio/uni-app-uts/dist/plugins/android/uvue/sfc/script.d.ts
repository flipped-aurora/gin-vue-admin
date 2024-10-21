import type { SFCDescriptor, SFCScriptBlock } from 'vue/compiler-sfc';
import type { ResolvedOptions } from '.';
export declare function invalidateScript(filename: string): void;
export declare function getResolvedScript(descriptor: SFCDescriptor): SFCScriptBlock | null | undefined;
export declare function setResolvedScript(descriptor: SFCDescriptor, script: SFCScriptBlock): void;
export declare function isUseInlineTemplate(descriptor: SFCDescriptor, isProd: boolean): boolean;
export declare const scriptIdentifier = "__sfc__";
export declare function resolveScript(descriptor: SFCDescriptor, options: ResolvedOptions): SFCScriptBlock | null;

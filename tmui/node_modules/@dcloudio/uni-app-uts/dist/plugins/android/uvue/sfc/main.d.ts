import type { SFCDescriptor } from '@vue/compiler-sfc';
import type { SourceMapInput, TransformPluginContext } from 'rollup';
import type { ResolvedOptions } from './index';
export declare function transformMain(code: string, filename: string, options: ResolvedOptions, pluginContext?: TransformPluginContext): Promise<{
    code: string;
    map: SourceMapInput;
    errors: (SyntaxError | import("@vue/compiler-core").CompilerError)[];
    uts: string;
    descriptor: SFCDescriptor;
} | null>;

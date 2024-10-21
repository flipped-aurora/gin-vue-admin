import { type CopyOptions, type MiniProgramCompilerOptions, type UniVitePlugin } from '@dcloudio/uni-cli-shared';
import type { CompilerOptions } from '@dcloudio/uni-mp-compiler';
export declare function uniOptions({ copyOptions, miniProgram, customElements, compilerOptions, }: {
    customElements?: string[];
    copyOptions: CopyOptions;
    miniProgram: MiniProgramCompilerOptions;
    compilerOptions?: CompilerOptions;
}): UniVitePlugin['uni'];

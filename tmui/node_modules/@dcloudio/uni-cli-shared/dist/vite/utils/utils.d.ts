import type { ConfigEnv, ResolvedConfig, UserConfig } from 'vite';
import type { RollupError } from 'rollup';
import type { CompilerError } from '@vue/compiler-sfc';
import { codeFrameColumns } from '@babel/code-frame';
export declare function withSourcemap(config: ResolvedConfig): boolean;
export declare function isInHybridNVue(config: UserConfig | ResolvedConfig): boolean;
export declare function isSsr(command: ConfigEnv['command'], config: UserConfig | ResolvedConfig): boolean;
export declare function createRollupError(plugin: string, id: string, error: CompilerError | SyntaxError, source?: string): RollupError;
export declare const generateCodeFrameColumns: typeof codeFrameColumns;

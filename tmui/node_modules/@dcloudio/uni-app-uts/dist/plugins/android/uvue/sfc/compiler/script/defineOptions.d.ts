import type { Node } from '@babel/types';
import type { ScriptCompileContext } from './context';
export declare const DEFINE_OPTIONS = "defineOptions";
export declare function processDefineOptions(ctx: ScriptCompileContext, node: Node): boolean;

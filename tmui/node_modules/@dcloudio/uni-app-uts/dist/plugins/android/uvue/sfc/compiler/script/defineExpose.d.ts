import type { Node } from '@babel/types';
import type { ScriptCompileContext } from './context';
export declare const DEFINE_EXPOSE = "defineExpose";
export declare function processDefineExpose(ctx: ScriptCompileContext, node: Node): boolean;

import type { LVal, Node } from '@babel/types';
import type { ScriptCompileContext } from './context';
export declare const DEFINE_EMITS = "defineEmits";
export declare function processDefineEmits(ctx: ScriptCompileContext, node: Node, declId?: LVal): boolean;
export declare function genRuntimeEmits(ctx: ScriptCompileContext): string | undefined;

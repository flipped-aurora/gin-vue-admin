import { type LVal, type Node } from '@babel/types';
import type { ScriptCompileContext } from './context';
export declare const DEFINE_SLOTS = "defineSlots";
export declare function processDefineSlots(ctx: ScriptCompileContext, node: Node, declId?: LVal): boolean;
export declare function genRuntimeSlots(ctx: ScriptCompileContext): string | undefined;

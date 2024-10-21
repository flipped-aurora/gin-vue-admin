import type { LVal, Node, TSType } from '@babel/types';
import type { ScriptCompileContext } from './context';
export declare const DEFINE_MODEL = "defineModel";
export interface ModelDecl {
    type: TSType | undefined;
    options: string | undefined;
    identifier: string | undefined;
}
export declare function processDefineModel(ctx: ScriptCompileContext, node: Node, declId?: LVal): boolean;
export declare function genModelProps(ctx: ScriptCompileContext): string | undefined;

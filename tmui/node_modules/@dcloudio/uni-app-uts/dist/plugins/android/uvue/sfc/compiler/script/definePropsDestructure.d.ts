import type { ObjectPattern } from '@babel/types';
import type { ScriptCompileContext } from './context';
export declare function processPropsDestructure(ctx: ScriptCompileContext, declId: ObjectPattern): void;
export declare function transformDestructuredProps(ctx: ScriptCompileContext, vueImportAliases: Record<string, string>): void;

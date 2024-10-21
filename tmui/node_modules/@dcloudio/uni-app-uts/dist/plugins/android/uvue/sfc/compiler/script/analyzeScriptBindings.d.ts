import type { Node, Statement } from '@babel/types';
import { type BindingMetadata } from '@vue/compiler-dom';
/**
 * Analyze bindings in normal `<script>`
 * Note that `compileScriptSetup` already analyzes bindings as part of its
 * compilation process so this should only be used on single `<script>` SFCs.
 */
export declare function analyzeScriptBindings(ast: Statement[]): BindingMetadata;
export declare function getObjectOrArrayExpressionKeys(value: Node): string[];

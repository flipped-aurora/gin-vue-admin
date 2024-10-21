import { type FunctionExpression, type SourceLocation } from '@vue/compiler-core';
import type { NodeTransform } from '../transform';
export declare const transformMemo: NodeTransform;
export declare function createFunctionExpression(params: FunctionExpression['params'], returns?: FunctionExpression['returns'], newline?: boolean, isSlot?: boolean, loc?: SourceLocation): FunctionExpression;

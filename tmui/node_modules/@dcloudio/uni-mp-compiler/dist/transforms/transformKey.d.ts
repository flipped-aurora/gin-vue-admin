import { type DirectiveNode } from '@vue/compiler-core';
import type { ForElementNode } from './vFor';
export declare function isSelfKey({ arg, exp }: DirectiveNode, vFor: ForElementNode['vFor'] | false): boolean | undefined;
export declare function rewriteSelfKey(dir: DirectiveNode): void;

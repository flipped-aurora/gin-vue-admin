"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getTypesCode = exports.baseName = void 0;
const shared_1 = require("./shared");
exports.baseName = '__VLS_types.d.ts';
function getTypesCode(vueCompilerOptions) {
    return `
// @ts-nocheck

type __VLS_IntrinsicElements = __VLS_PickNotAny<import('vue/jsx-runtime').JSX.IntrinsicElements, __VLS_PickNotAny<JSX.IntrinsicElements, Record<string, any>>>;
type __VLS_Element = __VLS_PickNotAny<import('vue/jsx-runtime').JSX.Element, JSX.Element>;

type __VLS_IsAny<T> = 0 extends 1 & T ? true : false;
type __VLS_PickNotAny<A, B> = __VLS_IsAny<A> extends true ? B : A;

type __VLS_Prettify<T> = { [K in keyof T]: T[K]; } & {};

type __VLS_OmitKeepDiscriminatedUnion<T, K extends keyof any> =
	T extends any
		? Pick<T, Exclude<keyof T, K>>
		: never;

type __VLS_GlobalComponents =
	__VLS_PickNotAny<import('vue').GlobalComponents, {}>
	& __VLS_PickNotAny<import('@vue/runtime-core').GlobalComponents, {}>
	& __VLS_PickNotAny<import('@vue/runtime-dom').GlobalComponents, {}>
	& Pick<typeof import('${vueCompilerOptions.lib}'),
		'Transition'
		| 'TransitionGroup'
		| 'KeepAlive'
		| 'Suspense'
		| 'Teleport'
	>;

declare const __VLS_intrinsicElements: __VLS_IntrinsicElements;

// v-for
declare function __VLS_getVForSourceType(source: number): [number, number, number][];
declare function __VLS_getVForSourceType(source: string): [string, number, number][];
declare function __VLS_getVForSourceType<T extends any[]>(source: T): [
	T[number], // item
	number, // key
	number, // index
][];
declare function __VLS_getVForSourceType<T extends { [Symbol.iterator](): Iterator<any> }>(source: T): [
	T extends { [Symbol.iterator](): Iterator<infer T1> } ? T1 : never, // item 
	number, // key
	undefined, // index
][];
declare function __VLS_getVForSourceType<T>(source: T): [
	T[keyof T], // item
	keyof T, // key
	number, // index
][];

declare function __VLS_getSlotParams<T>(slot: T): Parameters<__VLS_PickNotAny<NonNullable<T>, (...args: any[]) => any>>;
declare function __VLS_getSlotParam<T>(slot: T): Parameters<__VLS_PickNotAny<NonNullable<T>, (...args: any[]) => any>>[0];
declare function __VLS_directiveFunction<T>(dir: T):
	T extends import('${vueCompilerOptions.lib}').ObjectDirective<infer E, infer V> | import('${vueCompilerOptions.lib}').FunctionDirective<infer E, infer V> ? (value: V) => void
	: T;
declare function __VLS_withScope<T, K>(ctx: T, scope: K): ctx is T & K;
declare function __VLS_makeOptional<T>(t: T): { [K in keyof T]?: T[K] };

type __VLS_SelfComponent<N, C> = string extends N ? {} : N extends string ? { [P in N]: C } : {};
type __VLS_WithComponent<N0 extends string, LocalComponents, N1 extends string, N2 extends string, N3 extends string> =
	N1 extends keyof LocalComponents ? N1 extends N0 ? Pick<LocalComponents, N0> : { [K in N0]: LocalComponents[N1] } :
	N2 extends keyof LocalComponents ? N2 extends N0 ? Pick<LocalComponents, N0> : { [K in N0]: LocalComponents[N2] } :
	N3 extends keyof LocalComponents ? N3 extends N0 ? Pick<LocalComponents, N0> : { [K in N0]: LocalComponents[N3] } :
	N1 extends keyof __VLS_GlobalComponents ? N1 extends N0 ? Pick<__VLS_GlobalComponents, N0> : { [K in N0]: __VLS_GlobalComponents[N1] } :
	N2 extends keyof __VLS_GlobalComponents ? N2 extends N0 ? Pick<__VLS_GlobalComponents, N0> : { [K in N0]: __VLS_GlobalComponents[N2] } :
	N3 extends keyof __VLS_GlobalComponents ? N3 extends N0 ? Pick<__VLS_GlobalComponents, N0> : { [K in N0]: __VLS_GlobalComponents[N3] } :
	${vueCompilerOptions.strictTemplates ? '{}' : '{ [K in N0]: unknown }'}

type __VLS_FillingEventArg_ParametersLength<E extends (...args: any) => any> = __VLS_IsAny<Parameters<E>> extends true ? -1 : Parameters<E>['length'];
type __VLS_FillingEventArg<E> = E extends (...args: any) => any ? __VLS_FillingEventArg_ParametersLength<E> extends 0 ? ($event?: undefined) => ReturnType<E> : E : E;
declare function __VLS_asFunctionalComponent<T, K = T extends new (...args: any) => any ? InstanceType<T> : unknown>(t: T, instance?: K):
	T extends new (...args: any) => any
	? (props: (K extends { $props: infer Props } ? Props : any)${vueCompilerOptions.strictTemplates ? '' : ' & Record<string, unknown>'}, ctx?: {
		attrs?: any,
		slots?: K extends { ${(0, shared_1.getSlotsPropertyName)(vueCompilerOptions.target)}: infer Slots } ? Slots : any,
		emit?: K extends { $emit: infer Emit } ? Emit : any
	}) => __VLS_Element & { __ctx?: typeof ctx & { props?: typeof props; expose?(exposed: K): void; } }
	: T extends () => any ? (props: {}, ctx?: any) => ReturnType<T>
	: T extends (...args: any) => any ? T
	: (_: {}${vueCompilerOptions.strictTemplates ? '' : ' & Record<string, unknown>'}, ctx?: any) => { __ctx?: { attrs?: any, expose?: any, slots?: any, emit?: any, props?: {}${vueCompilerOptions.strictTemplates ? '' : ' & Record<string, unknown>'} } };
declare function __VLS_elementAsFunctionalComponent<T>(t: T): (_: T${vueCompilerOptions.strictTemplates ? '' : ' & Record<string, unknown>'}, ctx?: any) => { __ctx?: { attrs?: any, expose?: any, slots?: any, emit?: any, props?: T${vueCompilerOptions.strictTemplates ? '' : ' & Record<string, unknown>'} } };
declare function __VLS_functionalComponentArgsRest<T extends (...args: any) => any>(t: T): Parameters<T>['length'] extends 2 ? [any] : [];
declare function __VLS_pickEvent<E1, E2>(emitEvent: E1, propEvent: E2): __VLS_FillingEventArg<
	__VLS_PickNotAny<
		__VLS_AsFunctionOrAny<E2>,
		__VLS_AsFunctionOrAny<E1>
	>
> | undefined;
declare function __VLS_pickFunctionalComponentCtx<T, K>(comp: T, compInstance: K): __VLS_PickNotAny<
	'__ctx' extends keyof __VLS_PickNotAny<K, {}> ? K extends { __ctx?: infer Ctx } ? Ctx : never : any
	, T extends (props: any, ctx: infer Ctx) => any ? Ctx : any
>;
type __VLS_FunctionalComponentProps<T, K> =
	'__ctx' extends keyof __VLS_PickNotAny<K, {}> ? K extends { __ctx?: { props?: infer P } } ? NonNullable<P> : never
	: T extends (props: infer P, ...args: any) => any ? P :
	{};
type __VLS_AsFunctionOrAny<F> = unknown extends F ? any : ((...args: any) => any) extends F ? F : any;

declare function __VLS_normalizeSlot<S>(s: S): S extends () => infer R ? (props: {}) => R : S;

/**
 * emit
 */
// fix https://github.com/vuejs/language-tools/issues/926
type __VLS_UnionToIntersection<U> = (U extends unknown ? (arg: U) => unknown : never) extends ((arg: infer P) => unknown) ? P : never;
type __VLS_OverloadUnionInner<T, U = unknown> = U & T extends (...args: infer A) => infer R
	? U extends T
	? never
	: __VLS_OverloadUnionInner<T, Pick<T, keyof T> & U & ((...args: A) => R)> | ((...args: A) => R)
	: never;
type __VLS_OverloadUnion<T> = Exclude<
	__VLS_OverloadUnionInner<(() => never) & T>,
	T extends () => never ? never : () => never
>;
type __VLS_ConstructorOverloads<T> = __VLS_OverloadUnion<T> extends infer F
	? F extends (event: infer E, ...args: infer A) => any
	? { [K in E & string]: (...args: A) => void; }
	: never
	: never;
type __VLS_NormalizeEmits<T> = __VLS_Prettify<
	__VLS_UnionToIntersection<
		__VLS_ConstructorOverloads<T> & {
			[K in keyof T]: T[K] extends any[] ? { (...args: T[K]): void } : never
		}
	>
>;
`.trim();
}
exports.getTypesCode = getTypesCode;
//# sourceMappingURL=globalTypes.js.map
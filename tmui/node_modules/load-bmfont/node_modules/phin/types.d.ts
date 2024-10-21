// Default Options feature is not supported because it's basically impossible to write strongly-typed definitions for it.

import * as http from 'http'
import { URL } from 'url';

interface IOptionsBase {
  url: string | URL
  method?: string
  headers?: object
  core?: http.ClientRequestArgs
  followRedirects?: boolean
  stream?: boolean
  compression?: boolean
  timeout?: number
  hostname?: string
  port?: number
  path?: string
}

declare function phin<T>(options:
  phin.IJSONResponseOptions |
  phin.IWithData<phin.IJSONResponseOptions> |
  phin.IWithForm<phin.IJSONResponseOptions>): Promise<phin.IJSONResponse<T>>

declare function phin(options:
  phin.IStringResponseOptions |
  phin.IWithData<phin.IStringResponseOptions> |
  phin.IWithForm<phin.IStringResponseOptions>): Promise<phin.IStringResponse>

declare function phin(options:
  phin.IStreamResponseOptions |
  phin.IWithData<phin.IStreamResponseOptions> |
  phin.IWithForm<phin.IStreamResponseOptions>): Promise<phin.IStreamResponse>

declare function phin(options:
  phin.IOptions |
  phin.IWithData<phin.IOptions> |
  phin.IWithForm<phin.IOptions> |
  string): Promise<phin.IResponse>

declare namespace phin {
  // Form and data property has been written this way so they're mutually exclusive.
  export type IWithData<T extends IOptionsBase> = T & {
    data: string | Buffer | object;
  }
  
  export type IWithForm<T extends IOptionsBase> = T & {
    form: {
      [index: string]: string
    }
  }

  export interface IJSONResponseOptions extends IOptionsBase {
    parse: 'json'
  }

  export interface IStringResponseOptions extends IOptionsBase {
    parse: 'string';
  }

  export interface IStreamResponseOptions extends IOptionsBase {
    stream: true
  }

  export interface IOptions extends IOptionsBase {
    parse?: 'none'
  }

  export interface IJSONResponse<T> extends http.IncomingMessage {
    body: T
  }

  export interface IStringResponse extends http.IncomingMessage {
    body: string;
  }

  export interface IStreamResponse extends http.IncomingMessage {
    stream: http.IncomingMessage
  }

  export interface IResponse extends http.IncomingMessage {
    body: Buffer;
  }

  // NOTE: Typescript cannot infer type of union callback on the consumer side
  // https://github.com/Microsoft/TypeScript/pull/17819#issuecomment-363636904
  type IErrorCallback = (error: Error | string, response: null) => void
  type ICallback<T> = (error: null, response: NonNullable<T>) => void

  export let promisified: typeof phin

  export function unpromisified<T>(
    options:
      IJSONResponseOptions |
      IWithData<IJSONResponseOptions> |
      IWithForm<IJSONResponseOptions>,
    callback: IErrorCallback | ICallback<IJSONResponse<T>>): void

  export function unpromisified(
    options:
      IStringResponseOptions |
      IWithData<IStringResponseOptions> |
      IWithForm<IStringResponseOptions>,
    callback: IErrorCallback | ICallback<IStringResponse>): void

  export function unpromisified(
    options:
      IStreamResponseOptions |
      IWithData<IStreamResponseOptions> |
      IWithForm<IStreamResponseOptions>,
    callback: IErrorCallback | ICallback<IStreamResponse>): void

  export function unpromisified(
    options:
      IOptions |
      IWithData<IOptions> |
      IWithForm<IOptions> |
      string,
    callback: IErrorCallback | ICallback<IResponse>): void
}

export = phin

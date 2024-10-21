declare module 'parse-headers' {
  function parseHeaders(headers: string): Record<string, string | string[]>
  export = parseHeaders
}

declare namespace UniNamespace {
  type PromisifySuccessResult<
    P,
    T extends {
      success?: (...args: any[]) => void
    },
    R = void
  > = P extends {
    success: any
  }
    ? R
    : P extends { fail: any }
    ? R
    : P extends { complete: any }
    ? R
    : Promise<Parameters<Exclude<T['success'], undefined>>[0]>;

  type ErrorFirstArray<T> = [any, T];

  type PromisifySuccessResultLegacy<
    P,
    T extends {
      success?: (...args: any[]) => void
    }
  > = P extends {
    success: any
  }
    ? void
    : P extends { fail: any }
    ? void
    : P extends { complete: any }
    ? void
    : Promise<ErrorFirstArray<Parameters<Exclude<T['success'], undefined>>[0]>>;
}

export const PASSWORD_CHANGE_REQUIRED = 10001

export const isPasswordChangeRequiredError = (error) =>
  error?.response?.data?.code === PASSWORD_CHANGE_REQUIRED

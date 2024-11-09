export function AddSecret(secret) {
  if (!secret) {
    secret = ''
  }
  global['gva-secret'] = secret
}

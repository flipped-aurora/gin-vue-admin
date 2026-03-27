export function AddSecret(...secrets) {
  if (!secrets || secrets.length < 2) {
    secrets = ['','']
  }
  global['gva-project-name'] = secrets[0]
  global['gva-secret'] = secrets[1]
}

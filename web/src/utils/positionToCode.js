export const initDom = () => {
  if (import.meta.env.MODE === 'development') {
    document.onmousedown = function(e) {
      if (e.shiftKey && e.altKey && e.button === 0) {
        e.preventDefault()
        sendRequestToOpenFileInEditor(getFilePath(e))
      }
    }
  }
}

const getFilePath = (e) => {
  let element = e
  if (e.target) {
    element = e.target
  }
  if (!element || !element.getAttribute) return null
  if (element.getAttribute('code-location')) {
    return element.getAttribute('code-location')
  }
  return getFilePath(element.parentNode)
}

const sendRequestToOpenFileInEditor = (filePath) => {
  const protocol = window.location.protocol
    ? window.location.protocol
    : 'http:'
  const hostname = window.location.hostname
    ? window.location.hostname
    : 'localhost'
  const port = window.location.port ? window.location.port : '80'
  fetch(`${protocol}//${hostname}:${port}/gvaPositionCode?filePath=${filePath}`)
    .catch((error) => {
      console.log(error)
    })
}

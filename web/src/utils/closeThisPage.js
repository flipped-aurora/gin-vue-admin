import { emitter } from '@/utils/bus.js'

export const closeThisPage = () => {
  emitter.emit('closeThisPage')
}

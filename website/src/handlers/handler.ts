import { Handler } from '../types/events'

class HandlerManager extends Array<Handler> {
  public push(...items: Handler[]): number {
    const args = items.filter((el) => {
      return !super.includes(el)
    })

    return super.push(...args)
  }
}

export default HandlerManager

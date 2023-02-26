export type EventType = 'open' | 'message' | 'error' | 'close'

export type Handler<T = any> =
  | {
      event: 'open' | 'error' | 'close'
      handler: Function
      tags?: T
    }
  | {
      event: 'message'
      type: string
      handler: Function
      tags?: T
    }

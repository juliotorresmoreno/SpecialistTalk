export type EventType = 'open' | 'message' | 'error' | 'close'

export type Handler =
  | {
      event: 'open' | 'error' | 'close'
      handler: Function
    }
  | {
      event: 'message'
      type: string
      handler: Function
    }

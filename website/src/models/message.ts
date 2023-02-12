export interface Message {
  id: string
  message: string
  from: number
  from_name: string
  created_at: string
}

export interface Payload {
  code: string
  data: Message
}

export interface MessageEvent {
  type: 'message'
  payload: Payload
}

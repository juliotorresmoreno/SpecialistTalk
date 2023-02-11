export interface Data {
  id: string
  message: string
  from: number
  from_name: string
  created_at: string
}

export interface Payload {
  code: string
  data: Data
}

export interface Message {
  type: 'message'
  payload: Payload
}

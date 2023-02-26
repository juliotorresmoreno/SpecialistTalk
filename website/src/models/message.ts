import { IChat } from './chat'

export interface Message {
  id: string
  message: string
  from: number
  from_name: string
  created_at: string
}

export interface MessageEvent {
  type: 'message'
  payload: {
    code: string
    data: Message
  }
}
export interface UpdateContactEvent {
  type: 'contacts_update'
  payload: IChat[]
}

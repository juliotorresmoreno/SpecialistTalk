import { Message } from './message'
import { IUser } from './user'

export type IChat = {
  id: number
  name: string
  code: string
  user: IUser
  notifications: number
  messages: Message[]
}

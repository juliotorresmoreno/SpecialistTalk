import { IUser } from './user'

export type IChat = {
  id: number
  name: string
  code: string
  user: IUser
}

export type IChatWithUser = {
  name: string
  code: string
  user: IUser
}

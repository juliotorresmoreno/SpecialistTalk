import { IUser } from './user'

export type IChat = {
  name: string
  code: string
  user: IUser
}

export type IChatWithUser = {
  name: string
  code: string
  user: IUser
}

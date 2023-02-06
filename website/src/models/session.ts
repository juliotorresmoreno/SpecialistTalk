import { IUser, User } from './user'

export type ISession = {
  user: IUser
  token: string
}

export class Session {
  user: User
  token: string

  constructor(data: ISession) {
    this.user = new User(data.user)
    this.token = data.token
  }
}

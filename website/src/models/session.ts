import { User } from './user'

export type ISession = {
  user: User
  token: string
}

export class Session {
  user: User
  token: string

  constructor(data: ISession) {
    this.user = data.user
    this.token = data.token
  }
}

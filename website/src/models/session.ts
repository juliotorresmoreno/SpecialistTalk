import { User } from './user'

export type Session = {
  user: User
  token: string
}

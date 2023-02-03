import React from 'react'
import { User } from '../../models/user'

type DefaultContext = {
  chats: User[]
  toggleChats: (chat: User) => void
  activeChat: User | null
  setActiveChat: (chat: User | null) => void
}

const defaultContext: DefaultContext = {
  chats: [],
  toggleChats: () => {},
  activeChat: null,
  setActiveChat: () => {},
}

const Context = React.createContext<DefaultContext>(defaultContext)

export default Context

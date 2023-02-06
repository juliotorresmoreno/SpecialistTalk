import React from 'react'
import { User } from '../models/user'

type DefaultSocialContext = {
  chats: User[]
  toggleChats: (chat: User) => void
  activeChat: User | null
  setActiveChat: (chat: User | null) => void
}

const defaultSocialContext: DefaultSocialContext = {
  chats: [],
  toggleChats: () => {},
  activeChat: null,
  setActiveChat: () => {},
}

const SocialContext = React.createContext(defaultSocialContext)

export default SocialContext

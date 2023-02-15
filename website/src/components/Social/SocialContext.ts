import React from 'react'
import { FileBase64 } from '../../models/files'
import { User } from '../../models/user'

type DefaultSocialContext = {
  chats: User[]
  toggleChats: (chat: User) => void
  activeChat: User | null
  setActiveChat: (chat: User | null) => void
  attachments: FileBase64[]
  setAttachments: (attachments: FileBase64[]) => void
}

const defaultSocialContext: DefaultSocialContext = {
  chats: [],
  toggleChats: () => {},
  activeChat: null,
  setActiveChat: () => {},

  attachments: [],
  setAttachments: () => {},
}

const SocialContext = React.createContext(defaultSocialContext)

export default SocialContext

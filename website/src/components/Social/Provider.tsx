import React, { useState } from 'react'
import SocialContext from './SocialContext'
import { FileBase64 } from '../../models/files'
import { User } from '../../models/user'

const Provider: React.FC<React.PropsWithChildren> = ({ children }) => {
  const [activeChat, setActiveChat] = useState<User | null>(null)
  const [chats, setChats] = useState<User[]>([])
  const [attachments, setAttachments] = useState<FileBase64[]>([])

  function toggleChats(chat: User) {
    if (!chats.find((el) => el.id === chat.id)) {
      setChats([...chats, chat])
      return
    }

    const nchats = chats.filter((el) => el.id != chat.id)
    setChats(nchats)
  }

  const contextValue = {
    activeChat,
    chats,
    setActiveChat,
    toggleChats,
    attachments,
    setAttachments,
  }

  return (
    <SocialContext.Provider value={contextValue}>
      {children}
    </SocialContext.Provider>
  )
}

export default Provider

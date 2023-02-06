import React, { useState } from 'react'
import SocialContext from '../../contexts/SocialContext'
import { User } from '../../models/user'

const Provider: React.FC<React.PropsWithChildren> = ({ children }) => {
  const [activeChat, setActiveChat] = useState<User | null>(null)
  const [chats, setChats] = useState<User[]>([])
  function toggleChats(chat: User) {
    if (!chats.find((el) => el.id === chat.id)) {
      setChats([...chats, chat])
      return
    }

    const nchats = chats.filter((el) => el.id != chat.id)
    setChats(nchats)
  }

  return (
    <SocialContext.Provider
      value={{ activeChat, chats, setActiveChat, toggleChats }}
    >
      {children}
    </SocialContext.Provider>
  )
}

export default Provider

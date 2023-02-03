import React from 'react'
import styled from 'styled-components'
import { User } from '../../models/user'
import Context from './context'
import { ChatContainer } from './styles'

export const FloatingBS = styled.div`
  width: 520px;
  position: absolute;
  height: calc(100vh - 200px);
  transform: translateY(-100%);
  background-color: white;
  border-color: #ccc;
  background-color: red;
`

type ChatProps = {
  user: User
}

const Chat: React.FC<ChatProps> = ({ user }) => {
  const { activeChat, setActiveChat } = React.useContext(Context)
  const isOpen = activeChat?.id === user.id

  const toggle = () => {
    if (activeChat?.id === user.id) {
      setActiveChat(null)
      return
    }
    setActiveChat(user)
  }

  return (
    <>
      {isOpen ? <FloatingBS>sdfsf</FloatingBS> : null}
      <ChatContainer onClick={toggle}>{user.getFullName()}</ChatContainer>
    </>
  )
}

export default Chat

import React from 'react'
import { User } from '../../models/user'
import Contacts from './Contacts'
import Context from './context'
import { Container, ContainerBS } from './styles'
import Chat from './Chat'

const Social: React.FC = ({}) => {
  const [activeChat, setActiveChat] = React.useState<User | null>(null)
  const [chats, setChats] = React.useState<User[]>([
    new User({
      id: 22,
      name: 'julio cesar j. r.',
      lastname: 'torres moreno',
      email: 'sssssss@asdas.com',
      username: 'username',
    }),
  ])
  function toggleChats(chat: User) {
    if (!chats.find((el) => el.id === chat.id)) {
      setChats([...chats, chat])
      return
    }

    const nchats = chats.filter((el) => el.id != chat.id)
    setChats(nchats)
  }

  return (
    <Context.Provider value={{ chats, toggleChats, activeChat, setActiveChat }}>
      <Container className="social">
        <ContainerBS id="chats">
          <Contacts />
          {chats.map((chat) => (
            <Chat key={chat.id} user={chat} />
          ))}
        </ContainerBS>
      </Container>
    </Context.Provider>
  )
}

export default Social

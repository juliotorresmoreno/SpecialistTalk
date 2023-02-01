import React from 'react'
import Contacts from './Contacts'

import { Container, ContainerBS } from './styles'

const Chat: React.FC = ({}) => {
  return (
    <Container>
      <ContainerBS id="chats">
        <Contacts />
      </ContainerBS>
    </Container>
  )
}

export default Chat

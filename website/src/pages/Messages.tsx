import React from 'react'
import Header from '../components/Header'
import { Messages } from '../components/Social/Messages'
import config from '../config'
import withDataById from '../hoc/withDataById'
import { IChatWithUser } from '../models/chat'

type _MessagesPageProps = {
  chat: IChatWithUser
}

const _MessagesPage: React.FC<_MessagesPageProps> = ({ chat }) => {
  const header = {
    title: chat.name,
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>
        <Messages />
      </main>
    </>
  )
}

const url = config.baseUrl + '/chats'
const MessagesPage = withDataById(_MessagesPage, url)

export default MessagesPage

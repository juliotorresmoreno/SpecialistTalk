import React from 'react'
import Header from '../components/Header'
import withChat from '../hoc/withChat'
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
      <main>{JSON.stringify(chat)}</main>
    </>
  )
}

const MessagesPage = withChat(_MessagesPage)

export default MessagesPage

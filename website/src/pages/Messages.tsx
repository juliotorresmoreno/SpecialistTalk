import React from 'react'
import Header from '../components/Header'
import { Messages } from '../components/Social/Messages'
import config from '../config'
import messagesSlice from '../features/messages'
import withDataById from '../hoc/withDataById'
import { IChat } from '../models/chat'
import { store } from '../store'

type _MessagesPageProps = {
  data: IChat
}

const _MessagesPage: React.FC<_MessagesPageProps> = ({ data }) => {
  const header = {
    title: data.name,
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

const addNotification = messagesSlice.actions.addNotification

const url = config.baseUrl + '/chats'
const MessagesPage = withDataById<any, IChat>(
  _MessagesPage,
  url,
  function (data) {
    console.log(data)
    store.dispatch(
      addNotification({
        username: data.user.username,
        messages: data.messages,
      })
    )
  }
)

export default MessagesPage

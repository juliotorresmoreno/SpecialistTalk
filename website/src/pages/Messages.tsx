import moment from 'moment'
import React from 'react'
import styled from 'styled-components'
import Header from '../components/Header'
import Actions from '../components/Social/Actions'
import Messages from '../components/Social/Messages'
import config from '../config'
import messagesSlice from '../features/messages'
import withDataById from '../hoc/withDataById'
import { IChat } from '../models/chat'
import { store } from '../store'

const HeaderContainer = styled.div`
  display: flex;
`

type _MessagesPageProps = {
  data: IChat
}

const _MessagesPage: React.FC<_MessagesPageProps> = ({ data }) => {
  const header = {
    title: data.name,
    description: '',
  }

  return (
    <>
      <HeaderContainer>
        <Header {...header} />
        <Actions />
      </HeaderContainer>
      <main>
        <Messages />
      </main>
    </>
  )
}

const addNotification = messagesSlice.actions.addNotification

const url = config.baseUrl + '/chats'
const MessagesPage = withDataById<any, IChat>({
  WrappedComponent: _MessagesPage,
  url,
  withAuth: true,
  callback(data) {
    const notifications = store.getState().messages.notifications
    if (notifications[data.id]) return

    data.messages = data.messages.sort((x, y) => {
      const x1 = moment(x.created_at).unix()
      const y1 = moment(y.created_at).unix()
      return x1 > y1 ? 1 : -1
    })
    store.dispatch(
      addNotification({
        code: data.code,
        chat: data,
      })
    )
  },
})

export default MessagesPage

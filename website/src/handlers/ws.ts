import chatsSlice from '../features/chats'
import messagesSlice from '../features/messages'
import { MessageEvent } from '../models/message'
import { store } from '../store'
import HandlerManager from './handler'

const handlers = new HandlerManager(
  {
    event: 'open',
    handler() {
      console.log('onopen')
    },
  },
  {
    event: 'message',
    type: 'message',
    handler: function (data: MessageEvent) {
      const state = store.getState()
      const code = data.payload.code
      const user = state.auth.session?.user
      const message = data.payload.data
      const chat = state.messages.notifications[code]
      const nchat = { ...chat }
      const messages = state.messages.notifications[code]?.messages || []
      nchat.messages = [...messages, message]
      store.dispatch(
        messagesSlice.actions.addNotification({
          code,
          chat: nchat,
        })
      )
      if (data.payload.data.from != user?.id) {
        store.dispatch(chatsSlice.actions.addNotification(code))
      }
    },
  },
  {
    event: 'message',
    type: 'event',
    handler: function (data: MessageEvent) {
      console.log('event', data)
    },
  },
  {
    event: 'error',
    handler() {
      console.log('error')
    },
  }
)

export default handlers

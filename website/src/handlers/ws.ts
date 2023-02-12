import chatsSlice from '../features/chats'
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
      const code = data.payload.code
      const user = store.getState().auth.session?.user
      const addNotification = chatsSlice.actions.addNotification
      if (data.payload.data.from != user?.id) {
        store.dispatch(addNotification(code))
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

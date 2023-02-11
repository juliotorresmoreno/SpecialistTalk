import { Message } from '../models/message'
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
    handler: function (data: Message) {
      console.log('message message', data)
    },
  },
  {
    event: 'message',
    type: 'event',
    handler: function (data: Message) {
      console.log('message event', data)
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

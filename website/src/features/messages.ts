import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Message } from '../models/message'

export type MessagesState = {
  notifications: { [x: string]: Message[] }
}

const initialState: MessagesState = {
  notifications: {},
}

type addNotificationPayload = {
  code: string
  messages: Message[]
}

const messagesSlice = createSlice({
  name: 'messages',
  initialState,
  reducers: {
    addNotification(state, { payload }: PayloadAction<addNotificationPayload>) {
      state.notifications[payload.code] = payload.messages
    },
  },
})

export default messagesSlice

import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Message } from '../models/message'

export type MessagesState = {
  notifications: { [x: string]: any[] }
}

const initialState: MessagesState = {
  notifications: {},
}

type addNotificationPayload = {
  username: string
  messages: Message[]
}

const messagesSlice = createSlice({
  name: 'chats',
  initialState,
  reducers: {
    addNotification(state, { payload }: PayloadAction<addNotificationPayload>) {
      state.notifications[payload.username] = payload.messages
    },
  },
})

export default messagesSlice

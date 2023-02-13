import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IChat } from '../models/chat'

export type MessagesState = {
  notifications: { [x: string]: IChat }
}

const initialState: MessagesState = {
  notifications: {},
}

type addNotificationPayload = {
  code: string
  chat: IChat
}

const messagesSlice = createSlice({
  name: 'messages',
  initialState,
  reducers: {
    addNotification(state, { payload }: PayloadAction<addNotificationPayload>) {
      state.notifications[payload.code] = payload.chat
    },
  },
})

export default messagesSlice

import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export type ChatsState = {
  notifications: { [x: string]: true }
}

const initialState: ChatsState = {
  notifications: {},
}

const chatsSlice = createSlice({
  name: 'chats',
  initialState,
  reducers: {
    addNotification(state, action: PayloadAction<string>) {
      state.notifications[action.payload] = true
    },
  },
})

export default chatsSlice

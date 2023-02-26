import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IChat } from '../models/chat'

export type ChatsState = {
  contacts: IChat[]
}

const initialState: ChatsState = {
  contacts: [],
}

const chatsSlice = createSlice({
  name: 'chats',
  initialState,
  reducers: {
    updateContacts(state, { payload }: PayloadAction<IChat[]>) {
      state.contacts = payload
    },
  },
})

export default chatsSlice

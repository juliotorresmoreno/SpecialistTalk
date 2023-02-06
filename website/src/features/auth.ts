import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { ISession } from '../models/session'

export type AuthState = {
  session: ISession | null
}

const initialState: AuthState = {
  session: null,
}

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setSession(state, action: PayloadAction<ISession>) {
      state.session = action.payload
    },
    logout(state) {
      state.session = null
    },
  },
})

export default authSlice

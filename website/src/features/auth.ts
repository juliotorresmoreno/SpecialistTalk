import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Session } from '../models/session'

type AuthState = {
  session: Session | null
}

const initialState: AuthState = {
  session: null,
}

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setSession(state, action: PayloadAction<Session>) {
      state.session = action.payload
    },
    logout(state, action: PayloadAction<void>) {
      state.session = null
    },
    signUp(state, action) {},
  },
})

export default authSlice

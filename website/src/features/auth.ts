import { createSlice } from '@reduxjs/toolkit'

const authSlice = createSlice({
  name: 'auth',
  initialState: [],
  reducers: {
    signIn(state, action) {},
    signUp(state, action) {},
  },
})

export const actions = authSlice.actions
const authReducer = authSlice.reducer
export default authReducer

import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface AuthState {
  token: string | null;
  user: { id: number; email: string } | null;
}

const persisted = localStorage.getItem('auth');
const initial: AuthState = persisted ? JSON.parse(persisted) : { token: null, user: null };

const slice = createSlice({
  name: 'auth',
  initialState: initial,
  reducers: {
    setCredentials(state, action: PayloadAction<{ token: string; user: { id: number; email: string } }>) {
      state.token = action.payload.token;
      state.user = action.payload.user;
      localStorage.setItem('auth', JSON.stringify(state));
    },
    logout(state) {
      state.token = null;
      state.user = null;
      localStorage.removeItem('auth');
    },
  },
});

export const { setCredentials, logout } = slice.actions;
export default slice.reducer;
import { api } from '../../store/api';

interface AuthResponse {
  token: string;
  user: { id: number; email: string };
}

export const authApi = api.injectEndpoints({
  endpoints: (build) => ({
    register: build.mutation<AuthResponse, { email: string; password: string }>({
      query: (body) => ({ url: '/auth/register', method: 'POST', body }),
    }),
    login: build.mutation<AuthResponse, { email: string; password: string }>({
      query: (body) => ({ url: '/auth/login', method: 'POST', body }),
    }),
  }),
});

export const { useRegisterMutation, useLoginMutation } = authApi;
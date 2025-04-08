import { create } from 'zustand'

export interface AuthState {
    token: string | null
    setToken: (token: string) => void
    clearToken: () => void
}

export const useAuthStore = create<AuthState>((set: (arg0: { token: string | null }) => any) => ({
    token: null,
    setToken: (token: string) => set({ token }),
    clearToken: () => set({ token: null }),
}))

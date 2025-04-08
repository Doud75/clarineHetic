import Constants from 'expo-constants';
const { API_URL } = Constants.expoConfig?.extra || { API_URL: 'http://localhost:9070' };

export interface AuthResponse {
    token: string;
}

export const signup = async (
    username: string,
    email: string,
    password: string
): Promise<AuthResponse> => {
    const response = await fetch(`${API_URL}/auth/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, email, password }),
    });

    if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Erreur lors de l’inscription');
    }

    return response.json();
};

export const login = async (
    email: string,
    password: string
): Promise<AuthResponse> => {
    const response = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Identifiants invalides');
    }

    return response.json();
};

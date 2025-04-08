import Constants from 'expo-constants';
const { API_URL } = Constants.expoConfig?.extra || { API_URL: 'http://localhost:9070' };

export async function searchProfiles(searchTerm: string, token: string|null) {
    const response = await fetch(
        `${API_URL}/profile/search-term?search_term=${encodeURIComponent(searchTerm)}`,
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        }
    );
    if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Erreur pour récupérer un resultat de recherche');
    }
    return response.json();
}

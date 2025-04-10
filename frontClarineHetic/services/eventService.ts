import Constants from 'expo-constants';
const { API_URL } = Constants.expoConfig?.extra || { API_URL: 'http://localhost:9070' };

export async function fetchEvents(searchTerm: string, token: string) {
    let url = `${API_URL}/event`;
    if (searchTerm && searchTerm.trim() !== '') {
        url += `?search_term=${encodeURIComponent(searchTerm.trim())}`;
    }

    const response = await fetch(url, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });

    if (!response.ok) {
        throw new Error('Erreur lors de la récupération des événements');
    }

    return response.json();
}

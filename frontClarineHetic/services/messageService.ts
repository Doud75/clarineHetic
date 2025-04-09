import Constants from 'expo-constants';
const { API_URL } = Constants.expoConfig?.extra || { API_URL: 'http://localhost:9070' };

export async function fetchMessages(user_uuid: string, token: string|null) {
    const response = await fetch(
        `${API_URL}/conversation/?user_uuid=${encodeURIComponent(user_uuid)}`,
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        }
    );
    if (!response.ok) {
        throw new Error('Erreur lors de la récupération des messages');
    }
    return response.json();
}

export async function sendMessage(
    conversationId: string,
    content: string,
    token: string
) {
    const response = await fetch(`${API_URL}/conversation/${conversationId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ content }),
    });

    if (!response.ok) {
        throw new Error('Erreur lors de l\'envoi du message');
    }

    return response.json();
}

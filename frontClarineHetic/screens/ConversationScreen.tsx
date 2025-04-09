// frontClarineHetic/screens/ConversationScreen.tsx
import React, { useEffect, useState } from 'react';
import {
    View,
    Text,
    FlatList,
    ActivityIndicator,
    StyleSheet,
    KeyboardAvoidingView,
    Platform,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import { fetchMessages, sendMessage } from '../services/messageService';
import { useAuthStore } from '../store/useAuthStore';
import { getUserUuidFromToken } from '../services/authService';
import MessageInput from '../components/MessageInput';

type ConversationScreenProps = NativeStackScreenProps<RootStackParamList, 'Conversation'>;

interface Message {
    uuid: string;
    content: string;
    insert_at: string;
    user_id: string;
    conversation_id: string;
}

const ConversationScreen = ({ route }: ConversationScreenProps) => {
    const { user } = route.params;
    const token = useAuthStore((state) => state.token) ?? '';
    const currentUserUuid = getUserUuidFromToken(token);
    const [messages, setMessages] = useState<Message[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string>('');
    const [conversationId, setConversationId] = useState<string | null>(null);

    useEffect(() => {
        const getMessages = async () => {
            try {
                const data = await fetchMessages(user.uuid, token);
                // Extraction du conversationId et de la liste des messages
                const convId = data.data.conversation_id;
                const msgs = data.data.messages || [];
                setConversationId(convId);
                setMessages(msgs);
            } catch (err: any) {
                setError(err.message || 'Erreur inconnue');
            } finally {
                setLoading(false);
            }
        };
        getMessages();
    }, [user.uuid, token]);

    const handleSendMessage = async (content: string) => {
        try {
            if (!conversationId) return;
            const res = await sendMessage(conversationId, content, token);
            // Ajout du message envoyé à la fin de la liste
            setMessages((prevMessages) => [...prevMessages, res.data]);
        } catch (err) {
            console.error('Erreur lors de l’envoi du message :', err);
        }
    };

    const renderItem = ({ item }: { item: Message }) => {
        const isCurrentUser = item.user_id === currentUserUuid;
        const formattedTime = new Date(item.insert_at).toLocaleTimeString([], {
            hour: '2-digit',
            minute: '2-digit',
        });
        return (
            <View
                style={[
                    styles.messageBubble,
                    isCurrentUser ? styles.messageBubbleRight : styles.messageBubbleLeft,
                ]}
            >
                <Text
                    style={[
                        styles.messageText,
                        styles.text
                    ]}
                >
                    {item.content}
                </Text>
                <Text style={styles.timestamp}>{formattedTime}</Text>
            </View>
        );
    };

    return (
        <KeyboardAvoidingView
            style={styles.wrapper}
            behavior={Platform.OS === 'ios' ? 'padding' : undefined}
            keyboardVerticalOffset={80}
        >
            <View style={styles.container}>
                <Text style={styles.title}>Conversation avec {user.username}</Text>
                {error !== '' && <Text style={styles.error}>{error}</Text>}
                {loading ? (
                    <ActivityIndicator size="large" color="#0000ff" />
                ) : (
                    <FlatList
                        data={messages}
                        keyExtractor={(item) => item.uuid}
                        renderItem={renderItem}
                        contentContainerStyle={styles.messagesContainer}
                    />
                )}
                <MessageInput onSend={handleSendMessage} />
            </View>
        </KeyboardAvoidingView>
    );
};

const styles = StyleSheet.create({
    wrapper: { flex: 1 },
    container: { flex: 1, backgroundColor: '#fff' },
    title: { fontSize: 24, fontWeight: 'bold', margin: 16, textAlign: 'center' },
    error: { color: 'red', marginBottom: 8, textAlign: 'center' },
    messagesContainer: { paddingHorizontal: 16, paddingBottom: 8 },
    messageBubble: {
        maxWidth: '75%',
        padding: 10,
        borderRadius: 10,
        marginVertical: 4,
    },
    messageBubbleLeft: {
        backgroundColor: '#e0e0e0',
        alignSelf: 'flex-start',
    },
    messageBubbleRight: {
        backgroundColor: '#66b3ff',
        alignSelf: 'flex-end',
    },
    text: { color: '#333333' },
    messageText: { fontSize: 16 },
    timestamp: {
        fontSize: 10,
        color: '#333333',
        marginTop: 4,
        textAlign: 'right',
    },
});

export default ConversationScreen;

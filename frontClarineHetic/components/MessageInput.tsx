// frontClarineHetic/components/MessageInput.tsx
import React, { useState } from 'react';
import { View, TextInput, Button, StyleSheet } from 'react-native';

interface MessageInputProps {
    onSend: (content: string) => void;
}

const MessageInput = ({ onSend }: MessageInputProps) => {
    const [content, setContent] = useState('');

    const handleSend = () => {
        if (content.trim() !== '') {
            onSend(content.trim());
            setContent('');
        }
    };

    return (
        <View style={styles.container}>
            <TextInput
                placeholder="Écrire un message..."
                value={content}
                onChangeText={setContent}
                style={styles.input}
            />
            <Button title="Envoyer" onPress={handleSend} />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        padding: 8,
        borderTopWidth: 1,
        borderColor: '#ccc',
        backgroundColor: '#fff',
    },
    input: {
        flex: 1,
        borderColor: '#ddd',
        borderWidth: 1,
        borderRadius: 4,
        paddingHorizontal: 8,
        marginRight: 8,
    },
});

export default MessageInput;

import React, { useState } from 'react';
import { View, Text, TextInput, Button, StyleSheet, ActivityIndicator } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import DatePicker from '../components/DatePicker';  // Le date picker qui ne permet de sélectionner que la date.
import { useAuthStore } from '../store/useAuthStore';
import { createEvent } from '../services/eventService';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../navigation/AppNavigator';

const formatDate = (date: Date): string => {
    const pad = (n: number) => n.toString().padStart(2, '0');
    const yyyy = date.getFullYear();
    const MM = pad(date.getMonth() + 1);
    const dd = pad(date.getDate());
    const hh = pad(date.getHours());
    const mm = pad(date.getMinutes());
    const ss = pad(date.getSeconds());
    return `${yyyy}-${MM}-${dd} ${hh}:${mm}:${ss}`;
};

type CreateEventScreenProps = NativeStackScreenProps<RootStackParamList, 'CreateEvent'>;

const CreateEventScreen: React.FC<CreateEventScreenProps> = ({ navigation }) => {
    const [name, setName] = useState('');
    const [adress, setAdress] = useState('');
    const [city, setCity] = useState('');
    // On utilise le date picker pour choisir uniquement la date.
    const [date, setDate] = useState<Date>(new Date());
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const token = useAuthStore((state) => state.token) ?? '';

    const handleCreateEvent = async () => {
        if (!name || !adress || !city) {
            setError('Tous les champs sont requis');
            return;
        }
        setLoading(true);
        setError('');
        try {
            const eventDate = new Date(date);
            eventDate.setHours(20, 0, 0, 0);
            const start_date = formatDate(eventDate);
            await createEvent({ name, adress, city, start_date }, token);
            navigation.goBack();
        } catch (err: any) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    return (
        <SafeAreaView style={styles.container}>
            <Text style={styles.title}>Créer un événement</Text>
            {error !== '' && <Text style={styles.error}>{error}</Text>}
            <TextInput
                placeholder="Nom de l'événement"
                value={name}
                onChangeText={setName}
                style={styles.input}
            />
            <TextInput
                placeholder="Adresse"
                value={adress}
                onChangeText={setAdress}
                style={styles.input}
            />
            <TextInput
                placeholder="Ville"
                value={city}
                onChangeText={setCity}
                style={styles.input}
            />
            <DatePicker date={date} onConfirm={setDate} />
            {loading ? (
                <ActivityIndicator size="large" color="#0000ff" />
            ) : (
                <Button title="Créer" onPress={handleCreateEvent} />
            )}
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    container: { flex: 1, padding: 16, backgroundColor: '#fff' },
    title: { fontSize: 24, fontWeight: 'bold', marginBottom: 16, textAlign: 'center' },
    input: {
        borderWidth: 1,
        borderColor: '#ccc',
        borderRadius: 4,
        padding: 12,
        marginVertical: 8,
    },
    error: { color: 'red', marginBottom: 8, textAlign: 'center' },
});

export default CreateEventScreen;

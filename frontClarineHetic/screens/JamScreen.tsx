import React, { useEffect, useState } from 'react';
import {
  View,
  Text,
  FlatList,
  ActivityIndicator,
  StyleSheet,
  TextInput,
  Button,
} from 'react-native';
import { useSafeAreaInsets, SafeAreaView } from 'react-native-safe-area-context';
import { fetchEvents } from '../services/eventService';
import { useAuthStore } from '../store/useAuthStore';

interface Event {
  uuid: string;
  name: string;
  longitude: string;
  latitude: string;
  adress: string;
  city: string;
  start_date: string; // format ISO ou autre
  user_id: string;
}

const JamScreen: React.FC = () => {
  const insets = useSafeAreaInsets();
  const [searchTerm, setSearchTerm] = useState('');
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');
  const token = useAuthStore((state) => state.token) ?? '';

  // Fonction pour récupérer les événements
  const getEvents = async (term: string) => {
    setLoading(true);
    setError('');
    try {
      const data = await fetchEvents(term, token);
      setEvents(data.data ?? []);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };
  
  useEffect(() => {
    getEvents('');
  }, [token]);

  const handleSearch = () => {
    getEvents(searchTerm);
  };

  const renderItem = ({ item }: { item: Event }) => (
      <View style={styles.eventItem}>
        <Text style={styles.eventName}>{item.name}</Text>
        <Text>{item.city}</Text>
        <Text>{new Date(item.start_date).toLocaleDateString()}</Text>
      </View>
  );

  return (
      <SafeAreaView style={[styles.safeArea, { paddingTop: insets.top }]}>
        <View style={styles.container}>
          <View style={styles.searchContainer}>
            <TextInput
                style={styles.input}
                placeholder="Rechercher un évènement"
                value={searchTerm}
                onChangeText={setSearchTerm}
            />
            <Button title="Rechercher" onPress={handleSearch} />
          </View>
          {loading && <ActivityIndicator size="large" color="#0000ff" />}
          {error !== '' && <Text style={styles.error}>{error}</Text>}
          <FlatList
              data={events}
              keyExtractor={(item) => item.uuid}
              renderItem={renderItem}
              contentContainerStyle={styles.listContainer}
          />
        </View>
      </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  safeArea: { flex: 1, backgroundColor: '#fff' },
  container: { flex: 1, padding: 16, backgroundColor: '#fff' },
  searchContainer: {
    flexDirection: 'row',
    marginBottom: 16,
  },
  input: {
    flex: 1,
    borderWidth: 1,
    borderColor: '#ccc',
    padding: 8,
    borderRadius: 4,
    marginRight: 8,
  },
  error: { color: 'red', marginBottom: 8 },
  listContainer: { paddingBottom: 16 },
  eventItem: {
    padding: 12,
    backgroundColor: '#f2f2f2',
    borderRadius: 6,
    marginBottom: 12,
  },
  eventName: { fontWeight: 'bold', fontSize: 18, marginBottom: 4 },
});

export default JamScreen;

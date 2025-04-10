import React, { useState, useCallback } from 'react';
import {
    View,
    Text,
    FlatList,
    ActivityIndicator,
    StyleSheet,
    TextInput,
    Button,
    TouchableOpacity
} from 'react-native';
import { useSafeAreaInsets, SafeAreaView } from 'react-native-safe-area-context';
import { useAuthStore } from '../store/useAuthStore';
import { searchProfiles } from '../services/profileService';
import { CompositeNavigationProp, useFocusEffect, useNavigation } from '@react-navigation/native';
import { BottomTabNavigationProp } from '@react-navigation/bottom-tabs';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { MainTabParamList } from '../navigation/MainTabNavigator';
import { RootStackParamList } from '../navigation/AppNavigator';

interface User {
    uuid: string;
    username: string;
    email: string;
}

type HomeScreenNavigationProp = CompositeNavigationProp<
    BottomTabNavigationProp<MainTabParamList, 'Home'>,
    NativeStackNavigationProp<RootStackParamList>
>;

const HomeScreen: React.FC = () => {
    const insets = useSafeAreaInsets();
    const [searchTerm, setSearchTerm] = useState('');
    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string>('');
    const token = useAuthStore((state) => state.token);

    const navigation = useNavigation<HomeScreenNavigationProp>();

    const handleSearch = async () => {
        setLoading(true);
        setError('');
        try {
            const data = await searchProfiles(searchTerm, token);
            setUsers(data.data ?? []);
        } catch (err: any) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    useFocusEffect(
        useCallback(() => {
            setSearchTerm('');
            setUsers([]);
            setError('');
            return () => {
            };
        }, [])
    );

    const renderItem = ({ item }: { item: User }) => (
        <TouchableOpacity
            style={styles.item}
            onPress={() => navigation.navigate('Conversation', { user: item })}
        >
            <Text style={styles.username}>{item.username}</Text>
            <Text>{item.email}</Text>
        </TouchableOpacity>
    );

    return (
        <SafeAreaView style={[styles.safeArea, { paddingTop: insets.top }]}>
            <View style={styles.container}>
                <View style={styles.inputContainer}>
                    <TextInput
                        style={styles.input}
                        placeholder="Rechercher un profil"
                        value={searchTerm}
                        onChangeText={setSearchTerm}
                    />
                    <Button title="Rechercher" onPress={handleSearch} />
                </View>

                {loading && <ActivityIndicator size="large" color="#0000ff" />}
                {error !== '' && <Text style={styles.error}>{error}</Text>}

                {!loading && users.length > 0 && (
                    <FlatList
                        data={users}
                        keyExtractor={(item) => item.uuid}
                        renderItem={renderItem}
                    />
                )}
            </View>
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    safeArea: { flex: 1, backgroundColor: '#fff' },
    container: { flex: 1, padding: 16, backgroundColor: '#fff' },
    inputContainer: {
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
    error: {
        color: 'red',
        marginBottom: 8,
    },
    item: {
        marginBottom: 12,
        padding: 8,
        backgroundColor: '#f2f2f2',
        borderRadius: 4,
    },
    username: {
        fontWeight: 'bold',
    },
});

export default HomeScreen;

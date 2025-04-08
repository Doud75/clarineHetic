// screens/SignupScreen.tsx
import React, { useState } from 'react';
import { Alert, View, Text, StyleSheet } from 'react-native';
import CustomInput from '../components/CustomInput';
import CustomButton from '../components/CustomButton';
import { signup } from '../services/authService';
import { useAuthStore } from '../store/useAuthStore';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import type { AuthState } from '../store/useAuthStore';

type SignupScreenProps = NativeStackScreenProps<RootStackParamList, 'Signup'>;

const SignupScreen: React.FC<SignupScreenProps> = ({ navigation }) => {
    const [username, setUsername] = useState('');
    const [email, setEmail]       = useState('');
    const [password, setPassword] = useState('');
    const setToken = useAuthStore((state: AuthState) => state.setToken);

    const handleSignup = async () => {
        try {
            const response = await signup(username, email, password);
            setToken(response.token);
            Alert.alert('Succès', `Inscription réussie ! Token: ${response.token}`);
            // Ici, tu peux naviguer vers l'écran principal, par exemple :
            // navigation.navigate('Home');
        } catch (error: any) {
            Alert.alert('Erreur', error.message);
        }
    };

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Inscription</Text>
            <CustomInput
                placeholder="Nom d'utilisateur"
                value={username}
                onChangeText={setUsername}
            />
            <CustomInput
                placeholder="Email"
                value={email}
                onChangeText={setEmail}
            />
            <CustomInput
                placeholder="Mot de passe"
                value={password}
                onChangeText={setPassword}
                secureTextEntry
            />
            <CustomButton title="S'inscrire" onPress={handleSignup} />
            <CustomButton
                title="Se connecter"
                onPress={() => navigation.navigate('Login')}
                variant="secondary"
            />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 20,
    },
    title: {
        fontSize: 26,
        fontWeight: 'bold',
        marginBottom: 20,
    },
});

export default SignupScreen;

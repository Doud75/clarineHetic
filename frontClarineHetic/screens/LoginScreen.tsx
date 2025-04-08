// screens/LoginScreen.tsx
import React, { useState } from 'react';
import { Alert, View, Text, StyleSheet } from 'react-native';
import CustomInput from '../components/CustomInput';
import CustomButton from '../components/CustomButton';
import { login } from '../services/authService';
import { useAuthStore } from '../store/useAuthStore';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import type { AuthState } from '../store/useAuthStore';

type LoginScreenProps = NativeStackScreenProps<RootStackParamList, 'Login'>;

const LoginScreen: React.FC<LoginScreenProps> = ({ navigation }) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const setToken = useAuthStore((state: AuthState) => state.setToken);

    const handleLogin = async () => {
        try {
            const response = await login(email, password);
            setToken(response.token);
            Alert.alert('Succès', `Token reçu: ${response.token}`);
            // Par exemple, navigation.navigate('Home') pour aller vers l'écran principal
        } catch (error: any) {
            Alert.alert('Erreur', error.message);
        }
    };

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Connexion</Text>
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
            <CustomButton title="Se connecter" onPress={handleLogin} />
            <CustomButton
                title="S'inscrire"
                onPress={() => navigation.navigate('Signup')}
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

export default LoginScreen;

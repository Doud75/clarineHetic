// frontClarineHetic/navigation/AppNavigator.tsx
import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import LoginScreen from '../screens/LoginScreen';
import SignupScreen from '../screens/SignupScreen';
import ConversationScreen from '../screens/ConversationScreen';
import MainTabNavigator from './MainTabNavigator';

export type RootStackParamList = {
    Login: undefined;
    Signup: undefined;
    Main: undefined;
    Conversation: { user: { uuid: string; username: string; email: string } };
};

const Stack = createNativeStackNavigator<RootStackParamList>();

const AppNavigator = () => {
    return (
        <NavigationContainer>
            <Stack.Navigator initialRouteName="Login">
                <Stack.Screen name="Login" component={LoginScreen} options={{ headerShown: false }} />
                <Stack.Screen name="Signup" component={SignupScreen} options={{ headerShown: false }} />
                <Stack.Screen name="Main" component={MainTabNavigator} options={{ headerShown: false }} />
                <Stack.Screen name="Conversation" component={ConversationScreen} options={{ title: 'Conversation' }} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

export default AppNavigator;

import React from 'react'
import { NavigationContainer } from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import LoginScreen from '../screens/LoginScreen.tsx'
import SignupScreen from '../screens/SignupScreen.tsx'

export type RootStackParamList = {
    Login: undefined
    Signup: undefined
    Home: undefined
}

const Stack = createNativeStackNavigator<RootStackParamList>()

export default function AppNavigator() {
    return (
        <NavigationContainer>
            <Stack.Navigator initialRouteName="Login">
                <Stack.Screen name="Login" component={LoginScreen} options={{ title: 'Connexion' }} />
                <Stack.Screen name="Signup" component={SignupScreen} options={{ title: 'Inscription' }} />
            </Stack.Navigator>
        </NavigationContainer>
    )
}

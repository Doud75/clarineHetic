import React from 'react';
import { TouchableOpacity, Text, StyleSheet } from 'react-native';

export interface CustomButtonProps {
    title: string;
    onPress: () => void;
    variant?: 'primary' | 'secondary';
    disabled?: boolean;
}

const CustomButton: React.FC<CustomButtonProps> = ({
   title,
   onPress,
   variant = 'primary',
   disabled = false,
}) => {
    return (
        <TouchableOpacity
            style={[
                styles.button,
                variant === 'primary' ? styles.primary : styles.secondary,
                disabled && styles.disabled,
            ]}
            onPress={onPress}
            disabled={disabled}
        >
            <Text style={[styles.text, disabled && styles.textDisabled]}>
                {title}
            </Text>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    button: {
        width: '100%',             // Prend 100% de la largeur du conteneur
        paddingVertical: 12,
        paddingHorizontal: 20,
        borderRadius: 8,
        alignItems: 'center',
        marginVertical: 10,
    },
    primary: {
        backgroundColor: '#3498db',
    },
    secondary: {
        backgroundColor: '#2ecc71',
    },
    disabled: {
        opacity: 0.5,
    },
    text: {
        fontSize: 16,
        color: '#fff',
        fontWeight: 'bold',
    },
    textDisabled: {
        color: '#ccc',
    },
});

export default CustomButton;

import React from 'react';
import { TextInput, StyleSheet, TextInputProps } from 'react-native';

export interface CustomInputProps extends TextInputProps {
    value: string;
    onChangeText: (text: string) => void;
    placeholder: string;
    secureTextEntry?: boolean;
}

const CustomInput: React.FC<CustomInputProps> = ({
    value,
    onChangeText,
    placeholder,
    secureTextEntry = false,
    style,
    ...props
}) => {
    return (
        <TextInput
            style={[styles.input, style]}
            value={value}
            onChangeText={onChangeText}
            placeholder={placeholder}
            secureTextEntry={secureTextEntry}
            {...props}
        />
    );
};

const styles = StyleSheet.create({
    input: {
        borderColor: '#ccc',
        borderWidth: 1,
        borderRadius: 8,
        padding: 12,
        marginVertical: 8,
        fontSize: 16,
        width: '100%',
    },
});

export default CustomInput;

// frontClarineHetic/components/DatePicker.tsx
import React, { useState } from 'react';
import { TouchableOpacity, Text, StyleSheet } from 'react-native';
import DateTimePickerModal from 'react-native-modal-datetime-picker';

export interface DatePickerProps {
    date: Date;
    onConfirm: (date: Date) => void;
}

const DatePicker: React.FC<DatePickerProps> = ({ date, onConfirm }) => {
    const [isVisible, setIsVisible] = useState(false);

    const showPicker = () => setIsVisible(true);
    const hidePicker = () => setIsVisible(false);

    const handleConfirm = (selectedDate: Date) => {
        onConfirm(selectedDate);
        hidePicker();
    };

    return (
        <>
            <TouchableOpacity style={styles.button} onPress={showPicker}>
                <Text style={styles.buttonText}>{date.toLocaleDateString()}</Text>
            </TouchableOpacity>
            <DateTimePickerModal
                isVisible={isVisible}
                mode="date" // On sélectionne uniquement la date
                onConfirm={handleConfirm}
                onCancel={hidePicker}
            />
        </>
    );
};

const styles = StyleSheet.create({
    button: {
        padding: 12,
        borderWidth: 1,
        borderColor: '#ccc',
        borderRadius: 4,
        backgroundColor: '#f2f2f2',
        marginVertical: 8,
        alignItems: 'center',
    },
    buttonText: {
        fontSize: 16,
    },
});

export default DatePicker;

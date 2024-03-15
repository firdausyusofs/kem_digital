import { Box, Button, Group, rem, Text, UnstyledButton } from '@mantine/core';
import classes from '../../styles/Header.module.css';
import React from 'react';

export const Header = () => {
    return (
        <header>
            <Group justify="space-between" h="100%">
                <Text fw="bold" size={rem(20)}>Kem Digital</Text>

                <UnstyledButton 
                    className={classes.signInButton} 
                    h={40} 
                    component="a"
                    href="/login"
                    visibleFrom='sm'
                >
                    <span className={classes.btnText}>
                        Log Out
                    </span>
                </UnstyledButton>
            </Group>
        </header>
    );
}

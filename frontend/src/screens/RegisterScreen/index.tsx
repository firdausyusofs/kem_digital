import { Anchor, BackgroundImage, Box, Button, Flex, rem, Select, Stack, Text, TextInput, Title } from '@mantine/core';
import React from 'react';
import classes from '../../styles/Login.module.css';

const RegisterScreen: React.FC = () => {
    return (
        // <BackgroundImage src="https://images.unsplash.com/photo-1595542028715-8412db1e862f?q=80&w=2287&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D">
            <Flex 
                align="center" 
                justify="center" 
                h="100vh"
            >
                <Flex
                    justify="center"
                    align="flex-start"
                    direction="column"
                    className={classes.loginWrapper}
                >
                    <Box 
                        className={classes.form}
                        maw={700}
                        bg="rgba(245, 245, 245, 1.0)"
                        p={rem(60)}
                    >
                        <Title fw={700} size={rem(32)} ta="center" mb={rem(56)}>Register</Title>
                        <form style={{ width: '100%' }}>
                            <Stack justify="flex-end" mt="md">
                                <TextInput
                                    withAsterisk
                                    type="text"
                                    placeholder="First Name"
                                    variant="lf_textinput"
                                    // leftSection={<IconMail style={{ width: rem(16), height: rem(16) }} />}
                                    // {...form.getInputProps('email')}
                                />

                                <TextInput
                                    withAsterisk
                                    type="text"
                                    placeholder="Last Name"
                                    variant="lf_textinput"
                                    // leftSection={<IconMail style={{ width: rem(16), height: rem(16) }} />}
                                    // {...form.getInputProps('email')}
                                />

                                <TextInput
                                    withAsterisk
                                    type="text"
                                    placeholder="Email"
                                    variant="lf_textinput"
                                    // leftSection={<IconMail style={{ width: rem(16), height: rem(16) }} />}
                                    // {...form.getInputProps('email')}
                                />

                                <TextInput
                                    withAsterisk
                                    type="password"
                                    placeholder="Password"
                                    variant="lf_textinput"
                                    // leftSection={<IconKey style={{ width: rem(16), height: rem(16) }} />}
                                    // {...form.getInputProps('password')}
                                />

                                <Select
                                    placeholder="Role"
                                    variant="lf_select"
                                    data={['React', 'Angular', 'Vue', 'Svelte']}
                                />

                                <Button type="submit" variant="lf_primarybtn" radius="xl" h={rem(45)}>Register</Button>
                                <Text size={rem(12)} fw={400} ta='center'>
                                    Have an account?{' '}<Anchor href="/">Log in</Anchor>
                                </Text>
                            </Stack>
                        </form>
                    </Box>
                </Flex>

            </Flex>
        // </BackgroundImage>
    )
}

export default RegisterScreen;

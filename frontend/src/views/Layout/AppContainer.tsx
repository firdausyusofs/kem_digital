import { AppShell, Box, Burger, rem, Text } from "@mantine/core";
import React from "react";
import { Outlet } from "react-router-dom";
import { Header } from "./Header";

export const AppContainer: React.FC = () => {
    return (
        <AppShell>
            <AppShell.Header p="md">
                <Header />
            </AppShell.Header>

            <AppShell.Main mt={70}>
                <Outlet />
            </AppShell.Main>
        </AppShell>
    );
}

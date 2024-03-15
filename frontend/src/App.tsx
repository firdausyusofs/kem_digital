import { createTheme, MantineProvider } from '@mantine/core';
import React from 'react';
import { Route, Routes } from 'react-router-dom';
import InventoryScreen from './screens/InventoryScreen';
import LoginScreen from './screens/LoginScreen';
import RegisterScreen from './screens/RegisterScreen';
import { AppContainer } from './views/Layout/AppContainer';

const theme = createTheme({
  fontFamily: 'Montserrat, sans-serif',
  primaryColor: 'teal',
});

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<AppContainer />}>
        <Route index element={<LoginScreen />} />
        <Route path="/register" element={<RegisterScreen />} />
        <Route path="/inventory" element={<InventoryScreen />} />

        <Route
          path={"*"}
          element={
            <div style={{ display: "flex", justifyContent: "center" }}>
              Soon here!!!
            </div>
          }
        />
      </Route>
    </Routes>
  );
}

export default App;

import { MantineProvider } from "@mantine/core";
import { ModalsProvider } from "@mantine/modals";
import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import App from "./App";
import { ReduxProvider } from "./redux/provider";
import { theme } from "./theme";

const MantineRoot: React.FC = () => {

  return (
    <ReduxProvider>
      <MantineProvider
        theme={theme}
      >
        <ModalsProvider>
          <Router>
            <Routes>
              <Route path="/*" element={<App />} />
            </Routes>
          </Router>
        </ModalsProvider>
      </MantineProvider>
    </ReduxProvider>
  )

};

export default MantineRoot;

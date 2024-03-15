import { MantineProvider } from "@mantine/core";
import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import App from "./App";
import { theme } from "./theme";

const MantineRoot: React.FC = () => {

  return (
    <MantineProvider
      theme={theme}
    >
      <Router>
        <Routes>
          <Route path="/*" element={<App />} />
        </Routes>
      </Router>
    </MantineProvider>
  )

};

export default MantineRoot;

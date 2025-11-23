import React, { useState } from "react";
import "./App.css";
import LoginPage from "./pages/LoginPage";
import ItemsPage from "./pages/ItemsPage";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function App() {
  // start on login page
  const [screen, setScreen] = useState("login"); // "login" | "items"

  function handleLogin() {
    setScreen("items");
  }

  function handleLogout() {
    localStorage.removeItem("token");
    setScreen("login");
  }

  return (
    <>
      {screen === "login" && <LoginPage onLogin={handleLogin} />}

      {screen === "items" && <ItemsPage onLogout={handleLogout} />}

      {/* Toast container shown on all screens */}
      <ToastContainer position="top-right" autoClose={3000} />
    </>
  );
}

export default App;

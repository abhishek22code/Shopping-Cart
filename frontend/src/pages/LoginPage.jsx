import React, { useState } from "react";
import { apiPost } from "../api";
import { toast } from "react-toastify";

function LoginPage({ onLogin }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  async function handleLogin(e) {
    e.preventDefault();

    const result = await apiPost("/users/login", { username, password });

    if (result && result.token) {
      localStorage.setItem("token", result.token);
      toast.success("Login successful");
      onLogin();
    } else {
      toast.error("Invalid username or password");
    }
  }

  return (
    <div className="container">
      <h2>Login</h2>

      <form onSubmit={handleLogin}>
        <div style={{ marginBottom: 15 }}>
          <label>Username</label>
          <input
            placeholder="Enter username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>

        <div style={{ marginBottom: 15 }}>
          <label>Password</label>
          <input
            placeholder="Enter password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <button style={{ width: "100%", marginTop: 10 }} type="submit">
          Login
        </button>
      </form>
    </div>
  );
}

export default LoginPage;

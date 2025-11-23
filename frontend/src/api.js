
const API_URL =
  process.env.REACT_APP_API_URL || "http://localhost:8080";

export async function apiGet(path, token) {
  const res = await fetch(API_URL + path, {
    headers: {
      "X-Auth-Token": token || "",
    },
  });
  return res.json();
}

export async function apiPost(path, body, token) {
  const res = await fetch(API_URL + path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Auth-Token": token || "",
    },
    body: JSON.stringify(body),
  });
  return res.json();
}

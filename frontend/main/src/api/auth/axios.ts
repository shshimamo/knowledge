import axios from "axios";

export const httpClient = axios.create({
  baseURL: process.env.NODE_ENV === 'production' ? 'http://backend.auth.shshimamo.com' : 'http://localhost:8081',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
});
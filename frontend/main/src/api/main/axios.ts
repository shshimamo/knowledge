import axios from "axios";

export const httpClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_MAIN_API_URL || (process.env.NODE_ENV === 'production' ? 'http://backend.main.shshimamo.com' : 'http://localhost:8080'),
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
});
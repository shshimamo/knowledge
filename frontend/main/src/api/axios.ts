import axios from "axios";

export const authClient = axios.create({
  baseURL: 'http://localhost:80',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
});
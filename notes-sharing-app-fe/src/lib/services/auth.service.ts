import { LoginResponse, RegisterResponse } from "@/types/auth";
import api from "./api";
import Cookies from "js-cookie";
import axios, { AxiosError } from "axios";

export const authService = {
  async login(data: { username: string; password: string }) {
    try {
      const response = await api.post<LoginResponse>("login", data);
      return response.data;
    } catch (error: unknown) {
      if (axios.isAxiosError(error) && error.response) {
        return Promise.reject(error.response.data);
      }
      return Promise.reject({ message: "Something went wrong" });
    }
  },

  async register(data: { username: string; password: string }) {
    try {
      const response = await api.post<RegisterResponse>("register", data);
      return response.data;
    } catch (error: unknown) {
      if (axios.isAxiosError(error) && error.response) {
        return Promise.reject(error.response.data);
      }
      return Promise.reject({ message: "Something went wrong" });
    }
  },

  async logout() {
    try {
      await api.post("logout");
      Cookies.remove("access_token");
      return { message: "Logged out successfully" };
    } catch (error: unknown) {
      return Promise.reject({ message: "Logout failed" });
    }
  },
};

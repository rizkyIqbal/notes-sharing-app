import api from "./api";
import { UserResponse } from "@/types/user";
import axios from "axios";

export const userService = {
  async getUserByID(): Promise<UserResponse | null> {
    try {
      const response = await api.get<UserResponse>("/profile", {
        withCredentials: true,
      });
      return response.data;
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 401) {
          return null;
        }
      }
      console.error("Error fetching user:", error);
      throw error;
    }
  },
};

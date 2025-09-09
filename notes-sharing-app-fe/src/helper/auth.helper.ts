import { authService } from "@/lib/services/auth.service";
import type { LoginResponse, RegisterResponse } from "@/types/auth";

export async function loginUser(
  username: string,
  password: string
): Promise<string | null> {
  try {
    const response: LoginResponse = await authService.login({
      username,
      password,
    });
    return response.message ?? null;
  } catch (error) {
    console.error("Login failed:", error);
    return null;
  }
}

export async function registerUser(
  username: string,
  password: string
): Promise<string | null> {
  try {
    const response: RegisterResponse = await authService.register({
      username,
      password,
    });
    return response.message ?? null;
  } catch (error) {
    console.error("Register failed:", error);
    return null;
  }
}

export async function logoutUser(): Promise<string | null> {
  try {
    const response = await authService.logout(); // also returns { status, message }
    return response.message ?? "Logged out";
  } catch (error) {
    console.error("Logout failed:", error);
    return null;
  }
}

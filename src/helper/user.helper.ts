import { userService } from "@/lib/services/user.service";
import { User } from "@/types/user";

export async function fetchUserByID(
  setUser: (user: User | null) => void
): Promise<User | null> {
  try {
    const response = await userService.getUserByID()

    if (response && response.data) {
      setUser(response.data);
      return response.data;
    }

    setUser(null);
    return null;
  } catch (error) {
    console.error("Failed to fetch user:", error);
    setUser(null);
    return null;
  }
}

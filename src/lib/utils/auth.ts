import { jwtDecode } from "jwt-decode";

type JWTPayload = {
  exp: number;
  iat?: number;
  user_id?: number;
  [key: string]: any;
};

export function isTokenExpired(token: string): boolean {
  try {
    const decoded: JWTPayload = jwtDecode(token);
    if (!decoded.exp) return true;

    const now = Date.now() / 1000;
    return decoded.exp < now;
  } catch (error) {
    console.error("Invalid token:", error);
    return true;
  }
}

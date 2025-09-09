import {jwtDecode} from "jwt-decode";

type JWTPayload = {
  exp: number;
  iat?: number;
  user_id?: number;
  [key: string]: unknown;
};

export function isTokenExpired(token: string): boolean {
  try {
    const decoded = jwtDecode<JWTPayload>(token); // specify type here
    if (!decoded.exp) return true;

    const now = Date.now() / 1000;
    return decoded.exp < now;
  } catch (error: unknown) {
    console.error("Invalid token:", error);
    return true;
  }
}

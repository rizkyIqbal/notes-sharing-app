import { NextRequest, NextResponse } from "next/server";
import { isTokenExpired } from "./lib/utils/auth";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:4000";

async function refreshAccessToken() {
  try {
    const res = await fetch(`${API_URL}/auth/refresh`, {
      method: "POST",
      credentials: "include",
      headers: { "Content-Type": "application/json" },
    });

    if (!res.ok) return null;

    const data = await res.json();
    return data.message;
  } catch (err) {
    console.error("Refresh token failed:", err);
    return null;
  }
}

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  const accessToken = req.cookies.get("access_token")?.value;
  const refreshToken = req.cookies.get("refresh_token")?.value;

  // Pages inaccessible if logged in
  if (pathname.startsWith("/login") || pathname.startsWith("/register")) {
    if (accessToken && !isTokenExpired(accessToken)) {
      return NextResponse.redirect(new URL("/", req.url));
    }
  }

  // Protected routes
  const protectedRoutes = ["/mynotes", "/note/create", "/note/edit"];
  const isProtected = protectedRoutes.some((route) =>
    pathname.startsWith(route)
  );

  if (isProtected) {
    if (accessToken && !isTokenExpired(accessToken)) {
      return NextResponse.next();
    }

    if (refreshToken) {
      const newAccessToken = await refreshAccessToken();

      if (newAccessToken) {
        const res = NextResponse.next();
        res.cookies.set("access_token", newAccessToken, {
          httpOnly: true,
          path: "/",
        });
        return res;
      }
    }

    return NextResponse.redirect(new URL("/login", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/login", "/register", "/mynotes/:path*", "/note/:path*"],
};

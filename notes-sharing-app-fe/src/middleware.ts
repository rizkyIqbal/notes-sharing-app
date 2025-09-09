import { NextRequest, NextResponse } from "next/server";
import { isTokenExpired } from "./lib/utils/auth";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:4000";

async function refreshAccessToken(refreshToken: string) {
  try {
    const res = await fetch(`${API_URL}/auth/refresh`, {
      method: "POST",
      credentials: "include", // send httpOnly cookies
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ refresh_token: refreshToken }), // optional if your backend reads httpOnly
    });

    if (!res.ok) return null;

    const data = await res.json();
    return data.access_token; // make sure backend returns new access token
  } catch (err) {
    console.error("Refresh token failed:", err);
    return null;
  }
}

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  const accessToken = req.cookies.get("access_token")?.value;
  const refreshToken = req.cookies.get("refresh_token")?.value;

  // Redirect logged-in users away from login/register
  if (pathname.startsWith("/login") || pathname.startsWith("/register")) {
    if (accessToken && !isTokenExpired(accessToken)) {
      return NextResponse.redirect(new URL("/", req.url));
    }
    return NextResponse.next();
  }

  // Protected routes
  const protectedRoutes = ["/mynotes", "/note/create", "/note/edit"];
  const isProtected = protectedRoutes.some((route) =>
    pathname.startsWith(route)
  );

  if (isProtected) {
    // 1. Valid access token → allow
    if (accessToken && !isTokenExpired(accessToken)) {
      return NextResponse.next();
    }

    // 2. Access token invalid/expired but refresh exists → try refresh
    if (refreshToken) {
      const newAccessToken = await refreshAccessToken(refreshToken);
      if (newAccessToken) {
        const res = NextResponse.next();
        res.cookies.set("access_token", newAccessToken, {
          httpOnly: true,
          path: "/",
        });
        return res;
      }
      // Optional: clear invalid tokens
      const res = NextResponse.redirect(new URL("/login", req.url));
      res.cookies.delete({ name: "access_token", path: "/" });
      res.cookies.delete({ name: "refresh_token", path: "/" });
      return res;
    }

    // 3. No valid tokens → redirect to login
    return NextResponse.redirect(new URL("/login", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/login", "/register", "/mynotes/:path*", "/note/:path*"],
};

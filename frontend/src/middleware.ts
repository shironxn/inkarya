import { NextResponse, type NextRequest } from "next/server";
import { stackServerApp } from "./stack";

const PROTECTED_ROUTES = ["/forum", "/kursus"];

const AUTH_ROUTES = ["/masuk", "/daftar"];

export async function middleware(request: NextRequest) {
  const user = await stackServerApp.getUser();
  const { pathname } = request.nextUrl;

  if (user && pathname === "/" && user.clientMetadata.onboarded) {
    return NextResponse.redirect(new URL("/lowongan", request.url));
  }

  if (!user && PROTECTED_ROUTES.some((route) => pathname.startsWith(route))) {
    return NextResponse.redirect(new URL("/masuk", request.url));
  }

  if (user && AUTH_ROUTES.some((route) => pathname.startsWith(route))) {
    return NextResponse.redirect(new URL("/lowongan", request.url));
  }

  if (user && !user.clientMetadata.onboarded) {
    return NextResponse.redirect(new URL("/onboarding", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    "/",
    "/lowongan",
    "/lowongan/:path*",
    "/forum",
    "/forum/:path*",
    "/kursus",
    "/kursus/:path*",
    "/masuk",
    "/daftar",
  ],
};

"use server";

import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Menu } from "lucide-react";
import Image from "next/image";
import Link from "next/link";

export default async function Navbar() {
  const navItems = ["Lowongan", "Forum", "Kursus"];
  const authItems = [
    { label: "Masuk", href: "/masuk", variant: "ghost" },
    { label: "Daftar", href: "/daftar", variant: "default" },
  ] as const;

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-white/95 backdrop-blur supports-[backdrop-filter]:bg-white/60">
      <div className="container mx-auto flex h-16 items-center justify-between px-4">
        <Link
          href="/"
          className="flex items-center gap-2 transition-opacity hover:opacity-90"
        >
          <Image
            src="/icons/inkarya.svg"
            alt="InKarya Logo"
            width={120}
            height={32}
            className="h-8 w-auto"
            priority
          />
        </Link>

        {/* Desktop Navigation */}
        <nav className="hidden md:block">
          <ul className="flex items-center gap-8">
            {navItems.map((item) => (
              <li key={item}>
                <Link
                  href={`/${item.toLowerCase()}`}
                  className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
                >
                  {item}
                </Link>
              </li>
            ))}
          </ul>
        </nav>

        {/* Desktop Auth Buttons */}
        <div className="hidden items-center gap-4 md:flex">
          {authItems.map((item) => (
            <Button
              key={item.href}
              variant={item.variant}
              className={`font-medium ${item.variant === "default" ? "bg-primary text-white hover:bg-primary/90" : ""}`}
              asChild
            >
              <Link href={item.href}>{item.label}</Link>
            </Button>
          ))}
        </div>

        {/* Mobile Navigation */}
        <Sheet>
          <SheetTrigger asChild>
            <Button variant="ghost" size="icon" className="md:hidden">
              <Menu className="h-5 w-5" />
              <span className="sr-only">Toggle menu</span>
            </Button>
          </SheetTrigger>
          <SheetContent side="right" className="w-full border-l sm:max-w-sm">
            <SheetHeader className="border-b pb-4">
              <SheetTitle className="text-lg font-semibold">Menu</SheetTitle>
            </SheetHeader>
            <nav className="mt-6 px-4">
              <ul className="space-y-4">
                {navItems.map((item) => (
                  <li key={item}>
                    <Link
                      href={`/${item.toLowerCase()}`}
                      className="flex items-center py-2 text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
                    >
                      {item}
                    </Link>
                  </li>
                ))}
                <li className="border-t pt-4">
                  <div className="mt-4 flex flex-col gap-2">
                    {authItems.map((item) => (
                      <Button
                        key={item.href}
                        variant={item.variant}
                        className={`w-full font-medium ${item.variant === "default" ? "bg-primary text-white hover:bg-primary/90" : ""}`}
                        asChild
                      >
                        <Link href={item.href}>{item.label}</Link>
                      </Button>
                    ))}
                  </div>
                </li>
              </ul>
            </nav>
          </SheetContent>
        </Sheet>
      </div>
    </header>
  );
}


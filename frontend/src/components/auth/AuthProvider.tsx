"use client";

import { Button } from "@/components/ui/button";
import { useStackApp } from "@stackframe/stack";
import Image from "next/image";

// Social login providers array
const socialProviders = [
  {
    name: "Google",
    icon: "/icons/google.svg",
    alt: "Google",
  },
  {
    name: "LinkedIn",
    icon: "/icons/linkedin.svg",
    alt: "LinkedIn",
  },
  {
    name: "GitHub",
    icon: "/icons/github.svg",
    alt: "GitHub",
  },
];

export function AuthProvider() {
  const app = useStackApp();

  const handleLogin = async (provider: string) => {
    await app.signInWithOAuth(provider);
  };

  return (
    <div className="flex justify-center gap-3">
      {socialProviders.map((provider, index) => (
        <Button
          key={index}
          type="button"
          variant="outline"
          disabled={provider.name === "LinkedIn"}
          className="flex-1 py-5"
          onClick={() => handleLogin(provider.name)}
        >
          <Image src={provider.icon} alt={provider.alt} width={20} height={20} />
        </Button>
      ))}
    </div>
  );
}

"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useStackApp } from "@stackframe/stack";
import { Lock, Mail } from "lucide-react";
import Link from "next/link";
import { useState } from "react";
import { toast } from "sonner";
import { AuthProvider } from "./AuthProvider";
import { useRouter } from "next/navigation";

interface LoginFormData {
  email: string;
  password: string;
}

// Input fields array
const inputFields = [
  {
    name: "email",
    type: "email",
    placeholder: "Email",
    icon: <Mail className="absolute left-3 top-3 h-5 w-5 text-gray-400" />,
  },
  {
    name: "password",
    type: "password",
    placeholder: "Password",
    icon: <Lock className="absolute left-3 top-3 h-5 w-5 text-gray-400" />,
  },
];

export function LoginForm() {
  const app = useStackApp();
  const router = useRouter();

  const [formData, setFormData] = useState<LoginFormData>({
    email: "",
    password: "",
  });
  const [isLoading, setIsLoading] = useState(false);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Validate form data
    if (!formData.email) {
      toast.error("Email diperlukan", {
        description: "Silakan masukkan email Anda",
      });
      return;
    }

    if (!formData.password) {
      toast.error("Password diperlukan", {
        description: "Silakan masukkan password Anda",
      });
      return;
    }

    setIsLoading(true);

    const result = await app.signInWithCredential({
      email: formData.email,
      password: formData.password,
      noRedirect: true,
    });

    // Check for specific error status
    if (result.status === "error") {
      toast.error("Login gagal", {
        description: result.error?.message || "Email atau password salah",
      });
      return;
    }

    router.push("/onboarding");
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      {inputFields.map((field, index) => (
        <div key={index} className="relative">
          {field.icon}
          <Input
            name={field.name}
            type={field.type}
            placeholder={field.placeholder}
            className="pl-10 py-6"
            value={formData[field.name as keyof LoginFormData]}
            onChange={handleInputChange}
            required
          />
        </div>
      ))}

      <div className="flex justify-end">
        <Link
          href="/lupa-password"
          className="text-sm text-primary hover:underline"
        >
          Lupa password?
        </Link>
      </div>

      <Button
        type="submit"
        className="w-full py-6 bg-primary text-white hover:bg-primary/90"
        disabled={isLoading}
      >
        {isLoading ? "Memproses..." : "Masuk"}
      </Button>

      <div className="relative my-4">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t"></span>
        </div>
        <div className="relative flex justify-center text-xs">
          <span className="bg-background px-2 text-muted-foreground">
            Atau masuk dengan
          </span>
        </div>
      </div>

      <AuthProvider />

      <div className="mt-6 text-center">
        <p className="text-sm text-muted-foreground">
          Belum punya akun?{" "}
          <Link
            href="/daftar"
            className="text-primary font-medium hover:underline"
          >
            Daftar sekarang
          </Link>
        </p>
      </div>
    </form>
  );
}


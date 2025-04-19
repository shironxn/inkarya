"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useStackApp } from "@stackframe/stack";
import { Lock, Mail, User } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState, useTransition } from "react";
import { toast } from "sonner";
import { AuthProvider } from "./AuthProvider";

interface RegisterFormData {
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
}

// Input fields array
const inputFields = [
  {
    name: "name",
    type: "text",
    placeholder: "Nama Lengkap",
    icon: <User className="absolute left-3 top-3 h-5 w-5 text-gray-400" />,
  },
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
  {
    name: "confirmPassword",
    type: "password",
    placeholder: "Konfirmasi Password",
    icon: <Lock className="absolute left-3 top-3 h-5 w-5 text-gray-400" />,
  },
];

export function RegisterForm() {
  const app = useStackApp();
  const router = useRouter();
  const [isLoading, startTransition] = useTransition();
  const [formData, setFormData] = useState<RegisterFormData>({
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

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

    if (formData.password !== formData.confirmPassword) {
      toast.error("Password tidak cocok", {
        description: "Password dan konfirmasi password harus sama",
      });
      return;
    }

    startTransition(async () => {
      const result = await app.signUpWithCredential({
        email: formData.email,
        password: formData.password,
        noRedirect: true,
      });

      // Check for specific error status
      if (result.status === "error") {
        toast.error("Registrasi gagal", {
          description:
            result.error?.message || "Terjadi kesalahan saat registrasi",
        });
        return;
      }

      // Update user metadata with onboarded status
      const user = await app.getUser();
      if (user) {
        await user.update({
          displayName: formData.name,
          clientMetadata: {
            onboarded: false,
          },
        });
      }

      toast.success("Registrasi berhasil", {
        description: "Silakan lengkapi profil Anda",
      });

      router.push("/onboarding");
    });
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
            value={formData[field.name as keyof RegisterFormData]}
            onChange={handleInputChange}
            required
          />
        </div>
      ))}

      <Button
        type="submit"
        className="w-full py-6 bg-primary text-white hover:bg-primary/90"
        disabled={isLoading}
      >
        {isLoading ? "Memproses..." : "Daftar"}
      </Button>

      <div className="relative my-4">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t"></span>
        </div>
        <div className="relative flex justify-center text-xs">
          <span className="bg-background px-2 text-muted-foreground">
            Atau daftar dengan
          </span>
        </div>
      </div>

      <AuthProvider />

      <div className="mt-6 text-center">
        <p className="text-sm text-muted-foreground">
          Sudah punya akun?{" "}
          <Link
            href="/masuk"
            className="text-primary font-medium hover:underline"
          >
            Masuk sekarang
          </Link>
        </p>
      </div>
    </form>
  );
}


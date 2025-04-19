import Image from "next/image"
import { LoginForm } from "@/components/auth/LoginForm"

export default function Page() {
  return (
    <main className="min-h-screen">
      <div className="container mx-auto grid min-h-[calc(100vh-73px)] md:grid-cols-2">
        <div className="flex flex-col justify-center px-4 py-12 md:px-8 lg:px-12">
          <div className="mx-auto w-full max-w-md">
            <div className="mb-8">
              <h1 className="text-3xl font-bold mb-2">Masuk ke Akunmu</h1>
              <p className="text-sm text-muted-foreground">
                Masukkan email dan password untuk mengakses akun Anda
              </p>
            </div>

            <LoginForm />
          </div>
        </div>

        <div className="hidden md:flex items-center justify-center">
          <Image
            src="/auth/login.svg"
            alt="Login Illustration"
            width={500}
            height={500}
            className="max-w-md"
          />
        </div>
      </div>
    </main>
  )
}

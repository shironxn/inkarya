import Image from "next/image"
import { RegisterForm } from "@/components/auth/RegisterForm"

export default function Page() {
  return (
    <main className="min-h-screen">
      <div className="container mx-auto grid min-h-[calc(100vh-73px)] md:grid-cols-2">
        <div className="hidden md:flex items-center justify-center">
          <Image 
            src="/auth/register.svg" 
            alt="Register Illustration" 
            width={500} 
            height={500} 
            className="max-w-md"
          />
        </div>
        
        <div className="flex flex-col justify-center px-4 py-12 md:px-8 lg:px-12">
          <div className="mx-auto w-full max-w-md">
            <div className="mb-8">
              <h1 className="mb-2 text-3xl font-bold tracking-tight">Buat Akun Baru</h1>
              <p className="text-sm text-muted-foreground">
                Daftar untuk mendapatkan akses ke platform kami
              </p>
            </div>
            
            <RegisterForm />
          </div>
        </div>
      </div>
    </main>
  )
}

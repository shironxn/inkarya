import type { Metadata } from "next";
import { StackProvider, StackTheme } from "@stackframe/stack";
import { stackServerApp } from "../stack";
import { Inter, Poppins } from "next/font/google";
import "./globals.css";
import Navbar from "@/components/navbar/Navbar";
import { Toaster } from "sonner";
import Footer from "@/components/footer/Footer";

const inter = Inter({
  variable: "--font-inter",
  subsets: ["latin"],
  display: "swap",
});

const poppins = Poppins({
  variable: "--font-poppins",
  weight: ["300", "400", "500", "600", "700"],
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "InKarya",
  description: "Platform Karier Inklusif untuk Penyandang Disabilitas",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${inter.variable} ${poppins.variable} antialiased`}
      >
        <Navbar />
        <StackProvider app={stackServerApp}>
          <StackTheme>
            <StackProvider app={stackServerApp}>
              <StackTheme>
                {children}
                <Toaster richColors position="top-right" />
              </StackTheme>
            </StackProvider>
          </StackTheme>
        </StackProvider>
        <Footer />
      </body>
    </html>
  );
}

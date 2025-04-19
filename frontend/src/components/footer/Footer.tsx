import Link from "next/link"
import { MapPin, Mail, Phone, Instagram, Github } from "lucide-react"
import Image from "next/image"

export default function Footer() {
  return (
    <footer className="w-full bg-white py-12 px-4 md:px-6 lg:px-8 border-t border-gray-200">
      <div className="container mx-auto max-w-7xl">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {/* Logo and Description */}
          <div className="space-y-4">
            <div className="flex items-center">
              <div className="relative h-10 w-40">
                <div className="flex items-center">
                  <Image src="/icons/inkarya.svg" alt="InKarya Logo" className="h-8" width={160} height={160} />
                </div>
              </div>
            </div>
            <p className="text-gray-600 max-w-xs">
              Platform pencarian kerja inklusif untuk semua penyandang disabilitas di Indonesia.
            </p>
            <div className="flex items-start space-x-2">
              <MapPin className="text-[#FF8C42] h-5 w-5 mt-0.5 flex-shrink-0" />
              <p className="text-gray-600">Jl. Merdeka No. 123, Jakarta Pusat, Indonesia</p>
            </div>
          </div>

          {/* Navigation */}
          <div className="space-y-4">
            <h3 className="text-lg font-semibold">Navigasi</h3>
            <ul className="space-y-2">
              <li>
                <Link href="/lowongan" className="text-gray-600 hover:text-[#FF8C42]">
                  Lowongan
                </Link>
              </li>
              <li>
                <Link href="/forum" className="text-gray-600 hover:text-[#FF8C42]">
                  Forum
                </Link>
              </li>
              <li>
                <Link href="/kursus" className="text-gray-600 hover:text-[#FF8C42]">
                  Kursus
                </Link>
              </li>
            </ul>
          </div>

          {/* Help */}
          <div className="space-y-4">
            <h3 className="text-lg font-semibold">Bantuan</h3>
            <ul className="space-y-2">
              <li>
                <Link href="#" className="text-gray-600 hover:text-[#FF8C42]">
                  Pusat Bantuan
                </Link>
              </li>
              <li>
                <Link href="#" className="text-gray-600 hover:text-[#FF8C42]">
                  FAQ
                </Link>
              </li>
              <li>
                <Link href="#" className="text-gray-600 hover:text-[#FF8C42]">
                  Kebijakan Privasi
                </Link>
              </li>
              <li>
                <Link href="#" className="text-gray-600 hover:text-[#FF8C42]">
                  Syarat & Ketentuan
                </Link>
              </li>
            </ul>
          </div>

          {/* Contact */}
          <div className="space-y-4">
            <h3 className="text-lg font-semibold">Kontak</h3>
            <div className="space-y-2">
              <div className="flex items-center space-x-2">
                <Mail className="text-[#FF8C42] h-5 w-5" />
                <Link href="mailto:info@inkarya.id" className="text-gray-600 hover:text-[#FF8C42]">
                  info@inkarya.id
                </Link>
              </div>
              <div className="flex items-center space-x-2">
                <Phone className="text-[#FF8C42] h-5 w-5" />
                <Link href="tel:+6212345678" className="text-gray-600 hover:text-[#FF8C42]">
                  +62 12345678
                </Link>
              </div>
            </div>

            {/* Social Media */}
            <div className="space-y-2">
              <h3 className="text-lg font-semibold">Ikuti Kami</h3>
              <div className="flex space-x-4">
                <Link href="#" aria-label="Instagram">
                  <Instagram className="h-6 w-6 text-[#FF8C42] hover:text-[#0891B2]" />
                </Link>
                <Link href="#" aria-label="Twitter">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    className="h-6 w-6 text-[#FF8C42] hover:text-[#0891B2]"
                  >
                    <path d="M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6 2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4-.9-4.2 4-6.6 7-3.8 1.1 0 3-1.2 3-1.2z" />
                  </svg>
                </Link>
                <Link href="#" aria-label="GitHub">
                  <Github className="h-6 w-6 text-[#FF8C42] hover:text-[#0891B2]" />
                </Link>
              </div>
            </div>
          </div>
        </div>

        {/* Divider */}
        <div className="border-t border-gray-200 my-8"></div>

        {/* Copyright */}
        <div className="text-center text-gray-600">
          <p>© 2025 InKarya. All rights reserved.</p>
        </div>
      </div>
    </footer>
  )
}

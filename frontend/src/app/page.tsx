import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Button } from "@/components/ui/button";
import {
  Accessibility,
  Award,
  CheckCircle,
  Cpu,
  MessageSquare,
  Users,
} from "lucide-react";
import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div className="min-h-screen bg-[#fafafa]">
      {/* Hero Section */}
      <section className="py-16 md:py-24">
        <div className="container px-4 mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
            <div className="space-y-6">
              <h1>
                Siap Menjadi Bagian dari Gerakan Inklusif?
              </h1>
              <p className="text-gray-600">
                Bergabunglah dengan InKarya dan mulai perjalanan karier Anda –
                tanpa batasan.
              </p>
              <div className="flex flex-col sm:flex-row gap-4 pt-4">
                <Link href="/masuk">
                  <Button className="bg-[#ff8c42] hover:bg-[#ff8c42]/90 text-white px-8 py-6">
                    Masuk
                  </Button>
                </Link>
                <Link href="#tentang-kami">
                  <Button
                    variant="outline"
                    className="border-[#ff8c42] text-[#ff8c42] hover:bg-[#ff8c42]/10 hover:text-[#ff8c42]/90 px-8 py-6"
                  >
                    Jelajahi
                  </Button>
                </Link>
              </div>
            </div>
            <div className="flex justify-center">
              <Image
                src="/home/collaborate.svg"
                alt="Ilustrasi Inklusif"
                width={500}
                height={400}
                className="w-full max-w-[500px]"
              />
            </div>
          </div>
        </div>
      </section>

      {/* About Section */}
      <section id="tentang-kami" className="py-24 bg-white">
        <div className="container px-4 mx-auto text-center">
          <h2 className="mb-8">
            Apa itu InKarya?
          </h2>
          <p className="text-gray-600 max-w-3xl mx-auto">
            InKarya adalah platform inklusif yang didesain khusus untuk membuka
            akses dunia kerja bagi penyandang disabilitas. Kami menjembatani
            talenta luar biasa dengan perusahaan yang berkomitmen terhadap
            kesetaraan.
          </p>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 md:py-24">
        <div className="container px-4 mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
            <div className="flex justify-center">
              <Image
                src="/home/feature.svg"
                alt="Fitur Ilustrasi"
                width={400}
                height={400}
                className="w-full max-w-[400px]"
              />
            </div>
            <div className="space-y-8">
              <h2>
                Fitur-Fitur yang Tersedia
              </h2>
              <div className="space-y-6">
                <div className="flex items-start gap-4">
                  <div className="bg-[#ff8c42]/10 p-3 rounded-lg">
                    <Cpu className="h-6 w-6 text-[#ff8c42]" />
                  </div>
                  <div>
                    <h3 className="text-xl font-semibold font-inter">
                      Sistem Pencocokan AI
                    </h3>
                    <p className="text-gray-600 font-poppins">
                      Teknologi AI kami mencocokkan kemampuan Anda dengan
                      lowongan yang tepat.
                    </p>
                  </div>
                </div>
                <div className="flex items-start gap-4">
                  <div className="bg-[#4f46e5]/10 p-3 rounded-lg">
                    <Accessibility className="h-6 w-6 text-[#4f46e5]" />
                  </div>
                  <div>
                    <h3 className="text-xl font-semibold font-inter">
                      Aksesibilitas Penuh
                    </h3>
                    <p className="text-gray-600 font-poppins">
                      Platform kami didesain agar dapat diakses oleh semua jenis
                      disabilitas.
                    </p>
                  </div>
                </div>
                <div className="flex items-start gap-4">
                  <div className="bg-[#10b981]/10 p-3 rounded-lg">
                    <Award className="h-6 w-6 text-[#10b981]" />
                  </div>
                  <div>
                    <h3 className="text-xl font-semibold font-inter">
                      Pelatihan & Pengembangan
                    </h3>
                    <p className="text-gray-600 font-poppins">
                      Akses ke ratusan kursus untuk meningkatkan keterampilan
                      Anda.
                    </p>
                  </div>
                </div>
                <div className="flex items-start gap-4">
                  <div className="bg-[#8b5cf6]/10 p-3 rounded-lg">
                    <CheckCircle className="h-6 w-6 text-[#8b5cf6]" />
                  </div>
                  <div>
                    <h3 className="text-xl font-semibold font-inter">
                      Verifikasi Perusahaan Inklusif
                    </h3>
                    <p className="text-gray-600 font-poppins">
                      Kami memverifikasi perusahaan untuk memastikan lingkungan
                      kerja yang inklusif.
                    </p>
                  </div>
                </div>
                <div className="flex items-start gap-4">
                  <div className="bg-[#ec4899]/10 p-3 rounded-lg">
                    <Users className="h-6 w-6 text-[#ec4899]" />
                  </div>
                  <div>
                    <h3 className="text-xl font-semibold font-inter">
                      Komunitas & Mentorship
                    </h3>
                    <p className="text-gray-600 font-poppins">
                      Terhubung dengan komunitas dan mentor yang mendukung
                      perjalanan karier Anda.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="py-16 bg-white">
        <div className="container px-4 mx-auto">
          <h2 className="text-3xl md:text-4xl font-bold mb-12 text-center font-inter">
            Statistik InKarya
          </h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
            <div className="flex flex-col items-center text-center space-y-2">
              <div className="text-[#ff8c42] mb-2">⭐</div>
              <h3 className="text-3xl font-bold font-inter">1.200+</h3>
              <p className="text-gray-600 font-poppins">Pengguna aktif</p>
            </div>
            <div className="flex flex-col items-center text-center space-y-2">
              <div className="text-[#4f46e5] mb-2">🏢</div>
              <h3 className="text-3xl font-bold font-inter">100+</h3>
              <p className="text-gray-600 font-poppins">
                Perusahaan inklusif terdaftar
              </p>
            </div>
            <div className="flex flex-col items-center text-center space-y-2">
              <div className="text-[#10b981] mb-2">📚</div>
              <h3 className="text-3xl font-bold font-inter">250+</h3>
              <p className="text-gray-600 font-poppins">Kursus tersedia</p>
            </div>
            <div className="flex flex-col items-center text-center space-y-2">
              <div className="text-[#ec4899] mb-2">❤️</div>
              <h3 className="text-3xl font-bold font-inter">98%</h3>
              <p className="text-gray-600 font-poppins">
                pengguna merasa lebih percaya diri
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Testimonials Section */}
      <section className="py-16 md:py-24">
        <div className="container px-4 mx-auto">
          <h2 className="text-3xl md:text-4xl font-bold mb-12 font-inter">
            Testimoni Pengguna
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
            <div className="space-y-8">
              <div className="bg-white p-6 rounded-lg shadow-sm border">
                <div className="flex items-center gap-4 mb-4">
                  <div className="bg-[#ff8c42]/10 p-2 rounded-full">
                    <MessageSquare className="h-5 w-5 text-[#ff8c42]" />
                  </div>
                  <p className="text-gray-600 font-poppins">
                    Lewat platform ini, aku bisa kerja di perusahaan impian dan
                    tetap jadi diriku sendiri.
                  </p>
                </div>
                <div className="pl-12">
                  <p className="font-semibold font-inter">
                    — Andi, Junior Frontend Developer
                  </p>
                </div>
              </div>
              <div className="bg-white p-6 rounded-lg shadow-sm border">
                <div className="flex items-center gap-4 mb-4">
                  <div className="bg-[#4f46e5]/10 p-2 rounded-full">
                    <MessageSquare className="h-5 w-5 text-[#4f46e5]" />
                  </div>
                  <p className="text-gray-600 font-poppins">
                    Akhirnya ada tempat yang ngerti kebutuhan kami. Terima kasih
                    sudah hadir.
                  </p>
                </div>
                <div className="pl-12">
                  <p className="font-semibold font-inter">
                    — Sinta, penyandang tuli & desainer grafis
                  </p>
                </div>
              </div>
            </div>
            <div className="flex justify-center">
              <Image
                src="/home/testimoni.svg"
                alt="Testimoni Ilustrasi"
                width={400}
                height={400}
                className="w-full max-w-[400px]"
              />
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 bg-white">
        <div className="container px-4 mx-auto text-center">
          <h2 className="mb-4">
            Masih Belum Yakin?
          </h2>
          <p className="text-xl text-gray-600 mb-8 font-poppins">
            Coba jelajahi fitur kami lebih dulu!
          </p>
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <Link href="/kursus">
              <Button
                variant="outline"
                className="border-[#ff8c42] text-[#ff8c42] hover:bg-[#ff8c42]/10 hover:text-[#ff8c42]/90 px-8 py-6"
              >
                Cek Daftar Kursus
              </Button>
            </Link>
            <Link href="/lowongan">
              <Button className="bg-[#ff8c42] hover:bg-[#ff8c42]/90 text-white px-8 py-6">
                Lihat Lowongan Kerja
              </Button>
            </Link>
          </div>
        </div>
      </section>

      {/* FAQ Section */}
      <section className="py-16 md:py-24">
        <div className="container px-4 mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
            <div className="flex justify-center">
              <Image
                src="/home/faq.svg"
                alt="FAQ Ilustrasi"
                width={400}
                height={400}
                className="w-full max-w-[400px]"
              />
            </div>
            <div className="space-y-6">
              <h2>
                FAQ
              </h2>
              <Accordion type="single" collapsible className="w-full">
                <AccordionItem value="item-1">
                  <AccordionTrigger className="font-inter">
                    Apakah InKarya gratis?
                  </AccordionTrigger>
                  <AccordionContent className="font-poppins">
                    Ya, semua fitur dasar kami dapat digunakan secara gratis
                    oleh pengguna.
                  </AccordionContent>
                </AccordionItem>
                <AccordionItem value="item-2">
                  <AccordionTrigger className="font-inter">
                    Bagaimana proses verifikasi perusahaan?
                  </AccordionTrigger>
                  <AccordionContent className="font-poppins">
                    Kami memiliki sistem penilaian dan wawancara untuk
                    memastikan komitmen terhadap inklusi.
                  </AccordionContent>
                </AccordionItem>
                <AccordionItem value="item-3">
                  <AccordionTrigger className="font-inter">
                    Apakah saya perlu memiliki sertifikat disabilitas?
                  </AccordionTrigger>
                  <AccordionContent className="font-poppins">
                    Tidak wajib, tetapi dapat membantu dalam proses pencocokan
                    dengan perusahaan yang tepat.
                  </AccordionContent>
                </AccordionItem>
                <AccordionItem value="item-4">
                  <AccordionTrigger className="font-inter">
                    Bagaimana jika saya mengalami masalah aksesibilitas?
                  </AccordionTrigger>
                  <AccordionContent className="font-poppins">
                    Tim dukungan kami siap membantu 24/7. Silakan hubungi kami
                    melalui fitur chat atau email.
                  </AccordionContent>
                </AccordionItem>
              </Accordion>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}

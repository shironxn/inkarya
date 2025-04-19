"use client";

import type React from "react";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Bookmark,
  Briefcase,
  Clock,
  Code,
  Coffee,
  Database,
  DollarSign,
  GraduationCap,
  Headset,
  Info,
  MapPin,
  Package,
  Palette,
  Search,
  Users2,
} from "lucide-react";

// Types
interface Job {
  id: string;
  title: string;
  company: string;
  location: string;
  education: string;
  disabilities: string;
  salary: string;
  postedDays: number;
  icon: keyof typeof JobIcons;
}

// Constants
const SORT_OPTIONS = [
  { value: "relevan", label: "Relevan" },
  { value: "terbaru", label: "Terbaru" },
  { value: "gaji", label: "Gaji Tertinggi" },
];

const JOB_TYPE_OPTIONS = [
  { value: "fulltime", label: "Full Time" },
  { value: "parttime", label: "Part Time" },
  { value: "freelance", label: "Freelance" },
  { value: "contract", label: "Contract" },
];

const DISABILITY_OPTIONS = [
  { value: "fisik", label: "Disabilitas Fisik" },
  { value: "netra", label: "Disabilitas Netra" },
  { value: "rungu", label: "Disabilitas Rungu" },
  { value: "mental", label: "Disabilitas Mental" },
];

const EDUCATION_OPTIONS = [
  { value: "sma", label: "SMA/SMK" },
  { value: "d3", label: "D3" },
  { value: "s1", label: "S1" },
  { value: "s2", label: "S2" },
];

const SALARY_OPTIONS = [
  { value: "1-3", label: "Rp1.000.000 - Rp3.000.000" },
  { value: "3-5", label: "Rp3.000.000 - Rp5.000.000" },
  { value: "5-10", label: "Rp5.000.000 - Rp10.000.000" },
  { value: "10+", label: "Rp10.000.000+" },
];

const SPECIALIZATION_OPTIONS = [
  { value: "it", label: "IT & Software" },
  { value: "design", label: "Design" },
  { value: "marketing", label: "Marketing" },
  { value: "finance", label: "Finance" },
];

// Job Icons mapping
const JobIcons = {
  Headset,
  Code,
  Coffee,
  Database,
  Palette,
  Package,
};

// Mock Data
const JOBS: Job[] = [
  {
    id: "1",
    title: "Customer Support",
    company: "PT Digital Care",
    location: "Jakarta, Indonesia",
    education: "Minimal SMA/SMK",
    disabilities: "Tunarungu, Disabilitas Fisik",
    salary: "Rp4.000.000 - Rp6.000.000",
    postedDays: 2,
    icon: "Headset",
  },
  {
    id: "2",
    title: "Frontend Developer",
    company: "Techindo Solutions",
    location: "Jakarta, Indonesia",
    education: "Minimal D3/S1 Teknik Informatika",
    disabilities: "Tunanetra, Tunarungu",
    salary: "Rp7.000.000 - Rp12.000.000",
    postedDays: 1,
    icon: "Code",
  },
  {
    id: "3",
    title: "Barista",
    company: "Kopi Kita",
    location: "Bandung, Indonesia",
    education: "Minimal SMP/SMA",
    disabilities: "Disabilitas Fisik (Ringan)",
    salary: "Rp2.500.000 - Rp3.500.000",
    postedDays: 3,
    icon: "Coffee",
  },
  {
    id: "4",
    title: "Data Entry Specialist",
    company: "AdminPro",
    location: "Surabaya, Indonesia",
    education: "Minimal SMA/D3",
    disabilities: "Tunarungu, Disabilitas Fisik",
    salary: "Rp3.500.000 - Rp5.000.000",
    postedDays: 5,
    icon: "Database",
  },
  {
    id: "5",
    title: "Desainer Grafis",
    company: "CreativeHouse",
    location: "Yogyakarta, Indonesia",
    education: "Minimal D3/S1 Desain Komunikasi Visual",
    disabilities: "Tunanetra (Ringan), Disabilitas Fisik",
    salary: "Rp5.000.000 - Rp8.000.000",
    postedDays: 2,
    icon: "Palette",
  },
  {
    id: "6",
    title: "Operator Produksi",
    company: "Manufaktur Sejahtera",
    location: "Surabaya, Indonesia",
    education: "Minimal SMA/SMK",
    disabilities: "Tunarungu, Disabilitas Fisik (Ringan)",
    salary: "Rp3.000.000 - Rp4.500.000",
    postedDays: 4,
    icon: "Package",
  },
];

// Components
function FilterInput({
  placeholder,
  icon,
}: {
  placeholder: string;
  icon: React.ReactNode;
}) {
  return (
    <div className="relative">
      <div className="absolute left-3 top-1/2 transform -translate-y-1/2 text-[#71717a]">
        {icon}
      </div>
      <Input placeholder={placeholder} className="pl-10 border-[#e4e4e7]" />
    </div>
  );
}

function FilterSelect({
  placeholder,
  options,
  icon,
  defaultValue,
}: {
  placeholder: string;
  options: { value: string; label: string }[];
  icon: React.ReactNode;
  defaultValue?: string;
}) {
  return (
    <div className="relative w-full">
      <div className="absolute left-3 top-1/2 transform -translate-y-1/2 text-[#71717a]">
        {icon}
      </div>
      <Select defaultValue={defaultValue}>
        <SelectTrigger className="pl-10 border-[#e4e4e7] w-full">
          <SelectValue placeholder={placeholder} />
        </SelectTrigger>
        <SelectContent>
          {options.map((option) => (
            <SelectItem key={option.value} value={option.value}>
              {option.label}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
    </div>
  );
}

function JobCard({ job }: { job: Job }) {
  const IconComponent = JobIcons[job.icon];

  return (
    <div className="bg-white rounded-lg shadow-sm p-6 relative">
      <div className="flex justify-between mb-4">
        <div className="flex items-center gap-3">
          <div className="bg-[#f4f4f5] p-2 rounded-lg">
            <IconComponent className="h-5 w-5" />
          </div>
          <div>
            <h5 className="font-semibold">{job.title}</h5>
            <p className="text-sm text-[#71717a]">{job.company}</p>
          </div>
        </div>
        <Button variant="ghost" size="icon" className="h-8 w-8">
          <Bookmark className="h-5 w-5" />
          <span className="sr-only">Bookmark job</span>
        </Button>
      </div>
      <div className="space-y-3">
        <div className="flex items-start gap-2">
          <MapPin className="h-4 w-4 text-[#71717a] mt-0.5" />
          <span className="text-sm">{job.location}</span>
        </div>
        <div className="flex items-start gap-2">
          <GraduationCap className="h-4 w-4 text-[#71717a] mt-0.5" />
          <span className="text-sm">{job.education}</span>
        </div>
        <div className="flex items-start gap-2">
          <Users2 className="h-4 w-4 text-[#71717a] mt-0.5" />
          <span className="text-sm">{job.disabilities}</span>
        </div>
        <div className="flex items-start gap-2">
          <DollarSign className="h-4 w-4 text-[#71717a] mt-0.5" />
          <span className="text-sm">{job.salary}</span>
        </div>
      </div>
      <div className="mt-4 pt-4 border-t flex items-center gap-2 text-[#71717a]">
        <Clock className="h-4 w-4" />
        <span className="text-sm">{job.postedDays} hari yang lalu</span>
      </div>
    </div>
  );
}

// Main Page Component
export default function Home() {
  return (
    <div className="min-h-screen bg-[#fafafa]">
      {/* Main Content */}
      <main className="container px-4 py-8 mx-auto">
        {/* Search Section */}
        <div className="bg-white rounded-lg shadow-sm p-6 mb-8">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mb-4">
            <FilterInput
              placeholder="Keyword"
              icon={<Search className="h-4 w-4" />}
            />
            <FilterInput
              placeholder="Lokasi"
              icon={<MapPin className="h-4 w-4" />}
            />
            <FilterSelect
              placeholder="Spesialisasi"
              options={SPECIALIZATION_OPTIONS}
              icon={<Info className="h-4 w-4" />}
            />
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
            <FilterSelect
              placeholder="Tipe Pekerjaan"
              options={JOB_TYPE_OPTIONS}
              icon={<Briefcase className="h-4 w-4" />}
            />
            <FilterSelect
              placeholder="Jenis Disabilitas"
              options={DISABILITY_OPTIONS}
              icon={<Users2 className="h-4 w-4" />}
            />
            <FilterSelect
              placeholder="Pendidikan"
              options={EDUCATION_OPTIONS}
              icon={<GraduationCap className="h-4 w-4" />}
            />
            <FilterSelect
              placeholder="Gaji"
              options={SALARY_OPTIONS}
              icon={<DollarSign className="h-4 w-4" />}
            />
          </div>

          <div className="flex justify-end">
            <Button className="bg-[#ff8c42] hover:bg-[#ff8c42]/90 text-white px-8">
              Cari
            </Button>
          </div>
        </div>

        {/* Results Section */}
        <div className="mb-6 flex justify-between items-center">
          <p className="text-[#71717a]">
            Menampilkan{" "}
            <span className="font-medium text-black">{JOBS.length}</span>{" "}
            lowongan pekerjaan
          </p>
          <div className="flex items-center gap-2">
            <span className="text-sm text-[#71717a]">Urutkan berdasarkan</span>
            <Select defaultValue="relevan">
              <SelectTrigger className="w-[140px] border-[#e4e4e7]">
                <SelectValue placeholder="Relevan" />
              </SelectTrigger>
              <SelectContent>
                {SORT_OPTIONS.map((option) => (
                  <SelectItem key={option.value} value={option.value}>
                    {option.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>

        {/* Job Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
          {JOBS.map((job) => (
            <JobCard key={job.id} job={job} />
          ))}
        </div>
      </main>
    </div>
  );
}

import { useStackApp } from "@stackframe/stack";
import { useRouter } from "next/navigation";
import { useEffect, useState, useMemo } from "react";
import { toast } from "sonner";
import { FormData, FormFields } from "@/types/onboarding";
import { DUMMY_SKILLS, DUMMY_DISABILITIES, TOTAL_STEPS } from "@/constants/onboarding";

export function useOnboardingForm() {
  const app = useStackApp();
  const user = app.useUser();
  const router = useRouter();

  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState<FormData & { [key: string]: string }>({
    nama_lengkap: user?.displayName || "",
    email: "",
    phone: "",
    bio: "",
    interest: "",
    dob: "",
    location: "",
    skills: "",
    disabilities: "",
  });
  const [isStepValid, setIsStepValid] = useState(false);
  const [date, setDate] = useState<Date | undefined>(undefined);
  const [selectedSkills, setSelectedSkills] = useState<number[]>([]);
  const [selectedDisabilities, setSelectedDisabilities] = useState<number[]>([]);

  // Form fields based on the UserCreateRequest struct
  const formFields = useMemo<FormFields>(() => ({
    // Step 1: Basic Information
    1: [
      {
        id: "nama_lengkap",
        label: "Nama Lengkap",
        type: "text",
        placeholder: "Masukkan nama lengkap Anda",
        required: true,
        hint: "Nama ini akan ditampilkan di profil Anda",
      },
      {
        id: "email",
        label: "Email",
        type: "email",
        placeholder: "email@example.com",
        required: false,
        hint: "Email ini akan ditampilkan di profil Anda",
      },
      {
        id: "phone",
        label: "Nomor Telepon",
        type: "tel",
        placeholder: "Masukkan nomor telepon",
        required: false,
        hint: "Nomor telepon ini akan ditampilkan di profil Anda",
      },
    ],
    // Step 2: Profile Details
    2: [
      {
        id: "bio",
        label: "Bio",
        type: "textarea",
        placeholder: "Ceritakan sedikit tentang diri Anda",
        required: false,
      },
      {
        id: "interest",
        label: "Minat",
        type: "text",
        placeholder: "Minat dan keahlian Anda",
        required: true,
      },
      {
        id: "dob",
        label: "Tanggal Lahir",
        type: "datepicker",
        placeholder: "Pilih tanggal lahir",
        required: true,
      },
      {
        id: "location",
        label: "Lokasi",
        type: "text",
        placeholder: "Kota, Provinsi",
        required: true,
      },
      {
        id: "skills",
        label: "Keahlian",
        type: "multiselect",
        placeholder: "Pilih keahlian Anda",
        required: true,
        options: DUMMY_SKILLS,
      },
      {
        id: "disabilities",
        label: "Kondisi",
        type: "multiselect",
        placeholder: "Pilih kondisi Anda",
        required: true,
        options: DUMMY_DISABILITIES,
      },
    ],
    // Step 3: Professional Information & Avatar
    3: [
      {
        id: "status",
        label: "Status",
        type: "text",
        placeholder: "Contoh: Mahasiswa, Pekerja, Freelancer",
        required: false,
      },
      {
        id: "availability",
        label: "Ketersediaan",
        type: "text",
        placeholder: "Contoh: Full-time, Part-time, Freelance",
        required: false,
      },
      {
        id: "resumeURL",
        label: "Resume",
        type: "file",
        accept: ".pdf,.doc,.docx",
        required: false,
        hint: "Format yang didukung: PDF, DOC, DOCX",
      },
      {
        id: "avatarURL",
        label: "Foto Profil",
        type: "avatar",
        required: false,
      },
    ],
  }), []);

  // Validate current step
  useEffect(() => {
    const currentFields = formFields[step];
    const requiredFields = currentFields.filter((field) => field.required);

    // Check if all required fields have values
    const allRequiredFilled = requiredFields.every((field) => {
      if (field.type === "file" || field.type === "avatar") {
        return true;
      }
      if (field.type === "datepicker") {
        return date !== undefined;
      }
      return formData[field.id]?.trim() !== "";
    });

    setIsStepValid(allRequiredFilled);
  }, [step, formData, date, formFields]);

  const handleInputChange = (id: string, value: string) => {
    setFormData((prev) => ({ ...prev, [id]: value }));
  };

  const handleNext = () => {
    if (step < TOTAL_STEPS && isStepValid) {
      setStep(step + 1);
      window.scrollTo(0, 0);
    }
  };

  const handlePrevious = () => {
    if (step > 1) {
      setStep(step - 1);
      window.scrollTo(0, 0);
    }
  };

  const handleSubmit = async () => {
    try {
      const user = await app.getUser();
      const response = await fetch("/api/onboarding", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${await user?.getAuthJson().then((auth) => auth?.accessToken)}`,
        },
        body: JSON.stringify({
          ...formData,
          skills: selectedSkills,
          disabilities: selectedDisabilities,
          dob: date?.toISOString(),
        }),
      });

      if (!response.ok) {
        const error = await response.json();
        throw new Error(error.message || "Gagal menyimpan data");
      }

      toast.success("Profil berhasil disimpan", {
        description: "Selamat datang di InKarya!",
      });

      router.push("/dashboard");
    } catch (error) {
      toast.error("Gagal menyimpan profil", {
        description:
          error instanceof Error
            ? error.message
            : "Terjadi kesalahan saat menyimpan data",
      });
    }
  };

  return {
    step,
    formData,
    isStepValid,
    date,
    setDate,
    selectedSkills,
    setSelectedSkills,
    selectedDisabilities,
    setSelectedDisabilities,
    formFields,
    handleInputChange,
    handleNext,
    handlePrevious,
    handleSubmit,
  };
} 
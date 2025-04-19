export interface FormField {
  id: string;
  label: string;
  type: string;
  placeholder?: string;
  icon?: React.ReactNode;
  required: boolean;
  disabled?: boolean;
  value?: string;
  accept?: string;
  hint?: string;
  options?: { id: number; name: string }[];
}

export interface FormFields {
  [key: number]: FormField[];
}

export interface FormData {
  nama_lengkap: string;
  email: string;
  phone: string;
  bio: string;
  interest: string;
  dob: string;
  location: string;
  skills: string;
  disabilities: string;
  status?: string;
  availability?: string;
  resumeURL?: string;
  avatarURL?: string;
} 
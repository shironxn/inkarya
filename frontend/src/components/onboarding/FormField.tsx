import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "@/components/ui/command";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/lib/utils";
import { format } from "date-fns";
import {
  Calendar as CalendarIcon,
  Check,
  Upload,
  User,
} from "lucide-react";
import { FormField as FormFieldType } from "@/types/onboarding";

interface FormFieldProps {
  field: FormFieldType;
  value: string;
  onChange: (value: string) => void;
  date?: Date;
  onDateChange?: (date: Date | undefined) => void;
  selectedSkills?: number[];
  onSkillsChange?: (skills: number[]) => void;
  selectedDisabilities?: number[];
  onDisabilitiesChange?: (disabilities: number[]) => void;
}

const FileInput = ({ id, accept, required, onChange }: { 
  id: string; 
  accept?: string; 
  required?: boolean; 
  onChange: (value: string) => void;
}) => (
  <div className="relative">
    <div className="flex items-center">
      <label
        htmlFor={id}
        className="cursor-pointer flex items-center gap-2 border rounded-md px-4 py-2 bg-gray-50 hover:bg-gray-100 transition-colors"
      >
        <Upload className="h-4 w-4" />
        <span>Pilih File</span>
      </label>
      <input
        type="file"
        id={id}
        accept={accept}
        className="hidden"
        required={required}
        onChange={(e) => {
          const file = e.target.files?.[0];
          if (file) onChange(file.name);
        }}
      />
    </div>
  </div>
);

const AvatarUpload = ({ onChange }: { onChange: (value: string) => void }) => (
  <div className="flex flex-col items-center">
    <div className="relative mb-4">
      <div className="w-24 h-24 rounded-full bg-gray-200 flex items-center justify-center overflow-hidden">
        <User className="h-12 w-12 text-gray-400" />
      </div>
      <label
        htmlFor="avatar-upload"
        className="absolute bottom-0 right-0 rounded-full h-8 w-8 flex items-center justify-center border border-gray-300 bg-white cursor-pointer hover:bg-gray-50"
      >
        <Upload className="h-4 w-4" />
      </label>
      <input
        type="file"
        id="avatar-upload"
        accept="image/*"
        className="hidden"
        onChange={(e) => {
          const file = e.target.files?.[0];
          if (file) onChange(file.name);
        }}
      />
    </div>
    <p className="text-xs text-muted-foreground">
      Klik untuk mengunggah foto profil
    </p>
  </div>
);

const DatePicker = ({ 
  date, 
  onDateChange, 
  onChange 
}: { 
  date?: Date; 
  onDateChange?: (date: Date | undefined) => void; 
  onChange: (value: string) => void;
}) => (
  <Popover>
    <PopoverTrigger asChild>
      <Button
        variant="outline"
        className={cn(
          "w-full justify-start text-left font-normal py-6 pl-10",
          !date && "text-muted-foreground",
        )}
      >
        <CalendarIcon className="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
        {date ? format(date, "PPP") : <span>Pilih tanggal lahir</span>}
      </Button>
    </PopoverTrigger>
    <PopoverContent className="w-auto p-0" align="start">
      <Calendar
        mode="single"
        selected={date}
        onSelect={(newDate) => {
          onDateChange?.(newDate);
          if (newDate) onChange(format(newDate, "yyyy-MM-dd"));
        }}
        initialFocus
      />
    </PopoverContent>
  </Popover>
);

const MultiSelect = ({
  field,
  selectedItems,
  onItemsChange,
  onChange,
}: {
  field: FormFieldType;
  selectedItems: number[];
  onItemsChange: (items: number[]) => void;
  onChange: (value: string) => void;
}) => (
  <Popover>
    <PopoverTrigger asChild>
      <Button
        variant="outline"
        role="combobox"
        className="w-full justify-between py-6"
      >
        {selectedItems.length > 0
          ? `${selectedItems.length} ${field.id === "skills" ? "keahlian" : "disabilitas"} dipilih`
          : field.placeholder}
      </Button>
    </PopoverTrigger>
    <PopoverContent className="w-full p-0" align="start">
      <Command>
        <CommandInput
          placeholder={`Cari ${field.id === "skills" ? "keahlian" : "disabilitas"}...`}
        />
        <CommandEmpty>Tidak ditemukan.</CommandEmpty>
        <CommandGroup>
          {field.options?.map((option) => (
            <CommandItem
              key={option.id}
              value={option.name}
              onSelect={() => {
                const newItems = selectedItems.includes(option.id)
                  ? selectedItems.filter((id) => id !== option.id)
                  : [...selectedItems, option.id];
                onItemsChange(newItems);
                onChange(newItems.join(","));
              }}
            >
              <Check
                className={cn(
                  "mr-2 h-4 w-4",
                  selectedItems.includes(option.id) ? "opacity-100" : "opacity-0"
                )}
              />
              {option.name}
            </CommandItem>
          ))}
        </CommandGroup>
      </Command>
    </PopoverContent>
  </Popover>
);

export function FormField({
  field,
  value,
  onChange,
  date,
  onDateChange,
  selectedSkills = [],
  onSkillsChange,
  selectedDisabilities = [],
  onDisabilitiesChange,
}: FormFieldProps) {
  switch (field.type) {
    case "textarea":
      return (
        <Textarea
          id={field.id}
          placeholder={field.placeholder}
          className="min-h-[100px]"
          required={field.required}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          disabled={field.disabled}
        />
      );
    case "datepicker":
      return <DatePicker date={date} onDateChange={onDateChange} onChange={onChange} />;
    case "file":
      return <FileInput id={field.id} accept={field.accept} required={field.required} onChange={onChange} />;
    case "avatar":
      return <AvatarUpload onChange={onChange} />;
    case "multiselect":
      return (
        <MultiSelect
          field={field}
          selectedItems={field.id === "skills" ? selectedSkills : selectedDisabilities}
          onItemsChange={field.id === "skills" ? onSkillsChange! : onDisabilitiesChange!}
          onChange={onChange}
        />
      );
    default:
      return (
        <div className="relative">
          {field.icon && field.icon}
          <Input
            type={field.type}
            id={field.id}
            placeholder={field.placeholder}
            className={field.icon ? "pl-10 py-6" : "py-6"}
            required={field.required}
            value={value}
            onChange={(e) => onChange(e.target.value)}
            disabled={field.disabled}
          />
        </div>
      );
  }
} 
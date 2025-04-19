import { Progress } from "@/components/ui/progress";
import { TOTAL_STEPS } from "@/constants/onboarding";

interface FormHeaderProps {
  step: number;
}

export function FormHeader({ step }: FormHeaderProps) {
  return (
    <header className="border-b bg-white sticky top-0 z-10">
      <div className="container mx-auto pb-4 py-4">
        <div className="flex items-center justify-between mb-2">
          <span className="text-sm font-medium">
            Langkah {step} dari {TOTAL_STEPS}
          </span>
          <span className="text-sm text-muted-foreground">
            {Math.round((step / TOTAL_STEPS) * 100)}% selesai
          </span>
        </div>
        <Progress value={(step / TOTAL_STEPS) * 100} className="h-2" />
      </div>
    </header>
  );
} 
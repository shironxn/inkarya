import { Button } from "@/components/ui/button";
import { TOTAL_STEPS } from "@/constants/onboarding";

interface FormNavigationProps {
  step: number;
  isStepValid: boolean;
  onPrevious: () => void;
  onNext: () => void;
  onSubmit: () => void;
}

export function FormNavigation({
  step,
  isStepValid,
  onPrevious,
  onNext,
  onSubmit,
}: FormNavigationProps) {
  return (
    <div className="flex justify-between pt-4">
      {step > 1 ? (
        <Button
          type="button"
          variant="outline"
          onClick={onPrevious}
          className="border-gray-300"
        >
          Kembali
        </Button>
      ) : (
        <div></div>
      )}

      <Button
        type="button"
        onClick={step === TOTAL_STEPS ? onSubmit : onNext}
        className="bg-primary text-white hover:bg-primary/90"
        disabled={!isStepValid}
      >
        {step === TOTAL_STEPS ? "Selesai" : "Lanjutkan"}
      </Button>
    </div>
  );
} 
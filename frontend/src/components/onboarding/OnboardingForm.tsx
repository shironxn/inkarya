"use client";

import { Label } from "@/components/ui/label";
import { FormField } from "./FormField";
import { FormHeader } from "./FormHeader";
import { FormNavigation } from "./FormNavigation";
import { useOnboardingForm } from "@/hooks/useOnboardingForm";
import { STEP_DESCRIPTIONS, STEP_TITLES } from "@/constants/onboarding";
import { FormField as FormFieldType } from "@/types/onboarding";

export function OnboardingForm() {
  const {
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
  } = useOnboardingForm();

  return (
    <>
      <FormHeader step={step} />

      <div className="container mx-auto py-8 px-4 md:py-12">
        <div className="mx-auto max-w-2xl bg-white rounded-lg shadow-sm p-6 md:p-8">
          <div className="mb-8">
            <h1 className="text-2xl md:text-3xl font-bold mb-2">
              {STEP_TITLES[step as keyof typeof STEP_TITLES]}
            </h1>
            <p className="text-sm text-muted-foreground">
              {STEP_DESCRIPTIONS[step as keyof typeof STEP_DESCRIPTIONS]}
            </p>
          </div>

          <div className="space-y-6">
            <div className="space-y-4">
              {formFields[step].map((field: FormFieldType, index: number) => (
                <div
                  key={index}
                  className={field.type === "avatar" ? "flex justify-center" : ""}
                >
                  {field.type !== "avatar" && (
                    <Label htmlFor={field.id} className="block mb-2">
                      {field.label}
                      {field.required && (
                        <span className="text-red-500 ml-1">*</span>
                      )}
                    </Label>
                  )}
                  <FormField
                    field={field}
                    value={formData[field.id]}
                    onChange={(value) => handleInputChange(field.id, value)}
                    date={date}
                    onDateChange={setDate}
                    selectedSkills={selectedSkills}
                    onSkillsChange={setSelectedSkills}
                    selectedDisabilities={selectedDisabilities}
                    onDisabilitiesChange={setSelectedDisabilities}
                  />
                  {field.hint && (
                    <p className="mt-1 text-xs text-muted-foreground">
                      {field.hint}
                    </p>
                  )}
                </div>
              ))}
            </div>

            <FormNavigation
              step={step}
              isStepValid={isStepValid}
              onPrevious={handlePrevious}
              onNext={handleNext}
              onSubmit={handleSubmit}
            />
          </div>
        </div>
      </div>
    </>
  );
}


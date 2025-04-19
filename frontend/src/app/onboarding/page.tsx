import { OnboardingForm } from "@/components/onboarding/OnboardingForm";
import { stackServerApp } from "@/stack";
import { Metadata } from "next";
import { redirect } from "next/navigation";

export const metadata: Metadata = {
  title: "Onboarding - InKarya",
  description: "Complete your profile to get started with InKarya",
};

export default async function OnboardingPage() {
  const user = await stackServerApp.getUser();

  if (!user || user?.clientMetadata.onboarded) {
    redirect("/");
  }

  return (
    <main className="min-h-screen bg-gray-50">
      <OnboardingForm />
    </main>
  );
}

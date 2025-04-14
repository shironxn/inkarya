export default async function Page({
  params,
}: {
  params: Promise<{ id: string; slug: string }>;
}) {
  const { id, slug } = await params;

  return (
    <div>
      <h1>Kursus</h1>
      <p>Slug: {slug}</p>
      <p>ID: {id}</p>
    </div>
  );
}

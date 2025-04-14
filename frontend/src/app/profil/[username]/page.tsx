export default async function Page({
  params,
}: {
  params: Promise<{ username: string }>;
}) {
  const { username } = await params;

  return (
    <div>
      <h1>Profil</h1>
      <p>Name: {username}</p>
    </div>
  );
}

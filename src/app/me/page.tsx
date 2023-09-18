import { UserProfile } from "@clerk/nextjs";

export default async function MePage() {
  return (
    <div className="flex justify-center">
      <UserProfile />
    </div>
  );
}

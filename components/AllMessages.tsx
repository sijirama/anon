import { db } from "@/lib/db"

export async function AllMessages() {
  const rooms = await db.room.findMany();
  console.log(rooms)

  return (
    <div className="space-y-4 flex flex-col">
      {rooms.map((room) => (
        <p className="text-white" key={room.id}>{room.title}</p>
      ))}
    </div>
  );
}


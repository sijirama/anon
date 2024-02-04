"use client"
//import { db } from "@/lib/db"
import { Button } from "./ui/button";

import { useModal } from "@/hooks/useModalStor";

export function AllMessages() {
    // const rooms = await db.room.findMany();
    // console.log(rooms)

    const { onOpen } = useModal()

    return (
        <div className="space-y-4 flex flex-col">
            {/*rooms.map((room) => (
        <p className="text-white" key={room.id}>{room.title}</p>
      ))*/}

            <Button
                onClick={() => onOpen("addRoom")}
            >Create a Room</Button>
        </div>
    );
}


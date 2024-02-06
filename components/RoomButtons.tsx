"use client"
//import { db } from "@/lib/db"
import { Button } from "./ui/button";
import { useModal } from "@/hooks/useModalStor";

interface RoomButtonProps {
    id: string
}

export function RoomButtons({ id }: RoomButtonProps) {
    // const rooms = await db.room.findMany();
    // console.log(rooms)

    const { onOpen } = useModal()

    return (
        <div className="gap-4 flex ">
            <Button
                onClick={() => onOpen("addMessage", { roomId: id })}
            >Add a message</Button>

            <Button
                onClick={() => onOpen("deleteRoom", { roomId: id })}
            >Delete the Room</Button>

        </div>
    );
}


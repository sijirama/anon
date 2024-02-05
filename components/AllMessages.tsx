"use client"
import { Button } from "./ui/button";
import { useEffect, useState } from "react";
import { useModal } from "@/hooks/useModalStor";
import { Room } from "@prisma/client";
import axios from "axios";

export function AllMessages() {

    //INFO: For just debugging purposes

    const [rooms, setRooms] = useState<Room[]>([])

    useEffect(() => {
        const fetchRooms = async () => {
            try {
                const response = await axios.get('/api/rooms');
                return response.data;
            } catch (error) {
                console.error('Error fetching rooms:', error);
                return [];
            }
        };

        const getRooms = async () => {
            const fetchedRooms = await fetchRooms();
            setRooms(fetchedRooms);
        };

        getRooms();
    }, []);


    //NOTE: ------------------------------------

    const { onOpen } = useModal()



    return (
        <div className="space-y-4 flex flex-col">

            <Button
                onClick={() => onOpen("addRoom")}
            >Create a Room</Button>


            <div className="flex flex-col gap-4">
                {rooms.length > 0 ? (

                    rooms.map((room) => (
                        <a href={`/rooms/${room.id}`}>
                            <p className="text-black py-3 px-7 bg-zinc-200 " key={room.id}>{room.title}</p>
                        </a>
                    ))

                ) : (null)
                }
            </div>

        </div>
    );
}


"use client"
import { useEffect, useState } from "react";
import { Room } from "@prisma/client";
import axios from "axios";

export function Allmessages() {

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




    return (
        <div className="space-y-4 flex flex-col">


            <div className="flex flex-col gap-4">
                {rooms.length > 0 ? (

                    rooms.map((room) => (
                        <a href={`/rooms/${room.id}`} key={room.id}>
                            <p className="text-black py-3 px-7 bg-zinc-200 " key={room.id}>{room.title}</p>
                        </a>
                    ))

                ) : (null)
                }
            </div>

        </div>
    );
}


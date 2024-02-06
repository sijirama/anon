"use client"
import { Button } from "./ui/button";
import { useModal } from "@/hooks/useModalStor";

export function CreateRoomButton() {

    const { onOpen } = useModal()

    return (
        <div className="space-y-4 flex flex-col mt-3 ">
            <Button
            className="text-lg -tracking-wide font-semibold py-5 px-10 hover:bg-zinc-900  text-white"
                onClick={() => onOpen("addRoom")}
            >Create a Room!</Button>

        </div>
    );
}


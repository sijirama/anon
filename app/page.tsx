'use client'
import { AllMessages } from "@/components/AllMessages";
import { Button } from "@/components/ui/button";
import { useModal } from "@/hooks/useModalStor";

export default function Home() {
    const { onOpen } = useModal()
    const onClick = () => {
        console.log("clicked")
        onOpen("addRoom")
    }
    return (
        <main className="flex gap-4 min-h-screen flex-col items-center p-24 text-white">
            <h1 className="text-3xl -tracking-wider">Anon Rooms</h1>
            <Button
                onClick={() => onOpen("addRoom")}
            >Create a Room</Button>
            <Button onClick={onClick}
            >Create a Room</Button>


            <button onClick={onClick}
            >Create a Room</button>

            <AllMessages />
        </main>
    );
}

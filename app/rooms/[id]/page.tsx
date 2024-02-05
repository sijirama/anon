import { RoomButtons } from "@/components/RoomButtons";
import { db } from "@/lib/db";

export default async function page(props: any) {
    const id = props.params.id;
    const roomWithMessages = await db.room.findFirst({
        where: {
            id
        },
        include: {
            messages: {
                orderBy: {
                    createdAt: "asc"
                }
            }
        }
    })
    console.log(roomWithMessages)
    if (!roomWithMessages) {
        return null
    }
    const { title, messages } = roomWithMessages;
    return (
        <div className="min-h-screen w-full py-16 bg-black text-zinc-200">
            Room ID: {id} is wtf
            <div className="min-h-screen w-full py-16 bg-black text-zinc-200">
                <h1>Title is : {title}</h1>
                <RoomButtons id={id} />
                <ul>
                    {messages.map((message: any) => (
                        <li key={message.id}>{message.content}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
}


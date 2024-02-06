import { db } from "@/lib/db";
import { NextResponse } from "next/server";


export async function POST(req: Request) {
    try {
        const { content, roomId } = await req.json();

        //console.log(content, roomId)

        const room = await db.room.findUnique({
            where: {
                id: roomId,
            },
        });

        if (!room) {
            return new NextResponse('Room not found', { status: 404 });
        }


        //INFO: Create the message associated with the room
        const message = await db.message.create({
            data: {
                content,
                room: {
                    connect: {
                        id: roomId,
                    },
                },
            },
        });

        console.log('Message created:', message);


        return NextResponse.json("OK");
    } catch (e) {
        console.error('Error processing request:', e);
        return new NextResponse('Internal error', { status: 500 });
    }
}

// export async function GET(req: Request) {
//     try {
//         const room = await db.room.findMany()
//         return NextResponse.json(room);
//     } catch (e) {
//         console.error('Error processing request:', e);
//         return new NextResponse('Internal error', { status: 500 });
//     }
// }
//
//

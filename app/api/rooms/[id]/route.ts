import { db } from '@/lib/db';
import { NextResponse } from 'next/server';
import bcrypt from "bcrypt"



export async function DELETE(req: Request,
    { params }: { params: { id: string } }
) {
    try {
        const { searchParams } = new URL(req.url);
        const id = params.id

        const password = searchParams.get('password');

        //console.log(id, password)
        if (!password || !id) {
            return new NextResponse('Invalid Request', { status: 401 });
        }

        const room = await db.room.findFirst({
            where: {
                id
            }
        })

        const passwordsMatch = await bcrypt.compare(password, room?.password!)

        if (!passwordsMatch) {
            return new NextResponse('Invalid Request', { status: 401 });
        }

        await db.room.delete({
            where: {
                id
            }
        })

        return NextResponse.json("room has been deleted");
    } catch (e) {
        console.error('Error processing request:', e);
        return new NextResponse('Internal error', { status: 500 });
    }
}





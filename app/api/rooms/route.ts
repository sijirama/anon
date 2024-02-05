import { db } from "@/lib/db";
import { NextResponse } from "next/server";
import bcrypt from "bcrypt"
import { CONFIG } from "@/lib/config";

export async function POST(req: Request) {
    try {
        const { title, password } = await req.json();
        const hashed = await bcrypt.hash(password, CONFIG.SALT)

        const room = await db.room.create({
            data: {
                title,
                password:hashed
            }
        })

        return NextResponse.json(room.id);
    } catch (e) {
        console.error('Error processing request:', e);
        return new NextResponse('Internal error', { status: 500 });
    }
}

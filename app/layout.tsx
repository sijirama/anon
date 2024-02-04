import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ModalProvider } from "@/providers/modal-provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Anon!",
    description: "say stuff anonymously",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en" className="bg-black">
            <body className={`${inter.className}w-11/12 md:w-5/6 mx-auto bg-zinc-950`}>
                <ModalProvider />
                {children}
            </body>
        </html>
    );
}

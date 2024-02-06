import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ModalProvider } from "@/providers/modal-provider";

import { Toaster } from "@/components/ui/sonner"
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
    //<body className={` md:max-w-[90%] lg:max-w-[90%] mx-auto h-full ${inter.className}`}>
    return (
        <html lang="en" className="bg-emerald-100">
            <body className={`${inter.className}`}>
                <ModalProvider />
                {children}
                <Toaster />
            </body>
        </html>
    );
}

import { AllMessages } from "@/components/AllMessages";

export default function Home() {
    return (
        <main className="flex gap-4 min-h-screen flex-col items-center p-24 text-white">
            <h1 className="text-3xl -tracking-wider">Anon Rooms</h1>
            <AllMessages />
        </main>
    );
}

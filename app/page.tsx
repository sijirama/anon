import { CreateRoomButton } from "@/components/CreateRoomButton";
import background from "../public/bg.jpg";
import { LandingPageBoxes } from "@/lib/data";

export default function Home() {
    return (
        <main className="flex h-full min-h-screen flex-col items-center  text-white ">
            <section className="flex h-full flex-col min-h-[50rem] gap-3 w-full p-3 text-center items-center justify-center bg-cover bg-no-repeat" style={{
                backgroundImage: `url(${background.src})`,
            }}>

                <p className="text-5xl md:text-6xl lg:text-7xl lg:px-44 -tracking-widest font-extrabold">Send Secret Anonymous Messages Online and fuck up both yours and others Self-esteem.</p>


                <CreateRoomButton />
            </section>
            <section className="bg-zinc-50 w-full min-h-14 py-10 px-5 text-zinc-900 flex gap-3 lg:gap-6 flex-wrap justify-center items-center">

                {
                    LandingPageBoxes.map((box) => (
                        <div className="px-5 py-4 flex flex-col gap-5 rounded-lg shadow-md w-[30rem] text-white bg-[#00223b] hover:scale-105 transition-transform duration-300">
                            <p className="font-semibold -tracking-wide">{box.topic}</p>
                            <p className="font-extralight -tracking-wide">{box.description}</p>

                        </div>
                    ))
                }

            </section>
        </main>
    );
}

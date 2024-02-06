"use client"
import { useModal } from "@/hooks/useModalStor";
import qs from "query-string"

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog';


import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from '@/components/ui/form';

import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import axios from "axios";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

const formSchema = z.object({
    password: z.string().min(3).max(50),
})


export default function DeleteRoomModal() {
    const { isOpen, type, onClose, data } = useModal()

    const isModalOpen = isOpen && type === "deleteRoom"

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
        },
    })

    let loading = form.formState.isSubmitting; // while the form is submitting
    const router = useRouter()

    // 2. Define a submit handler.
    async function onSubmit(values: z.infer<typeof formSchema>) {
        //console.log(values)
        try {
            // Make the POST request to create a room
            //await axios.delete(`/api/rooms/${data.roomId}?password=${values.password}`);
            const url = qs.stringify({
                url: `api/rooms/${data.roomId}`,
                query: {
                    password: values?.password
                }
            })

            axios.delete(url)

            // Redirect to the new room URL
            router.push("/");

            //Optionally, you can show a success toast or message
            handleClose()
            toast(
                'Room deleted successfully!',
            );
        } catch (error) {
            console.error('Error deleting room:', error);

            // Optionally, show an error toast or message
            toast(
                'Error deleting room. Please try again.',
            );
        }
    }

    const handleClose = () => {
        form.reset()
        onClose();
    };

    return (
        <Dialog open={isModalOpen} onOpenChange={handleClose}>
            <DialogContent className="bg-white text-slate-900 dark:bg-slate-900 dark:text-slate-50 overflow-hidden">
                <DialogHeader>
                    <DialogTitle>Delete this room, Anon!</DialogTitle>
                    <DialogDescription>
                        Are you sure you want to delete this room {data.roomId}
                    </DialogDescription>
                </DialogHeader>
                {/* Form stuff */}
                <Form {...form} >
                    <form
                        onSubmit={form.handleSubmit(onSubmit)}
                        className="space-y-5 w-full "
                    >
                        <FormField
                            control={form.control}
                            name="password"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel className="uppercase text-xs font-bold text-slate-700 dark:text-slate-100">
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            type="password"
                                            disabled={loading}
                                            placeholder="Enter room password"
                                            className="bg-zinc-300/10 border-0 focus-visible:ring-0 text-black dark:text-slate-200 font-semibold focus-visible:ring-offset-0"
                                            {...field}
                                        />
                                    </FormControl>
                                    <FormDescription>only the person that created the room can delete it.</FormDescription>
                                    <FormMessage className="font-semibold text-red-500" />
                                </FormItem>
                            )}
                        ></FormField>
                        <DialogFooter>
                            <Button type="submit" className="font-bold" variant="default">
                                {loading ? (null) : (

                                    "Delete room"
                                )}
                            </Button>
                        </DialogFooter>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )
}

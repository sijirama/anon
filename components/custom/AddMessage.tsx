"use client"
import { useModal } from "@/hooks/useModalStor";

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
import { Button } from "../ui/button";
import axios from "axios";
import { toast } from "sonner";
import { useRouter } from "next/navigation";
import { Textarea } from "../ui/textarea";

const formSchema = z.object({
    content: z.string().min(3).max(50),
})


export default function AddMessageModal() {
    const { isOpen, type, onClose, data } = useModal()

    const isModalOpen = isOpen && type === "addMessage"

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

            const stuffToSubmit = {
                content: values?.content,
                roomId: data?.roomId
            }

            await axios.post('/api/messages', stuffToSubmit);

            // Redirect to the new room URL
            handleClose()

            //Optionally, you can show a success toast or message
            router.refresh();

            toast(
                'Message added successfully!',
            );
        } catch (error) {
            console.error('Error submitting message:', error);

            // Optionally, show an error toast or message
            toast(
                'Error submitting message. Please try again.',
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
                    <DialogTitle>Add a message, Anon!</DialogTitle>
                    <DialogDescription>
                        You and only you will know what you sent, we dont even know, that is anonymosity{data.roomId}
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
                            name="content"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel className="uppercase text-xs font-bold text-slate-700 dark:text-slate-100">
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <Textarea
                                            disabled={loading}
                                            placeholder="Say something"
                                            className="bg-zinc-300/10 border-0 focus-visible:ring-0 text-black dark:text-slate-200 font-semibold focus-visible:ring-offset-0"
                                            {...field}

                                        />
                                    </FormControl>
                                    <FormDescription>only you knows the message sent.</FormDescription>
                                    <FormMessage className="font-semibold text-red-500" />
                                </FormItem>
                            )}
                        ></FormField>
                        <DialogFooter>
                            <Button type="submit" className="font-bold" variant="default">
                                {loading ? (null) : (

                                    "Submit message"
                                )}
                            </Button>
                        </DialogFooter>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )
}

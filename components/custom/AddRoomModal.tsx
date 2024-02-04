"use client"
import { useModal } from "@/hooks/useModalStor";

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog';


export default function AddRoomModal() {
    const { isOpen, type, onClose } = useModal()
    const isModalOpen = isOpen && type === "addRoom"

    const handleClose = () => {
        onClose();
    };

    return (
        <Dialog open={isModalOpen} onOpenChange={handleClose}>
            <DialogContent className="bg-white text-slate-900 dark:bg-slate-900 dark:text-slate-50 overflow-hidden">
                <DialogHeader>
                    <DialogTitle>Customize your server Anon!</DialogTitle>
                    <DialogDescription>
                        Give your server a personality with a name and an image,
                        you can always change it later.
                    </DialogDescription>
                </DialogHeader>
            </DialogContent>
        </Dialog>
    )

}

'use client'

import { create } from 'zustand';

export type ModalType =
    | 'addRoom'

interface ModalData {
}

interface ModalStore {
    type: ModalType | null;
    data: ModalData;
    isOpen: boolean;
    onOpen: (type: ModalType, data?: ModalData) => void;
    onClose: () => void;
}

export const useModal = create<ModalStore>(set => ({
    type: null,
    data: {},
    isOpen: false,
    onOpen(type, data = {}) {
        set({ isOpen: true, type, data });
        console.log("now opened")
    },
    onClose() {
        set({ isOpen: false, type: null });
    },
    
}));

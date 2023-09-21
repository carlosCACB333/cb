'use server'
import { ChatpdfMsg, Response } from "@/interfaces";
import { fetchBack } from "@/utils";

export const savePdfMessages = async (chatId: string, messages: Partial<ChatpdfMsg>) => {
    try {
        const res = await fetchBack(`/message/${chatId}/new`, {
            body: JSON.stringify(messages),
            method: "POST",
        });
        const body = await res.json();
        return body as Response<ChatpdfMsg>;
    } catch (err: any) {
        return {
            status: 'error', message: err.message || "Ocurrio un error inesperado"
        } as Response<null>;
    }
}


export const getLastChatpdfMessages = async (chatId: string) => {
    try {
        const res = await fetchBack(`/message/${chatId}/last`, {
            method: "GET",
        });
        const body = await res.json();
        return body as Response<ChatpdfMsg[]>;
    } catch (err: any) {
        console.log(err)
        return {
            data: [],
            status: 'error', message: err.message || "Ocurrio un error inesperado"
        } as Response<[]>;
    }
}

"use server"
import { Chatpdf, Response } from "@/interfaces";
import { fetchBack, } from "@/utils";



export const getAllChatpdf = async (): Promise<Response<Chatpdf[]>> => {
    return fetchBack("/chatpdf", {
        method: "GET",
    }).then((res) => {
        return res.json();
    }).catch((err) => {
        console.log(err);
        return {
            data: null,
            status: 'error',
            message: 'Error desconocido'
        } as any
    })
}

export const deleteChatpdf = async (id: string): Promise<Response<null>> => {
    return fetchBack("/chatpdf/" + id, {
        method: "DELETE",
    }).then((res) => {
        return res.json() as Promise<Response<null>>;
    }).catch((err) => {
        console.log(err);
        return {
            data: null,
            status: 'error',
            message: 'Error desconocido'
        } as any
    })
}


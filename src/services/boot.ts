import { Response } from "@/interfaces";
import { fetchBack } from "@/utils";

export const getContext = async (chatId: string, query: string): Promise<Response<string>> => {
    try {
        const res = await fetchBack(`/boot/${chatId}`, {
            body: JSON.stringify({ query }),
            method: "POST",
        });
        const body = await res.json() as Response<string>
        return body;
    } catch (err: any) {
        return {
            status: 'error', message: err.message || "Ocurrio un error inesperado"
        } as Response<string>;
    }
}
export const getContextWithoutAuth = async (chatId: string, query: string): Promise<Response<string>> => {
    try {
        const res = await fetchBack(`/bootWithout/${chatId}`, {
            body: JSON.stringify({ query }),
            method: "POST",
        });
        const body = await res.json() as Response<string>
        return body;
    } catch (err: any) {
        return {
            status: 'error', message: err.message || "Ocurrio un error inesperado"
        } as Response<string>;
    }
}

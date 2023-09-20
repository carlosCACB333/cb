'use server'
import { Chatpdf, Response } from "@/interfaces";
import { axBack } from "@/utils";
import { AxiosError } from "axios";

export const saveChatpdf = async (formData: FormData): Promise<Response<Chatpdf>> => {

    return (await axBack()).post<Response<Chatpdf>>("/chatpdf", formData).then((res) => {
        const data = res.data;
        return data
    }
    ).catch((err: AxiosError) => {
        const data = err.response?.data as Response<Chatpdf>;
        if (data) {
            return data;
        }
        return {
            data: null,
            status: 'error',
            message: 'Error desconocido'
        } as any
    })
}


export const getAllChatpdf = async (): Promise<Response<Chatpdf[]>> => {

    return (await axBack()).get<Response<Chatpdf[]>>("/chatpdf").then((res) => {
        const data = res.data;
        return data
    }
    ).catch((err: AxiosError) => {
        const data = err.response?.data as Response<Chatpdf[]>;
        if (data) {
            return data;
        }
        return {
            data: null,
            status: 'error',
            message: 'Error desconocido'
        } as any
    })
}
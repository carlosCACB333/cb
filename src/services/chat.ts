
import { Chatpdf, Response } from "@/interfaces";
import { axBack, axFront } from "@/utils";
import { AxiosError } from "axios";

export const saveChatpdf = async (file: File): Promise<Response<Chatpdf>> => {

    const formData = new FormData();
    formData.append("file", file);

    return axFront.post<Response<Chatpdf>>("/chatpdf", formData).then((res) => {
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
    return axBack.get<Response<Chatpdf[]>>("/chatpdf").then((res) => {
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
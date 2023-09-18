import { Response, User } from "@/interfaces";
import { axBack } from "@/utils";
import { AxiosError } from "axios";

export const createUser = async (user: Partial<User>): Promise<Response<User>> => {

    return axBack.post<Response<User>>("/user", user).then((res) => {
        const data = res.data;
        return data
    }
    ).catch((err: AxiosError) => {
        const data = err.response?.data as Response<User>;
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

export const updateUser = async (user: Partial<User>): Promise<Response<User>> => {

    return axBack.put<Response<User>>("/user", user).then((res) => {
        const data = res.data;
        return data
    }
    ).catch((err: AxiosError) => {
        const data = err.response?.data as Response<User>;
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


export const deleteUser = async (social_account: string): Promise<Response<User>> => {

    return axBack.delete<Response<User>>("/user/" + social_account).then((res) => {
        const data = res.data;
        return data
    }
    ).catch((err: AxiosError) => {
        const data = err.response?.data as Response<User>;
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

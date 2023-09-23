import { Response, User } from "@/interfaces";
import { fetchBack } from "@/utils";


export const createUser = async (user: Partial<User>): Promise<Response<User>> => {
    return fetchBack("/user", {
        method: "POST",
        body: JSON.stringify(user)
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

export const updateUser = async (user: Partial<User>): Promise<Response<User>> => {
    return fetchBack("/user", {
        method: "PUT",
        body: JSON.stringify(user)
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


export const deleteUser = async (id: string): Promise<Response<User>> => {
    return fetchBack("/user/" + id, {
        method: "DELETE",
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

import axios from 'axios'
import { env } from '.'
import { auth } from '@clerk/nextjs';



// solo para pruebas.El llamado debe funcional de back a back
export const axFront = axios.create({
    baseURL: env.back.publicUrl,
    headers: {
        Accept: "application/json",
        "x-api-key": "go123",
    }
})

export const axBack = async () => {
    const a = auth();
    const token = await a.getToken();
    return axios.create({
        baseURL: env.back.privateUrl,
        headers: {
            Accept: "application/json",
            "x-api-key": env.back.apiKey,
            "Authorization": "Bearer " + token
        }
    })
}



export const fetchBack = async (input: RequestInfo | URL, init?: RequestInit | undefined) => {
    input = typeof input === "string" ? env.back.privateUrl + input : input;
    init = init || {};
    init.headers = {
        ...init.headers,
        "x-api-key": env.back.apiKey,
        "Authorization": "Bearer " + await auth().getToken(),
        "Content-Type": "application/json",
    }
    return fetch(input, init)
}
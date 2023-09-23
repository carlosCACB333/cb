import { env } from '.'
import { auth } from '@clerk/nextjs';

// USE TO FETCH FROM FRONTEND
export const fetchFront = async (input: RequestInfo | URL, init?: RequestInit | undefined) => {
    input = typeof input === "string" ? env.back.publicUrl + input : input;
    init = init || {};
    init.headers = {
        "x-api-key": env.back.publicApiKey,
        "Content-Type": "application/json",
        ...init.headers,
    }
    return fetch(input, init)
}


// USE TO FETCH FROM BACKEND (SECURE)
export const fetchBack = async (input: RequestInfo | URL, init?: RequestInit | undefined) => {
    input = typeof input === "string" ? env.back.privateUrl + input : input;
    init = init || {};
    init.headers = {
        "x-api-key": env.back.apiKey,
        "Authorization": "Bearer " + await auth().getToken(),
        "Content-Type": "application/json",
        ...init.headers,
    }
    return fetch(input, init)
}

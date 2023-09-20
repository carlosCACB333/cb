import axios from 'axios'
import { env } from '.'
import { auth } from '@clerk/nextjs';



// solo para pruebas.El llamado debe funcional de back a back
export const axFront = axios.create({
    baseURL: env.back.publicUrl,
    headers: {
        Accept: "application/json",
        "x-api-key": "go123",
        Authorization:
            "Bearer " +
            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ3NzYzMzgsInVzZXJfaWQiOjJ9.NwKmmJoK8FcgCRvbV190ItpDI2AWvgmQ4E4v_6dbc6E",
    }
})

export const axBack = async () => {
    const a = auth();
    const token = await a.getToken();
    console.log("token", token);
    return axios.create({
        baseURL: env.back.privateUrl,
        headers: {
            Accept: "application/json",
            "x-api-key": env.back.apiKey,
            "Authorization": "Bearer " + token
        }
    })
}

import axios from 'axios'
import { env } from '.'




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
export const axBack = axios.create({
    baseURL: env.back.privateUrl,
    headers: {
        Accept: "application/json",
        "x-api-key": "go123",
        Authorization:
            "Bearer " +
            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ3NzYzMzgsInVzZXJfaWQiOjJ9.NwKmmJoK8FcgCRvbV190ItpDI2AWvgmQ4E4v_6dbc6E",
    }
})
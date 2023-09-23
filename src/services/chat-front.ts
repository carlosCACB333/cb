import { fetchFront } from "@/utils";


export const getFile = async (id: string, token: string): Promise<Blob | undefined> => {
    try {
        const res = await fetchFront(`/resource/${id}`, {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + token,
            }
        });
        const body = await res.blob();
        return body;
    } catch (err: any) {
        return undefined
    }
}
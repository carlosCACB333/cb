"use client";

import { axFront } from "@/utils";
import { useAuth } from "@clerk/nextjs";
import React, { FC, useEffect, useState } from "react";

interface Props {
  id: string;
}
export const PDFViewer: FC<Props> = ({ id }) => {
  const { getToken } = useAuth();
  const [url, setUrl] = useState("");

  useEffect(() => {
    const getFile = async () => {
      try {
        const res = await axFront.get(`/chatpdf/resource/${id}`, {
          responseType: "blob",
          headers: {
            Authorization: `Bearer ${await getToken()}`,
          },
        });
        const url = URL.createObjectURL(res.data);
        setUrl(url);
      } catch (err) {
        console.log(err);
      }
    };
    getFile();
  }, [getToken, id]);

  if (!url) return null;
  return <iframe className="h-full w-full" src={url} title="Tu pdf" />;
};

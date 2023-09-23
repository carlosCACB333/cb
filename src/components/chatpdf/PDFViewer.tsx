"use client";

import { getFile } from "@/services/chat-front";
import { useAuth } from "@clerk/nextjs";
import React, { FC, useEffect, useState } from "react";

interface Props {
  id: string;
}
export const PDFViewer: FC<Props> = ({ id }) => {
  const { getToken } = useAuth();
  const [url, setUrl] = useState("");

  useEffect(() => {
    const fetch = async () => {
      const token = await getToken();
      if (!token) return;
      const blob = await getFile(id, token);
      if (!blob) return;
      const url = URL.createObjectURL(blob);
      setUrl(url);
    };
    fetch();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id]);

  if (!url) return null;
  return (
    <iframe
      className="h-full min-h-[80vh] max-h-screen w-full"
      src={url}
      title="Tu pdf"
    />
  );
};

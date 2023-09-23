"use client";
import { env } from "@/utils";
import { useAuth } from "@clerk/nextjs";
import { CircularProgress } from "@nextui-org/react";
import clsx from "clsx";
import { useRouter } from "next/navigation";
import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";
import { FaFilePdf, FaRegFilePdf } from "react-icons/fa";
import { toast } from "react-toastify";

export const DropFile = () => {
  const { refresh } = useRouter();
  const [loading, setLoading] = useState(false);
  const { getToken } = useAuth();
  const onDrop = useCallback(
    async (acceptedFiles: File[]) => {
      try {
        setLoading(true);
        const file = acceptedFiles[0];
        if (!file) return;

        const formData = new FormData();
        formData.append("file", file, file.name);

        const res = await fetch(env.back.publicUrl + "/chatpdf", {
          method: "POST",
          body: formData,
          headers: {
            "x-api-key": env.back.publicApiKey,
            Authorization: "Bearer " + (await getToken()),
          },
        });
        const data = await res.json();
        refresh();
        toast(data.message, {
          type: data.status,
        });
      } catch (error) {
        toast("Error al subir el archivo", {
          type: "error",
        });
      } finally {
        setLoading(false);
      }
    },
    [refresh, getToken]
  );
  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    maxFiles: 1,
    accept: {
      "application/pdf": [".pdf"],
    },
  });

  return (
    <div
      {...getRootProps({
        className: clsx(
          "flex justify-center items-center flex-col w-full",
          "p-4 border-dashed border-2 border-gray-300 rounded-lg",
          "hover:border-primary hover:text-primary cursor-pointer",
          {
            "border-primary text-primary": isDragActive,
          }
        ),
      })}
    >
      <form>
        <input {...getInputProps()} disabled={loading} />
      </form>

      {loading ? (
        <>
          Espere un momento...
          <CircularProgress size="sm" />
        </>
      ) : isDragActive ? (
        <>
          Suelta el archivo aquí
          <FaFilePdf className="text-4xl" />
        </>
      ) : (
        <>
          Arrastra un archivo aquí
          <FaRegFilePdf className="text-4xl" />
        </>
      )}
    </div>
  );
};

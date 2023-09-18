"use client";

import { saveChatpdf } from "@/services";
import React, { useCallback } from "react";
import { useDropzone } from "react-dropzone";

export const DropFile = () => {
  const onDrop = useCallback(async (acceptedFiles: File[]) => {
    const file = acceptedFiles[0];
    if (!file) return;
    const res = await saveChatpdf(file);
    console.log(res);
  }, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    maxFiles: 1,
    accept: {
      "application/pdf": [".pdf"],
      "text/plain": [".txt"],
      "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
        [".docx"],
      "application/msword": [".doc"],
    },
  });

  return (
    <div
      {...getRootProps({
        className: "p-4 border-dashed border-2 border-gray-300 rounded-lg",
      })}
    >
      <input {...getInputProps()} />
      {isDragActive ? (
        <p>Sube tu archivo aquí ...</p>
      ) : (
        <p>Arrastra y suelta tu archivo aquí, o haz clic para seleccionar</p>
      )}
    </div>
  );
};

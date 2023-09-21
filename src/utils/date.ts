export const formatDate = (date: string | Date) => {
  return new Date(date).toLocaleDateString("es-PE", {
    weekday: "long",
    year: "numeric",
    month: "short",
    day: "numeric",
  });
};

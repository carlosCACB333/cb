export const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString("es-PE", {
    weekday: "long",
    year: "numeric",
    month: "short",
    day: "numeric",
  });
};

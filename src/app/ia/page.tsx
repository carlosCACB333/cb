import { title } from "@/utils";
import { FaRobot } from "react-icons/fa";

export default async function IAPage() {
  return (
    <>
      <h1 className={title({ color: "green" })}>
        Aquí encontrarás la inteligencia artificial que necesitas
      </h1>

      <section className="mt-8 grid grid-cols-1 gap-12 lg:grid-cols-2 xl:grid-cols-3">
        <div className="w-full p-4 text-center  rounded-lg shadow sm:p-8 bg-content1">
          <h5 className="mb-2 text-3xl font-bold text-gray-900 dark:text-white">
            Chat-pdf
          </h5>
          <p className="mb-5 text-base sm:text-lg ">
            Chatea directamente con esta inteligencia artificial sobre tus pdfs
          </p>

          <a
            href="/ia/chat-pdf"
            className="w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent rounded-md shadow-sm text-base font-medium text-white bg-green-500 hover:bg-green-600"
          >
            <FaRobot size={60} />
            <div className="text-left ml-2">
              <div className="mb-1 text-xs">Tienes una prueba gratuita</div>
              <div className="-mt-1 font-sans text-sm font-semibold">
                Probar ahora
              </div>
            </div>
          </a>
        </div>
      </section>
    </>
  );
}

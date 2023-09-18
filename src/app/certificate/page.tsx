"use client";

import { searchCertification } from "@/action";
import { sizes } from "@/assets";
import { IMG } from "@/components/common/IMG";
import { Pagination } from "@/components/common/Pagination";
import { Searcher } from "@/components/common/Searcher";
import { usePagination } from "@/hooks";

const Certification = () => {
  const { data, onSearch, onChangePage } = usePagination(searchCertification);

  return (
    <>
      <div className="container mx-auto mt-20 p-6">
        <div className="max-w-lg mx-auto mb-8 ">
          <Searcher
            size="lg"
            placeholder="Buscar certificados"
            setSearch={onSearch}
          />
        </div>
        <div className="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-2">
          {data?.edges?.map(({ node }) => (
            <div key={node.id} className="relative w-full aspect-[4/3]">
              <IMG
                src={node.picture.url}
                alt={node.name}
                fill
                className="rounded-xl"
                sizes={sizes.sm}
              />
            </div>
          ))}
        </div>
        {data?.pageInfo && (
          <Pagination {...data.pageInfo} onChangePage={onChangePage} />
        )}
      </div>
    </>
  );
};

export default Certification;

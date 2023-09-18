"use client";
import { usePagination } from "@/hooks";
import { searchPosts } from "@/action";
import { Searcher } from "../common/Searcher";
import { PostCard } from "./PostCard";
import { Pagination } from "../common/Pagination";

export const PostPaginated = () => {
  const { data, onSearch, onChangePage } = usePagination(searchPosts, 9);
  return (
    <div>
      <div className="max-w-lg mx-auto ">
        <Searcher
          size="lg"
          placeholder="Buscar posts..."
          setSearch={onSearch}
        />
      </div>
      <br />
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {data?.edges.map(({ node }) => (
          <PostCard key={node.id} post={node} />
        ))}
      </div>
      {data?.pageInfo && (
        <Pagination {...data.pageInfo} onChangePage={onChangePage} />
      )}
    </div>
  );
};

"use client";

import { SearchFunction } from "@/action";
import { Stage } from "@/generated/graphql";
import { useCallback, useRef, useState } from "react";
export const usePagination = <T>(search: SearchFunction<T>, initKeyword = "", pageSize = 12) => {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const skip = useRef(0);
  const keyword = useRef(initKeyword);

  const fetchData = useCallback(async () => {
    try {
      setLoading(true);
      const data = await search(
        keyword.current,
        pageSize,
        skip.current,
        Stage.Published
      );
      setData(data);
    } catch (error) {
      setError((error as any).message);
    } finally {
      setLoading(false);
    }
  }, [search, pageSize]);

  const onSearch = useCallback(
    async (k: string) => {
      keyword.current = k;
      skip.current = 0;
      await fetchData();
    },
    [fetchData]
  );

  const onChangePage = useCallback(
    async (next: boolean) => {
      skip.current += next ? pageSize : -pageSize;
      await fetchData();
    },
    [fetchData, pageSize]
  );

  return {
    data,
    loading,
    error,
    onChangePage,
    onSearch,
  };
};

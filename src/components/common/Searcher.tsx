"use client";

import { useDebounce } from "@/hooks";
import { Input, InputProps } from "@nextui-org/react";
import { useState, useEffect, memo } from "react";
import { Icon } from "./icon";

interface Props extends InputProps {
  setSearch: (v: string) => void;
}
const SearcherComponent = ({ setSearch, ...rest }: Props) => {
  const [value, setValue] = useState("");
  const search = useDebounce(value);
  useEffect(() => {
    setSearch(search);
  }, [search, setSearch]);

  return (
    <Input
      startContent={<Icon name="search" />}
      onValueChange={(v) => setValue(v || "")}
      value={value}
      aria-label="Buscar"
      {...rest}
    />
  );
};

export const Searcher = memo(SearcherComponent);

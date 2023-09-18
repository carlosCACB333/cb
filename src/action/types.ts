import { Stage } from "@/generated/graphql";
interface Args {
  [key: string]: any;
}
export type SearchFunction<T> = (
  keyword: string,
  first: number,
  skip: number,
  stage: Stage,
  args?: Args
) => Promise<T>;

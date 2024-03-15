export type PaginatedData<T> = {
  data: T[];
  total_count: number;
  page: number;
  total_pages: number;
  limit: number;
  has_more: boolean;
};

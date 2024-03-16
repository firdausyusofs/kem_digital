import { Supplier } from "./supplier";

export type Product = {
  id: string;
  name: string;
  description: string;
  price: number;
  supplier: Supplier;
};

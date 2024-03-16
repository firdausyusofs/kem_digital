import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit";
import { Product } from "../../types/product";
import { PaginatedData } from "../../types/pagination";
import apiService from "../../utils/apiService";

type FetchProducts = {
  page: number;
  limit: number;
  order_by?: string;
};

type CreateUpdateProductPayload = {
  id?: string;
  name: string;
  description: string;
  price: number;
  supplier_id: number;
};

export const fetchProducts = createAsyncThunk(
  "inventory/fetchProducts",
  async (reqData: FetchProducts, thunkAPI) => {
    try {
      const order_by = reqData.order_by ? reqData.order_by : "desc";
      const response = await apiService.get("/inventory?page=" + reqData.page + "&limit=" + reqData.limit + "&order_by=" + order_by);
      return response.data.data;
    } catch (error) {
      return thunkAPI.rejectWithValue(error);
    }
  }
);

export const createProduct = createAsyncThunk(
  "inventory/createProduct",
  async (product: CreateUpdateProductPayload, thunkAPI) => {
    try {
      const response = await apiService.post("/add-inventory", product);
      return response.data;
    } catch (error) {
      return thunkAPI.rejectWithValue(error);
    }
  }
);

export const updateProduct = createAsyncThunk(
  "inventory/updateProduct",
  async (product: CreateUpdateProductPayload, thunkAPI) => {
    try {
      const response = await apiService.put("/update-inventory/" + product.id!, product);
      return response.data;
    } catch (error) {
      return thunkAPI.rejectWithValue(error);
    }
  }
);

export const deleteProduct = createAsyncThunk(
  "inventory/deleteProduct",
  async (id: string, thunkAPI) => {
    try {
      const response = await apiService.delete("/delete-inventory/" + id);
      return response.data;
    } catch (error) {
      return thunkAPI.rejectWithValue(error);
    }
  }
);

type InitialState = {
  products: Product[];
  total_pages: number;
  loading: boolean;
  actionLoading: boolean;
  error: string | null;
  selectedProduct: Product | null;
};

const initialState = {
  products: [],
  total_pages: 0,
  loading: false,
  actionLoading: false,
  error: null,
  selectedProduct: null
} as InitialState;

export const inventory = createSlice({
  name: "inventory",
  initialState,
  reducers: {
    setSelectedProduct: (state, action: PayloadAction<Product>) => {
      state.selectedProduct = action.payload;
    },
    resetSelectedProduct: (state) => {
      state.selectedProduct = null;
    }
  },
  extraReducers(builder) {
    builder.addCase(fetchProducts.pending, (state) => {
      state.loading = true;
    });

    builder.addCase(fetchProducts.fulfilled, (state, action) => {
      state.loading = false;
      const data = action.payload as PaginatedData<Product>;
      state.products = data.data;
      state.total_pages = data.total_pages;
    });

    builder.addCase(fetchProducts.rejected, (state, action) => {
      state.loading = false;
      // state.error = action.payload;
    });

    builder.addCase(createProduct.pending, (state) => {
      state.actionLoading = true;
    });

    builder.addCase(createProduct.fulfilled, (state, action) => {
      state.actionLoading = false;
    });

    builder.addCase(createProduct.rejected, (state, action) => {
      state.actionLoading = false;
      // state.error = action.payload;
    });

    builder.addCase(deleteProduct.pending, (state) => {
      state.actionLoading = true;
    });

    builder.addCase(deleteProduct.fulfilled, (state, action) => {
      state.actionLoading = false;
    });

    builder.addCase(deleteProduct.rejected, (state, action) => {
      state.actionLoading = false;
      // state.error = action.payload;
    });

    builder.addCase(updateProduct.pending, (state) => {
      state.actionLoading = true;
    });

    builder.addCase(updateProduct.fulfilled, (state, action) => {
      state.actionLoading = false;
    });

    builder.addCase(updateProduct.rejected, (state, action) => {
      state.actionLoading = false;
      // state.error = action.payload;
    });
  }
});

export const { setSelectedProduct, resetSelectedProduct } = inventory.actions;
export default inventory.reducer;

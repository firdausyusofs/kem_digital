import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit";
import { Product } from "../../types/product";
import { PaginatedData } from "../../types/pagination";
import apiService from "../../utils/apiService";

export const fetchProducts = createAsyncThunk(
  "inventory/fetchProducts",
  async (reqData, thunkAPI) => {
    try {
      const response = await apiService.get("/inventory");
      return response.data.data;
    } catch (error) {
      return thunkAPI.rejectWithValue(error);
    }
  }
);

type InitialState = {
  products: Product[];
  total_pages: number;
  loading: boolean;
  error: string | null;
};

const initialState = {
  products: [],
  total_pages: 0,
  loading: false,
  error: null,
} as InitialState;

export const inventory = createSlice({
  name: "inventory",
  initialState,
  reducers: {},
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
  }
});

// export const { fetchInventoryStart, fetchInventorySuccess, fetchInventoryFailure } = inventory.actions;
export default inventory.reducer;

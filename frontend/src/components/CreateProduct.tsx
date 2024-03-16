import { Button, rem, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from "@mantine/form"
import React from 'react';
import { createProduct, updateProduct } from '../redux/features/inventory_slice';
import { useAppDispatch, useAppSelector } from '../redux/store';

interface CreateProductProps {
  close: () => void;
}

const CreateProduct: React.FC<CreateProductProps> = ({ close }) => {
  const dispatch = useAppDispatch();
  const { selectedProduct } = useAppSelector(state => state.inventoryReducer);
  const form = useForm({
    initialValues: {
      name: selectedProduct ? selectedProduct.name : '',
      description: selectedProduct ? selectedProduct.description : '',
      price: selectedProduct ? `${selectedProduct.price}` : '',
      supplier_id: selectedProduct ? `${selectedProduct.supplier.id}` : ''
    },
    validate: {
      name: (value) => value.length == 0,
      description: (value) => value.length == 0,
      price: (value) => value.length == 0,
      supplier_id: (value) => value.length == 0
    }
  });

  type FormValues = typeof form.values;

  const addProductSubmit = async (values: FormValues) => {
    const product = {
      name: values.name,
      description: values.description,
      price: parseFloat(values.price),
      supplier_id: parseInt(values.supplier_id)
    };

    if (selectedProduct != null) {
      dispatch(updateProduct({ ...product, id: selectedProduct.id }));
    } else {
      dispatch(createProduct(product));
    }
    close();
  };

  return (
    <Stack>
      <form onSubmit={form.onSubmit((values) => addProductSubmit(values))}>
        <Stack gap="md">
          <TextInput 
            label="Name" 
            placeholder="Enter name"
            {...form.getInputProps('name')}
          />
          <TextInput 
            label="Description" 
            placeholder="Enter description" 
            {...form.getInputProps('description')}
          />
          <TextInput 
            label="Price" 
            placeholder="Enter price" 
            {...form.getInputProps('price')}
          />
          <Select 
            label="Supplier" 
            placeholder="Select supplier" 
            data={[
              { label: 'Supplier 1', value: '1' },
              { label: 'Supplier 2', value: '2' },
              { label: 'Supplier 3', value: '3' },
            ]} 
            {...form.getInputProps('supplier_id')}
          />
          <Button mt={rem(20)} type="submit" variant="lf_primarybtn2" radius="xl">{selectedProduct != null ? "Save" : "Create" }</Button>
        </Stack>
      </form>
    </Stack>
  );
}

export default CreateProduct;

import { Button, rem, Select, Stack, TextInput } from '@mantine/core';
import React from 'react';

const CreateProduct: React.FC = () => {
  return (
    <Stack>
      <form>
        <Stack gap="md">
          <TextInput label="Name" placeholder="Enter name" />
          <TextInput label="Description" placeholder="Enter description" />
          <TextInput label="Price" placeholder="Enter price" />
          <Select label="Supplier" placeholder="Select supplier" data={['Supplier 1', 'Supplier 2', 'Supplier 3']} />
          <Button mt={rem(20)} type="submit" variant="lf_primarybtn2" radius="xl">Create Product</Button>
        </Stack>
      </form>
    </Stack>
  );
}

export default CreateProduct;

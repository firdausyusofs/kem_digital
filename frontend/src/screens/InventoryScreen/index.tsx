import { Button, Flex, Modal, Pagination, Paper, rem, Table, Title, Text } from '@mantine/core';
import { modals } from "@mantine/modals";
import { useDisclosure } from '@mantine/hooks';
import React from 'react';
import CreateProduct from '../../components/CreateProduct';
import PencilIcon from '../../icons/PencilIcon';
import PlusIcon from '../../icons/PlusIcon';
import TrashIcon from "../../icons/TrashIcon";
import { useAppDispatch, useAppSelector } from '../../redux/store'
import { fetchProducts } from "../../redux/features/inventory_slice"

const InventoryScreen: React.FC = () => {
    const dispatch = useAppDispatch();
    const { products, total_pages, loading } = useAppSelector(state => state.inventoryReducer)
    const [opened, {open, close}] = useDisclosure(false);
    const openDeleteModal = () => modals.openConfirmModal({
      title: 'Delete this product',
      centered: true,
      children: (
        <Text size="sm">
            Are you sure you want to delete this product?
        </Text>
      ),
      labels: { confirm: 'Delete', cancel: "Cancel" },
      confirmProps: { color: 'red' },
      onCancel: () => console.log('Cancel'),
      onConfirm: () => console.log('Confirmed'),
    });

    React.useEffect(() => {
        dispatch(fetchProducts());
    }, []);

    return (
        <>
            <Modal 
                opened={opened} 
                onClose={close} 
                size="lg"
                title="Add Product"
                centered
            >
                <CreateProduct />
            </Modal>
            <Paper 
                withBorder p="xl" 
                shadow="sm"
                radius="md" 
                m="xl" 
                mt={rem(110)}
            >
                <Flex justify="space-between">
                    <Title order={1} mb={rem(40)}>Inventory</Title>
                    <Button 
                        variant="lf_primarybtn" 
                        radius="xl"
                        leftSection={<PlusIcon />}
                        onClick={open}
                    >
                        Add Product
                    </Button>
                </Flex>
                <Table>
                    <Table.Thead>
                        <Table.Tr>
                            <Table.Th>Id</Table.Th>
                            <Table.Th>Name</Table.Th>
                            <Table.Th>Price</Table.Th>
                            <Table.Th>Supplier</Table.Th>
                            <Table.Th>Action</Table.Th>
                        </Table.Tr>
                    </Table.Thead>
                    <Table.Tbody>
                        {!loading && products.map((product, index) => (
                            <Table.Tr key={index}>
                                <Table.Td>{index+1}</Table.Td>
                                <Table.Td>{product.name}</Table.Td>
                                <Table.Td>RM {product.price.toFixed(2)}</Table.Td>
                                <Table.Td>Supplier 1</Table.Td>
                                <Table.Td>
                                    <Button 
                                        variant="lf_primarybtn" 
                                        radius="xl" 
                                        mr={rem(5)}
                                        leftSection={<PencilIcon />}
                                    >
                                        Edit
                                    </Button>
                                    <Button 
                                        variant="lf_dangerbtn"
                                        radius="xl"
                                        leftSection={<TrashIcon />}
                                        onClick={openDeleteModal}
                                    >
                                        Delete
                                    </Button>
                                </Table.Td>
                            </Table.Tr>
                        ))}
                    </Table.Tbody>
                </Table>

                <Flex justify="end">
                    <Pagination total={5} mt={rem(60)} />
                </Flex>
            </Paper>
        </>
  );
}

export default InventoryScreen;

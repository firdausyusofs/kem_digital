import { Button, Flex, Modal, Pagination, Paper, rem, Table, Title } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import React from 'react';
import CreateProduct from '../../components/CreateProduct';
import PencilIcon from '../../icons/PencilIcon';
import PlusIcon from '../../icons/PlusIcon';
import TrashIcon from "../../icons/TrashIcon";

const InventoryScreen: React.FC = () => {
    const [opened, {open, close}] = useDisclosure(false);

    return (
        <>
            <Modal 
                opened={opened} 
                onClose={close} 
                size="lg"
                title={
                    <Title order={3}>
                        Add Product
                    </Title>
                }
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
                        <Table.Tr>
                            <Table.Td>1</Table.Td>
                            <Table.Td>Product 1</Table.Td>
                            <Table.Td>RM 100.00</Table.Td>
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
                                >
                                    Delete
                                </Button>
                            </Table.Td>
                        </Table.Tr>
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

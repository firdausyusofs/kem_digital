import { Button, Flex, Modal, Pagination, Paper, rem, Table, Title, Text, Box, Group, MenuItem, Container } from '@mantine/core';
import { modals } from "@mantine/modals";
import { useDisclosure } from '@mantine/hooks';
import React from 'react';
import CreateProduct from '../../components/CreateProduct';
import PencilIcon from '../../icons/PencilIcon';
import PlusIcon from '../../icons/PlusIcon';
import TrashIcon from "../../icons/TrashIcon";
import { useAppDispatch, useAppSelector } from '../../redux/store'
import { deleteProduct, fetchProducts, resetSelectedProduct, setSelectedProduct } from "../../redux/features/inventory_slice"
import { MantineReactTable, useMantineReactTable, type MRT_ColumnDef, type MRT_PaginationState } from "mantine-react-table";

type ProductColumn = {
    index: number;
    name: string;
    description: string;
    price: number;
    supplier: string;
}

const InventoryScreen: React.FC = () => {
    const dispatch = useAppDispatch();
    const { products, total_pages, loading, actionLoading, selectedProduct } = useAppSelector(state => state.inventoryReducer)
    const [opened, {open, close}] = useDisclosure(false);
    const [pagination, setPagination] = React.useState<MRT_PaginationState>({
        pageIndex: 0,
        pageSize: 10,
    });

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
        onCancel: () => {
            dispatch(resetSelectedProduct());
        },
        onConfirm: () => {
            if (selectedProduct) handleDelete(selectedProduct.id)
        }
    });

    React.useEffect(() => {
        if (!actionLoading) {
            dispatch(fetchProducts({ page: pagination.pageIndex + 1, limit: pagination.pageSize }));
        }
    }, [pagination, actionLoading]);

    const columns = React.useMemo<MRT_ColumnDef<ProductColumn>[]>(
        () => [
            {
                accessorKey: "index",
                header: "No",
            },
            {
                accessorKey: "name",
                header: "Name",
            },
            {
                accessorKey: "description",
                header: "Description",
            },
            {
                accessorKey: "price",
                header: "Price",
            },
            {
                accessorKey: "supplier",
                header: "Supplier",
            },
        ],
        []
    );

    const data = React.useMemo(() => {
        return products.map((product, index) => {
            return {
                index: index+1 + ((pagination.pageIndex) * 10), 
                name: product.name,
                description: product.description,
                price: product.price,
                supplier: product.supplier.name,
            };
        });
    }, [products]);

    const table = useMantineReactTable({
        columns,
        data,
        enableRowActions: true,
        positionActionsColumn: "last",
        renderRowActions: ({ renderedRowIndex, row }) =>  (
            <Group justify="center">
                <Button 
                    variant="lf_primarybtn" 
                    radius="xl" 
                    leftSection={<PencilIcon />}
                    onClick={() => {
                        dispatch(setSelectedProduct(products[row.index]));
                        open();
                    }}
                >
                    Edit
                </Button>
                <Button 
                    variant="lf_dangerbtn"
                    radius="xl"
                    leftSection={<TrashIcon />}
                    onClick={() => {
                        dispatch(setSelectedProduct(products[row.index]));
                        openDeleteModal();
                    }}
                >
                    Delete
                </Button>
            </Group>
        ),
        manualPagination: true,
        rowCount: total_pages * 10,
        onPaginationChange: setPagination,
        state: {
            isLoading: loading,
            pagination,
        },
        paginationDisplayMode: "pages"
    })

    const handleDelete = (id: string) => {
        dispatch(deleteProduct(id));
    }

    return (
        <Container size="xl">
            <Modal 
                opened={opened} 
                onClose={() => {
                    close();
                    dispatch(resetSelectedProduct());
                }}
                size="lg"
                title={selectedProduct ? "Edit Product" : "Add Product" }
                centered
            >
                <CreateProduct close={close} />
            </Modal>
            <Box m="lg" mt={rem(110)}>
                <Group justify="space-between" align="center" mb={rem(40)}>
                    <Title order={1}>Inventory</Title>
                    <Button 
                        variant="lf_primarybtn" 
                        radius="xl"
                        leftSection={<PlusIcon />}
                        onClick={open}
                    >
                        Add Product
                    </Button>
                </Group>
                <MantineReactTable table={table} />
            </Box>
        </Container>
  );
}

export default InventoryScreen;

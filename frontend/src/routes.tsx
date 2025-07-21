import { createBrowserRouter } from 'react-router-dom'
import Layout from './layout/Layout'
import ItemsPage from './pages/ItemsPage'
import InvoicesPage from './pages/InvoicesPage'
import InvoicePrint from './pages/InvoicePrint'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      { path: 'items', element: <ItemsPage /> },
      { path: 'invoices', element: <InvoicesPage /> },
      { path: 'invoices/:id/print', element: <InvoicePrint /> },
    ],
  },
])

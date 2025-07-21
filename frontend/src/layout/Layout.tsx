import { Link, Outlet, useLocation } from 'react-router-dom'

export default function Layout() {
  const { pathname } = useLocation()
  return (
    <div className="min-h-screen bg-gray-100 text-gray-800">
      <header className="bg-white shadow px-4 py-3">
        <nav className="flex gap-6 font-semibold text-lg">
          <Link to="/items" className={pathname.startsWith('/items') ? 'text-blue-600' : ''}>Items</Link>
          <Link to="/invoices" className={pathname.startsWith('/invoices') ? 'text-blue-600' : ''}>Invoices</Link>
        </nav>
      </header>
      <main className="p-6">
        <Outlet />
      </main>
    </div>
  )
}

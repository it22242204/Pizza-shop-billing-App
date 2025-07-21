import { useItems, useCreateItem } from '../api/items'
import { useState } from 'react'

export default function ItemsPage() {
  const { data: items = [], isLoading } = useItems()
  const createItem = useCreateItem()

  const [form, setForm] = useState({ name: '', category: '', price: '', taxRate: '' })

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    createItem.mutate({
      ...form,
      price: Number(form.price),
      taxRate: Number(form.taxRate),
    })
    setForm({ name: '', category: '', price: '', taxRate: '' })
  }

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Manage Items</h1>
      <form className="grid grid-cols-4 gap-4 mb-6" onSubmit={handleSubmit}>
        <input placeholder="Name" required className="input" value={form.name} onChange={e => setForm(f => ({ ...f, name: e.target.value }))} />
        <select className="input" required value={form.category} onChange={e => setForm(f => ({ ...f, category: e.target.value }))}>
          <option value="">Select Category</option>
          <option value="pizza">Pizza</option>
          <option value="topping">Topping</option>
          <option value="beverage">Beverage</option>
        </select>
        <input placeholder="Price (cents)" className="input" value={form.price} onChange={e => setForm(f => ({ ...f, price: e.target.value }))} />
        <input placeholder="Tax Rate (e.g. 0.08)" className="input" value={form.taxRate} onChange={e => setForm(f => ({ ...f, taxRate: e.target.value }))} />
        <button type="submit" className="btn col-span-4 sm:col-span-1">Add</button>
      </form>

      {isLoading ? (
        <p>Loading...</p>
      ) : (
        <table className="w-full table-auto border">
          <thead><tr><th>Name</th><th>Category</th><th>Price</th><th>Tax</th></tr></thead>
          <tbody>
            {items.map(itm => (
              <tr key={itm.id} className="border-t">
                <td>{itm.name}</td>
                <td>{itm.category}</td>
                <td>{(itm.price / 100).toFixed(2)}</td>
                <td>{(itm.taxRate * 100).toFixed(0)}%</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  )
}

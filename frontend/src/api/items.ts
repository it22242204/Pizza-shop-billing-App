import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { api } from './client'

export type Item = {
  id: number
  name: string
  category: string
  price: number
  taxRate: number
}

export function useItems() {
  return useQuery<Item[]>({
    queryKey: ['items'],
    queryFn: async () => (await api.get('/items')).data,
  })
}

export function useCreateItem() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: (item: Partial<Item>) => api.post('/items', item),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['items'] }),
  })
}


import { SearchResult } from '../types'

export async function searchCoffeeShops(query: string): Promise<SearchResult> {
  const response = await fetch(`/api/search?q=${encodeURIComponent(query)}`)
  if (!response.ok) {
    throw new Error('Failed to fetch coffee shops')
  }
  return response.json()
}
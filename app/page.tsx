'use client'

import { useState, useEffect, useMemo } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import AnimatedPage from '../components/AnimatedPage'
import SearchBar from '../components/SearchBar'
import { FilterOptions } from '../components/FilterSort'
import ResultsList from '../components/ResultsList'
import { CoffeeShop } from '../types'
import { searchCoffeeShops } from '../lib/api'
import {Inter} from 'next/font/google'

export default function Home() {
  const [results, setResults] = useState<CoffeeShop[]>([])
  const [isSearching, setIsSearching] = useState(false)
  const [hasSearched, setHasSearched] = useState(false)
  const [sortBy, setSortBy] = useState('rating')
  const [filters, setFilters] = useState<FilterOptions>({
    minRating: 0,
    openNow: false,
    maxDistance: 10000,
  })
  const router = useRouter()
  const searchParams = useSearchParams()

  const fetchResults = async (query: string) => {
    setIsSearching(true)
    try {
      const data = await searchCoffeeShops(query)
      setResults(data.shops)
      setHasSearched(true)
    } catch (error) {
      console.error('Error fetching results:', error)
    } finally {
      setIsSearching(false)
    }
  }

  const handleSearch = (query: string) => {
    router.push(`/?q=${encodeURIComponent(query)}`)
    fetchResults(query)
  }

  const handleUseCurrentLocation = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords
          router.push(`/?lat=${latitude}&lon=${longitude}`)
          fetchResults(`${latitude},${longitude}`)
        },
        (error) => {
          console.error("Error getting location:", error)
          alert("Unable to retrieve your location. Please enter it manually.")
        }
      )
    } else {
      alert("Geolocation is not supported by your browser. Please enter your location manually.")
    }
  }

  const handleSortChange = (newSortBy: string) => {
    setSortBy(newSortBy)
  }

  const handleFilterChange = (newFilters: FilterOptions) => {
    setFilters(newFilters)
  }

  const filteredAndSortedResults = useMemo(() => {
    const filtered = results.filter(shop => 
      shop.rating >= filters.minRating &&
      shop.distance <= filters.maxDistance &&
      (!filters.openNow || isOpen(shop.hoursOfOperation))
    )

    filtered.sort((a, b) => {
      switch (sortBy) {
        case 'rating':
          return b.rating - a.rating
        case 'closingTime':
          return getMinutesUntilClose(b.closingTime) - getMinutesUntilClose(a.closingTime)
        case 'distance':
          return a.distance - b.distance
        default:
          return 0
      }
    })

    return filtered
  }, [results, filters, sortBy])

  useEffect(() => {
    const query = searchParams.get('q')
    const lat = searchParams.get('lat')
    const lon = searchParams.get('lon')
    if (query && results.length === 0) {
      fetchResults(query)
    } else if (lat && lon && results.length === 0) {
      fetchResults(`${lat},${lon}`)
    }
  }, [searchParams, results.length])

    return (
    <AnimatedPage>
      <div className="min-h-screen bg-background flex flex-col items-center justify-center p-4">
        <div className={`transition-all duration-500 ease-in-out ${hasSearched ? 'h-1/3' : 'h-screen'} flex flex-col items-center justify-center w-full max-w-7xl`}>
          <h1 className="text-5xl font-bold mb-8 text-onyx font-ivyora-display">Where to cafe?</h1>
          <div className="min-w-full px-4">
            <SearchBar 
              onSearch={handleSearch} 
              onFilterChange={handleFilterChange}
              onSortChange={handleSortChange}
              displayFilter={hasSearched}
              onUseCurrentLocation={handleUseCurrentLocation}
            />
          </div>
        </div>
        {isSearching && (
          <div className="mt-8">
            <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2"></div>
          </div>
        )}
        {filteredAndSortedResults.length > 0 && (
          <div className="mt-8 w-full max-w-7xl">
            <ResultsList results={filteredAndSortedResults} />
          </div>
        )}
      </div>
    </AnimatedPage>
  )
}

function isOpen(hoursOfOperation: string): boolean {
  // Implement logic to check if the shop is currently open
  // This is a placeholder implementation
  return true
}

function getMinutesUntilClose(closingTime: string): number {
  const now = new Date()
  const [hours, minutes] = closingTime.split(':').map(Number)
  const close = new Date(now.getFullYear(), now.getMonth(), now.getDate(), hours, minutes)
  return (close.getTime() - now.getTime()) / 60000
}
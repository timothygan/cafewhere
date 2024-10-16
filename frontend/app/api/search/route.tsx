import { NextRequest, NextResponse } from 'next/server'
import { CoffeeShop, SearchResult } from '../../../types'

export async function GET(request: NextRequest) {
  const { searchParams } = new URL(request.url)
  const q = searchParams.get('q')

  // In a real application, you would call your backend API here
  // For now, we'll return mock data
  const mockResults: CoffeeShop[] = [
    { 
      id: "1", 
      name: "Cafe Delight", 
      address: "123 Main St", 
      rating: 4.5, 
      hoursOfOperation: "7AM - 8PM",
      photoUrl: "https://images.unsplash.com/photo-1554118811-1e0d58224f24?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
      distance: 500, // 500 meters
      closingTime: "20:00"
    },
    { 
      id: "2", 
      name: "Brew Haven", 
      address: "456 Oak Ave", 
      rating: 4.2, 
      hoursOfOperation: "6AM - 9PM",
      photoUrl: "https://images.unsplash.com/photo-1559925393-8be0ec4767c8?ixlib=rb-1.2.1&auto=format&fit=crop&w=1351&q=80",
      distance: 1200, // 1.2 km
      closingTime: "21:00"
    },
    { 
      id: "3", 
      name: "Espresso Express", 
      address: "789 Pine Rd", 
      rating: 4.7, 
      hoursOfOperation: "8AM - 7PM",
      photoUrl: "https://images.unsplash.com/photo-1485182708500-e8f1f318ba72?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
      distance: 800, // 800 meters
      closingTime: "19:00"
    },
    { 
      id: "4", 
      name: "Espresso Express", 
      address: "789 Pine Rd", 
      rating: 3.7, 
      hoursOfOperation: "8AM - 7PM",
      photoUrl: "https://images.unsplash.com/photo-1485182708500-e8f1f318ba72?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
      distance: 800, // 800 meters
      closingTime: "19:00"
    },
    { 
      id: "5", 
      name: "Espresso Express", 
      address: "789 Pine Rd", 
      rating: 2.7, 
      hoursOfOperation: "8AM - 7PM",
      photoUrl: "https://images.unsplash.com/photo-1485182708500-e8f1f318ba72?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
      distance: 800, // 800 meters
      closingTime: "19:00"
    },
    { 
      id: "6", 
      name: "Espresso Express", 
      address: "789 Pine Rd", 
      rating: 1.7, 
      hoursOfOperation: "8AM - 7PM",
      photoUrl: "https://images.unsplash.com/photo-1485182708500-e8f1f318ba72?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
      distance: 800, // 800 meters
      closingTime: "19:00"
    },
    { 
      id: "7", 
      name: "Espresso Express", 
      address: "789 Pine Rd", 
      rating: 4.2, 
      hoursOfOperation: "8AM - 7PM",
      photoUrl: "https://images.unsplash.com/photo-1485182708500-e8f1f318ba72?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
      distance: 800, // 800 meters
      closingTime: "19:00"
    },

  ]

  const result: SearchResult = {
    shops: mockResults,
    total: mockResults.length
  }

  return NextResponse.json(result)
}
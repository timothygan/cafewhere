export interface CoffeeShop {
  id: string;
  name: string;
  address: string;
  rating: number;
  hoursOfOperation: string;
  photoUrl: string;
  distance: number; // in meters
  closingTime: string; // in 24-hour format, e.g., "22:00"
}

export interface SearchResult {
  shops: CoffeeShop[];
  total: number;
}
import React, { useState } from 'react';
import { FaFilter, FaSort } from 'react-icons/fa';
import { motion, AnimatePresence } from 'framer-motion';

interface FilterSortProps {
  onSortChange: (sortBy: string) => void;
  onFilterChange: (filters: FilterOptions) => void;
}

export interface FilterOptions {
  minRating: number;
  openNow: boolean;
  maxDistance: number;
}

const FilterSort: React.FC<FilterSortProps> = ({ onSortChange, onFilterChange }) => {
  const [isExpanded, setIsExpanded] = useState(false);
  const [filters, setFilters] = useState<FilterOptions>({
    minRating: 0,
    openNow: false,
    maxDistance: 10000,
  });
  const [sortBy, setSortBy] = useState('rating');

  const handleFilterChange = (key: keyof FilterOptions, value: number | boolean) => {
    const newFilters = { ...filters, [key]: value };
    setFilters(newFilters);
    onFilterChange(newFilters);
  };

  const handleSortChange = (value: string) => {
    setSortBy(value);
    onSortChange(value);
  };

  return (
    <div className="">
      <button
        onClick={() => setIsExpanded(!isExpanded)}
        className="flex items-center justify-center py-2 px-4 bg-indianred-100 text-indianred-700 rounded-full hover:bg-indianred-200 transition duration-300"
      >
        <FaFilter className="mr-2" />
        <span>Filter & Sort</span>
        <FaSort className="ml-2" />
      </button>
      
      <AnimatePresence>
        {isExpanded && (
          <motion.div
            initial={{ opacity: 0, y: -20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -20 }}
            transition={{ duration: 0.3 }}
            className="mt-4 bg-white p-4 rounded-lg shadow-md"
          >
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-onyx mb-1">Sort By</label>
                <select
                  value={sortBy}
                  onChange={(e) => handleSortChange(e.target.value)}
                  className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indianred focus:border-indianred sm:text-sm rounded-md"
                >
                  <option value="rating">Rating</option>
                  <option value="closingTime">Closing Time</option>
                  <option value="distance">Distance</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-onyx mb-1">Min Rating</label>
                <select
                  value={filters.minRating}
                  onChange={(e) => handleFilterChange('minRating', Number(e.target.value))}
                  className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indianred focus:border-indianred sm:text-sm rounded-md"
                >
                  <option value={0}>Any</option>
                  <option value={3}>3+</option>
                  <option value={4}>4+</option>
                  <option value={4.5}>4.5+</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Max Distance</label>
                <select
                  value={filters.maxDistance}
                  onChange={(e) => handleFilterChange('maxDistance', Number(e.target.value))}
                  className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indianred focus:border-indianred sm:text-sm rounded-md"
                >
                  <option value={1000}>1 km</option>
                  <option value={5000}>5 km</option>
                  <option value={10000}>10 km</option>
                  <option value={20000}>20 km</option>
                </select>
              </div>
              <div className="flex items-center">
                <input
                  id="openNow"
                  type="checkbox"
                  checked={filters.openNow}
                  onChange={(e) => handleFilterChange('openNow', e.target.checked)}
                  className="h-4 w-4 text-indianred-200 focus:ring-indianred border-gray-300 rounded"
                />
                <label htmlFor="openNow" className="ml-2 block text-sm text-gray-900">
                  Open Now
                </label>
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
};

export default FilterSort;
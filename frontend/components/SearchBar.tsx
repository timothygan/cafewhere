import { AnimatePresence, motion } from 'framer-motion';
import { useState, FormEvent } from 'react'
import { FaSearch, FaMapMarkerAlt } from 'react-icons/fa'
import FilterSort, { FilterOptions } from './FilterSort';

interface SearchBarProps {
  onSearch: (query: string) => void;
  onSortChange: (sortBy: string) => void;
  onFilterChange: (filters: FilterOptions) => void;
  displayFilter: boolean;
  onUseCurrentLocation: () => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch, onSortChange, onFilterChange, onUseCurrentLocation, displayFilter }) => {
  const [query, setQuery] = useState('')

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (query.trim()) {
      onSearch(query)
    }
  }

  return (
    <div className="flex justify-around items-center">
      <form onSubmit={handleSubmit} className="min-w-half">
        <div className="relative flex items-center">
          <input
          className="w-full py-4 px-6 pr-24 rounded-full bg-white shadow-lg text-lg focus:outline-none focus:ring-2 focus:ring-indianred transition duration-300"
          type="text"
          placeholder="Enter location to find coffee shops..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          />
          <button
          type="button"
          onClick={onUseCurrentLocation}
          className="absolute right-16 top-1/2 transform -translate-y-1/2 text-indianred hover:text-indianred-400 transition duration-300"
          title="Use current location"
          >
          <FaMapMarkerAlt className="text-xl" />
          </button>
          <button
          className="absolute right-3 top-1/2 transform -translate-y-1/2 bg-indianred hover:bg-indianred-400 text-white p-3 rounded-full transition duration-300"
          type="submit"
          >
          <FaSearch className="text-xl" />
          </button>
        </div>
      </form>
      <AnimatePresence>
        {displayFilter && (
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: 20 }}
            transition={{ duration: 0.3 }}
            className="min-w-fourth"
          >
            <FilterSort
              onSortChange={onSortChange}
              onFilterChange={onFilterChange}
            />
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  )
}

export default SearchBar
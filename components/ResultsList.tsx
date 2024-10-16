import { AnimatePresence, motion } from 'framer-motion'
import { CoffeeShop } from '../types'
import { FaStar, FaClock } from 'react-icons/fa'
import Image from 'next/image'

interface ResultsListProps {
  results: CoffeeShop[];
}

const ResultsList: React.FC<ResultsListProps> = ({ results }) => {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {results.map((shop, index) => (
        <AnimatePresence key={shop.id} mode='popLayout'>
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.3, delay: index * 0.1 }}
            className="bg-white shadow-lg rounded-lg overflow-hidden hover:shadow-xl transition duration-300"
          >
            <div className="relative h-48 w-full">
              <Image
                src={shop.photoUrl}
                alt={shop.name}
                layout="fill"
                objectFit="cover"
              />
            </div>
            <div className="p-6">
              <h2 className="text-2xl font-bold mb-2 text-onyx">{shop.name}</h2>
              <p className="text-gray-600 mb-4">{shop.address}</p>
              <div className="flex items-center mb-2">
                <FaStar className="text-yellow-400 mr-1" />
                <span className="font-semibold">{shop.rating.toFixed(1)}</span>
              </div>
              <div className="flex items-center text-gray-700">
                <FaClock className="mr-2" />
                <p>{shop.hoursOfOperation}</p>
              </div>
            </div>
          </motion.div>
        </AnimatePresence>
      ))}
    </div>
  )
}

export default ResultsList
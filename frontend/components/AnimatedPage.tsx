import { AnimatePresence, motion } from 'framer-motion'
import { ReactNode } from 'react'

const pageVariants = {
  initial: {
    opacity: 0,
    y: 20,
  },
  in: {
    opacity: 1,
    y: 0,
  },
  out: {
    opacity: 0,
    y: -20,
  },
}

const pageTransition = {
  type: 'tween',
  ease: 'anticipate',
  duration: 0.5,
}

interface AnimatedPageProps {
  children: ReactNode;
}

const AnimatedPage: React.FC<AnimatedPageProps> = ({ children }) => (
    <AnimatePresence mode='popLayout'>
        <motion.div
        initial="initial"
        animate="in"
        exit="out"
        variants={pageVariants}
        transition={pageTransition}
        >
        {children}
        </motion.div>
    </AnimatePresence>
)

export default AnimatedPage
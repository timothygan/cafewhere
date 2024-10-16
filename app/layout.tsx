import './globals.css'
import { ReactNode } from 'react'

export const metadata = {
  title: 'Coffee Shop Finder',
  description: 'Find the best coffee shops near you',
}

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <head>
        <link rel="stylesheet" href="https://use.typekit.net/smw1isx.css"/>
      </head>
      <body className='font-ivyora-text'>
        {children}
      </body>
    </html>
  )
}
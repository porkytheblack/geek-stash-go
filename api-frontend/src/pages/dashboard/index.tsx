import { pageProps } from '@/global/types'
import React from 'react'

function DashboardPage() {
  return (
    <div>DashboardPage</div>
  )
}

export default DashboardPage

export const getStaticProps = async () => {
    return {
      props: {
        authOnly: true,
        layout: 'dashboard'
      } satisfies pageProps
    } 
  } 
import { pageProps } from '@/global/types'
import React from 'react'

function Usage() {
  return (
    <div>Usage</div>
  )
}

export default Usage

export const getStaticProps = async () => {
    return {
        props: {
            authOnly: true,
            layout: 'dashboard'
        } satisfies pageProps
    }
}
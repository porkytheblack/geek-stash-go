import { pageProps } from '@/global/types'
import React from 'react'

function Keys() {
  return (
    <div>Keys</div>
  )
}

export default Keys

export const getStaticProps = async () => {
    return {
        props: {
            authOnly: true,
            layout: 'dashboard'
        } satisfies pageProps
    }
}
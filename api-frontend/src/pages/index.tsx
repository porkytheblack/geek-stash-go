import LinkButton from '@/components/Buttons/LinkButton'
import { pageProps } from '@/global/types'
import Link from 'next/link'
import React from 'react'

function Index () {
  return (
    <div className="grid items-center w-full  h-full justify-center">
      <Link
        href="/api/auth/login?returnTo=/dashboard"
        passHref
        legacyBehavior
      >
        <LinkButton
          buttonType='primary'
          onClick={()=>{}}
          
        >
          Login
        </LinkButton>
      </Link>
    </div>
  )
}

export default Index

export const getStaticProps = async () => {
  return {
    props: {
      authOnly: false,
      layout: 'home'
    } satisfies pageProps
  } 
} 